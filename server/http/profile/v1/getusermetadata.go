package v1

import (
	"database/sql"
	"harmony-server/server/http/hm"
	"harmony-server/server/http/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type GetUserMetadataData struct {
	appID string `validate:"required"`
}

func (h Handlers) GetUserMetadata(c echo.Context) error {
	ctx := c.(hm.HarmonyContext)
	data := ctx.Data.(GetUserMetadataData)
	meta, err := h.Deps.DB.GetUserMetadata(ctx.UserID, data.appID)
	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusNotFound, responses.MetadataNotFound)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, responses.UnknownError)
	}
	return ctx.JSON(http.StatusOK, GetUserMetadataResponse{
		Metadata: meta,
	})
}
