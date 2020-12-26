package attachments

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gregjones/httpcache/diskcache"
	"github.com/harmony-development/legato/server/http/attachments/backend"
	"github.com/harmony-development/legato/server/http/hm"
	"github.com/harmony-development/legato/server/http/responses"
	"github.com/harmony-development/legato/server/http/routing"
	"github.com/labstack/echo/v4"
	"github.com/peterbourgon/diskv"
	"willnorris.com/go/imageproxy"
)

type Dependencies struct {
	APIGroup    *echo.Group
	Router      routing.IRouter
	FileBackend backend.AttachmentBackend
}

type API struct {
	*echo.Group
	Dependencies
	ImageProxy *imageproxy.Proxy
}

type UploadData struct {
	Filename    string `validate:"required"`
	ContentType string `validate:"required"`
}

type UploadResponse struct {
	ID string `json:"id"`
}

func (a *API) UploadHandler(c echo.Context) error {
	ctx := c.(hm.HarmonyContext)
	data := ctx.Data.(UploadData)

	if data.Filename == "" || data.ContentType == "" {
		return echo.NewHTTPError(http.StatusBadRequest, responses.MissingFilename)
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		fmt.Println("multipart form error: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	var handle io.Reader

	storedFiles, ok := form.Value["file"]
	if ok {
		if len(storedFiles) < 1 {
			return echo.NewHTTPError(http.StatusBadRequest, responses.MissingFiles)
		} else if len(storedFiles) > 1 {
			return echo.NewHTTPError(http.StatusBadRequest, responses.TooManyFiles)
		}

		handle = strings.NewReader(storedFiles[0])
	} else {
		files, ok := form.File["file"]
		if !ok || len(files) < 1 {
			return echo.NewHTTPError(http.StatusBadRequest, responses.MissingFiles)
		} else if len(files) > 1 {
			return echo.NewHTTPError(http.StatusBadRequest, responses.TooManyFiles)
		}

		file := files[0]

		handle, err = file.Open()
		if err != nil {
			fmt.Println("unable to open file: ", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
	}

	id, err := a.FileBackend.SaveFile(data.Filename, data.ContentType, handle)
	if err != nil {
		fmt.Println("unable to save file: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return ctx.JSON(http.StatusOK, UploadResponse{
		ID: id,
	})
}

func (a *API) DownloadHandler(c echo.Context) error {
	ctx := c.(hm.HarmonyContext)
	if ctx.Param("file_id") == "" {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	fileID := ctx.Param("file_id")

	if strings.HasPrefix(fileID, "http") {
		decoded, err := url.QueryUnescape(fileID)
		if err != nil {
			return err
		}

		fakeReq, err := http.NewRequest(http.MethodGet, "/"+decoded, nil)
		if err != nil {
			return err
		}
		a.ImageProxy.ServeHTTP(c.Response(), fakeReq)
		return nil
	}

	contentType, filename, _, handle, err := a.FileBackend.ReadFile(fileID)
	if err != nil {
		if err != backend.NotFound {
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return echo.NewHTTPError(http.StatusNotFound)
	}

	viewMode := "attachment"

	if strings.HasPrefix(contentType, "image/") || strings.HasPrefix(contentType, "video/") || strings.HasPrefix(contentType, "audio/") {
		viewMode = "inline"
	}

	defer handle.Close()

	c.Response().Header().Set("Content-Type", contentType)
	c.Response().Header().Set("Content-Disposition", fmt.Sprintf("%s; filename=\"%s\"", viewMode, filename))

	fileData, err := ioutil.ReadAll(handle)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	reader := bytes.NewReader(fileData)

	http.ServeContent(c.Response(), c.Request(), filename, time.Unix(0, 0), reader)

	return nil
}

func New(deps Dependencies) (*API, error) {
	api := &API{
		Group:        deps.APIGroup,
		Dependencies: deps,
		ImageProxy: imageproxy.NewProxy(nil, diskcache.NewWithDiskv(diskv.New(diskv.Options{
			BasePath:  "./imageproxycache",
			Transform: func(s string) []string { return []string{s[0:2], s[2:4]} },
		}))),
	}

	api.Router.BindRoutes(api.Group, []routing.Route{
		{
			Path:    "/upload",
			Handler: api.UploadHandler,
			Auth:    true,
			RateLimit: &routing.RateLimit{
				Duration: 1 * time.Second,
				Burst:    3,
			},
			Method: routing.POST,
			Schema: UploadData{},
		},
		{
			Path:    "/download/:file_id",
			Handler: api.DownloadHandler,
			Auth:    false,
			RateLimit: &routing.RateLimit{
				Duration: 1 * time.Second,
				Burst:    10,
			},
			Method: routing.GET,
		},
	})
	return api, nil
}
