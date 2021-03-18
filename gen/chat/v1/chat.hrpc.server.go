package v1

import (
	"io/ioutil"
	"net/http"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/gorilla/websocket"
	"github.com/harmony-development/hrpc/server"
	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func BindPB(obj interface{}, c echo.Context) error {
	buf, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}

	ct := c.Request().Header.Get("Content-Type")
	switch ct {
	case "application/hrpc", "application/octet-stream":
		if err = proto.Unmarshal(buf, obj.(proto.Message)); err != nil {
			return err
		}
	case "application/hrpc-json":
		if err = protojson.Unmarshal(buf, obj.(proto.Message)); err != nil {
			return err
		}
	}

	return nil
}

var Chatᐳv1ᐳchat *descriptorpb.FileDescriptorProto = new(descriptorpb.FileDescriptorProto)

func init() {
	data := []byte("\n\x12chat/v1/chat.proto\x12\x10protocol.chat.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1bharmonytypes/v1/types.proto\x1a\x15chat/v1/profile.proto\x1a\x14chat/v1/guilds.proto\x1a\x16chat/v1/channels.proto\x1a\x16chat/v1/messages.proto\x1a\x14chat/v1/emotes.proto\x1a\x19chat/v1/permissions.proto\x1a\x17chat/v1/streaming.proto\x1a\x15chat/v1/postbox.proto2\xb9,\n\vChatService\x12f\n\vCreateGuild\x12$.protocol.chat.v1.CreateGuildRequest\x1a%.protocol.chat.v1.CreateGuildResponse\"\n\x9aD\x02\b\x01\x9aD\x02\x1a\x00\x12~\n\fCreateInvite\x12%.protocol.chat.v1.CreateInviteRequest\x1a&.protocol.chat.v1.CreateInviteResponse\"\x1f\x9aD\x02\b\x01\x9aD\x17\x1a\x15invites.manage.create\x12\x82\x01\n\rCreateChannel\x12&.protocol.chat.v1.CreateChannelRequest\x1a'.protocol.chat.v1.CreateChannelResponse\" \x9aD\x02\b\x01\x9aD\x18\x1a\x16channels.manage.create\x12m\n\x0fCreateEmotePack\x12(.protocol.chat.v1.CreateEmotePackRequest\x1a).protocol.chat.v1.CreateEmotePackResponse\"\x05\x9aD\x02\b\x01\x12d\n\fGetGuildList\x12%.protocol.chat.v1.GetGuildListRequest\x1a&.protocol.chat.v1.GetGuildListResponse\"\x05\x9aD\x02\b\x01\x12~\n\x13AddGuildToGuildList\x12,.protocol.chat.v1.AddGuildToGuildListRequest\x1a-.protocol.chat.v1.AddGuildToGuildListResponse\"\n\x9aD\x02\b\x01\x9aD\x02\x10\x01\x12\x8d\x01\n\x18RemoveGuildFromGuildList\x121.protocol.chat.v1.RemoveGuildFromGuildListRequest\x1a2.protocol.chat.v1.RemoveGuildFromGuildListResponse\"\n\x9aD\x02\b\x01\x9aD\x02\x10\x01\x12X\n\bGetGuild\x12!.protocol.chat.v1.GetGuildRequest\x1a\".protocol.chat.v1.GetGuildResponse\"\x05\x9aD\x02\b\x01\x12~\n\x0fGetGuildInvites\x12(.protocol.chat.v1.GetGuildInvitesRequest\x1a).protocol.chat.v1.GetGuildInvitesResponse\"\x16\x9aD\x02\b\x01\x9aD\x0e\x1a\finvites.view\x12m\n\x0fGetGuildMembers\x12(.protocol.chat.v1.GetGuildMembersRequest\x1a).protocol.chat.v1.GetGuildMembersResponse\"\x05\x9aD\x02\b\x01\x12k\n\x10GetGuildChannels\x12).protocol.chat.v1.GetGuildChannelsRequest\x1a*.protocol.chat.v1.GetGuildChannelsResponse\"\x00\x12\x88\x01\n\x12GetChannelMessages\x12+.protocol.chat.v1.GetChannelMessagesRequest\x1a,.protocol.chat.v1.GetChannelMessagesResponse\"\x17\x9aD\x02\b\x01\x9aD\x0f\x1a\rmessages.view\x12p\n\nGetMessage\x12#.protocol.chat.v1.GetMessageRequest\x1a$.protocol.chat.v1.GetMessageResponse\"\x17\x9aD\x02\b\x01\x9aD\x0f\x1a\rmessages.view\x12g\n\rGetEmotePacks\x12&.protocol.chat.v1.GetEmotePacksRequest\x1a'.protocol.chat.v1.GetEmotePacksResponse\"\x05\x9aD\x02\b\x01\x12v\n\x12GetEmotePackEmotes\x12+.protocol.chat.v1.GetEmotePackEmotesRequest\x1a,.protocol.chat.v1.GetEmotePackEmotesResponse\"\x05\x9aD\x02\b\x01\x12\x8c\x01\n\x16UpdateGuildInformation\x12/.protocol.chat.v1.UpdateGuildInformationRequest\x1a\x16.google.protobuf.Empty\")\x9aD\x02\b\x01\x9aD!\x1a\x1fguild.manage.change-information\x12\x93\x01\n\x18UpdateChannelInformation\x121.protocol.chat.v1.UpdateChannelInformationRequest\x1a\x16.google.protobuf.Empty\",\x9aD\x02\b\x01\x9aD$\x1a\"channels.manage.change-information\x12y\n\x12UpdateChannelOrder\x12+.protocol.chat.v1.UpdateChannelOrderRequest\x1a\x16.google.protobuf.Empty\"\x1e\x9aD\x02\b\x01\x9aD\x16\x1a\x14channels.manage.move\x12l\n\rUpdateMessage\x12*.protocol.chat.v1.UpdateMessageTextRequest\x1a\x16.google.protobuf.Empty\"\x17\x9aD\x02\b\x01\x9aD\x0f\x1a\rmessages.send\x12X\n\x0eAddEmoteToPack\x12'.protocol.chat.v1.AddEmoteToPackRequest\x1a\x16.google.protobuf.Empty\"\x05\x9aD\x02\b\x01\x12j\n\vDeleteGuild\x12$.protocol.chat.v1.DeleteGuildRequest\x1a\x16.google.protobuf.Empty\"\x1d\x9aD\x02\b\x01\x9aD\x15\x1a\x13guild.manage.delete\x12n\n\fDeleteInvite\x12%.protocol.chat.v1.DeleteInviteRequest\x1a\x16.google.protobuf.Empty\"\x1f\x9aD\x02\b\x01\x9aD\x17\x1a\x15invites.manage.delete\x12q\n\rDeleteChannel\x12&.protocol.chat.v1.DeleteChannelRequest\x1a\x16.google.protobuf.Empty\" \x9aD\x02\b\x01\x9aD\x18\x1a\x16channels.manage.delete\x12V\n\rDeleteMessage\x12&.protocol.chat.v1.DeleteMessageRequest\x1a\x16.google.protobuf.Empty\"\x05\x9aD\x02\b\x01\x12b\n\x13DeleteEmoteFromPack\x12,.protocol.chat.v1.DeleteEmoteFromPackRequest\x1a\x16.google.protobuf.Empty\"\x05\x9aD\x02\b\x01\x12Z\n\x0fDeleteEmotePack\x12(.protocol.chat.v1.DeleteEmotePackRequest\x1a\x16.google.protobuf.Empty\"\x05\x9aD\x02\b\x01\x12Z\n\x0fDequipEmotePack\x12(.protocol.chat.v1.DequipEmotePackRequest\x1a\x16.google.protobuf.Empty\"\x05\x9aD\x02\b\x01\x12[\n\tJoinGuild\x12\".protocol.chat.v1.JoinGuildRequest\x1a#.protocol.chat.v1.JoinGuildResponse\"\x05\x9aD\x02\b\x01\x12P\n\nLeaveGuild\x12#.protocol.chat.v1.LeaveGuildRequest\x1a\x16.google.protobuf.Empty\"\x05\x9aD\x02\b\x01\x12j\n\rTriggerAction\x12&.protocol.chat.v1.TriggerActionRequest\x1a\x16.google.protobuf.Empty\"\x19\x9aD\x02\b\x01\x9aD\x11\x1a\x0factions.trigger\x12s\n\vSendMessage\x12$.protocol.chat.v1.SendMessageRequest\x1a%.protocol.chat.v1.SendMessageResponse\"\x17\x9aD\x02\b\x01\x9aD\x0f\x1a\rmessages.send\x12r\n\x12QueryHasPermission\x12).protocol.chat.v1.QueryPermissionsRequest\x1a*.protocol.chat.v1.QueryPermissionsResponse\"\x05\x9aD\x02\b\x01\x12s\n\x0eSetPermissions\x12'.protocol.chat.v1.SetPermissionsRequest\x1a\x16.google.protobuf.Empty\" \x9aD\x02\b\x01\x9aD\x18\x1a\x16permissions.manage.set\x12\x85\x01\n\x0eGetPermissions\x12'.protocol.chat.v1.GetPermissionsRequest\x1a(.protocol.chat.v1.GetPermissionsResponse\" \x9aD\x02\b\x01\x9aD\x18\x1a\x16permissions.manage.get\x12i\n\bMoveRole\x12!.protocol.chat.v1.MoveRoleRequest\x1a\".protocol.chat.v1.MoveRoleResponse\"\x16\x9aD\x02\b\x01\x9aD\x0e\x1a\froles.manage\x12u\n\rGetGuildRoles\x12&.protocol.chat.v1.GetGuildRolesRequest\x1a'.protocol.chat.v1.GetGuildRolesResponse\"\x13\x9aD\x02\b\x01\x9aD\v\x1a\troles.get\x12u\n\fAddGuildRole\x12%.protocol.chat.v1.AddGuildRoleRequest\x1a&.protocol.chat.v1.AddGuildRoleResponse\"\x16\x9aD\x02\b\x01\x9aD\x0e\x1a\froles.manage\x12k\n\x0fModifyGuildRole\x12(.protocol.chat.v1.ModifyGuildRoleRequest\x1a\x16.google.protobuf.Empty\"\x16\x9aD\x02\b\x01\x9aD\x0e\x1a\froles.manage\x12k\n\x0fDeleteGuildRole\x12(.protocol.chat.v1.DeleteGuildRoleRequest\x1a\x16.google.protobuf.Empty\"\x16\x9aD\x02\b\x01\x9aD\x0e\x1a\froles.manage\x12p\n\x0fManageUserRoles\x12(.protocol.chat.v1.ManageUserRolesRequest\x1a\x16.google.protobuf.Empty\"\x1b\x9aD\x02\b\x01\x9aD\x13\x1a\x11roles.user.manage\x12d\n\fGetUserRoles\x12%.protocol.chat.v1.GetUserRolesRequest\x1a&.protocol.chat.v1.GetUserRolesResponse\"\x05\x9aD\x02\b\x01\x12Y\n\fStreamEvents\x12%.protocol.chat.v1.StreamEventsRequest\x1a\x17.protocol.chat.v1.Event\"\x05\x9aD\x02\b\x01(\x010\x01\x12F\n\x04Sync\x12\x1d.protocol.chat.v1.SyncRequest\x1a\x1b.protocol.chat.v1.SyncEvent\"\x000\x01\x12U\n\aGetUser\x12 .protocol.chat.v1.GetUserRequest\x1a!.protocol.chat.v1.GetUserResponse\"\x05\x9aD\x02\b\x01\x12a\n\vGetUserBulk\x12$.protocol.chat.v1.GetUserBulkRequest\x1a%.protocol.chat.v1.GetUserBulkResponse\"\x05\x9aD\x02\b\x01\x12m\n\x0fGetUserMetadata\x12(.protocol.chat.v1.GetUserMetadataRequest\x1a).protocol.chat.v1.GetUserMetadataResponse\"\x05\x9aD\x02\b\x01\x12V\n\rProfileUpdate\x12&.protocol.chat.v1.ProfileUpdateRequest\x1a\x16.google.protobuf.Empty\"\x05\x9aD\x02\b\x01\x12Z\n\x06Typing\x12\x1f.protocol.chat.v1.TypingRequest\x1a\x16.google.protobuf.Empty\"\x17\x9aD\x02\b\x01\x9aD\x0f\x1a\rmessages.send\x12d\n\fPreviewGuild\x12%.protocol.chat.v1.PreviewGuildRequest\x1a&.protocol.chat.v1.PreviewGuildResponse\"\x05\x9aD\x02\b\x00\x12^\n\aBanUser\x12 .protocol.chat.v1.BanUserRequest\x1a\x16.google.protobuf.Empty\"\x19\x9aD\x02\b\x01\x9aD\x11\x1a\x0fuser.manage.ban\x12a\n\bKickUser\x12!.protocol.chat.v1.KickUserRequest\x1a\x16.google.protobuf.Empty\"\x1a\x9aD\x02\b\x01\x9aD\x12\x1a\x10user.manage.kick\x12d\n\tUnbanUser\x12\".protocol.chat.v1.UnbanUserRequest\x1a\x16.google.protobuf.Empty\"\x1b\x9aD\x02\b\x01\x9aD\x13\x1a\x11user.manage.unbanB3Z1github.com/harmony-development/legato/gen/chat/v1J\x993\n\a\x12\x05\x00\x00\x9a\x02\x01\n\b\n\x01\f\x12\x03\x00\x00\x12\n\t\n\x02\x03\x00\x12\x03\x02\x00%\n.\n\x02\x03\x01\x12\x03\x04\x00%\x1a# import \"validate/validate.proto\";\n\n\t\n\x02\x03\x02\x12\x03\x05\x00\x1f\n\t\n\x02\x03\x03\x12\x03\x06\x00\x1e\n\t\n\x02\x03\x04\x12\x03\a\x00 \n\t\n\x02\x03\x05\x12\x03\b\x00 \n\t\n\x02\x03\x06\x12\x03\t\x00\x1e\n\t\n\x02\x03\a\x12\x03\n\x00#\n\t\n\x02\x03\b\x12\x03\v\x00!\n\t\n\x02\x03\t\x12\x03\f\x00\x1f\n\b\n\x01\x02\x12\x03\x0e\x00\x19\n\b\n\x01\b\x12\x03\x10\x00H\n\t\n\x02\b\v\x12\x03\x10\x00H\n\v\n\x02\x06\x00\x12\x05\x12\x00\x9a\x02\x01\n\n\n\x03\x06\x00\x01\x12\x03\x12\b\x13\n\f\n\x04\x06\x00\x02\x00\x12\x04\x13\x02\x16\x03\n\f\n\x05\x06\x00\x02\x00\x01\x12\x03\x13\x06\x11\n\f\n\x05\x06\x00\x02\x00\x02\x12\x03\x13\x12$\n\f\n\x05\x06\x00\x02\x00\x03\x12\x03\x13/B\n\f\n\x05\x06\x00\x02\x00\x04\x12\x03\x14\x04E\n\x0f\n\b\x06\x00\x02\x00\x04\xc3\b\x01\x12\x03\x14\x04E\n\f\n\x05\x06\x00\x02\x00\x04\x12\x03\x15\x04D\n\x0f\n\b\x06\x00\x02\x00\x04\xc3\b\x03\x12\x03\x15\x04D\n\f\n\x04\x06\x00\x02\x01\x12\x04\x18\x02\x1c\x03\n\f\n\x05\x06\x00\x02\x01\x01\x12\x03\x18\x06\x12\n\f\n\x05\x06\x00\x02\x01\x02\x12\x03\x18\x13&\n\f\n\x05\x06\x00\x02\x01\x03\x12\x03\x181E\n\f\n\x05\x06\x00\x02\x01\x04\x12\x03\x19\x04E\n\x0f\n\b\x06\x00\x02\x01\x04\xc3\b\x01\x12\x03\x19\x04E\n\r\n\x05\x06\x00\x02\x01\x04\x12\x04\x1a\x04\x1b \n\x10\n\b\x06\x00\x02\x01\x04\xc3\b\x03\x12\x04\x1a\x04\x1b \n\f\n\x04\x06\x00\x02\x02\x12\x04\x1e\x02\"\x03\n\f\n\x05\x06\x00\x02\x02\x01\x12\x03\x1e\x06\x13\n\f\n\x05\x06\x00\x02\x02\x02\x12\x03\x1e\x14(\n\f\n\x05\x06\x00\x02\x02\x03\x12\x03\x1e3H\n\f\n\x05\x06\x00\x02\x02\x04\x12\x03\x1f\x04E\n\x0f\n\b\x06\x00\x02\x02\x04\xc3\b\x01\x12\x03\x1f\x04E\n\r\n\x05\x06\x00\x02\x02\x04\x12\x04 \x04!!\n\x10\n\b\x06\x00\x02\x02\x04\xc3\b\x03\x12\x04 \x04!!\n\f\n\x04\x06\x00\x02\x03\x12\x04#\x02&\x03\n\f\n\x05\x06\x00\x02\x03\x01\x12\x03#\x06\x15\n\f\n\x05\x06\x00\x02\x03\x02\x12\x03#\x16,\n\f\n\x05\x06\x00\x02\x03\x03\x12\x03$\x0f&\n\f\n\x05\x06\x00\x02\x03\x04\x12\x03%\x04E\n\x0f\n\b\x06\x00\x02\x03\x04\xc3\b\x01\x12\x03%\x04E\n\f\n\x04\x06\x00\x02\x04\x12\x04(\x02*\x03\n\f\n\x05\x06\x00\x02\x04\x01\x12\x03(\x06\x12\n\f\n\x05\x06\x00\x02\x04\x02\x12\x03(\x13&\n\f\n\x05\x06\x00\x02\x04\x03\x12\x03(1E\n\f\n\x05\x06\x00\x02\x04\x04\x12\x03)\x04E\n\x0f\n\b\x06\x00\x02\x04\x04\xc3\b\x01\x12\x03)\x04E\n\f\n\x04\x06\x00\x02\x05\x12\x04+\x02/\x03\n\f\n\x05\x06\x00\x02\x05\x01\x12\x03+\x06\x19\n\f\n\x05\x06\x00\x02\x05\x02\x12\x03+\x1a4\n\f\n\x05\x06\x00\x02\x05\x03\x12\x03,\x0f*\n\f\n\x05\x06\x00\x02\x05\x04\x12\x03-\x04E\n\x0f\n\b\x06\x00\x02\x05\x04\xc3\b\x01\x12\x03-\x04E\n\f\n\x05\x06\x00\x02\x05\x04\x12\x03.\x04<\n\x0f\n\b\x06\x00\x02\x05\x04\xc3\b\x02\x12\x03.\x04<\n\f\n\x04\x06\x00\x02\x06\x12\x040\x024\x03\n\f\n\x05\x06\x00\x02\x06\x01\x12\x030\x06\x1e\n\f\n\x05\x06\x00\x02\x06\x02\x12\x030\x1f>\n\f\n\x05\x06\x00\x02\x06\x03\x12\x031\x0f/\n\f\n\x05\x06\x00\x02\x06\x04\x12\x032\x04E\n\x0f\n\b\x06\x00\x02\x06\x04\xc3\b\x01\x12\x032\x04E\n\f\n\x05\x06\x00\x02\x06\x04\x12\x033\x04<\n\x0f\n\b\x06\x00\x02\x06\x04\xc3\b\x02\x12\x033\x04<\n\f\n\x04\x06\x00\x02\a\x12\x046\x028\x03\n\f\n\x05\x06\x00\x02\a\x01\x12\x036\x06\x0e\n\f\n\x05\x06\x00\x02\a\x02\x12\x036\x0f\x1e\n\f\n\x05\x06\x00\x02\a\x03\x12\x036)9\n\f\n\x05\x06\x00\x02\a\x04\x12\x037\x04E\n\x0f\n\b\x06\x00\x02\a\x04\xc3\b\x01\x12\x037\x04E\n<\n\x04\x06\x00\x02\b\x12\x04:\x02>\x03\x1a. This requires the \"invites.view\" permission.\n\n\f\n\x05\x06\x00\x02\b\x01\x12\x03:\x06\x15\n\f\n\x05\x06\x00\x02\b\x02\x12\x03:\x16,\n\f\n\x05\x06\x00\x02\b\x03\x12\x03;\x0f&\n\f\n\x05\x06\x00\x02\b\x04\x12\x03<\x04E\n\x0f\n\b\x06\x00\x02\b\x04\xc3\b\x01\x12\x03<\x04E\n\f\n\x05\x06\x00\x02\b\x04\x12\x03=\x04P\n\x0f\n\b\x06\x00\x02\b\x04\xc3\b\x03\x12\x03=\x04P\n\f\n\x04\x06\x00\x02\t\x12\x04?\x02B\x03\n\f\n\x05\x06\x00\x02\t\x01\x12\x03?\x06\x15\n\f\n\x05\x06\x00\x02\t\x02\x12\x03?\x16,\n\f\n\x05\x06\x00\x02\t\x03\x12\x03@\x0f&\n\f\n\x05\x06\x00\x02\t\x04\x12\x03A\x04E\n\x0f\n\b\x06\x00\x02\t\x04\xc3\b\x01\x12\x03A\x04E\nc\n\x04\x06\x00\x02\n\x12\x04F\x02G+\x1aU You will only be informed of channels you have the \"messages.view\"\n permission for.\n\n\f\n\x05\x06\x00\x02\n\x01\x12\x03F\x06\x16\n\f\n\x05\x06\x00\x02\n\x02\x12\x03F\x17.\n\f\n\x05\x06\x00\x02\n\x03\x12\x03G\x0f'\n\f\n\x04\x06\x00\x02\v\x12\x04I\x02N\x03\n\f\n\x05\x06\x00\x02\v\x01\x12\x03I\x06\x18\n\f\n\x05\x06\x00\x02\v\x02\x12\x03I\x192\n\f\n\x05\x06\x00\x02\v\x03\x12\x03J\x0f)\n\f\n\x05\x06\x00\x02\v\x04\x12\x03K\x04E\n\x0f\n\b\x06\x00\x02\v\x04\xc3\b\x01\x12\x03K\x04E\n\r\n\x05\x06\x00\x02\v\x04\x12\x04L\x04M\x18\n\x10\n\b\x06\x00\x02\v\x04\xc3\b\x03\x12\x04L\x04M\x18\n\f\n\x04\x06\x00\x02\f\x12\x04P\x02T\x03\n\f\n\x05\x06\x00\x02\f\x01\x12\x03P\x06\x10\n\f\n\x05\x06\x00\x02\f\x02\x12\x03P\x11\"\n\f\n\x05\x06\x00\x02\f\x03\x12\x03P-?\n\f\n\x05\x06\x00\x02\f\x04\x12\x03Q\x04E\n\x0f\n\b\x06\x00\x02\f\x04\xc3\b\x01\x12\x03Q\x04E\n\r\n\x05\x06\x00\x02\f\x04\x12\x04R\x04S\x18\n\x10\n\b\x06\x00\x02\f\x04\xc3\b\x03\x12\x04R\x04S\x18\n\f\n\x04\x06\x00\x02\r\x12\x04U\x02W\x03\n\f\n\x05\x06\x00\x02\r\x01\x12\x03U\x06\x13\n\f\n\x05\x06\x00\x02\r\x02\x12\x03U\x14(\n\f\n\x05\x06\x00\x02\r\x03\x12\x03U3H\n\f\n\x05\x06\x00\x02\r\x04\x12\x03V\x04E\n\x0f\n\b\x06\x00\x02\r\x04\xc3\b\x01\x12\x03V\x04E\n\f\n\x04\x06\x00\x02\x0e\x12\x04X\x02[\x03\n\f\n\x05\x06\x00\x02\x0e\x01\x12\x03X\x06\x18\n\f\n\x05\x06\x00\x02\x0e\x02\x12\x03X\x192\n\f\n\x05\x06\x00\x02\x0e\x03\x12\x03Y\x0f)\n\f\n\x05\x06\x00\x02\x0e\x04\x12\x03Z\x04E\n\x0f\n\b\x06\x00\x02\x0e\x04\xc3\b\x01\x12\x03Z\x04E\n\f\n\x04\x06\x00\x02\x0f\x12\x04]\x02b\x03\n\f\n\x05\x06\x00\x02\x0f\x01\x12\x03]\x06\x1c\n\f\n\x05\x06\x00\x02\x0f\x02\x12\x03]\x1d:\n\f\n\x05\x06\x00\x02\x0f\x03\x12\x03^\x0f$\n\f\n\x05\x06\x00\x02\x0f\x04\x12\x03_\x04E\n\x0f\n\b\x06\x00\x02\x0f\x04\xc3\b\x01\x12\x03_\x04E\n\r\n\x05\x06\x00\x02\x0f\x04\x12\x04`\x04a*\n\x10\n\b\x06\x00\x02\x0f\x04\xc3\b\x03\x12\x04`\x04a*\n\f\n\x04\x06\x00\x02\x10\x12\x04c\x02h\x03\n\f\n\x05\x06\x00\x02\x10\x01\x12\x03c\x06\x1e\n\f\n\x05\x06\x00\x02\x10\x02\x12\x03c\x1f>\n\f\n\x05\x06\x00\x02\x10\x03\x12\x03d\x0f$\n\f\n\x05\x06\x00\x02\x10\x04\x12\x03e\x04E\n\x0f\n\b\x06\x00\x02\x10\x04\xc3\b\x01\x12\x03e\x04E\n\r\n\x05\x06\x00\x02\x10\x04\x12\x04f\x04g-\n\x10\n\b\x06\x00\x02\x10\x04\xc3\b\x03\x12\x04f\x04g-\n\f\n\x04\x06\x00\x02\x11\x12\x04i\x02n\x03\n\f\n\x05\x06\x00\x02\x11\x01\x12\x03i\x06\x18\n\f\n\x05\x06\x00\x02\x11\x02\x12\x03i\x192\n\f\n\x05\x06\x00\x02\x11\x03\x12\x03j\x0f$\n\f\n\x05\x06\x00\x02\x11\x04\x12\x03k\x04E\n\x0f\n\b\x06\x00\x02\x11\x04\xc3\b\x01\x12\x03k\x04E\n\r\n\x05\x06\x00\x02\x11\x04\x12\x04l\x04m\x1f\n\x10\n\b\x06\x00\x02\x11\x04\xc3\b\x03\x12\x04l\x04m\x1f\n\f\n\x04\x06\x00\x02\x12\x12\x04o\x02s\x03\n\f\n\x05\x06\x00\x02\x12\x01\x12\x03o\x06\x13\n\f\n\x05\x06\x00\x02\x12\x02\x12\x03o\x14,\n\f\n\x05\x06\x00\x02\x12\x03\x12\x03o7L\n\f\n\x05\x06\x00\x02\x12\x04\x12\x03p\x04E\n\x0f\n\b\x06\x00\x02\x12\x04\xc3\b\x01\x12\x03p\x04E\n\r\n\x05\x06\x00\x02\x12\x04\x12\x04q\x04r\x18\n\x10\n\b\x06\x00\x02\x12\x04\xc3\b\x03\x12\x04q\x04r\x18\n\f\n\x04\x06\x00\x02\x13\x12\x04t\x02v\x03\n\f\n\x05\x06\x00\x02\x13\x01\x12\x03t\x06\x14\n\f\n\x05\x06\x00\x02\x13\x02\x12\x03t\x15*\n\f\n\x05\x06\x00\x02\x13\x03\x12\x03t5J\n\f\n\x05\x06\x00\x02\x13\x04\x12\x03u\x04E\n\x0f\n\b\x06\x00\x02\x13\x04\xc3\b\x01\x12\x03u\x04E\n\f\n\x04\x06\x00\x02\x14\x12\x04x\x02|\x03\n\f\n\x05\x06\x00\x02\x14\x01\x12\x03x\x06\x11\n\f\n\x05\x06\x00\x02\x14\x02\x12\x03x\x12$\n\f\n\x05\x06\x00\x02\x14\x03\x12\x03x/D\n\f\n\x05\x06\x00\x02\x14\x04\x12\x03y\x04E\n\x0f\n\b\x06\x00\x02\x14\x04\xc3\b\x01\x12\x03y\x04E\n\r\n\x05\x06\x00\x02\x14\x04\x12\x04z\x04{\x1e\n\x10\n\b\x06\x00\x02\x14\x04\xc3\b\x03\x12\x04z\x04{\x1e\n\r\n\x04\x06\x00\x02\x15\x12\x05}\x02\x81\x01\x03\n\f\n\x05\x06\x00\x02\x15\x01\x12\x03}\x06\x12\n\f\n\x05\x06\x00\x02\x15\x02\x12\x03}\x13&\n\f\n\x05\x06\x00\x02\x15\x03\x12\x03}1F\n\f\n\x05\x06\x00\x02\x15\x04\x12\x03~\x04E\n\x0f\n\b\x06\x00\x02\x15\x04\xc3\b\x01\x12\x03~\x04E\n\x0e\n\x05\x06\x00\x02\x15\x04\x12\x05\u007f\x04\x80\x01 \n\x11\n\b\x06\x00\x02\x15\x04\xc3\b\x03\x12\x05\u007f\x04\x80\x01 \n\x0e\n\x04\x06\x00\x02\x16\x12\x06\x82\x01\x02\x86\x01\x03\n\r\n\x05\x06\x00\x02\x16\x01\x12\x04\x82\x01\x06\x13\n\r\n\x05\x06\x00\x02\x16\x02\x12\x04\x82\x01\x14(\n\r\n\x05\x06\x00\x02\x16\x03\x12\x04\x82\x013H\n\r\n\x05\x06\x00\x02\x16\x04\x12\x04\x83\x01\x04E\n\x10\n\b\x06\x00\x02\x16\x04\xc3\b\x01\x12\x04\x83\x01\x04E\n\x0f\n\x05\x06\x00\x02\x16\x04\x12\x06\x84\x01\x04\x85\x01!\n\x12\n\b\x06\x00\x02\x16\x04\xc3\b\x03\x12\x06\x84\x01\x04\x85\x01!\nk\n\x04\x06\x00\x02\x17\x12\x06\x89\x01\x02\x8b\x01\x03\x1a[ This requires the \"messages.manage.delete\" permission if you are not the\n message author.\n\n\r\n\x05\x06\x00\x02\x17\x01\x12\x04\x89\x01\x06\x13\n\r\n\x05\x06\x00\x02\x17\x02\x12\x04\x89\x01\x14(\n\r\n\x05\x06\x00\x02\x17\x03\x12\x04\x89\x013H\n\r\n\x05\x06\x00\x02\x17\x04\x12\x04\x8a\x01\x04E\n\x10\n\b\x06\x00\x02\x17\x04\xc3\b\x01\x12\x04\x8a\x01\x04E\n\x0e\n\x04\x06\x00\x02\x18\x12\x06\x8c\x01\x02\x8f\x01\x03\n\r\n\x05\x06\x00\x02\x18\x01\x12\x04\x8c\x01\x06\x19\n\r\n\x05\x06\x00\x02\x18\x02\x12\x04\x8c\x01\x1a4\n\r\n\x05\x06\x00\x02\x18\x03\x12\x04\x8d\x01\x0f$\n\r\n\x05\x06\x00\x02\x18\x04\x12\x04\x8e\x01\x04E\n\x10\n\b\x06\x00\x02\x18\x04\xc3\b\x01\x12\x04\x8e\x01\x04E\n\x0e\n\x04\x06\x00\x02\x19\x12\x06\x90\x01\x02\x92\x01\x03\n\r\n\x05\x06\x00\x02\x19\x01\x12\x04\x90\x01\x06\x15\n\r\n\x05\x06\x00\x02\x19\x02\x12\x04\x90\x01\x16,\n\r\n\x05\x06\x00\x02\x19\x03\x12\x04\x90\x017L\n\r\n\x05\x06\x00\x02\x19\x04\x12\x04\x91\x01\x04E\n\x10\n\b\x06\x00\x02\x19\x04\xc3\b\x01\x12\x04\x91\x01\x04E\n\x0e\n\x04\x06\x00\x02\x1a\x12\x06\x93\x01\x02\x95\x01\x03\n\r\n\x05\x06\x00\x02\x1a\x01\x12\x04\x93\x01\x06\x15\n\r\n\x05\x06\x00\x02\x1a\x02\x12\x04\x93\x01\x16,\n\r\n\x05\x06\x00\x02\x1a\x03\x12\x04\x93\x017L\n\r\n\x05\x06\x00\x02\x1a\x04\x12\x04\x94\x01\x04E\n\x10\n\b\x06\x00\x02\x1a\x04\xc3\b\x01\x12\x04\x94\x01\x04E\n\x0e\n\x04\x06\x00\x02\x1b\x12\x06\x97\x01\x02\x99\x01\x03\n\r\n\x05\x06\x00\x02\x1b\x01\x12\x04\x97\x01\x06\x0f\n\r\n\x05\x06\x00\x02\x1b\x02\x12\x04\x97\x01\x10 \n\r\n\x05\x06\x00\x02\x1b\x03\x12\x04\x97\x01+<\n\r\n\x05\x06\x00\x02\x1b\x04\x12\x04\x98\x01\x04E\n\x10\n\b\x06\x00\x02\x1b\x04\xc3\b\x01\x12\x04\x98\x01\x04E\n\x0e\n\x04\x06\x00\x02\x1c\x12\x06\x9a\x01\x02\x9c\x01\x03\n\r\n\x05\x06\x00\x02\x1c\x01\x12\x04\x9a\x01\x06\x10\n\r\n\x05\x06\x00\x02\x1c\x02\x12\x04\x9a\x01\x11\"\n\r\n\x05\x06\x00\x02\x1c\x03\x12\x04\x9a\x01-B\n\r\n\x05\x06\x00\x02\x1c\x04\x12\x04\x9b\x01\x04E\n\x10\n\b\x06\x00\x02\x1c\x04\xc3\b\x01\x12\x04\x9b\x01\x04E\n\x0e\n\x04\x06\x00\x02\x1d\x12\x06\x9e\x01\x02\xa2\x01\x03\n\r\n\x05\x06\x00\x02\x1d\x01\x12\x04\x9e\x01\x06\x13\n\r\n\x05\x06\x00\x02\x1d\x02\x12\x04\x9e\x01\x14(\n\r\n\x05\x06\x00\x02\x1d\x03\x12\x04\x9e\x013H\n\r\n\x05\x06\x00\x02\x1d\x04\x12\x04\x9f\x01\x04E\n\x10\n\b\x06\x00\x02\x1d\x04\xc3\b\x01\x12\x04\x9f\x01\x04E\n\x0f\n\x05\x06\x00\x02\x1d\x04\x12\x06\xa0\x01\x04\xa1\x01\x1a\n\x12\n\b\x06\x00\x02\x1d\x04\xc3\b\x03\x12\x06\xa0\x01\x04\xa1\x01\x1a\n\x0e\n\x04\x06\x00\x02\x1e\x12\x06\xa4\x01\x02\xa8\x01\x03\n\r\n\x05\x06\x00\x02\x1e\x01\x12\x04\xa4\x01\x06\x11\n\r\n\x05\x06\x00\x02\x1e\x02\x12\x04\xa4\x01\x12$\n\r\n\x05\x06\x00\x02\x1e\x03\x12\x04\xa4\x01/B\n\r\n\x05\x06\x00\x02\x1e\x04\x12\x04\xa5\x01\x04E\n\x10\n\b\x06\x00\x02\x1e\x04\xc3\b\x01\x12\x04\xa5\x01\x04E\n\x0f\n\x05\x06\x00\x02\x1e\x04\x12\x06\xa6\x01\x04\xa7\x01\x18\n\x12\n\b\x06\x00\x02\x1e\x04\xc3\b\x03\x12\x06\xa6\x01\x04\xa7\x01\x18\n\x0e\n\x04\x06\x00\x02\x1f\x12\x06\xaa\x01\x02\xb2\x01\x03\n\r\n\x05\x06\x00\x02\x1f\x01\x12\x04\xaa\x01\x06\x18\n\r\n\x05\x06\x00\x02\x1f\x02\x12\x04\xaa\x01\x190\n\r\n\x05\x06\x00\x02\x1f\x03\x12\x04\xab\x01\x0f'\n\r\n\x05\x06\x00\x02\x1f\x04\x12\x04\xac\x01\x04E\n\xdb\x01\n\b\x06\x00\x02\x1f\x04\xc3\b\x01\x12\x04\xac\x01\x04E\"\xc8\x01 This permissions node is only required if you're trying\n to see if someone besides yourself has a permission.\n\n option (harmonytypes.v1.metadata).requires_permission_node =\n     \"permissions.query\";\n\n\x0e\n\x04\x06\x00\x02 \x12\x06\xb4\x01\x02\xb8\x01\x03\n\r\n\x05\x06\x00\x02 \x01\x12\x04\xb4\x01\x06\x14\n\r\n\x05\x06\x00\x02 \x02\x12\x04\xb4\x01\x15*\n\r\n\x05\x06\x00\x02 \x03\x12\x04\xb4\x015J\n\r\n\x05\x06\x00\x02 \x04\x12\x04\xb5\x01\x04E\n\x10\n\b\x06\x00\x02 \x04\xc3\b\x01\x12\x04\xb5\x01\x04E\n\x0f\n\x05\x06\x00\x02 \x04\x12\x06\xb6\x01\x04\xb7\x01!\n\x12\n\b\x06\x00\x02 \x04\xc3\b\x03\x12\x06\xb6\x01\x04\xb7\x01!\n\x0e\n\x04\x06\x00\x02!\x12\x06\xba\x01\x02\xbe\x01\x03\n\r\n\x05\x06\x00\x02!\x01\x12\x04\xba\x01\x06\x14\n\r\n\x05\x06\x00\x02!\x02\x12\x04\xba\x01\x15*\n\r\n\x05\x06\x00\x02!\x03\x12\x04\xba\x015K\n\r\n\x05\x06\x00\x02!\x04\x12\x04\xbb\x01\x04E\n\x10\n\b\x06\x00\x02!\x04\xc3\b\x01\x12\x04\xbb\x01\x04E\n\x0f\n\x05\x06\x00\x02!\x04\x12\x06\xbc\x01\x04\xbd\x01!\n\x12\n\b\x06\x00\x02!\x04\xc3\b\x03\x12\x06\xbc\x01\x04\xbd\x01!\n\x0e\n\x04\x06\x00\x02\"\x12\x06\xc0\x01\x02\xc3\x01\x03\n\r\n\x05\x06\x00\x02\"\x01\x12\x04\xc0\x01\x06\x0e\n\r\n\x05\x06\x00\x02\"\x02\x12\x04\xc0\x01\x0f\x1e\n\r\n\x05\x06\x00\x02\"\x03\x12\x04\xc0\x01)9\n\r\n\x05\x06\x00\x02\"\x04\x12\x04\xc1\x01\x04E\n\x10\n\b\x06\x00\x02\"\x04\xc3\b\x01\x12\x04\xc1\x01\x04E\n\r\n\x05\x06\x00\x02\"\x04\x12\x04\xc2\x01\x04P\n\x10\n\b\x06\x00\x02\"\x04\xc3\b\x03\x12\x04\xc2\x01\x04P\n\x0e\n\x04\x06\x00\x02#\x12\x06\xc5\x01\x02\xc8\x01\x03\n\r\n\x05\x06\x00\x02#\x01\x12\x04\xc5\x01\x06\x13\n\r\n\x05\x06\x00\x02#\x02\x12\x04\xc5\x01\x14(\n\r\n\x05\x06\x00\x02#\x03\x12\x04\xc5\x013H\n\r\n\x05\x06\x00\x02#\x04\x12\x04\xc6\x01\x04E\n\x10\n\b\x06\x00\x02#\x04\xc3\b\x01\x12\x04\xc6\x01\x04E\n\r\n\x05\x06\x00\x02#\x04\x12\x04\xc7\x01\x04M\n\x10\n\b\x06\x00\x02#\x04\xc3\b\x03\x12\x04\xc7\x01\x04M\n\x0e\n\x04\x06\x00\x02$\x12\x06\xca\x01\x02\xcd\x01\x03\n\r\n\x05\x06\x00\x02$\x01\x12\x04\xca\x01\x06\x12\n\r\n\x05\x06\x00\x02$\x02\x12\x04\xca\x01\x13&\n\r\n\x05\x06\x00\x02$\x03\x12\x04\xca\x011E\n\r\n\x05\x06\x00\x02$\x04\x12\x04\xcb\x01\x04E\n\x10\n\b\x06\x00\x02$\x04\xc3\b\x01\x12\x04\xcb\x01\x04E\n\r\n\x05\x06\x00\x02$\x04\x12\x04\xcc\x01\x04P\n\x10\n\b\x06\x00\x02$\x04\xc3\b\x03\x12\x04\xcc\x01\x04P\n\x0e\n\x04\x06\x00\x02%\x12\x06\xcf\x01\x02\xd2\x01\x03\n\r\n\x05\x06\x00\x02%\x01\x12\x04\xcf\x01\x06\x15\n\r\n\x05\x06\x00\x02%\x02\x12\x04\xcf\x01\x16,\n\r\n\x05\x06\x00\x02%\x03\x12\x04\xcf\x017L\n\r\n\x05\x06\x00\x02%\x04\x12\x04\xd0\x01\x04E\n\x10\n\b\x06\x00\x02%\x04\xc3\b\x01\x12\x04\xd0\x01\x04E\n\r\n\x05\x06\x00\x02%\x04\x12\x04\xd1\x01\x04P\n\x10\n\b\x06\x00\x02%\x04\xc3\b\x03\x12\x04\xd1\x01\x04P\n\x0e\n\x04\x06\x00\x02&\x12\x06\xd4\x01\x02\xd7\x01\x03\n\r\n\x05\x06\x00\x02&\x01\x12\x04\xd4\x01\x06\x15\n\r\n\x05\x06\x00\x02&\x02\x12\x04\xd4\x01\x16,\n\r\n\x05\x06\x00\x02&\x03\x12\x04\xd4\x017L\n\r\n\x05\x06\x00\x02&\x04\x12\x04\xd5\x01\x04E\n\x10\n\b\x06\x00\x02&\x04\xc3\b\x01\x12\x04\xd5\x01\x04E\n\r\n\x05\x06\x00\x02&\x04\x12\x04\xd6\x01\x04P\n\x10\n\b\x06\x00\x02&\x04\xc3\b\x03\x12\x04\xd6\x01\x04P\n\x0e\n\x04\x06\x00\x02'\x12\x06\xd9\x01\x02\xdd\x01\x03\n\r\n\x05\x06\x00\x02'\x01\x12\x04\xd9\x01\x06\x15\n\r\n\x05\x06\x00\x02'\x02\x12\x04\xd9\x01\x16,\n\r\n\x05\x06\x00\x02'\x03\x12\x04\xd9\x017L\n\r\n\x05\x06\x00\x02'\x04\x12\x04\xda\x01\x04E\n\x10\n\b\x06\x00\x02'\x04\xc3\b\x01\x12\x04\xda\x01\x04E\n\x0f\n\x05\x06\x00\x02'\x04\x12\x06\xdb\x01\x04\xdc\x01\x1c\n\x12\n\b\x06\x00\x02'\x04\xc3\b\x03\x12\x06\xdb\x01\x04\xdc\x01\x1c\n\x0e\n\x04\x06\x00\x02(\x12\x06\xdf\x01\x02\xe6\x01\x03\n\r\n\x05\x06\x00\x02(\x01\x12\x04\xdf\x01\x06\x12\n\r\n\x05\x06\x00\x02(\x02\x12\x04\xdf\x01\x13&\n\r\n\x05\x06\x00\x02(\x03\x12\x04\xdf\x011E\n\r\n\x05\x06\x00\x02(\x04\x12\x04\xe0\x01\x04E\n\xd1\x01\n\b\x06\x00\x02(\x04\xc3\b\x01\x12\x04\xe0\x01\x04E\"\xbe\x01 This permissions node is only required if you are trying to get the roles\n of someone other than yourself.\n\n option (harmonytypes.v1.metadata).requires_permission_node =\n \"roles.user.get\";\n\n\x0e\n\x04\x06\x00\x02)\x12\x06\xe8\x01\x02\xea\x01\x03\n\r\n\x05\x06\x00\x02)\x01\x12\x04\xe8\x01\x06\x12\n\r\n\x05\x06\x00\x02)\x05\x12\x04\xe8\x01\x13\x19\n\r\n\x05\x06\x00\x02)\x02\x12\x04\xe8\x01\x1a-\n\r\n\x05\x06\x00\x02)\x06\x12\x04\xe8\x018>\n\r\n\x05\x06\x00\x02)\x03\x12\x04\xe8\x01?D\n\r\n\x05\x06\x00\x02)\x04\x12\x04\xe9\x01\x04E\n\x10\n\b\x06\x00\x02)\x04\xc3\b\x01\x12\x04\xe9\x01\x04E\n\f\n\x04\x06\x00\x02*\x12\x04\xec\x01\x025\n\r\n\x05\x06\x00\x02*\x01\x12\x04\xec\x01\x06\n\n\r\n\x05\x06\x00\x02*\x02\x12\x04\xec\x01\v\x16\n\r\n\x05\x06\x00\x02*\x06\x12\x04\xec\x01!'\n\r\n\x05\x06\x00\x02*\x03\x12\x04\xec\x01(1\n\x0e\n\x04\x06\x00\x02+\x12\x06\xee\x01\x02\xf0\x01\x03\n\r\n\x05\x06\x00\x02+\x01\x12\x04\xee\x01\x06\r\n\r\n\x05\x06\x00\x02+\x02\x12\x04\xee\x01\x0e\x1c\n\r\n\x05\x06\x00\x02+\x03\x12\x04\xee\x01'6\n\r\n\x05\x06\x00\x02+\x04\x12\x04\xef\x01\x04E\n\x10\n\b\x06\x00\x02+\x04\xc3\b\x01\x12\x04\xef\x01\x04E\n\x0e\n\x04\x06\x00\x02,\x12\x06\xf2\x01\x02\xf4\x01\x03\n\r\n\x05\x06\x00\x02,\x01\x12\x04\xf2\x01\x06\x11\n\r\n\x05\x06\x00\x02,\x02\x12\x04\xf2\x01\x12$\n\r\n\x05\x06\x00\x02,\x03\x12\x04\xf2\x01/B\n\r\n\x05\x06\x00\x02,\x04\x12\x04\xf3\x01\x04E\n\x10\n\b\x06\x00\x02,\x04\xc3\b\x01\x12\x04\xf3\x01\x04E\n\x0e\n\x04\x06\x00\x02-\x12\x06\xf6\x01\x02\xf9\x01\x03\n\r\n\x05\x06\x00\x02-\x01\x12\x04\xf6\x01\x06\x15\n\r\n\x05\x06\x00\x02-\x02\x12\x04\xf6\x01\x16,\n\r\n\x05\x06\x00\x02-\x03\x12\x04\xf7\x01\x0f&\n\r\n\x05\x06\x00\x02-\x04\x12\x04\xf8\x01\x04E\n\x10\n\b\x06\x00\x02-\x04\xc3\b\x01\x12\x04\xf8\x01\x04E\n\x0e\n\x04\x06\x00\x02.\x12\x06\xfb\x01\x02\xfd\x01\x03\n\r\n\x05\x06\x00\x02.\x01\x12\x04\xfb\x01\x06\x13\n\r\n\x05\x06\x00\x02.\x02\x12\x04\xfb\x01\x14(\n\r\n\x05\x06\x00\x02.\x03\x12\x04\xfb\x013H\n\r\n\x05\x06\x00\x02.\x04\x12\x04\xfc\x01\x04E\n\x10\n\b\x06\x00\x02.\x04\xc3\b\x01\x12\x04\xfc\x01\x04E\n\x0e\n\x04\x06\x00\x02/\x12\x06\xff\x01\x02\x83\x02\x03\n\r\n\x05\x06\x00\x02/\x01\x12\x04\xff\x01\x06\f\n\r\n\x05\x06\x00\x02/\x02\x12\x04\xff\x01\r\x1a\n\r\n\x05\x06\x00\x02/\x03\x12\x04\xff\x01%:\n\r\n\x05\x06\x00\x02/\x04\x12\x04\x80\x02\x04E\n\x10\n\b\x06\x00\x02/\x04\xc3\b\x01\x12\x04\x80\x02\x04E\n\x0f\n\x05\x06\x00\x02/\x04\x12\x06\x81\x02\x04\x82\x02\x18\n\x12\n\b\x06\x00\x02/\x04\xc3\b\x03\x12\x06\x81\x02\x04\x82\x02\x18\n\x0e\n\x04\x06\x00\x020\x12\x06\x85\x02\x02\x87\x02\x03\n\r\n\x05\x06\x00\x020\x01\x12\x04\x85\x02\x06\x12\n\r\n\x05\x06\x00\x020\x02\x12\x04\x85\x02\x13&\n\r\n\x05\x06\x00\x020\x03\x12\x04\x85\x021E\n\r\n\x05\x06\x00\x020\x04\x12\x04\x86\x02\x04F\n\x10\n\b\x06\x00\x020\x04\xc3\b\x01\x12\x04\x86\x02\x04F\n\x0e\n\x04\x06\x00\x021\x12\x06\x89\x02\x02\x8d\x02\x03\n\r\n\x05\x06\x00\x021\x01\x12\x04\x89\x02\x06\r\n\r\n\x05\x06\x00\x021\x02\x12\x04\x89\x02\x0e\x1c\n\r\n\x05\x06\x00\x021\x03\x12\x04\x89\x02'<\n\r\n\x05\x06\x00\x021\x04\x12\x04\x8a\x02\x04E\n\x10\n\b\x06\x00\x021\x04\xc3\b\x01\x12\x04\x8a\x02\x04E\n\x0f\n\x05\x06\x00\x021\x04\x12\x06\x8b\x02\x04\x8c\x02\x1a\n\x12\n\b\x06\x00\x021\x04\xc3\b\x03\x12\x06\x8b\x02\x04\x8c\x02\x1a\n\x0e\n\x04\x06\x00\x022\x12\x06\x8f\x02\x02\x93\x02\x03\n\r\n\x05\x06\x00\x022\x01\x12\x04\x8f\x02\x06\x0e\n\r\n\x05\x06\x00\x022\x02\x12\x04\x8f\x02\x0f\x1e\n\r\n\x05\x06\x00\x022\x03\x12\x04\x8f\x02)>\n\r\n\x05\x06\x00\x022\x04\x12\x04\x90\x02\x04E\n\x10\n\b\x06\x00\x022\x04\xc3\b\x01\x12\x04\x90\x02\x04E\n\x0f\n\x05\x06\x00\x022\x04\x12\x06\x91\x02\x04\x92\x02\x1b\n\x12\n\b\x06\x00\x022\x04\xc3\b\x03\x12\x06\x91\x02\x04\x92\x02\x1b\n\x0e\n\x04\x06\x00\x023\x12\x06\x95\x02\x02\x99\x02\x03\n\r\n\x05\x06\x00\x023\x01\x12\x04\x95\x02\x06\x0f\n\r\n\x05\x06\x00\x023\x02\x12\x04\x95\x02\x10 \n\r\n\x05\x06\x00\x023\x03\x12\x04\x95\x02+@\n\r\n\x05\x06\x00\x023\x04\x12\x04\x96\x02\x04E\n\x10\n\b\x06\x00\x023\x04\xc3\b\x01\x12\x04\x96\x02\x04E\n\x0f\n\x05\x06\x00\x023\x04\x12\x06\x97\x02\x04\x98\x02\x1c\n\x12\n\b\x06\x00\x023\x04\xc3\b\x03\x12\x06\x97\x02\x04\x98\x02\x1cb\x06proto3")

	err := proto.Unmarshal(data, Chatᐳv1ᐳchat)
	if err != nil {
		panic(err)
	}
}

