// SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

// Code generated by protoc-gen-go-hrpc. DO NOT EDIT.

package batchv1

import (
	context "context"
	errors "errors"
	server "github.com/harmony-development/hrpc/server"
	proto "google.golang.org/protobuf/proto"
)

type BatchServiceServer interface {
	// Batch requests.
	// Does not support batching stream requests.
	// Batched requests should be verified and an error should be thrown if they
	// are invalid.
	Batch(context.Context, *BatchRequest) (*BatchResponse, error)
	// BatchSame allows batching for requests using the same endpoint.
	// This allows for additional network optimizations since the endpoint doesn't
	// have to be sent for every request.
	BatchSame(context.Context, *BatchSameRequest) (*BatchSameResponse, error)
}

type DefaultBatchService struct{}

func (DefaultBatchService) Batch(context.Context, *BatchRequest) (*BatchResponse, error) {
	return nil, errors.New("unimplemented")
}
func (DefaultBatchService) BatchSame(context.Context, *BatchSameRequest) (*BatchSameResponse, error) {
	return nil, errors.New("unimplemented")
}

type BatchServiceHandler struct {
	Server BatchServiceServer
}

func NewBatchServiceHandler(server BatchServiceServer) *BatchServiceHandler {
	return &BatchServiceHandler{Server: server}
}
func (h *BatchServiceHandler) Name() string {
	return "BatchService"
}
func (h *BatchServiceHandler) Routes() map[string]server.RawHandler {
	return map[string]server.RawHandler{
		"/protocol.batch.v1.BatchService.Batch/": server.NewUnaryHandler(&BatchRequest{}, func(c context.Context, req proto.Message) (proto.Message, error) {
			return h.Server.Batch(c, req.(*BatchRequest))
		}),
		"/protocol.batch.v1.BatchService.BatchSame/": server.NewUnaryHandler(&BatchSameRequest{}, func(c context.Context, req proto.Message) (proto.Message, error) {
			return h.Server.BatchSame(c, req.(*BatchSameRequest))
		}),
	}
}
