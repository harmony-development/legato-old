// SPDX-FileCopyrightText: 2021 None
//
// SPDX-License-Identifier: CC0-1.0

// Code generated by protoc-gen-go-hrpc. DO NOT EDIT.

package emotev1

import (
	bytes "bytes"
	context "context"
	proto "google.golang.org/protobuf/proto"
	ioutil "io/ioutil"
	http "net/http"
	httptest "net/http/httptest"
)

type EmoteServiceClient interface {
	// Endpoint to create an emote pack.
	CreateEmotePack(context.Context, *CreateEmotePackRequest) (*CreateEmotePackResponse, error)
	// Endpoint to get the emote packs you have equipped.
	GetEmotePacks(context.Context, *GetEmotePacksRequest) (*GetEmotePacksResponse, error)
	// Endpoint to get the emotes in an emote pack.
	GetEmotePackEmotes(context.Context, *GetEmotePackEmotesRequest) (*GetEmotePackEmotesResponse, error)
	// Endpoint to add an emote to an emote pack that you own.
	AddEmoteToPack(context.Context, *AddEmoteToPackRequest) (*AddEmoteToPackResponse, error)
	// Endpoint to delete an emote pack that you own.
	DeleteEmotePack(context.Context, *DeleteEmotePackRequest) (*DeleteEmotePackResponse, error)
	// Endpoint to delete an emote from an emote pack.
	DeleteEmoteFromPack(context.Context, *DeleteEmoteFromPackRequest) (*DeleteEmoteFromPackResponse, error)
	// Endpoint to dequip an emote pack that you have equipped.
	DequipEmotePack(context.Context, *DequipEmotePackRequest) (*DequipEmotePackResponse, error)
	// Endpoint to equip an emote pack.
	EquipEmotePack(context.Context, *EquipEmotePackRequest) (*EquipEmotePackResponse, error)
}

type HTTPEmoteServiceClient struct {
	Client  http.Client
	BaseURL string
}

