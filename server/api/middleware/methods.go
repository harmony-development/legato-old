package middleware

import (
	"context"
	"fmt"

	harmonytypesv1 "github.com/harmony-development/legato/gen/harmonytypes/v1"
	"github.com/harmony-development/legato/server/responses"
	"github.com/jhump/protoreflect/desc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

var Methods map[string]*desc.MethodDescriptor

func (m Middlewares) MethodMetadataInterceptor(c context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	desc, err := protoregistry.GlobalFiles.FindDescriptorByName(protoreflect.FullName(info.FullMethod))
	if err != nil {
		return nil, fmt.Errorf("unknown method: %w", err)
	}
	method := desc.(protoreflect.MethodDescriptor)
	opts := proto.GetExtension(method.Options(), harmonytypesv1.E_Metadata).(*harmonytypesv1.HarmonyMethodMetadata)

	ctx := c.(HarmonyContext)

	// Auth
	{
		if opts.RequiresAuthentication {
			goto afterAuth
		}

		userid, err := AuthHandler(m.DB, ctx)

		if err != nil {
			return nil, err
		}

		ctx.UserID = userid
	}
afterAuth:

	// Local
	{
		if !opts.RequiresLocal {
			goto afterLocal
		}

		err := m.DB.UserIsLocal(ctx.UserID)
		if err != nil {
			return nil, err
		}
	}
afterLocal:

	// Location
	{
		if err := LocationHandler(m.DB, req, info.FullMethod, ctx.UserID); err != nil {
			return nil, err
		}
	}

	// Permissions
	{
		if opts.RequiresPermissionNode == "" {
			if GetRPCConfig(info.FullMethod).WantsRoles {
				location, ok := req.(interface {
					GetGuildId() uint64
				})
				if !ok {
					panic("wants roles middleware used on message without a location")
				}
				roles, err := m.DB.RolesForUser(location.GetGuildId(), ctx.UserID)
				if err != nil {
					return nil, status.Error(codes.Internal, responses.InternalServerError)
				}
				ctx.UserRoles = roles
				owner, err := m.DB.GetOwner(location.GetGuildId())
				if err != nil {
					return nil, status.Error(codes.Internal, responses.InternalServerError)
				}
				ctx.IsOwner = owner == ctx.UserID
			}
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
			return nil, status.Error(codes.Internal, responses.InternalServerError)
		}
		if owner == ctx.UserID {
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
			return nil, status.Error(codes.Internal, responses.InternalServerError)
		}
		ctx.UserRoles = roles

		if !m.Perms.Check(opts.RequiresPermissionNode, roles, guildID, channelID) {
			return nil, status.Error(codes.PermissionDenied, responses.InsufficientPrivileges)
		}
	}
afterPermissions:

	return handler(ctx, req)
}