var ChatServiceData *descriptorpb.ServiceDescriptorProto = new(descriptorpb.ServiceDescriptorProto)

func init() {
	data := []byte("\n\vChatService\x12f\n\vCreateGuild\x12$.protocol.chat.v1.CreateGuildRequest\x1a%.protocol.chat.v1.CreateGuildResponse\"\n\x9aD\x02\b\x01\x9aD\x02\x1a\x00\x12~\n\fCreateInvite\x12%.protocol.chat.v1.CreateInviteRequest\x1a&.protocol.chat.v1.CreateInviteResponse\"\x1f\x9aD\x02\b\x01\x9aD\x17\x1a\x15invites.manage.create\x12\x82\x01\n\rCreateChannel\x12&.protocol.chat.v1.CreateChannelRequest\x1a'.protocol.chat.v1.CreateChannelResponse\" \x9aD\x02\b\x01\x9aD\x18\x1a\x16channels.manage.create\x12m\n\x0fCreateEmotePack\x12(.protocol.chat.v1.CreateEmotePackRequest\x1a).protocol.chat.v1.CreateEmotePackResponse\"\x05\x9aD\x02\b\x01\x12d\n\fGetGuildList\x12%.protocol.chat.v1.GetGuildListRequest\x1a&.protocol.chat.v1.GetGuildListResponse\"\x05\x9aD\x02\b\x01\x12~\n\x13AddGuildToGuildList\x12,.protocol.chat.v1.AddGuildToGuildListRequest\x1a-.protocol.chat.v1.AddGuildToGuildListResponse\"\n\x9aD\x02\b\x01\x9aD\x02\x10\x01\x12\x8d\x01\n\x18RemoveGuildFromGuildList\x121.protocol.chat.v1.RemoveGuildFromGuildListRequest\x1a2.protocol.chat.v1.RemoveGuildFromGuildListResponse\"\n\x9aD\x02\b\x01\x9aD\x02\x10\x01\x12X\n\bGetGuild\x12!.protocol.chat.v1.GetGuildRequest\x1a\".protocol.chat.v1.GetGuildResponse\"\x05\x9aD\x02\b\x01\x12~\n\x0fGetGuildInvites\x12(.protocol.chat.v1.GetGuildInvitesRequest\x1a).protocol.chat.v1.GetGuildInvitesResponse\"\x16\x9aD\x02\b\x01\x9aD\x0e\x1a\finvites.view\x12m\n\x0fGetGuildMembers\x12(.protocol.chat.v1.GetGuildMembersRequest\x1a).protocol.chat.v1.GetGuildMembersResponse\"\x05\x9aD\x02\b\x01\x12k\n\x10GetGuildChannels\x12).protocol.chat.v1.GetGuildChannelsRequest\x1a*.protocol.chat.v1.GetGuildChannelsResponse\"\x00\x12\x88\x01\n\x12GetChannelMessages\x12+.protocol.chat.v1.GetChannelMessagesRequest\x1a,.protocol.chat.v1.GetChannelMessagesResponse\"\x17\x9aD\x02\b\x01\x9aD\x0f\x1a\rmessages.view\x12p\n\nGetMessage\x12#.protocol.chat.v1.GetMessageRequest\x1a$.protocol.chat.v1.GetMessageResponse\"\x17\x9aD\x02\b\x01\x9aD\x0f\x1a\rmessages.view\x12g\n\rGetEmotePacks\x12&.protocol.chat.v1.GetEmotePacksRequest\x1a'.protocol.chat.v1.GetEmotePacksResponse\"\x05\x9aD\x02\b\x01\x12v\n\x12GetEmotePackEmotes\x12+.protocol.chat.v1.GetEmotePackEmotesRequest\x1a,.protocol.chat.v1.GetEmotePackEmotesResponse\"\x05\x9aD\x02\b\x01\x12\x8c\x01\n\x16UpdateGuildInformation\x12/.protocol.chat.v1.UpdateGuildInformationRequest\x1a\x16.google.protobuf.Empty\")\x9aD\x02\b\x01\x9aD!\x1a\x1fguild.manage.change-information\x12\x93\x01\n\x18UpdateChannelInformation\x121.protocol.chat.v1.UpdateChannelInformationRequest\x1a\x16.google.protobuf.Empty\",\x9aD\x02\b\x01\x9aD$\x1a\"channels.manage.change-information\x12y\n\x12UpdateChannelOrder\x12+.protocol.chat.v1.UpdateChannelOrderRequest\x1a\x16.google.protobuf.Empty\"\x1e\x9aD\x02\b\x01\x9aD\x16\x1a\x14channels.manage.move\x12l\n\rUpdateMessage\x12*.protocol.chat.v1.UpdateMessageTextRequest\x1a\x16.google.protobuf.Empty\"\x17\x9aD\x02\b\x01\x9aD\x0f\x1a\rmessages.send\x12X\n\x0eAddEmoteToPack\x12'.protocol.chat.v1.AddEmoteToPackRequest\x1a\x16.google.protobuf.Empty\"\x05\x9aD\x02\b\x01\x12j\n\vDeleteGuild\x12$.protocol.chat.v1.DeleteGuildRequest\x1a\x16.google.protobuf.Empty\"\x1d\x9aD\x02\b\x01\x9aD\x15\x1a\x13guild.manage.delete\x12n\n\fDeleteInvite\x12%.protocol.chat.v1.DeleteInviteRequest\x1a\x16.google.protobuf.Empty\"\x1f\x9aD\x02\b\x01\x9aD\x17\x1a\x15invites.manage.delete\x12q\n\rDeleteChannel\x12&.protocol.chat.v1.DeleteChannelRequest\x1a\x16.google.protobuf.Empty\" \x9aD\x02\b\x01\x9aD\x18\x1a\x16channels.manage.delete\x12V\n\rDeleteMessage\x12&.protocol.chat.v1.DeleteMessageRequest\x1a\x16.google.protobuf.Empty\"\x05\x9aD\x02\b\x01\x12b\n\x13DeleteEmoteFromPack\x12,.protocol.chat.v1.DeleteEmoteFromPackRequest\x1a\x16.google.protobuf.Empty\"\x05\x9aD\x02\b\x01\x12Z\n\x0fDeleteEmotePack\x12(.protocol.chat.v1.DeleteEmotePackRequest\x1a\x16.google.protobuf.Empty\"\x05\x9aD\x02\b\x01\x12Z\n\x0fDequipEmotePack\x12(.protocol.chat.v1.DequipEmotePackRequest\x1a\x16.google.protobuf.Empty\"\x05\x9aD\x02\b\x01\x12[\n\tJoinGuild\x12\".protocol.chat.v1.JoinGuildRequest\x1a#.protocol.chat.v1.JoinGuildResponse\"\x05\x9aD\x02\b\x01\x12P\n\nLeaveGuild\x12#.protocol.chat.v1.LeaveGuildRequest\x1a\x16.google.protobuf.Empty\"\x05\x9aD\x02\b\x01\x12j\n\rTriggerAction\x12&.protocol.chat.v1.TriggerActionRequest\x1a\x16.google.protobuf.Empty\"\x19\x9aD\x02\b\x01\x9aD\x11\x1a\x0factions.trigger\x12s\n\vSendMessage\x12$.protocol.chat.v1.SendMessageRequest\x1a%.protocol.chat.v1.SendMessageResponse\"\x17\x9aD\x02\b\x01\x9aD\x0f\x1a\rmessages.send\x12r\n\x12QueryHasPermission\x12).protocol.chat.v1.QueryPermissionsRequest\x1a*.protocol.chat.v1.QueryPermissionsResponse\"\x05\x9aD\x02\b\x01\x12s\n\x0eSetPermissions\x12'.protocol.chat.v1.SetPermissionsRequest\x1a\x16.google.protobuf.Empty\" \x9aD\x02\b\x01\x9aD\x18\x1a\x16permissions.manage.set\x12\x85\x01\n\x0eGetPermissions\x12'.protocol.chat.v1.GetPermissionsRequest\x1a(.protocol.chat.v1.GetPermissionsResponse\" \x9aD\x02\b\x01\x9aD\x18\x1a\x16permissions.manage.get\x12i\n\bMoveRole\x12!.protocol.chat.v1.MoveRoleRequest\x1a\".protocol.chat.v1.MoveRoleResponse\"\x16\x9aD\x02\b\x01\x9aD\x0e\x1a\froles.manage\x12u\n\rGetGuildRoles\x12&.protocol.chat.v1.GetGuildRolesRequest\x1a'.protocol.chat.v1.GetGuildRolesResponse\"\x13\x9aD\x02\b\x01\x9aD\v\x1a\troles.get\x12u\n\fAddGuildRole\x12%.protocol.chat.v1.AddGuildRoleRequest\x1a&.protocol.chat.v1.AddGuildRoleResponse\"\x16\x9aD\x02\b\x01\x9aD\x0e\x1a\froles.manage\x12k\n\x0fModifyGuildRole\x12(.protocol.chat.v1.ModifyGuildRoleRequest\x1a\x16.google.protobuf.Empty\"\x16\x9aD\x02\b\x01\x9aD\x0e\x1a\froles.manage\x12k\n\x0fDeleteGuildRole\x12(.protocol.chat.v1.DeleteGuildRoleRequest\x1a\x16.google.protobuf.Empty\"\x16\x9aD\x02\b\x01\x9aD\x0e\x1a\froles.manage\x12p\n\x0fManageUserRoles\x12(.protocol.chat.v1.ManageUserRolesRequest\x1a\x16.google.protobuf.Empty\"\x1b\x9aD\x02\b\x01\x9aD\x13\x1a\x11roles.user.manage\x12d\n\fGetUserRoles\x12%.protocol.chat.v1.GetUserRolesRequest\x1a&.protocol.chat.v1.GetUserRolesResponse\"\x05\x9aD\x02\b\x01\x12Y\n\fStreamEvents\x12%.protocol.chat.v1.StreamEventsRequest\x1a\x17.protocol.chat.v1.Event\"\x05\x9aD\x02\b\x01(\x010\x01\x12F\n\x04Sync\x12\x1d.protocol.chat.v1.SyncRequest\x1a\x1b.protocol.chat.v1.SyncEvent\"\x000\x01\x12U\n\aGetUser\x12 .protocol.chat.v1.GetUserRequest\x1a!.protocol.chat.v1.GetUserResponse\"\x05\x9aD\x02\b\x01\x12a\n\vGetUserBulk\x12$.protocol.chat.v1.GetUserBulkRequest\x1a%.protocol.chat.v1.GetUserBulkResponse\"\x05\x9aD\x02\b\x01\x12m\n\x0fGetUserMetadata\x12(.protocol.chat.v1.GetUserMetadataRequest\x1a).protocol.chat.v1.GetUserMetadataResponse\"\x05\x9aD\x02\b\x01\x12V\n\rProfileUpdate\x12&.protocol.chat.v1.ProfileUpdateRequest\x1a\x16.google.protobuf.Empty\"\x05\x9aD\x02\b\x01\x12Z\n\x06Typing\x12\x1f.protocol.chat.v1.TypingRequest\x1a\x16.google.protobuf.Empty\"\x17\x9aD\x02\b\x01\x9aD\x0f\x1a\rmessages.send\x12d\n\fPreviewGuild\x12%.protocol.chat.v1.PreviewGuildRequest\x1a&.protocol.chat.v1.PreviewGuildResponse\"\x05\x9aD\x02\b\x00\x12^\n\aBanUser\x12 .protocol.chat.v1.BanUserRequest\x1a\x16.google.protobuf.Empty\"\x19\x9aD\x02\b\x01\x9aD\x11\x1a\x0fuser.manage.ban\x12a\n\bKickUser\x12!.protocol.chat.v1.KickUserRequest\x1a\x16.google.protobuf.Empty\"\x1a\x9aD\x02\b\x01\x9aD\x12\x1a\x10user.manage.kick\x12d\n\tUnbanUser\x12\".protocol.chat.v1.UnbanUserRequest\x1a\x16.google.protobuf.Empty\"\x1b\x9aD\x02\b\x01\x9aD\x13\x1a\x11user.manage.unban")

	err := proto.Unmarshal(data, ChatServiceData)
	if err != nil {
		panic(err)
	}
}

