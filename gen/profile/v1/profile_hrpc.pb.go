// Code generated by protoc-gen-go-hrpc. DO NOT EDIT.

package profilev1

import (
	context "context"
	errors "errors"
	server "github.com/harmony-development/hrpc/server"
	proto "google.golang.org/protobuf/proto"
)

type ProfileServiceServer interface {
	// Gets a user's profile.
	GetProfile(context.Context, *GetProfileRequest) (*GetProfileResponse, error)
	// Updates the user's profile.
	UpdateProfile(context.Context, *UpdateProfileRequest) (*UpdateProfileResponse, error)
	// Gets app data for a user (this can be used to store user preferences which
	// is synchronized across devices).
	GetAppData(context.Context, *GetAppDataRequest) (*GetAppDataResponse, error)
	// Sets the app data for a user.
	SetAppData(context.Context, *SetAppDataRequest) (*SetAppDataResponse, error)
}

type DefaultProfileService struct{}

func (DefaultProfileService) GetProfile(context.Context, *GetProfileRequest) (*GetProfileResponse, error) {
	return nil, errors.New("unimplemented")
}
func (DefaultProfileService) UpdateProfile(context.Context, *UpdateProfileRequest) (*UpdateProfileResponse, error) {
	return nil, errors.New("unimplemented")
}
func (DefaultProfileService) GetAppData(context.Context, *GetAppDataRequest) (*GetAppDataResponse, error) {
	return nil, errors.New("unimplemented")
}
func (DefaultProfileService) SetAppData(context.Context, *SetAppDataRequest) (*SetAppDataResponse, error) {
	return nil, errors.New("unimplemented")
}

type ProfileServiceHandler struct {
	Server ProfileServiceServer
}

func NewProfileServiceHandler(server ProfileServiceServer) *ProfileServiceHandler {
	return &ProfileServiceHandler{Server: server}
}
func (h *ProfileServiceHandler) Name() string {
	return "ProfileService"
}
func (h *ProfileServiceHandler) Routes() map[string]server.RawHandler {
	return map[string]server.RawHandler{
		"/protocol.profile.v1.ProfileService.GetProfile/": server.NewUnaryHandler(&GetProfileRequest{}, func(c context.Context, req proto.Message) (proto.Message, error) {
			return h.Server.GetProfile(c, req.(*GetProfileRequest))
		}),
		"/protocol.profile.v1.ProfileService.UpdateProfile/": server.NewUnaryHandler(&UpdateProfileRequest{}, func(c context.Context, req proto.Message) (proto.Message, error) {
			return h.Server.UpdateProfile(c, req.(*UpdateProfileRequest))
		}),
		"/protocol.profile.v1.ProfileService.GetAppData/": server.NewUnaryHandler(&GetAppDataRequest{}, func(c context.Context, req proto.Message) (proto.Message, error) {
			return h.Server.GetAppData(c, req.(*GetAppDataRequest))
		}),
		"/protocol.profile.v1.ProfileService.SetAppData/": server.NewUnaryHandler(&SetAppDataRequest{}, func(c context.Context, req proto.Message) (proto.Message, error) {
			return h.Server.SetAppData(c, req.(*SetAppDataRequest))
		}),
	}
}
