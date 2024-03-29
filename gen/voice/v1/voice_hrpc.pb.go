// SPDX-FileCopyrightText: 2021 None
//
// SPDX-License-Identifier: CC0-1.0

// Code generated by protoc-gen-go-hrpc. DO NOT EDIT.

package voicev1

import (
	context "context"
	errors "errors"
	server "github.com/harmony-development/hrpc/server"
	proto "google.golang.org/protobuf/proto"
)

type VoiceServiceServer interface {
	// Endpoint to connect to a voice channel.
	Connect(context.Context, *ConnectRequest) (*ConnectResponse, error)
	// Endpoint to stream states of a voice connection.
	StreamState(context.Context, *StreamStateRequest) (chan *StreamStateResponse, error)
}

type DefaultVoiceService struct{}

func (DefaultVoiceService) Connect(context.Context, *ConnectRequest) (*ConnectResponse, error) {
	return nil, errors.New("unimplemented")
}
func (DefaultVoiceService) StreamState(context.Context, *StreamStateRequest) (chan *StreamStateResponse, error) {
	return nil, errors.New("unimplemented")
}

type VoiceServiceHandler struct {
	Server VoiceServiceServer
}

func NewVoiceServiceHandler(server VoiceServiceServer) *VoiceServiceHandler {
	return &VoiceServiceHandler{Server: server}
}
func (h *VoiceServiceHandler) Name() string {
	return "VoiceService"
}
func (h *VoiceServiceHandler) Routes() map[string]server.RawHandler {
	return map[string]server.RawHandler{
		"/protocol.voice.v1.VoiceService.Connect/": server.NewUnaryHandler(&ConnectRequest{}, func(c context.Context, req proto.Message) (proto.Message, error) {
			return h.Server.Connect(c, req.(*ConnectRequest))
		}),
	}
}