type ChatServiceServer interface {
	CreateGuild(ctx echo.Context, r *CreateGuildRequest) (resp *CreateGuildResponse, err error)

	CreateInvite(ctx echo.Context, r *CreateInviteRequest) (resp *CreateInviteResponse, err error)

	CreateChannel(ctx echo.Context, r *CreateChannelRequest) (resp *CreateChannelResponse, err error)

	CreateEmotePack(ctx echo.Context, r *CreateEmotePackRequest) (resp *CreateEmotePackResponse, err error)

	GetGuildList(ctx echo.Context, r *GetGuildListRequest) (resp *GetGuildListResponse, err error)

	AddGuildToGuildList(ctx echo.Context, r *AddGuildToGuildListRequest) (resp *AddGuildToGuildListResponse, err error)

	RemoveGuildFromGuildList(ctx echo.Context, r *RemoveGuildFromGuildListRequest) (resp *RemoveGuildFromGuildListResponse, err error)

	GetGuild(ctx echo.Context, r *GetGuildRequest) (resp *GetGuildResponse, err error)

	GetGuildInvites(ctx echo.Context, r *GetGuildInvitesRequest) (resp *GetGuildInvitesResponse, err error)

	GetGuildMembers(ctx echo.Context, r *GetGuildMembersRequest) (resp *GetGuildMembersResponse, err error)

	GetGuildChannels(ctx echo.Context, r *GetGuildChannelsRequest) (resp *GetGuildChannelsResponse, err error)

	GetChannelMessages(ctx echo.Context, r *GetChannelMessagesRequest) (resp *GetChannelMessagesResponse, err error)

	GetMessage(ctx echo.Context, r *GetMessageRequest) (resp *GetMessageResponse, err error)

	GetEmotePacks(ctx echo.Context, r *GetEmotePacksRequest) (resp *GetEmotePacksResponse, err error)

	GetEmotePackEmotes(ctx echo.Context, r *GetEmotePackEmotesRequest) (resp *GetEmotePackEmotesResponse, err error)

	UpdateGuildInformation(ctx echo.Context, r *UpdateGuildInformationRequest) (resp *empty.Empty, err error)

	UpdateChannelInformation(ctx echo.Context, r *UpdateChannelInformationRequest) (resp *empty.Empty, err error)

	UpdateChannelOrder(ctx echo.Context, r *UpdateChannelOrderRequest) (resp *empty.Empty, err error)

	UpdateMessage(ctx echo.Context, r *UpdateMessageTextRequest) (resp *empty.Empty, err error)

	AddEmoteToPack(ctx echo.Context, r *AddEmoteToPackRequest) (resp *empty.Empty, err error)

	DeleteGuild(ctx echo.Context, r *DeleteGuildRequest) (resp *empty.Empty, err error)

	DeleteInvite(ctx echo.Context, r *DeleteInviteRequest) (resp *empty.Empty, err error)

	DeleteChannel(ctx echo.Context, r *DeleteChannelRequest) (resp *empty.Empty, err error)

	DeleteMessage(ctx echo.Context, r *DeleteMessageRequest) (resp *empty.Empty, err error)

	DeleteEmoteFromPack(ctx echo.Context, r *DeleteEmoteFromPackRequest) (resp *empty.Empty, err error)

	DeleteEmotePack(ctx echo.Context, r *DeleteEmotePackRequest) (resp *empty.Empty, err error)

	DequipEmotePack(ctx echo.Context, r *DequipEmotePackRequest) (resp *empty.Empty, err error)

	JoinGuild(ctx echo.Context, r *JoinGuildRequest) (resp *JoinGuildResponse, err error)

	LeaveGuild(ctx echo.Context, r *LeaveGuildRequest) (resp *empty.Empty, err error)

	TriggerAction(ctx echo.Context, r *TriggerActionRequest) (resp *empty.Empty, err error)

	SendMessage(ctx echo.Context, r *SendMessageRequest) (resp *SendMessageResponse, err error)

	QueryHasPermission(ctx echo.Context, r *QueryPermissionsRequest) (resp *QueryPermissionsResponse, err error)

	SetPermissions(ctx echo.Context, r *SetPermissionsRequest) (resp *empty.Empty, err error)

	GetPermissions(ctx echo.Context, r *GetPermissionsRequest) (resp *GetPermissionsResponse, err error)

	MoveRole(ctx echo.Context, r *MoveRoleRequest) (resp *MoveRoleResponse, err error)

	GetGuildRoles(ctx echo.Context, r *GetGuildRolesRequest) (resp *GetGuildRolesResponse, err error)

	AddGuildRole(ctx echo.Context, r *AddGuildRoleRequest) (resp *AddGuildRoleResponse, err error)

	ModifyGuildRole(ctx echo.Context, r *ModifyGuildRoleRequest) (resp *empty.Empty, err error)

	DeleteGuildRole(ctx echo.Context, r *DeleteGuildRoleRequest) (resp *empty.Empty, err error)

	ManageUserRoles(ctx echo.Context, r *ManageUserRolesRequest) (resp *empty.Empty, err error)

	GetUserRoles(ctx echo.Context, r *GetUserRolesRequest) (resp *GetUserRolesResponse, err error)

	StreamEvents(ctx echo.Context, in chan *StreamEventsRequest, out chan *Event)

	Sync(ctx echo.Context, r *SyncRequest, out chan *SyncEvent)

	GetUser(ctx echo.Context, r *GetUserRequest) (resp *GetUserResponse, err error)

	GetUserBulk(ctx echo.Context, r *GetUserBulkRequest) (resp *GetUserBulkResponse, err error)

	GetUserMetadata(ctx echo.Context, r *GetUserMetadataRequest) (resp *GetUserMetadataResponse, err error)

	ProfileUpdate(ctx echo.Context, r *ProfileUpdateRequest) (resp *empty.Empty, err error)

	Typing(ctx echo.Context, r *TypingRequest) (resp *empty.Empty, err error)

	PreviewGuild(ctx echo.Context, r *PreviewGuildRequest) (resp *PreviewGuildResponse, err error)

	BanUser(ctx echo.Context, r *BanUserRequest) (resp *empty.Empty, err error)

	KickUser(ctx echo.Context, r *KickUserRequest) (resp *empty.Empty, err error)

	UnbanUser(ctx echo.Context, r *UnbanUserRequest) (resp *empty.Empty, err error)
}

