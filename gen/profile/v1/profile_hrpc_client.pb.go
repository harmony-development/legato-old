// SPDX-FileCopyrightText: 2021 None
//
// SPDX-License-Identifier: CC0-1.0

// Code generated by protoc-gen-go-hrpc. DO NOT EDIT.

package profilev1

import (
	bytes "bytes"
	context "context"
	proto "google.golang.org/protobuf/proto"
	ioutil "io/ioutil"
	http "net/http"
	httptest "net/http/httptest"
)

type ProfileServiceClient interface {
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

type HTTPProfileServiceClient struct {
	Client  http.Client
	BaseURL string
}

func (client *HTTPProfileServiceClient) GetProfile(req *GetProfileRequest) (*GetProfileResponse, error) {
	data, marshalErr := proto.Marshal(req)
	if marshalErr != nil {
		return nil, marshalErr
	}
	reader := bytes.NewReader(data)
	resp, err := client.Client.Post(client.BaseURL+"/protocol.profile.v1.ProfileService.GetProfile/", "application/hrpc", reader)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ret := &GetProfileResponse{}
	unmarshalErr := proto.Unmarshal(body, ret)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return ret, nil
}
func (client *HTTPProfileServiceClient) UpdateProfile(req *UpdateProfileRequest) (*UpdateProfileResponse, error) {
	data, marshalErr := proto.Marshal(req)
	if marshalErr != nil {
		return nil, marshalErr
	}
	reader := bytes.NewReader(data)
	resp, err := client.Client.Post(client.BaseURL+"/protocol.profile.v1.ProfileService.UpdateProfile/", "application/hrpc", reader)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ret := &UpdateProfileResponse{}
	unmarshalErr := proto.Unmarshal(body, ret)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return ret, nil
}
func (client *HTTPProfileServiceClient) GetAppData(req *GetAppDataRequest) (*GetAppDataResponse, error) {
	data, marshalErr := proto.Marshal(req)
	if marshalErr != nil {
		return nil, marshalErr
	}
	reader := bytes.NewReader(data)
	resp, err := client.Client.Post(client.BaseURL+"/protocol.profile.v1.ProfileService.GetAppData/", "application/hrpc", reader)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ret := &GetAppDataResponse{}
	unmarshalErr := proto.Unmarshal(body, ret)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return ret, nil
}
func (client *HTTPProfileServiceClient) SetAppData(req *SetAppDataRequest) (*SetAppDataResponse, error) {
	data, marshalErr := proto.Marshal(req)
	if marshalErr != nil {
		return nil, marshalErr
	}
	reader := bytes.NewReader(data)
	resp, err := client.Client.Post(client.BaseURL+"/protocol.profile.v1.ProfileService.SetAppData/", "application/hrpc", reader)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ret := &SetAppDataResponse{}
	unmarshalErr := proto.Unmarshal(body, ret)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return ret, nil
}

type HTTPTestProfileServiceClient struct {
	Client interface {
		Test(*http.Request, ...int) (*http.Response, error)
	}
}

func (client *HTTPTestProfileServiceClient) GetProfile(req *GetProfileRequest) (*GetProfileResponse, error) {
	data, marshalErr := proto.Marshal(req)
	if marshalErr != nil {
		return nil, marshalErr
	}
	reader := bytes.NewReader(data)
	testreq := httptest.NewRequest("POST", "/protocol.profile.v1.ProfileService.GetProfile/", reader)
	resp, err := client.Client.Test(testreq)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ret := &GetProfileResponse{}
	unmarshalErr := proto.Unmarshal(body, ret)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return ret, nil
}
func (client *HTTPTestProfileServiceClient) UpdateProfile(req *UpdateProfileRequest) (*UpdateProfileResponse, error) {
	data, marshalErr := proto.Marshal(req)
	if marshalErr != nil {
		return nil, marshalErr
	}
	reader := bytes.NewReader(data)
	testreq := httptest.NewRequest("POST", "/protocol.profile.v1.ProfileService.UpdateProfile/", reader)
	resp, err := client.Client.Test(testreq)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ret := &UpdateProfileResponse{}
	unmarshalErr := proto.Unmarshal(body, ret)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return ret, nil
}
func (client *HTTPTestProfileServiceClient) GetAppData(req *GetAppDataRequest) (*GetAppDataResponse, error) {
	data, marshalErr := proto.Marshal(req)
	if marshalErr != nil {
		return nil, marshalErr
	}
	reader := bytes.NewReader(data)
	testreq := httptest.NewRequest("POST", "/protocol.profile.v1.ProfileService.GetAppData/", reader)
	resp, err := client.Client.Test(testreq)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ret := &GetAppDataResponse{}
	unmarshalErr := proto.Unmarshal(body, ret)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return ret, nil
}
func (client *HTTPTestProfileServiceClient) SetAppData(req *SetAppDataRequest) (*SetAppDataResponse, error) {
	data, marshalErr := proto.Marshal(req)
	if marshalErr != nil {
		return nil, marshalErr
	}
	reader := bytes.NewReader(data)
	testreq := httptest.NewRequest("POST", "/protocol.profile.v1.ProfileService.SetAppData/", reader)
	resp, err := client.Client.Test(testreq)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ret := &SetAppDataResponse{}
	unmarshalErr := proto.Unmarshal(body, ret)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return ret, nil
}
