package v1

import (
	"harmony-server/server/http/hm"
	"harmony-server/util"

	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/labstack/echo/v4"
)

func (h Handlers) GetInvites(c echo.Context) error {
	ctx, _ := c.(hm.HarmonyContext)

	invites, err := h.Deps.DB.GetInvites(*ctx.Location.GuildID)
	if err != nil {
		sentry.CaptureException(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to get invites")
	}
	return ctx.JSON(http.StatusOK, GetInvitesResponse{
		Invites: func() (ret []Invite) {
			for _, invite := range invites {
				ret = append(ret, Invite{
					ID:      invite.InviteID,
					GuildID: util.U64TS(invite.GuildID),
					Uses: func() int32 {
						if invite.PossibleUses.Valid {
							return invite.PossibleUses.Int32
						}
						return -1
					}(),
					UsedCount: invite.Uses,
				})
			}
			return
		}(),
	})
}