var ChatServiceServerCreateGuildData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\vCreateGuild\x12$.protocol.chat.v1.CreateGuildRequest\x1a%.protocol.chat.v1.CreateGuildResponse\"\n\x9aD\x02\b\x01\x9aD\x02\x1a\x00")

	err := proto.Unmarshal(data, ChatServiceServerCreateGuildData)
	if err != nil {
		panic(err)
	}
}

var ChatServiceServerCreateInviteData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\fCreateInvite\x12%.protocol.chat.v1.CreateInviteRequest\x1a&.protocol.chat.v1.CreateInviteResponse\"\x1f\x9aD\x02\b\x01\x9aD\x17\x1a\x15invites.manage.create")

	err := proto.Unmarshal(data, ChatServiceServerCreateInviteData)
	if err != nil {
		panic(err)
	}
}

var ChatServiceServerCreateChannelData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\rCreateChannel\x12&.protocol.chat.v1.CreateChannelRequest\x1a'.protocol.chat.v1.CreateChannelResponse\" \x9aD\x02\b\x01\x9aD\x18\x1a\x16channels.manage.create")

	err := proto.Unmarshal(data, ChatServiceServerCreateChannelData)
	if err != nil {
		panic(err)
	}
}

var ChatServiceServerCreateEmotePackData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\x0fCreateEmotePack\x12(.protocol.chat.v1.CreateEmotePackRequest\x1a).protocol.chat.v1.CreateEmotePackResponse\"\x05\x9aD\x02\b\x01")

	err := proto.Unmarshal(data, ChatServiceServerCreateEmotePackData)
	if err != nil {
		panic(err)
	}
}

