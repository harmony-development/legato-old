package v1

import (
	"database/sql"
	"net/http"

	"github.com/harmony-development/legato/server/http/hm"
	"github.com/harmony-development/legato/server/http/responses"

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