func (client *HTTPEmoteServiceClient) CreateEmotePack(req *CreateEmotePackRequest) (*CreateEmotePackResponse, error) {
	data, marshalErr := proto.Marshal(req)
	if marshalErr != nil {
		return nil, marshalErr
	}
	reader := bytes.NewReader(data)
	resp, err := client.Client.Post(client.BaseURL+"/protocol.emote.v1.EmoteService.CreateEmotePack/", "application/hrpc", reader)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ret := &CreateEmotePackResponse{}
	unmarshalErr := proto.Unmarshal(body, ret)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return ret, nil
}
func (client *HTTPEmoteServiceClient) GetEmotePacks(req *GetEmotePacksRequest) (*GetEmotePacksResponse, error) {
	data, marshalErr := proto.Marshal(req)
	if marshalErr != nil {
		return nil, marshalErr
	}
	reader := bytes.NewReader(data)
	resp, err := client.Client.Post(client.BaseURL+"/protocol.emote.v1.EmoteService.GetEmotePacks/", "application/hrpc", reader)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ret := &GetEmotePacksResponse{}
	unmarshalErr := proto.Unmarshal(body, ret)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return ret, nil
}
func (client *HTTPEmoteServiceClient) GetEmotePackEmotes(req *GetEmotePackEmotesRequest) (*GetEmotePackEmotesResponse, error) {
	data, marshalErr := proto.Marshal(req)
	if marshalErr != nil {
		return nil, marshalErr
	}
	reader := bytes.NewReader(data)
	resp, err := client.Client.Post(client.BaseURL+"/protocol.emote.v1.EmoteService.GetEmotePackEmotes/", "application/hrpc", reader)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ret := &GetEmotePackEmotesResponse{}
	unmarshalErr := proto.Unmarshal(body, ret)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return ret, nil
}
func (client *HTTPEmoteServiceClient) AddEmoteToPack(req *AddEmoteToPackRequest) (*AddEmoteToPackResponse, error) {
	data, marshalErr := proto.Marshal(req)
	if marshalErr != nil {
		return nil, marshalErr
	}
	reader := bytes.NewReader(data)
	resp, err := client.Client.Post(client.BaseURL+"/protocol.emote.v1.EmoteService.AddEmoteToPack/", "application/hrpc", reader)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ret := &AddEmoteToPackResponse{}
	unmarshalErr := proto.Unmarshal(body, ret)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return ret, nil
}
func (client *HTTPEmoteServiceClient) DeleteEmotePack(req *DeleteEmotePackRequest) (*DeleteEmotePackResponse, error) {
	data, marshalErr := proto.Marshal(req)
	if marshalErr != nil {
		return nil, marshalErr
	}
	reader := bytes.NewReader(data)
	resp, err := client.Client.Post(client.BaseURL+"/protocol.emote.v1.EmoteService.DeleteEmotePack/", "application/hrpc", reader)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ret := &DeleteEmotePackResponse{}
	unmarshalErr := proto.Unmarshal(body, ret)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return ret, nil
}
func (client *HTTPEmoteServiceClient) DeleteEmoteFromPack(req *DeleteEmoteFromPackRequest) (*DeleteEmoteFromPackResponse, error) {
	data, marshalErr := proto.Marshal(req)
	if marshalErr != nil {
		return nil, marshalErr
	}
	reader := bytes.NewReader(data)
	resp, err := client.Client.Post(client.BaseURL+"/protocol.emote.v1.EmoteService.DeleteEmoteFromPack/", "application/hrpc", reader)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ret := &DeleteEmoteFromPackResponse{}
	unmarshalErr := proto.Unmarshal(body, ret)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return ret, nil
}
func (client *HTTPEmoteServiceClient) DequipEmotePack(req *DequipEmotePackRequest) (*DequipEmotePackResponse, error) {
	data, marshalErr := proto.Marshal(req)
	if marshalErr != nil {
		return nil, marshalErr
	}
	reader := bytes.NewReader(data)
	resp, err := client.Client.Post(client.BaseURL+"/protocol.emote.v1.EmoteService.DequipEmotePack/", "application/hrpc", reader)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ret := &DequipEmotePackResponse{}
	unmarshalErr := proto.Unmarshal(body, ret)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return ret, nil
}
func (client *HTTPEmoteServiceClient) EquipEmotePack(req *EquipEmotePackRequest) (*EquipEmotePackResponse, error) {
	data, marshalErr := proto.Marshal(req)
	if marshalErr != nil {
		return nil, marshalErr
	}
	reader := bytes.NewReader(data)
	resp, err := client.Client.Post(client.BaseURL+"/protocol.emote.v1.EmoteService.EquipEmotePack/", "application/hrpc", reader)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ret := &EquipEmotePackResponse{}
	unmarshalErr := proto.Unmarshal(body, ret)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return ret, nil
}

type HTTPTestEmoteServiceClient struct {
	Client interface {
		Test(*http.Request, ...int) (*http.Response, error)
	}
}