var ChatServiceServerGetGuildListData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\fGetGuildList\x12%.protocol.chat.v1.GetGuildListRequest\x1a&.protocol.chat.v1.GetGuildListResponse\"\x05\x9aD\x02\b\x01")

	err := proto.Unmarshal(data, ChatServiceServerGetGuildListData)
	if err != nil {
		panic(err)
	}
}

var ChatServiceServerAddGuildToGuildListData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\x13AddGuildToGuildList\x12,.protocol.chat.v1.AddGuildToGuildListRequest\x1a-.protocol.chat.v1.AddGuildToGuildListResponse\"\n\x9aD\x02\b\x01\x9aD\x02\x10\x01")

	err := proto.Unmarshal(data, ChatServiceServerAddGuildToGuildListData)
	if err != nil {
		panic(err)
	}
}

var ChatServiceServerRemoveGuildFromGuildListData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\x18RemoveGuildFromGuildList\x121.protocol.chat.v1.RemoveGuildFromGuildListRequest\x1a2.protocol.chat.v1.RemoveGuildFromGuildListResponse\"\n\x9aD\x02\b\x01\x9aD\x02\x10\x01")

	err := proto.Unmarshal(data, ChatServiceServerRemoveGuildFromGuildListData)
	if err != nil {
		panic(err)
	}
}

