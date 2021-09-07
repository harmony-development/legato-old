// SPDX-FileCopyrightText: 2021 None
//
// SPDX-License-Identifier: CC0-1.0

// Code generated by protoc-gen-go-hrpc. DO NOT EDIT.

package batchv1

import (
	bytes "bytes"
	context "context"
	proto "google.golang.org/protobuf/proto"
	ioutil "io/ioutil"
	http "net/http"
	httptest "net/http/httptest"
)

type BatchServiceClient interface {
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

type HTTPBatchServiceClient struct {
	Client  http.Client
	BaseURL string
}

func (client *HTTPBatchServiceClient) Batch(req *BatchRequest) (*BatchResponse, error) {
	data, marshalErr := proto.Marshal(req)
	if marshalErr != nil {
		return nil, marshalErr
	}
	reader := bytes.NewReader(data)
	resp, err := client.Client.Post(client.BaseURL+"/protocol.batch.v1.BatchService.Batch/", "application/hrpc", reader)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ret := &BatchResponse{}
	unmarshalErr := proto.Unmarshal(body, ret)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return ret, nil
}
func (client *HTTPBatchServiceClient) BatchSame(req *BatchSameRequest) (*BatchSameResponse, error) {
	data, marshalErr := proto.Marshal(req)
	if marshalErr != nil {
		return nil, marshalErr
	}
	reader := bytes.NewReader(data)
	resp, err := client.Client.Post(client.BaseURL+"/protocol.batch.v1.BatchService.BatchSame/", "application/hrpc", reader)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ret := &BatchSameResponse{}
	unmarshalErr := proto.Unmarshal(body, ret)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return ret, nil
}

type HTTPTestBatchServiceClient struct {
	Client interface {
		Test(*http.Request, ...int) (*http.Response, error)
	}
}

func (client *HTTPTestBatchServiceClient) Batch(req *BatchRequest) (*BatchResponse, error) {
	data, marshalErr := proto.Marshal(req)
	if marshalErr != nil {
		return nil, marshalErr
	}
	reader := bytes.NewReader(data)
	testreq := httptest.NewRequest("POST", "/protocol.batch.v1.BatchService.Batch/", reader)
	resp, err := client.Client.Test(testreq)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ret := &BatchResponse{}
	unmarshalErr := proto.Unmarshal(body, ret)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return ret, nil
}
func (client *HTTPTestBatchServiceClient) BatchSame(req *BatchSameRequest) (*BatchSameResponse, error) {
	data, marshalErr := proto.Marshal(req)
	if marshalErr != nil {
		return nil, marshalErr
	}
	reader := bytes.NewReader(data)
	testreq := httptest.NewRequest("POST", "/protocol.batch.v1.BatchService.BatchSame/", reader)
	resp, err := client.Client.Test(testreq)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ret := &BatchSameResponse{}
	unmarshalErr := proto.Unmarshal(body, ret)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return ret, nil
}
