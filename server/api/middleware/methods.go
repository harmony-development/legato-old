package middleware

import (
	"fmt"

	"github.com/harmony-development/hrpc/server"
	harmonytypesv1 "github.com/harmony-development/legato/gen/harmonytypes/v1"
	"github.com/harmony-development/legato/server/responses"
	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

var Methods map[string]*descriptorpb.MethodDescriptorProto

func (m Middlewares) MethodMetadataInterceptor(meth *descriptorpb.MethodDescriptorProto, serv *descriptorpb.ServiceDescriptorProto, d *descriptorpb.FileDescriptorProto, h server.Handler) server.Handler {
	return func(c echo.Context, req protoreflect.ProtoMessage) (protoreflect.ProtoMessage, error) {
		ctx := c.(HarmonyContext)
		opts := proto.GetExtension(meth.Options, harmonytypesv1.E_Metadata).(*harmonytypesv1.HarmonyMethodMetadata)
		if opts == nil {
			goto afterPermissions
		}

		{
			if !opts.RequiresAuthentication {
				goto afterAuth
			}
			userID, err := m.AuthHandler(ctx)
			if err != nil {
				return nil, err
			}
			ctx.UserID = userID
		}
	afterAuth:

		// Local
		{
			if !opts.RequiresLocal {
				goto afterLocal
			}

			isLocal, err := m.DB.UserIsLocal(ctx.UserID)
			if !isLocal || err != nil {
				return nil, err
			}
		}
	afterLocal:

		// Location
		{
			if err := m.LocationHandler(req, meth.GetName(), ctx.UserID); err != nil {
				return nil, err
			}
		}

		// Permissions
		{
			if opts.RequiresPermissionNode == "" {
				goto afterPermissions
			}

			location, ok := req.(interface {
				GetGuildId() uint64
			})
			if !ok {
				panic("guild permission middleware used on message without a location")
			}
			guildID := location.GetGuildId()
			owner, err := m.DB.GetOwner(guildID)
			if err != nil {
				return nil, err
			}
			if owner == ctx.UserID {
				ctx.IsOwner = true
				goto afterPermissions
			}

			channelID := uint64(0)
			channelLocation, ok := req.(interface {
				GetChannelId() uint64
			})
			if ok {
				channelID = channelLocation.GetChannelId()
			}

			roles, err := m.DB.RolesForUser(guildID, ctx.UserID)
			if err != nil {
				return nil, err
			}
			ctx.UserRoles = roles

			if !m.Perms.Check(opts.RequiresPermissionNode, roles, guildID, channelID) {
				return nil, responses.NewError(responses.NotEnoughPermissions)
			}
		}
	afterPermissions:
		if (opts != nil && opts.RequiresPermissionNode == "") || opts == nil {
			if GetRPCConfig(fmt.Sprintf("/%s.%s/%s", d.GetPackage(), serv.GetName(), meth.GetName())).WantsRoles {
				location, ok := req.(interface {
					GetGuildId() uint64
				})
				if !ok {
					panic("wants roles middleware used on message without a location")
				}
				roles, err := m.DB.RolesForUser(location.GetGuildId(), ctx.UserID)
				if err != nil {
					return nil, err
				}
				ctx.UserRoles = roles
				owner, err := m.DB.GetOwner(location.GetGuildId())
				if err != nil {
					return nil, err
				}

				ctx.IsOwner = owner == ctx.UserID
			}
		}
		return h(ctx, req)
	}
}