var ChatServiceServerGetGuildData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\bGetGuild\x12!.protocol.chat.v1.GetGuildRequest\x1a\".protocol.chat.v1.GetGuildResponse\"\x05\x9aD\x02\b\x01")

	err := proto.Unmarshal(data, ChatServiceServerGetGuildData)
	if err != nil {
		panic(err)
	}
}

var ChatServiceServerGetGuildInvitesData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\x0fGetGuildInvites\x12(.protocol.chat.v1.GetGuildInvitesRequest\x1a).protocol.chat.v1.GetGuildInvitesResponse\"\x16\x9aD\x02\b\x01\x9aD\x0e\x1a\finvites.view")

	err := proto.Unmarshal(data, ChatServiceServerGetGuildInvitesData)
	if err != nil {
		panic(err)
	}
}

var ChatServiceServerGetGuildMembersData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\x0fGetGuildMembers\x12(.protocol.chat.v1.GetGuildMembersRequest\x1a).protocol.chat.v1.GetGuildMembersResponse\"\x05\x9aD\x02\b\x01")

	err := proto.Unmarshal(data, ChatServiceServerGetGuildMembersData)
	if err != nil {
		panic(err)
	}
}

var ChatServiceServerGetGuildChannelsData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\x10GetGuildChannels\x12).protocol.chat.v1.GetGuildChannelsRequest\x1a*.protocol.chat.v1.GetGuildChannelsResponse\"\x00")

	err := proto.Unmarshal(data, ChatServiceServerGetGuildChannelsData)
	if err != nil {
		panic(err)
	}
}

var ChatServiceServerGetChannelMessagesData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\x12GetChannelMessages\x12+.protocol.chat.v1.GetChannelMessagesRequest\x1a,.protocol.chat.v1.GetChannelMessagesResponse\"\x17\x9aD\x02\b\x01\x9aD\x0f\x1a\rmessages.view")

	err := proto.Unmarshal(data, ChatServiceServerGetChannelMessagesData)
	if err != nil {
		panic(err)
	}
}

var ChatServiceServerGetMessageData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\nGetMessage\x12#.protocol.chat.v1.GetMessageRequest\x1a$.protocol.chat.v1.GetMessageResponse\"\x17\x9aD\x02\b\x01\x9aD\x0f\x1a\rmessages.view")

	err := proto.Unmarshal(data, ChatServiceServerGetMessageData)
	if err != nil {
		panic(err)
	}
}

var ChatServiceServerGetEmotePacksData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\rGetEmotePacks\x12&.protocol.chat.v1.GetEmotePacksRequest\x1a'.protocol.chat.v1.GetEmotePacksResponse\"\x05\x9aD\x02\b\x01")

	err := proto.Unmarshal(data, ChatServiceServerGetEmotePacksData)
	if err != nil {
		panic(err)
	}
}

var ChatServiceServerGetEmotePackEmotesData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\x12GetEmotePackEmotes\x12+.protocol.chat.v1.GetEmotePackEmotesRequest\x1a,.protocol.chat.v1.GetEmotePackEmotesResponse\"\x05\x9aD\x02\b\x01")

	err := proto.Unmarshal(data, ChatServiceServerGetEmotePackEmotesData)
	if err != nil {
		panic(err)
	}
}

var ChatServiceServerUpdateGuildInformationData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\x16UpdateGuildInformation\x12/.protocol.chat.v1.UpdateGuildInformationRequest\x1a\x16.google.protobuf.Empty\")\x9aD\x02\b\x01\x9aD!\x1a\x1fguild.manage.change-information")

	err := proto.Unmarshal(data, ChatServiceServerUpdateGuildInformationData)
	if err != nil {
		panic(err)
	}
}

var ChatServiceServerUpdateChannelInformationData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\x18UpdateChannelInformation\x121.protocol.chat.v1.UpdateChannelInformationRequest\x1a\x16.google.protobuf.Empty\",\x9aD\x02\b\x01\x9aD$\x1a\"channels.manage.change-information")

	err := proto.Unmarshal(data, ChatServiceServerUpdateChannelInformationData)
	if err != nil {
		panic(err)
	}
}

var ChatServiceServerUpdateChannelOrderData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\x12UpdateChannelOrder\x12+.protocol.chat.v1.UpdateChannelOrderRequest\x1a\x16.google.protobuf.Empty\"\x1e\x9aD\x02\b\x01\x9aD\x16\x1a\x14channels.manage.move")

	err := proto.Unmarshal(data, ChatServiceServerUpdateChannelOrderData)
	if err != nil {
		panic(err)
	}
}

var ChatServiceServerUpdateMessageData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\rUpdateMessage\x12*.protocol.chat.v1.UpdateMessageTextRequest\x1a\x16.google.protobuf.Empty\"\x17\x9aD\x02\b\x01\x9aD\x0f\x1a\rmessages.send")

	err := proto.Unmarshal(data, ChatServiceServerUpdateMessageData)
	if err != nil {
		panic(err)
	}
}

var ChatServiceServerAddEmoteToPackData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\x0eAddEmoteToPack\x12'.protocol.chat.v1.AddEmoteToPackRequest\x1a\x16.google.protobuf.Empty\"\x05\x9aD\x02\b\x01")

	err := proto.Unmarshal(data, ChatServiceServerAddEmoteToPackData)
	if err != nil {
		panic(err)
	}
}

var ChatServiceServerDeleteGuildData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\vDeleteGuild\x12$.protocol.chat.v1.DeleteGuildRequest\x1a\x16.google.protobuf.Empty\"\x1d\x9aD\x02\b\x01\x9aD\x15\x1a\x13guild.manage.delete")

	err := proto.Unmarshal(data, ChatServiceServerDeleteGuildData)
	if err != nil {
		panic(err)
	}
}

var ChatServiceServerDeleteInviteData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\fDeleteInvite\x12%.protocol.chat.v1.DeleteInviteRequest\x1a\x16.google.protobuf.Empty\"\x1f\x9aD\x02\b\x01\x9aD\x17\x1a\x15invites.manage.delete")

	err := proto.Unmarshal(data, ChatServiceServerDeleteInviteData)
	if err != nil {
		panic(err)
	}
}

var ChatServiceServerDeleteChannelData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\rDeleteChannel\x12&.protocol.chat.v1.DeleteChannelRequest\x1a\x16.google.protobuf.Empty\" \x9aD\x02\b\x01\x9aD\x18\x1a\x16channels.manage.delete")

	err := proto.Unmarshal(data, ChatServiceServerDeleteChannelData)
	if err != nil {
		panic(err)
	}
}

var ChatServiceServerDeleteMessageData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\rDeleteMessage\x12&.protocol.chat.v1.DeleteMessageRequest\x1a\x16.google.protobuf.Empty\"\x05\x9aD\x02\b\x01")

	err := proto.Unmarshal(data, ChatServiceServerDeleteMessageData)
	if err != nil {
		panic(err)
	}
}

var ChatServiceServerDeleteEmoteFromPackData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\x13DeleteEmoteFromPack\x12,.protocol.chat.v1.DeleteEmoteFromPackRequest\x1a\x16.google.protobuf.Empty\"\x05\x9aD\x02\b\x01")

	err := proto.Unmarshal(data, ChatServiceServerDeleteEmoteFromPackData)
	if err != nil {
		panic(err)
	}
}

var ChatServiceServerDeleteEmotePackData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\x0fDeleteEmotePack\x12(.protocol.chat.v1.DeleteEmotePackRequest\x1a\x16.google.protobuf.Empty\"\x05\x9aD\x02\b\x01")

	err := proto.Unmarshal(data, ChatServiceServerDeleteEmotePackData)
	if err != nil {
		panic(err)
	}
}

var ChatServiceServerDequipEmotePackData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\x0fDequipEmotePack\x12(.protocol.chat.v1.DequipEmotePackRequest\x1a\x16.google.protobuf.Empty\"\x05\x9aD\x02\b\x01")

	err := proto.Unmarshal(data, ChatServiceServerDequipEmotePackData)
	if err != nil {
		panic(err)
	}
}

var ChatServiceServerJoinGuildData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\tJoinGuild\x12\".protocol.chat.v1.JoinGuildRequest\x1a#.protocol.chat.v1.JoinGuildResponse\"\x05\x9aD\x02\b\x01")

	err := proto.Unmarshal(data, ChatServiceServerJoinGuildData)
	if err != nil {
		panic(err)
	}
}

var ChatServiceServerLeaveGuildData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\nLeaveGuild\x12#.protocol.chat.v1.LeaveGuildRequest\x1a\x16.google.protobuf.Empty\"\x05\x9aD\x02\b\x01")

	err := proto.Unmarshal(data, ChatServiceServerLeaveGuildData)
	if err != nil {
		panic(err)
	}
}

var ChatServiceServerTriggerActionData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\rTriggerAction\x12&.protocol.chat.v1.TriggerActionRequest\x1a\x16.google.protobuf.Empty\"\x19\x9aD\x02\b\x01\x9aD\x11\x1a\x0factions.trigger")

	err := proto.Unmarshal(data, ChatServiceServerTriggerActionData)
	if err != nil {
		panic(err)
	}
}

var ChatServiceServerSendMessageData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\vSendMessage\x12$.protocol.chat.v1.SendMessageRequest\x1a%.protocol.chat.v1.SendMessageResponse\"\x17\x9aD\x02\b\x01\x9aD\x0f\x1a\rmessages.send")

	err := proto.Unmarshal(data, ChatServiceServerSendMessageData)
	if err != nil {
		panic(err)
	}
}

var ChatServiceServerQueryHasPermissionData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\x12QueryHasPermission\x12).protocol.chat.v1.QueryPermissionsRequest\x1a*.protocol.chat.v1.QueryPermissionsResponse\"\x05\x9aD\x02\b\x01")

	err := proto.Unmarshal(data, ChatServiceServerQueryHasPermissionData)
	if err != nil {
		panic(err)
	}
}

var ChatServiceServerSetPermissionsData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\x0eSetPermissions\x12'.protocol.chat.v1.SetPermissionsRequest\x1a\x16.google.protobuf.Empty\" \x9aD\x02\b\x01\x9aD\x18\x1a\x16permissions.manage.set")

	err := proto.Unmarshal(data, ChatServiceServerSetPermissionsData)
	if err != nil {
		panic(err)
	}
}

var ChatServiceServerGetPermissionsData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\x0eGetPermissions\x12'.protocol.chat.v1.GetPermissionsRequest\x1a(.protocol.chat.v1.GetPermissionsResponse\" \x9aD\x02\b\x01\x9aD\x18\x1a\x16permissions.manage.get")

	err := proto.Unmarshal(data, ChatServiceServerGetPermissionsData)
	if err != nil {
		panic(err)
	}
}

var ChatServiceServerMoveRoleData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\bMoveRole\x12!.protocol.chat.v1.MoveRoleRequest\x1a\".protocol.chat.v1.MoveRoleResponse\"\x16\x9aD\x02\b\x01\x9aD\x0e\x1a\froles.manage")

	err := proto.Unmarshal(data, ChatServiceServerMoveRoleData)
	if err != nil {
		panic(err)
	}
}

var ChatServiceServerGetGuildRolesData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\rGetGuildRoles\x12&.protocol.chat.v1.GetGuildRolesRequest\x1a'.protocol.chat.v1.GetGuildRolesResponse\"\x13\x9aD\x02\b\x01\x9aD\v\x1a\troles.get")

	err := proto.Unmarshal(data, ChatServiceServerGetGuildRolesData)
	if err != nil {
		panic(err)
	}
}

var ChatServiceServerAddGuildRoleData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\fAddGuildRole\x12%.protocol.chat.v1.AddGuildRoleRequest\x1a&.protocol.chat.v1.AddGuildRoleResponse\"\x16\x9aD\x02\b\x01\x9aD\x0e\x1a\froles.manage")

	err := proto.Unmarshal(data, ChatServiceServerAddGuildRoleData)
	if err != nil {
		panic(err)
	}
}

var ChatServiceServerModifyGuildRoleData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\x0fModifyGuildRole\x12(.protocol.chat.v1.ModifyGuildRoleRequest\x1a\x16.google.protobuf.Empty\"\x16\x9aD\x02\b\x01\x9aD\x0e\x1a\froles.manage")

	err := proto.Unmarshal(data, ChatServiceServerModifyGuildRoleData)
	if err != nil {
		panic(err)
	}
}

var ChatServiceServerDeleteGuildRoleData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\x0fDeleteGuildRole\x12(.protocol.chat.v1.DeleteGuildRoleRequest\x1a\x16.google.protobuf.Empty\"\x16\x9aD\x02\b\x01\x9aD\x0e\x1a\froles.manage")

	err := proto.Unmarshal(data, ChatServiceServerDeleteGuildRoleData)
	if err != nil {
		panic(err)
	}
}

var ChatServiceServerManageUserRolesData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\x0fManageUserRoles\x12(.protocol.chat.v1.ManageUserRolesRequest\x1a\x16.google.protobuf.Empty\"\x1b\x9aD\x02\b\x01\x9aD\x13\x1a\x11roles.user.manage")

	err := proto.Unmarshal(data, ChatServiceServerManageUserRolesData)
	if err != nil {
		panic(err)
	}
}