func (client *HTTPTestEmoteServiceClient) CreateEmotePack(req *CreateEmotePackRequest) (*CreateEmotePackResponse, error) {
	data, marshalErr := proto.Marshal(req)
	if marshalErr != nil {
		return nil, marshalErr
	}
	reader := bytes.NewReader(data)
	testreq := httptest.NewRequest("POST", "/protocol.emote.v1.EmoteService.CreateEmotePack/", reader)
	resp, err := client.Client.Test(testreq)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ret := &CreateEmotePackResponse{}
	unmarshalErr := proto.Unmarshal(body, ret)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return ret, nil
}
func (client *HTTPTestEmoteServiceClient) GetEmotePacks(req *GetEmotePacksRequest) (*GetEmotePacksResponse, error) {
	data, marshalErr := proto.Marshal(req)
	if marshalErr != nil {
		return nil, marshalErr
	}
	reader := bytes.NewReader(data)
	testreq := httptest.NewRequest("POST", "/protocol.emote.v1.EmoteService.GetEmotePacks/", reader)
	resp, err := client.Client.Test(testreq)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ret := &GetEmotePacksResponse{}
	unmarshalErr := proto.Unmarshal(body, ret)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return ret, nil
}
func (client *HTTPTestEmoteServiceClient) GetEmotePackEmotes(req *GetEmotePackEmotesRequest) (*GetEmotePackEmotesResponse, error) {
	data, marshalErr := proto.Marshal(req)
	if marshalErr != nil {
		return nil, marshalErr
	}
	reader := bytes.NewReader(data)
	testreq := httptest.NewRequest("POST", "/protocol.emote.v1.EmoteService.GetEmotePackEmotes/", reader)
	resp, err := client.Client.Test(testreq)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ret := &GetEmotePackEmotesResponse{}
	unmarshalErr := proto.Unmarshal(body, ret)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return ret, nil
}
func (client *HTTPTestEmoteServiceClient) AddEmoteToPack(req *AddEmoteToPackRequest) (*AddEmoteToPackResponse, error) {
	data, marshalErr := proto.Marshal(req)
	if marshalErr != nil {
		return nil, marshalErr
	}
	reader := bytes.NewReader(data)
	testreq := httptest.NewRequest("POST", "/protocol.emote.v1.EmoteService.AddEmoteToPack/", reader)
	resp, err := client.Client.Test(testreq)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ret := &AddEmoteToPackResponse{}
	unmarshalErr := proto.Unmarshal(body, ret)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return ret, nil
}
func (client *HTTPTestEmoteServiceClient) DeleteEmotePack(req *DeleteEmotePackRequest) (*DeleteEmotePackResponse, error) {
	data, marshalErr := proto.Marshal(req)
	if marshalErr != nil {
		return nil, marshalErr
	}
	reader := bytes.NewReader(data)
	testreq := httptest.NewRequest("POST", "/protocol.emote.v1.EmoteService.DeleteEmotePack/", reader)
	resp, err := client.Client.Test(testreq)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ret := &DeleteEmotePackResponse{}
	unmarshalErr := proto.Unmarshal(body, ret)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return ret, nil
}
func (client *HTTPTestEmoteServiceClient) DeleteEmoteFromPack(req *DeleteEmoteFromPackRequest) (*DeleteEmoteFromPackResponse, error) {
	data, marshalErr := proto.Marshal(req)
	if marshalErr != nil {
		return nil, marshalErr
	}
	reader := bytes.NewReader(data)
	testreq := httptest.NewRequest("POST", "/protocol.emote.v1.EmoteService.DeleteEmoteFromPack/", reader)
	resp, err := client.Client.Test(testreq)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ret := &DeleteEmoteFromPackResponse{}
	unmarshalErr := proto.Unmarshal(body, ret)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return ret, nil
}
func (client *HTTPTestEmoteServiceClient) DequipEmotePack(req *DequipEmotePackRequest) (*DequipEmotePackResponse, error) {
	data, marshalErr := proto.Marshal(req)
	if marshalErr != nil {
		return nil, marshalErr
	}
	reader := bytes.NewReader(data)
	testreq := httptest.NewRequest("POST", "/protocol.emote.v1.EmoteService.DequipEmotePack/", reader)
	resp, err := client.Client.Test(testreq)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ret := &DequipEmotePackResponse{}
	unmarshalErr := proto.Unmarshal(body, ret)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return ret, nil
}
func (client *HTTPTestEmoteServiceClient) EquipEmotePack(req *EquipEmotePackRequest) (*EquipEmotePackResponse, error) {
	data, marshalErr := proto.Marshal(req)
	if marshalErr != nil {
		return nil, marshalErr
	}
	reader := bytes.NewReader(data)
	testreq := httptest.NewRequest("POST", "/protocol.emote.v1.EmoteService.EquipEmotePack/", reader)
	resp, err := client.Client.Test(testreq)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ret := &EquipEmotePackResponse{}
	unmarshalErr := proto.Unmarshal(body, ret)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return ret, nil
}
