package v1

import "net/http"
import "google.golang.org/protobuf/proto"
import "io/ioutil"
import "fmt"
import "github.com/gorilla/websocket"
import "net/url"
import "bytes"

import "github.com/golang/protobuf/ptypes/empty"

type ChatServiceClient struct {
	client    *http.Client
	serverURL string
}

func NewChatServiceClient(url string) *ChatServiceClient {
	return &ChatServiceClient{
		client:    &http.Client{},
		serverURL: url,
	}
}

func (client *ChatServiceClient) CreateGuild(r *CreateGuildRequest) (*CreateGuildResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/CreateGuild", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &CreateGuildResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) CreateInvite(r *CreateInviteRequest) (*CreateInviteResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/CreateInvite", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &CreateInviteResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) CreateChannel(r *CreateChannelRequest) (*CreateChannelResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/CreateChannel", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &CreateChannelResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) CreateEmotePack(r *CreateEmotePackRequest) (*CreateEmotePackResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/CreateEmotePack", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &CreateEmotePackResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) GetGuildList(r *GetGuildListRequest) (*GetGuildListResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/GetGuildList", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &GetGuildListResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) AddGuildToGuildList(r *AddGuildToGuildListRequest) (*AddGuildToGuildListResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/AddGuildToGuildList", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &AddGuildToGuildListResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) RemoveGuildFromGuildList(r *RemoveGuildFromGuildListRequest) (*RemoveGuildFromGuildListResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/RemoveGuildFromGuildList", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &RemoveGuildFromGuildListResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) GetGuild(r *GetGuildRequest) (*GetGuildResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/GetGuild", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &GetGuildResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) GetGuildInvites(r *GetGuildInvitesRequest) (*GetGuildInvitesResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/GetGuildInvites", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &GetGuildInvitesResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) GetGuildMembers(r *GetGuildMembersRequest) (*GetGuildMembersResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/GetGuildMembers", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &GetGuildMembersResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) GetGuildChannels(r *GetGuildChannelsRequest) (*GetGuildChannelsResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/GetGuildChannels", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &GetGuildChannelsResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) GetChannelMessages(r *GetChannelMessagesRequest) (*GetChannelMessagesResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/GetChannelMessages", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &GetChannelMessagesResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) GetMessage(r *GetMessageRequest) (*GetMessageResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/GetMessage", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &GetMessageResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) GetEmotePacks(r *GetEmotePacksRequest) (*GetEmotePacksResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/GetEmotePacks", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &GetEmotePacksResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) GetEmotePackEmotes(r *GetEmotePackEmotesRequest) (*GetEmotePackEmotesResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/GetEmotePackEmotes", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &GetEmotePackEmotesResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) UpdateGuildInformation(r *UpdateGuildInformationRequest) (*empty.Empty, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/UpdateGuildInformation", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &empty.Empty{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) UpdateChannelInformation(r *UpdateChannelInformationRequest) (*empty.Empty, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/UpdateChannelInformation", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &empty.Empty{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) UpdateChannelOrder(r *UpdateChannelOrderRequest) (*empty.Empty, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/UpdateChannelOrder", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &empty.Empty{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) UpdateMessage(r *UpdateMessageRequest) (*empty.Empty, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/UpdateMessage", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &empty.Empty{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) AddEmoteToPack(r *AddEmoteToPackRequest) (*empty.Empty, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/AddEmoteToPack", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &empty.Empty{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) DeleteGuild(r *DeleteGuildRequest) (*empty.Empty, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/DeleteGuild", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &empty.Empty{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) DeleteInvite(r *DeleteInviteRequest) (*empty.Empty, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/DeleteInvite", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &empty.Empty{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) DeleteChannel(r *DeleteChannelRequest) (*empty.Empty, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/DeleteChannel", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &empty.Empty{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) DeleteMessage(r *DeleteMessageRequest) (*empty.Empty, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/DeleteMessage", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &empty.Empty{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) DeleteEmoteFromPack(r *DeleteEmoteFromPackRequest) (*empty.Empty, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/DeleteEmoteFromPack", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &empty.Empty{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) DeleteEmotePack(r *DeleteEmotePackRequest) (*empty.Empty, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/DeleteEmotePack", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &empty.Empty{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) DequipEmotePack(r *DequipEmotePackRequest) (*empty.Empty, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/DequipEmotePack", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &empty.Empty{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) JoinGuild(r *JoinGuildRequest) (*JoinGuildResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/JoinGuild", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &JoinGuildResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) LeaveGuild(r *LeaveGuildRequest) (*empty.Empty, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/LeaveGuild", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &empty.Empty{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) TriggerAction(r *TriggerActionRequest) (*empty.Empty, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/TriggerAction", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &empty.Empty{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) SendMessage(r *SendMessageRequest) (*SendMessageResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/SendMessage", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &SendMessageResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) QueryHasPermission(r *QueryPermissionsRequest) (*QueryPermissionsResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/QueryHasPermission", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &QueryPermissionsResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) SetPermissions(r *SetPermissionsRequest) (*empty.Empty, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/SetPermissions", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &empty.Empty{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) GetPermissions(r *GetPermissionsRequest) (*GetPermissionsResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/GetPermissions", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &GetPermissionsResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) MoveRole(r *MoveRoleRequest) (*MoveRoleResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/MoveRole", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &MoveRoleResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) GetGuildRoles(r *GetGuildRolesRequest) (*GetGuildRolesResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/GetGuildRoles", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &GetGuildRolesResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) AddGuildRole(r *AddGuildRoleRequest) (*AddGuildRoleResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/AddGuildRole", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &AddGuildRoleResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) ModifyGuildRole(r *ModifyGuildRoleRequest) (*empty.Empty, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/ModifyGuildRole", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &empty.Empty{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) DeleteGuildRole(r *DeleteGuildRoleRequest) (*empty.Empty, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/DeleteGuildRole", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &empty.Empty{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) ManageUserRoles(r *ManageUserRolesRequest) (*empty.Empty, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/ManageUserRoles", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &empty.Empty{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) GetUserRoles(r *GetUserRolesRequest) (*GetUserRolesResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/GetUserRoles", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &GetUserRolesResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) StreamEvents() (in chan<- *StreamEventsRequest, out <-chan *Event, err error) {
	u := url.URL{Scheme: "ws", Host: client.serverURL, Path: "/protocol.chat.v1.ChatService/StreamEvents"}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	inC := make(chan *StreamEventsRequest)
	outC := make(chan *Event)

	go func() {
		defer c.Close()

		msgs := make(chan []byte)

		go func() {
			for {
				_, message, err := c.ReadMessage()
				if err != nil {
					close(msgs)
					break
				}
				msgs <- message
			}
		}()

		for {
			select {
			case msg, ok := <-msgs:
				if !ok {
					return
				}

				thing := new(Event)
				err = proto.Unmarshal(msg, thing)
				if err != nil {
					return
				}

				outC <- thing
			case send, ok := <-inC:
				if !ok {
					return
				}

				data, err := proto.Marshal(send)
				if err != nil {
					return
				}

				err = c.WriteMessage(websocket.BinaryMessage, data)
				if err != nil {
					return
				}
			}
		}
	}()

	return inC, outC, nil
}

func (client *ChatServiceClient) Sync(r *SyncRequest) (chan *SyncEvent, error) {
	panic("unimplemented")
}

func (client *ChatServiceClient) GetUser(r *GetUserRequest) (*GetUserResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/GetUser", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &GetUserResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) GetUserMetadata(r *GetUserMetadataRequest) (*GetUserMetadataResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/GetUserMetadata", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &GetUserMetadataResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) ProfileUpdate(r *ProfileUpdateRequest) (*empty.Empty, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/ProfileUpdate", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &empty.Empty{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) Typing(r *TypingRequest) (*empty.Empty, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/Typing", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &empty.Empty{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) PreviewGuild(r *PreviewGuildRequest) (*PreviewGuildResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/PreviewGuild", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &PreviewGuildResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}