var ChatServiceServerGetUserRolesData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\fGetUserRoles\x12%.protocol.chat.v1.GetUserRolesRequest\x1a&.protocol.chat.v1.GetUserRolesResponse\"\x05\x9aD\x02\b\x01")

	err := proto.Unmarshal(data, ChatServiceServerGetUserRolesData)
	if err != nil {
		panic(err)
	}
}

var ChatServiceServerStreamEventsData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\fStreamEvents\x12%.protocol.chat.v1.StreamEventsRequest\x1a\x17.protocol.chat.v1.Event\"\x05\x9aD\x02\b\x01(\x010\x01")

	err := proto.Unmarshal(data, ChatServiceServerStreamEventsData)
	if err != nil {
		panic(err)
	}
}

var ChatServiceServerSyncData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\x04Sync\x12\x1d.protocol.chat.v1.SyncRequest\x1a\x1b.protocol.chat.v1.SyncEvent\"\x000\x01")

	err := proto.Unmarshal(data, ChatServiceServerSyncData)
	if err != nil {
		panic(err)
	}
}

var ChatServiceServerGetUserData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\aGetUser\x12 .protocol.chat.v1.GetUserRequest\x1a!.protocol.chat.v1.GetUserResponse\"\x05\x9aD\x02\b\x01")

	err := proto.Unmarshal(data, ChatServiceServerGetUserData)
	if err != nil {
		panic(err)
	}
}

var ChatServiceServerGetUserBulkData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\vGetUserBulk\x12$.protocol.chat.v1.GetUserBulkRequest\x1a%.protocol.chat.v1.GetUserBulkResponse\"\x05\x9aD\x02\b\x01")

	err := proto.Unmarshal(data, ChatServiceServerGetUserBulkData)
	if err != nil {
		panic(err)
	}
}

var ChatServiceServerGetUserMetadataData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\x0fGetUserMetadata\x12(.protocol.chat.v1.GetUserMetadataRequest\x1a).protocol.chat.v1.GetUserMetadataResponse\"\x05\x9aD\x02\b\x01")

	err := proto.Unmarshal(data, ChatServiceServerGetUserMetadataData)
	if err != nil {
		panic(err)
	}
}

var ChatServiceServerProfileUpdateData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\rProfileUpdate\x12&.protocol.chat.v1.ProfileUpdateRequest\x1a\x16.google.protobuf.Empty\"\x05\x9aD\x02\b\x01")

	err := proto.Unmarshal(data, ChatServiceServerProfileUpdateData)
	if err != nil {
		panic(err)
	}
}

var ChatServiceServerTypingData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\x06Typing\x12\x1f.protocol.chat.v1.TypingRequest\x1a\x16.google.protobuf.Empty\"\x17\x9aD\x02\b\x01\x9aD\x0f\x1a\rmessages.send")

	err := proto.Unmarshal(data, ChatServiceServerTypingData)
	if err != nil {
		panic(err)
	}
}

var ChatServiceServerPreviewGuildData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\fPreviewGuild\x12%.protocol.chat.v1.PreviewGuildRequest\x1a&.protocol.chat.v1.PreviewGuildResponse\"\x05\x9aD\x02\b\x00")

	err := proto.Unmarshal(data, ChatServiceServerPreviewGuildData)
	if err != nil {
		panic(err)
	}
}

var ChatServiceServerBanUserData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\aBanUser\x12 .protocol.chat.v1.BanUserRequest\x1a\x16.google.protobuf.Empty\"\x19\x9aD\x02\b\x01\x9aD\x11\x1a\x0fuser.manage.ban")

	err := proto.Unmarshal(data, ChatServiceServerBanUserData)
	if err != nil {
		panic(err)
	}
}

var ChatServiceServerKickUserData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\bKickUser\x12!.protocol.chat.v1.KickUserRequest\x1a\x16.google.protobuf.Empty\"\x1a\x9aD\x02\b\x01\x9aD\x12\x1a\x10user.manage.kick")

	err := proto.Unmarshal(data, ChatServiceServerKickUserData)
	if err != nil {
		panic(err)
	}
}

var ChatServiceServerUnbanUserData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\tUnbanUser\x12\".protocol.chat.v1.UnbanUserRequest\x1a\x16.google.protobuf.Empty\"\x1b\x9aD\x02\b\x01\x9aD\x13\x1a\x11user.manage.unban")

	err := proto.Unmarshal(data, ChatServiceServerUnbanUserData)
	if err != nil {
		panic(err)
	}
}

type ChatServiceHandler struct {
	Server       ChatServiceServer
	ErrorHandler func(err error, w http.ResponseWriter)
	UnaryPre     server.HandlerTransformer
	upgrader     websocket.Upgrader
}

func NewChatServiceHandler(s ChatServiceServer) *ChatServiceHandler {
	return &ChatServiceHandler{
		Server: s,
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(_ *http.Request) bool {
				return true
			},
		},
	}
}

func (h *ChatServiceHandler) SetUnaryPre(s server.HandlerTransformer) {
	h.UnaryPre = s
}

func (h *ChatServiceHandler) Routes() map[string]echo.HandlerFunc {
	return map[string]echo.HandlerFunc{

		"/protocol.chat.v1.ChatService/CreateGuild": h.CreateGuildHandler,

		"/protocol.chat.v1.ChatService/CreateInvite": h.CreateInviteHandler,

		"/protocol.chat.v1.ChatService/CreateChannel": h.CreateChannelHandler,

		"/protocol.chat.v1.ChatService/CreateEmotePack": h.CreateEmotePackHandler,

		"/protocol.chat.v1.ChatService/GetGuildList": h.GetGuildListHandler,

		"/protocol.chat.v1.ChatService/AddGuildToGuildList": h.AddGuildToGuildListHandler,

		"/protocol.chat.v1.ChatService/RemoveGuildFromGuildList": h.RemoveGuildFromGuildListHandler,

		"/protocol.chat.v1.ChatService/GetGuild": h.GetGuildHandler,

		"/protocol.chat.v1.ChatService/GetGuildInvites": h.GetGuildInvitesHandler,

		"/protocol.chat.v1.ChatService/GetGuildMembers": h.GetGuildMembersHandler,

		"/protocol.chat.v1.ChatService/GetGuildChannels": h.GetGuildChannelsHandler,

		"/protocol.chat.v1.ChatService/GetChannelMessages": h.GetChannelMessagesHandler,

		"/protocol.chat.v1.ChatService/GetMessage": h.GetMessageHandler,

		"/protocol.chat.v1.ChatService/GetEmotePacks": h.GetEmotePacksHandler,

		"/protocol.chat.v1.ChatService/GetEmotePackEmotes": h.GetEmotePackEmotesHandler,

		"/protocol.chat.v1.ChatService/UpdateGuildInformation": h.UpdateGuildInformationHandler,

		"/protocol.chat.v1.ChatService/UpdateChannelInformation": h.UpdateChannelInformationHandler,

		"/protocol.chat.v1.ChatService/UpdateChannelOrder": h.UpdateChannelOrderHandler,

		"/protocol.chat.v1.ChatService/UpdateMessage": h.UpdateMessageHandler,

		"/protocol.chat.v1.ChatService/AddEmoteToPack": h.AddEmoteToPackHandler,

		"/protocol.chat.v1.ChatService/DeleteGuild": h.DeleteGuildHandler,

		"/protocol.chat.v1.ChatService/DeleteInvite": h.DeleteInviteHandler,

		"/protocol.chat.v1.ChatService/DeleteChannel": h.DeleteChannelHandler,

		"/protocol.chat.v1.ChatService/DeleteMessage": h.DeleteMessageHandler,

		"/protocol.chat.v1.ChatService/DeleteEmoteFromPack": h.DeleteEmoteFromPackHandler,

		"/protocol.chat.v1.ChatService/DeleteEmotePack": h.DeleteEmotePackHandler,

		"/protocol.chat.v1.ChatService/DequipEmotePack": h.DequipEmotePackHandler,

		"/protocol.chat.v1.ChatService/JoinGuild": h.JoinGuildHandler,

		"/protocol.chat.v1.ChatService/LeaveGuild": h.LeaveGuildHandler,

		"/protocol.chat.v1.ChatService/TriggerAction": h.TriggerActionHandler,

		"/protocol.chat.v1.ChatService/SendMessage": h.SendMessageHandler,

		"/protocol.chat.v1.ChatService/QueryHasPermission": h.QueryHasPermissionHandler,

		"/protocol.chat.v1.ChatService/SetPermissions": h.SetPermissionsHandler,

		"/protocol.chat.v1.ChatService/GetPermissions": h.GetPermissionsHandler,

		"/protocol.chat.v1.ChatService/MoveRole": h.MoveRoleHandler,

		"/protocol.chat.v1.ChatService/GetGuildRoles": h.GetGuildRolesHandler,

		"/protocol.chat.v1.ChatService/AddGuildRole": h.AddGuildRoleHandler,

		"/protocol.chat.v1.ChatService/ModifyGuildRole": h.ModifyGuildRoleHandler,

		"/protocol.chat.v1.ChatService/DeleteGuildRole": h.DeleteGuildRoleHandler,

		"/protocol.chat.v1.ChatService/ManageUserRoles": h.ManageUserRolesHandler,

		"/protocol.chat.v1.ChatService/GetUserRoles": h.GetUserRolesHandler,

		"/protocol.chat.v1.ChatService/StreamEvents": h.StreamEventsHandler,

		"/protocol.chat.v1.ChatService/Sync": h.SyncHandler,

		"/protocol.chat.v1.ChatService/GetUser": h.GetUserHandler,

		"/protocol.chat.v1.ChatService/GetUserBulk": h.GetUserBulkHandler,

		"/protocol.chat.v1.ChatService/GetUserMetadata": h.GetUserMetadataHandler,

		"/protocol.chat.v1.ChatService/ProfileUpdate": h.ProfileUpdateHandler,

		"/protocol.chat.v1.ChatService/Typing": h.TypingHandler,

		"/protocol.chat.v1.ChatService/PreviewGuild": h.PreviewGuildHandler,

		"/protocol.chat.v1.ChatService/BanUser": h.BanUserHandler,

		"/protocol.chat.v1.ChatService/KickUser": h.KickUserHandler,

		"/protocol.chat.v1.ChatService/UnbanUser": h.UnbanUserHandler,
	}
}

func (h *ChatServiceHandler) CreateGuildHandler(c echo.Context) error {

	requestProto := new(CreateGuildRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.CreateGuild(c, req.(*CreateGuildRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(ChatServiceServerCreateGuildData, ChatServiceData, Chatᐳv1ᐳchat, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}

func (h *ChatServiceHandler) CreateInviteHandler(c echo.Context) error {

	requestProto := new(CreateInviteRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.CreateInvite(c, req.(*CreateInviteRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(ChatServiceServerCreateInviteData, ChatServiceData, Chatᐳv1ᐳchat, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}

func (h *ChatServiceHandler) CreateChannelHandler(c echo.Context) error {

	requestProto := new(CreateChannelRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.CreateChannel(c, req.(*CreateChannelRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(ChatServiceServerCreateChannelData, ChatServiceData, Chatᐳv1ᐳchat, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}

func (h *ChatServiceHandler) CreateEmotePackHandler(c echo.Context) error {

	requestProto := new(CreateEmotePackRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.CreateEmotePack(c, req.(*CreateEmotePackRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(ChatServiceServerCreateEmotePackData, ChatServiceData, Chatᐳv1ᐳchat, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}

func (h *ChatServiceHandler) GetGuildListHandler(c echo.Context) error {

	requestProto := new(GetGuildListRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.GetGuildList(c, req.(*GetGuildListRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(ChatServiceServerGetGuildListData, ChatServiceData, Chatᐳv1ᐳchat, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}

func (h *ChatServiceHandler) AddGuildToGuildListHandler(c echo.Context) error {

	requestProto := new(AddGuildToGuildListRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.AddGuildToGuildList(c, req.(*AddGuildToGuildListRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(ChatServiceServerAddGuildToGuildListData, ChatServiceData, Chatᐳv1ᐳchat, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}

func (h *ChatServiceHandler) RemoveGuildFromGuildListHandler(c echo.Context) error {

	requestProto := new(RemoveGuildFromGuildListRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.RemoveGuildFromGuildList(c, req.(*RemoveGuildFromGuildListRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(ChatServiceServerRemoveGuildFromGuildListData, ChatServiceData, Chatᐳv1ᐳchat, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}

func (h *ChatServiceHandler) GetGuildHandler(c echo.Context) error {

	requestProto := new(GetGuildRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.GetGuild(c, req.(*GetGuildRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(ChatServiceServerGetGuildData, ChatServiceData, Chatᐳv1ᐳchat, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}

func (h *ChatServiceHandler) GetGuildInvitesHandler(c echo.Context) error {

	requestProto := new(GetGuildInvitesRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.GetGuildInvites(c, req.(*GetGuildInvitesRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(ChatServiceServerGetGuildInvitesData, ChatServiceData, Chatᐳv1ᐳchat, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}

func (h *ChatServiceHandler) GetGuildMembersHandler(c echo.Context) error {

	requestProto := new(GetGuildMembersRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.GetGuildMembers(c, req.(*GetGuildMembersRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(ChatServiceServerGetGuildMembersData, ChatServiceData, Chatᐳv1ᐳchat, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}

func (h *ChatServiceHandler) GetGuildChannelsHandler(c echo.Context) error {

	requestProto := new(GetGuildChannelsRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.GetGuildChannels(c, req.(*GetGuildChannelsRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(ChatServiceServerGetGuildChannelsData, ChatServiceData, Chatᐳv1ᐳchat, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}

func (h *ChatServiceHandler) GetChannelMessagesHandler(c echo.Context) error {

	requestProto := new(GetChannelMessagesRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.GetChannelMessages(c, req.(*GetChannelMessagesRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(ChatServiceServerGetChannelMessagesData, ChatServiceData, Chatᐳv1ᐳchat, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}

func (h *ChatServiceHandler) GetMessageHandler(c echo.Context) error {

	requestProto := new(GetMessageRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.GetMessage(c, req.(*GetMessageRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(ChatServiceServerGetMessageData, ChatServiceData, Chatᐳv1ᐳchat, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}

func (h *ChatServiceHandler) GetEmotePacksHandler(c echo.Context) error {

	requestProto := new(GetEmotePacksRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.GetEmotePacks(c, req.(*GetEmotePacksRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(ChatServiceServerGetEmotePacksData, ChatServiceData, Chatᐳv1ᐳchat, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}

func (h *ChatServiceHandler) GetEmotePackEmotesHandler(c echo.Context) error {

	requestProto := new(GetEmotePackEmotesRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.GetEmotePackEmotes(c, req.(*GetEmotePackEmotesRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(ChatServiceServerGetEmotePackEmotesData, ChatServiceData, Chatᐳv1ᐳchat, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}

func (h *ChatServiceHandler) UpdateGuildInformationHandler(c echo.Context) error {

	requestProto := new(UpdateGuildInformationRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.UpdateGuildInformation(c, req.(*UpdateGuildInformationRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(ChatServiceServerUpdateGuildInformationData, ChatServiceData, Chatᐳv1ᐳchat, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}

func (h *ChatServiceHandler) UpdateChannelInformationHandler(c echo.Context) error {

	requestProto := new(UpdateChannelInformationRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.UpdateChannelInformation(c, req.(*UpdateChannelInformationRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(ChatServiceServerUpdateChannelInformationData, ChatServiceData, Chatᐳv1ᐳchat, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}

func (h *ChatServiceHandler) UpdateChannelOrderHandler(c echo.Context) error {

	requestProto := new(UpdateChannelOrderRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.UpdateChannelOrder(c, req.(*UpdateChannelOrderRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(ChatServiceServerUpdateChannelOrderData, ChatServiceData, Chatᐳv1ᐳchat, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}

func (h *ChatServiceHandler) UpdateMessageHandler(c echo.Context) error {

	requestProto := new(UpdateMessageTextRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.UpdateMessage(c, req.(*UpdateMessageTextRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(ChatServiceServerUpdateMessageData, ChatServiceData, Chatᐳv1ᐳchat, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}

func (h *ChatServiceHandler) AddEmoteToPackHandler(c echo.Context) error {

	requestProto := new(AddEmoteToPackRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.AddEmoteToPack(c, req.(*AddEmoteToPackRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(ChatServiceServerAddEmoteToPackData, ChatServiceData, Chatᐳv1ᐳchat, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}

func (h *ChatServiceHandler) DeleteGuildHandler(c echo.Context) error {

	requestProto := new(DeleteGuildRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.DeleteGuild(c, req.(*DeleteGuildRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(ChatServiceServerDeleteGuildData, ChatServiceData, Chatᐳv1ᐳchat, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}

func (h *ChatServiceHandler) DeleteInviteHandler(c echo.Context) error {

	requestProto := new(DeleteInviteRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.DeleteInvite(c, req.(*DeleteInviteRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(ChatServiceServerDeleteInviteData, ChatServiceData, Chatᐳv1ᐳchat, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}

func (h *ChatServiceHandler) DeleteChannelHandler(c echo.Context) error {

	requestProto := new(DeleteChannelRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.DeleteChannel(c, req.(*DeleteChannelRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(ChatServiceServerDeleteChannelData, ChatServiceData, Chatᐳv1ᐳchat, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}

func (h *ChatServiceHandler) DeleteMessageHandler(c echo.Context) error {

	requestProto := new(DeleteMessageRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.DeleteMessage(c, req.(*DeleteMessageRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(ChatServiceServerDeleteMessageData, ChatServiceData, Chatᐳv1ᐳchat, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}

func (h *ChatServiceHandler) DeleteEmoteFromPackHandler(c echo.Context) error {

	requestProto := new(DeleteEmoteFromPackRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.DeleteEmoteFromPack(c, req.(*DeleteEmoteFromPackRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(ChatServiceServerDeleteEmoteFromPackData, ChatServiceData, Chatᐳv1ᐳchat, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}

func (h *ChatServiceHandler) DeleteEmotePackHandler(c echo.Context) error {

	requestProto := new(DeleteEmotePackRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.DeleteEmotePack(c, req.(*DeleteEmotePackRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(ChatServiceServerDeleteEmotePackData, ChatServiceData, Chatᐳv1ᐳchat, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}

func (h *ChatServiceHandler) DequipEmotePackHandler(c echo.Context) error {

	requestProto := new(DequipEmotePackRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.DequipEmotePack(c, req.(*DequipEmotePackRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(ChatServiceServerDequipEmotePackData, ChatServiceData, Chatᐳv1ᐳchat, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}

func (h *ChatServiceHandler) JoinGuildHandler(c echo.Context) error {

	requestProto := new(JoinGuildRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.JoinGuild(c, req.(*JoinGuildRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(ChatServiceServerJoinGuildData, ChatServiceData, Chatᐳv1ᐳchat, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}

func (h *ChatServiceHandler) LeaveGuildHandler(c echo.Context) error {

	requestProto := new(LeaveGuildRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.LeaveGuild(c, req.(*LeaveGuildRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(ChatServiceServerLeaveGuildData, ChatServiceData, Chatᐳv1ᐳchat, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}

func (h *ChatServiceHandler) TriggerActionHandler(c echo.Context) error {

	requestProto := new(TriggerActionRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.TriggerAction(c, req.(*TriggerActionRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(ChatServiceServerTriggerActionData, ChatServiceData, Chatᐳv1ᐳchat, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}

func (h *ChatServiceHandler) SendMessageHandler(c echo.Context) error {

	requestProto := new(SendMessageRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.SendMessage(c, req.(*SendMessageRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(ChatServiceServerSendMessageData, ChatServiceData, Chatᐳv1ᐳchat, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}

func (h *ChatServiceHandler) QueryHasPermissionHandler(c echo.Context) error {

	requestProto := new(QueryPermissionsRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.QueryHasPermission(c, req.(*QueryPermissionsRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(ChatServiceServerQueryHasPermissionData, ChatServiceData, Chatᐳv1ᐳchat, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}

func (h *ChatServiceHandler) SetPermissionsHandler(c echo.Context) error {

	requestProto := new(SetPermissionsRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.SetPermissions(c, req.(*SetPermissionsRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(ChatServiceServerSetPermissionsData, ChatServiceData, Chatᐳv1ᐳchat, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}

func (h *ChatServiceHandler) GetPermissionsHandler(c echo.Context) error {

	requestProto := new(GetPermissionsRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.GetPermissions(c, req.(*GetPermissionsRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(ChatServiceServerGetPermissionsData, ChatServiceData, Chatᐳv1ᐳchat, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}

func (h *ChatServiceHandler) MoveRoleHandler(c echo.Context) error {

	requestProto := new(MoveRoleRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.MoveRole(c, req.(*MoveRoleRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(ChatServiceServerMoveRoleData, ChatServiceData, Chatᐳv1ᐳchat, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}

func (h *ChatServiceHandler) GetGuildRolesHandler(c echo.Context) error {

	requestProto := new(GetGuildRolesRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.GetGuildRoles(c, req.(*GetGuildRolesRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(ChatServiceServerGetGuildRolesData, ChatServiceData, Chatᐳv1ᐳchat, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}

func (h *ChatServiceHandler) AddGuildRoleHandler(c echo.Context) error {

	requestProto := new(AddGuildRoleRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.AddGuildRole(c, req.(*AddGuildRoleRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(ChatServiceServerAddGuildRoleData, ChatServiceData, Chatᐳv1ᐳchat, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}

func (h *ChatServiceHandler) ModifyGuildRoleHandler(c echo.Context) error {

	requestProto := new(ModifyGuildRoleRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.ModifyGuildRole(c, req.(*ModifyGuildRoleRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(ChatServiceServerModifyGuildRoleData, ChatServiceData, Chatᐳv1ᐳchat, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}

func (h *ChatServiceHandler) DeleteGuildRoleHandler(c echo.Context) error {

	requestProto := new(DeleteGuildRoleRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.DeleteGuildRole(c, req.(*DeleteGuildRoleRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(ChatServiceServerDeleteGuildRoleData, ChatServiceData, Chatᐳv1ᐳchat, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}

func (h *ChatServiceHandler) ManageUserRolesHandler(c echo.Context) error {

	requestProto := new(ManageUserRolesRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.ManageUserRoles(c, req.(*ManageUserRolesRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(ChatServiceServerManageUserRolesData, ChatServiceData, Chatᐳv1ᐳchat, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}

func (h *ChatServiceHandler) GetUserRolesHandler(c echo.Context) error {

	requestProto := new(GetUserRolesRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.GetUserRoles(c, req.(*GetUserRolesRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(ChatServiceServerGetUserRolesData, ChatServiceData, Chatᐳv1ᐳchat, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}

func (h *ChatServiceHandler) StreamEventsHandler(c echo.Context) error {

	ws, err := h.upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		c.Logger().Error(err)
		return nil
	}
	defer ws.Close()

	in := make(chan *StreamEventsRequest, 100)

	out := make(chan *Event, 100)

	h.Server.StreamEvents(c, in, out)

	msgs := make(chan []byte)

	go func() {
		for {
			_, message, err := ws.ReadMessage()
			if err != nil {
				close(msgs)
				break
			}
			msgs <- message
		}
	}()

	defer ws.Close()

	for {
		select {
		case data, ok := <-msgs:
			if !ok {
				close(in)
				close(out)
				return nil
			}

			item := new(StreamEventsRequest)
			switch c.Request().Header.Get("Content-Type") {
			case "application/hrpc-json":
				if err = protojson.Unmarshal(data, item); err != nil {
					close(in)
					close(out)
					c.Logger().Error(err)
					return nil
				}
			default:
				if err = proto.Unmarshal(data, item); err != nil {
					close(in)
					close(out)
					c.Logger().Error(err)
					return nil
				}
			}

			in <- item
		case msg, ok := <-out:
			if !ok {
				close(in)
				close(out)
				return nil
			}

			w, err := ws.NextWriter(websocket.BinaryMessage)
			if err != nil {

				close(in)

				close(out)
				c.Logger().Error(err)
				return nil
			}

			var response []byte

			switch c.Request().Header.Get("Content-Type") {
			case "application/hrpc-json":
				response, err = protojson.Marshal(msg)
			default:
				response, err = proto.Marshal(msg)
			}

			if err != nil {

				close(in)

				close(out)
				c.Logger().Error(err)
				return nil
			}

			if _, err := w.Write(response); err != nil {

				close(in)

				close(out)
				c.Logger().Error(err)
				return nil
			}
			if err := w.Close(); err != nil {

				close(in)

				close(out)
				c.Logger().Error(err)
				return nil
			}
		}

	}

}

func (h *ChatServiceHandler) SyncHandler(c echo.Context) error {

	ws, err := h.upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		c.Logger().Error(err)
		return nil
	}
	defer ws.Close()

	in := new(SyncRequest)
	_, message, err := ws.ReadMessage()
	if err != nil {
		c.Logger().Error(err)
		return nil
	}
	if err := proto.Unmarshal(message, in); err != nil {
		c.Logger().Error(err)
		return nil
	}
	switch c.Request().Header.Get("Content-Type") {
	case "application/hrpc-json":
		if err = protojson.Unmarshal(message, in); err != nil {
			return err
		}
	default:
		if err = proto.Unmarshal(message, in); err != nil {
			return err
		}
	}

	out := make(chan *SyncEvent, 100)

	h.Server.Sync(c, in, out)

	defer ws.Close()

	for msg := range out {

		w, err := ws.NextWriter(websocket.BinaryMessage)
		if err != nil {

			close(out)
			c.Logger().Error(err)
			return nil
		}

		var response []byte

		switch c.Request().Header.Get("Content-Type") {
		case "application/hrpc-json":
			response, err = protojson.Marshal(msg)
		default:
			response, err = proto.Marshal(msg)
		}

		if err != nil {

			close(out)
			c.Logger().Error(err)
			return nil
		}

		if _, err := w.Write(response); err != nil {

			close(out)
			c.Logger().Error(err)
			return nil
		}
		if err := w.Close(); err != nil {

			close(out)
			c.Logger().Error(err)
			return nil
		}
	}

	return nil

}

func (h *ChatServiceHandler) GetUserHandler(c echo.Context) error {

	requestProto := new(GetUserRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.GetUser(c, req.(*GetUserRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(ChatServiceServerGetUserData, ChatServiceData, Chatᐳv1ᐳchat, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}

func (h *ChatServiceHandler) GetUserBulkHandler(c echo.Context) error {

	requestProto := new(GetUserBulkRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.GetUserBulk(c, req.(*GetUserBulkRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(ChatServiceServerGetUserBulkData, ChatServiceData, Chatᐳv1ᐳchat, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}

func (h *ChatServiceHandler) GetUserMetadataHandler(c echo.Context) error {

	requestProto := new(GetUserMetadataRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.GetUserMetadata(c, req.(*GetUserMetadataRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(ChatServiceServerGetUserMetadataData, ChatServiceData, Chatᐳv1ᐳchat, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}

func (h *ChatServiceHandler) ProfileUpdateHandler(c echo.Context) error {

	requestProto := new(ProfileUpdateRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.ProfileUpdate(c, req.(*ProfileUpdateRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(ChatServiceServerProfileUpdateData, ChatServiceData, Chatᐳv1ᐳchat, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}

func (h *ChatServiceHandler) TypingHandler(c echo.Context) error {

	requestProto := new(TypingRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.Typing(c, req.(*TypingRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(ChatServiceServerTypingData, ChatServiceData, Chatᐳv1ᐳchat, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}

func (h *ChatServiceHandler) PreviewGuildHandler(c echo.Context) error {

	requestProto := new(PreviewGuildRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.PreviewGuild(c, req.(*PreviewGuildRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(ChatServiceServerPreviewGuildData, ChatServiceData, Chatᐳv1ᐳchat, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}

func (h *ChatServiceHandler) BanUserHandler(c echo.Context) error {

	requestProto := new(BanUserRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.BanUser(c, req.(*BanUserRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(ChatServiceServerBanUserData, ChatServiceData, Chatᐳv1ᐳchat, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}

func (h *ChatServiceHandler) KickUserHandler(c echo.Context) error {

	requestProto := new(KickUserRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.KickUser(c, req.(*KickUserRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(ChatServiceServerKickUserData, ChatServiceData, Chatᐳv1ᐳchat, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}

func (h *ChatServiceHandler) UnbanUserHandler(c echo.Context) error {

	requestProto := new(UnbanUserRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.UnbanUser(c, req.(*UnbanUserRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(ChatServiceServerUnbanUserData, ChatServiceData, Chatᐳv1ᐳchat, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}
