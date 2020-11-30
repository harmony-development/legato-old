// Code generated by sqlc. DO NOT EDIT.

package queries

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.acquireEmotePackStmt, err = db.PrepareContext(ctx, acquireEmotePack); err != nil {
		return nil, fmt.Errorf("error preparing query AcquireEmotePack: %w", err)
	}
	if q.addEmoteToPackStmt, err = db.PrepareContext(ctx, addEmoteToPack); err != nil {
		return nil, fmt.Errorf("error preparing query AddEmoteToPack: %w", err)
	}
	if q.addFileHashStmt, err = db.PrepareContext(ctx, addFileHash); err != nil {
		return nil, fmt.Errorf("error preparing query AddFileHash: %w", err)
	}
	if q.addForeignUserStmt, err = db.PrepareContext(ctx, addForeignUser); err != nil {
		return nil, fmt.Errorf("error preparing query AddForeignUser: %w", err)
	}
	if q.addLocalUserStmt, err = db.PrepareContext(ctx, addLocalUser); err != nil {
		return nil, fmt.Errorf("error preparing query AddLocalUser: %w", err)
	}
	if q.addMessageStmt, err = db.PrepareContext(ctx, addMessage); err != nil {
		return nil, fmt.Errorf("error preparing query AddMessage: %w", err)
	}
	if q.addNonceStmt, err = db.PrepareContext(ctx, addNonce); err != nil {
		return nil, fmt.Errorf("error preparing query AddNonce: %w", err)
	}
	if q.addProfileStmt, err = db.PrepareContext(ctx, addProfile); err != nil {
		return nil, fmt.Errorf("error preparing query AddProfile: %w", err)
	}
	if q.addSessionStmt, err = db.PrepareContext(ctx, addSession); err != nil {
		return nil, fmt.Errorf("error preparing query AddSession: %w", err)
	}
	if q.addToGuildListStmt, err = db.PrepareContext(ctx, addToGuildList); err != nil {
		return nil, fmt.Errorf("error preparing query AddToGuildList: %w", err)
	}
	if q.addUserStmt, err = db.PrepareContext(ctx, addUser); err != nil {
		return nil, fmt.Errorf("error preparing query AddUser: %w", err)
	}
	if q.addUserToGuildStmt, err = db.PrepareContext(ctx, addUserToGuild); err != nil {
		return nil, fmt.Errorf("error preparing query AddUserToGuild: %w", err)
	}
	if q.addUserToRoleStmt, err = db.PrepareContext(ctx, addUserToRole); err != nil {
		return nil, fmt.Errorf("error preparing query AddUserToRole: %w", err)
	}
	if q.createChannelStmt, err = db.PrepareContext(ctx, createChannel); err != nil {
		return nil, fmt.Errorf("error preparing query CreateChannel: %w", err)
	}
	if q.createEmotePackStmt, err = db.PrepareContext(ctx, createEmotePack); err != nil {
		return nil, fmt.Errorf("error preparing query CreateEmotePack: %w", err)
	}
	if q.createGuildStmt, err = db.PrepareContext(ctx, createGuild); err != nil {
		return nil, fmt.Errorf("error preparing query CreateGuild: %w", err)
	}
	if q.createGuildInviteStmt, err = db.PrepareContext(ctx, createGuildInvite); err != nil {
		return nil, fmt.Errorf("error preparing query CreateGuildInvite: %w", err)
	}
	if q.createRoleStmt, err = db.PrepareContext(ctx, createRole); err != nil {
		return nil, fmt.Errorf("error preparing query CreateRole: %w", err)
	}
	if q.deleteChannelStmt, err = db.PrepareContext(ctx, deleteChannel); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteChannel: %w", err)
	}
	if q.deleteEmoteFromPackStmt, err = db.PrepareContext(ctx, deleteEmoteFromPack); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteEmoteFromPack: %w", err)
	}
	if q.deleteEmotePackStmt, err = db.PrepareContext(ctx, deleteEmotePack); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteEmotePack: %w", err)
	}
	if q.deleteGuildStmt, err = db.PrepareContext(ctx, deleteGuild); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteGuild: %w", err)
	}
	if q.deleteInviteStmt, err = db.PrepareContext(ctx, deleteInvite); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteInvite: %w", err)
	}
	if q.deleteMessageStmt, err = db.PrepareContext(ctx, deleteMessage); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteMessage: %w", err)
	}
	if q.deleteRoleStmt, err = db.PrepareContext(ctx, deleteRole); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteRole: %w", err)
	}
	if q.dequipEmotePackStmt, err = db.PrepareContext(ctx, dequipEmotePack); err != nil {
		return nil, fmt.Errorf("error preparing query DequipEmotePack: %w", err)
	}
	if q.emailExistsStmt, err = db.PrepareContext(ctx, emailExists); err != nil {
		return nil, fmt.Errorf("error preparing query EmailExists: %w", err)
	}
	if q.expireSessionsStmt, err = db.PrepareContext(ctx, expireSessions); err != nil {
		return nil, fmt.Errorf("error preparing query ExpireSessions: %w", err)
	}
	if q.getAvatarStmt, err = db.PrepareContext(ctx, getAvatar); err != nil {
		return nil, fmt.Errorf("error preparing query GetAvatar: %w", err)
	}
	if q.getChannelPositionStmt, err = db.PrepareContext(ctx, getChannelPosition); err != nil {
		return nil, fmt.Errorf("error preparing query GetChannelPosition: %w", err)
	}
	if q.getChannelsStmt, err = db.PrepareContext(ctx, getChannels); err != nil {
		return nil, fmt.Errorf("error preparing query GetChannels: %w", err)
	}
	if q.getEmotePackEmotesStmt, err = db.PrepareContext(ctx, getEmotePackEmotes); err != nil {
		return nil, fmt.Errorf("error preparing query GetEmotePackEmotes: %w", err)
	}
	if q.getEmotePacksStmt, err = db.PrepareContext(ctx, getEmotePacks); err != nil {
		return nil, fmt.Errorf("error preparing query GetEmotePacks: %w", err)
	}
	if q.getFileByHashStmt, err = db.PrepareContext(ctx, getFileByHash); err != nil {
		return nil, fmt.Errorf("error preparing query GetFileByHash: %w", err)
	}
	if q.getGuildDataStmt, err = db.PrepareContext(ctx, getGuildData); err != nil {
		return nil, fmt.Errorf("error preparing query GetGuildData: %w", err)
	}
	if q.getGuildListStmt, err = db.PrepareContext(ctx, getGuildList); err != nil {
		return nil, fmt.Errorf("error preparing query GetGuildList: %w", err)
	}
	if q.getGuildListPositionStmt, err = db.PrepareContext(ctx, getGuildListPosition); err != nil {
		return nil, fmt.Errorf("error preparing query GetGuildListPosition: %w", err)
	}
	if q.getGuildMembersStmt, err = db.PrepareContext(ctx, getGuildMembers); err != nil {
		return nil, fmt.Errorf("error preparing query GetGuildMembers: %w", err)
	}
	if q.getGuildOwnerStmt, err = db.PrepareContext(ctx, getGuildOwner); err != nil {
		return nil, fmt.Errorf("error preparing query GetGuildOwner: %w", err)
	}
	if q.getGuildPictureStmt, err = db.PrepareContext(ctx, getGuildPicture); err != nil {
		return nil, fmt.Errorf("error preparing query GetGuildPicture: %w", err)
	}
	if q.getLastGuildPositionInListStmt, err = db.PrepareContext(ctx, getLastGuildPositionInList); err != nil {
		return nil, fmt.Errorf("error preparing query GetLastGuildPositionInList: %w", err)
	}
	if q.getLocalUserIDStmt, err = db.PrepareContext(ctx, getLocalUserID); err != nil {
		return nil, fmt.Errorf("error preparing query GetLocalUserID: %w", err)
	}
	if q.getMessageStmt, err = db.PrepareContext(ctx, getMessage); err != nil {
		return nil, fmt.Errorf("error preparing query GetMessage: %w", err)
	}
	if q.getMessageAuthorStmt, err = db.PrepareContext(ctx, getMessageAuthor); err != nil {
		return nil, fmt.Errorf("error preparing query GetMessageAuthor: %w", err)
	}
	if q.getMessageDateStmt, err = db.PrepareContext(ctx, getMessageDate); err != nil {
		return nil, fmt.Errorf("error preparing query GetMessageDate: %w", err)
	}
	if q.getMessagesStmt, err = db.PrepareContext(ctx, getMessages); err != nil {
		return nil, fmt.Errorf("error preparing query GetMessages: %w", err)
	}
	if q.getNonceInfoStmt, err = db.PrepareContext(ctx, getNonceInfo); err != nil {
		return nil, fmt.Errorf("error preparing query GetNonceInfo: %w", err)
	}
	if q.getPackOwnerStmt, err = db.PrepareContext(ctx, getPackOwner); err != nil {
		return nil, fmt.Errorf("error preparing query GetPackOwner: %w", err)
	}
	if q.getPermissionsStmt, err = db.PrepareContext(ctx, getPermissions); err != nil {
		return nil, fmt.Errorf("error preparing query GetPermissions: %w", err)
	}
	if q.getPermissionsWithoutChannelStmt, err = db.PrepareContext(ctx, getPermissionsWithoutChannel); err != nil {
		return nil, fmt.Errorf("error preparing query GetPermissionsWithoutChannel: %w", err)
	}
	if q.getRolePositionStmt, err = db.PrepareContext(ctx, getRolePosition); err != nil {
		return nil, fmt.Errorf("error preparing query GetRolePosition: %w", err)
	}
	if q.getRolesForGuildStmt, err = db.PrepareContext(ctx, getRolesForGuild); err != nil {
		return nil, fmt.Errorf("error preparing query GetRolesForGuild: %w", err)
	}
	if q.getUserStmt, err = db.PrepareContext(ctx, getUser); err != nil {
		return nil, fmt.Errorf("error preparing query GetUser: %w", err)
	}
	if q.getUserByEmailStmt, err = db.PrepareContext(ctx, getUserByEmail); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserByEmail: %w", err)
	}
	if q.getUserMetadataStmt, err = db.PrepareContext(ctx, getUserMetadata); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserMetadata: %w", err)
	}
	if q.guildWithIDExistsStmt, err = db.PrepareContext(ctx, guildWithIDExists); err != nil {
		return nil, fmt.Errorf("error preparing query GuildWithIDExists: %w", err)
	}
	if q.guildsForUserStmt, err = db.PrepareContext(ctx, guildsForUser); err != nil {
		return nil, fmt.Errorf("error preparing query GuildsForUser: %w", err)
	}
	if q.guildsForUserWithDataStmt, err = db.PrepareContext(ctx, guildsForUserWithData); err != nil {
		return nil, fmt.Errorf("error preparing query GuildsForUserWithData: %w", err)
	}
	if q.incrementInviteStmt, err = db.PrepareContext(ctx, incrementInvite); err != nil {
		return nil, fmt.Errorf("error preparing query IncrementInvite: %w", err)
	}
	if q.isIPWhitelistedStmt, err = db.PrepareContext(ctx, isIPWhitelisted); err != nil {
		return nil, fmt.Errorf("error preparing query IsIPWhitelisted: %w", err)
	}
	if q.isUserWhitelistedStmt, err = db.PrepareContext(ctx, isUserWhitelisted); err != nil {
		return nil, fmt.Errorf("error preparing query IsUserWhitelisted: %w", err)
	}
	if q.messageWithIDExistsStmt, err = db.PrepareContext(ctx, messageWithIDExists); err != nil {
		return nil, fmt.Errorf("error preparing query MessageWithIDExists: %w", err)
	}
	if q.moveChannelStmt, err = db.PrepareContext(ctx, moveChannel); err != nil {
		return nil, fmt.Errorf("error preparing query MoveChannel: %w", err)
	}
	if q.moveGuildStmt, err = db.PrepareContext(ctx, moveGuild); err != nil {
		return nil, fmt.Errorf("error preparing query MoveGuild: %w", err)
	}
	if q.moveRoleStmt, err = db.PrepareContext(ctx, moveRole); err != nil {
		return nil, fmt.Errorf("error preparing query MoveRole: %w", err)
	}
	if q.numChannelsWithIDStmt, err = db.PrepareContext(ctx, numChannelsWithID); err != nil {
		return nil, fmt.Errorf("error preparing query NumChannelsWithID: %w", err)
	}
	if q.openInvitesStmt, err = db.PrepareContext(ctx, openInvites); err != nil {
		return nil, fmt.Errorf("error preparing query OpenInvites: %w", err)
	}
	if q.removeGuildFromListStmt, err = db.PrepareContext(ctx, removeGuildFromList); err != nil {
		return nil, fmt.Errorf("error preparing query RemoveGuildFromList: %w", err)
	}
	if q.removeUserFromGuildStmt, err = db.PrepareContext(ctx, removeUserFromGuild); err != nil {
		return nil, fmt.Errorf("error preparing query RemoveUserFromGuild: %w", err)
	}
	if q.removeUserFromRoleStmt, err = db.PrepareContext(ctx, removeUserFromRole); err != nil {
		return nil, fmt.Errorf("error preparing query RemoveUserFromRole: %w", err)
	}
	if q.resolveGuildIDStmt, err = db.PrepareContext(ctx, resolveGuildID); err != nil {
		return nil, fmt.Errorf("error preparing query ResolveGuildID: %w", err)
	}
	if q.rolesForUserStmt, err = db.PrepareContext(ctx, rolesForUser); err != nil {
		return nil, fmt.Errorf("error preparing query RolesForUser: %w", err)
	}
	if q.sessionToUserIDStmt, err = db.PrepareContext(ctx, sessionToUserID); err != nil {
		return nil, fmt.Errorf("error preparing query SessionToUserID: %w", err)
	}
	if q.setGuildNameStmt, err = db.PrepareContext(ctx, setGuildName); err != nil {
		return nil, fmt.Errorf("error preparing query SetGuildName: %w", err)
	}
	if q.setGuildPictureStmt, err = db.PrepareContext(ctx, setGuildPicture); err != nil {
		return nil, fmt.Errorf("error preparing query SetGuildPicture: %w", err)
	}
	if q.setPermissionsStmt, err = db.PrepareContext(ctx, setPermissions); err != nil {
		return nil, fmt.Errorf("error preparing query SetPermissions: %w", err)
	}
	if q.setStatusStmt, err = db.PrepareContext(ctx, setStatus); err != nil {
		return nil, fmt.Errorf("error preparing query SetStatus: %w", err)
	}
	if q.updateAvatarStmt, err = db.PrepareContext(ctx, updateAvatar); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateAvatar: %w", err)
	}
	if q.updateChannelNameStmt, err = db.PrepareContext(ctx, updateChannelName); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateChannelName: %w", err)
	}
	if q.updateMessageActionsStmt, err = db.PrepareContext(ctx, updateMessageActions); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateMessageActions: %w", err)
	}
	if q.updateMessageContentStmt, err = db.PrepareContext(ctx, updateMessageContent); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateMessageContent: %w", err)
	}
	if q.updateMessageEmbedsStmt, err = db.PrepareContext(ctx, updateMessageEmbeds); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateMessageEmbeds: %w", err)
	}
	if q.updateMessageOverridesStmt, err = db.PrepareContext(ctx, updateMessageOverrides); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateMessageOverrides: %w", err)
	}
	if q.updateUsernameStmt, err = db.PrepareContext(ctx, updateUsername); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateUsername: %w", err)
	}
	if q.userInGuildStmt, err = db.PrepareContext(ctx, userInGuild); err != nil {
		return nil, fmt.Errorf("error preparing query UserInGuild: %w", err)
	}
	if q.userIsLocalStmt, err = db.PrepareContext(ctx, userIsLocal); err != nil {
		return nil, fmt.Errorf("error preparing query UserIsLocal: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.acquireEmotePackStmt != nil {
		if cerr := q.acquireEmotePackStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing acquireEmotePackStmt: %w", cerr)
		}
	}
	if q.addEmoteToPackStmt != nil {
		if cerr := q.addEmoteToPackStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addEmoteToPackStmt: %w", cerr)
		}
	}
	if q.addFileHashStmt != nil {
		if cerr := q.addFileHashStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addFileHashStmt: %w", cerr)
		}
	}
	if q.addForeignUserStmt != nil {
		if cerr := q.addForeignUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addForeignUserStmt: %w", cerr)
		}
	}
	if q.addLocalUserStmt != nil {
		if cerr := q.addLocalUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addLocalUserStmt: %w", cerr)
		}
	}
	if q.addMessageStmt != nil {
		if cerr := q.addMessageStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addMessageStmt: %w", cerr)
		}
	}
	if q.addNonceStmt != nil {
		if cerr := q.addNonceStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addNonceStmt: %w", cerr)
		}
	}
	if q.addProfileStmt != nil {
		if cerr := q.addProfileStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addProfileStmt: %w", cerr)
		}
	}
	if q.addSessionStmt != nil {
		if cerr := q.addSessionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addSessionStmt: %w", cerr)
		}
	}
	if q.addToGuildListStmt != nil {
		if cerr := q.addToGuildListStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addToGuildListStmt: %w", cerr)
		}
	}
	if q.addUserStmt != nil {
		if cerr := q.addUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addUserStmt: %w", cerr)
		}
	}
	if q.addUserToGuildStmt != nil {
		if cerr := q.addUserToGuildStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addUserToGuildStmt: %w", cerr)
		}
	}
	if q.addUserToRoleStmt != nil {
		if cerr := q.addUserToRoleStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addUserToRoleStmt: %w", cerr)
		}
	}
	if q.createChannelStmt != nil {
		if cerr := q.createChannelStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createChannelStmt: %w", cerr)
		}
	}
	if q.createEmotePackStmt != nil {
		if cerr := q.createEmotePackStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createEmotePackStmt: %w", cerr)
		}
	}
	if q.createGuildStmt != nil {
		if cerr := q.createGuildStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createGuildStmt: %w", cerr)
		}
	}
	if q.createGuildInviteStmt != nil {
		if cerr := q.createGuildInviteStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createGuildInviteStmt: %w", cerr)
		}
	}
	if q.createRoleStmt != nil {
		if cerr := q.createRoleStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createRoleStmt: %w", cerr)
		}
	}
	if q.deleteChannelStmt != nil {
		if cerr := q.deleteChannelStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteChannelStmt: %w", cerr)
		}
	}
	if q.deleteEmoteFromPackStmt != nil {
		if cerr := q.deleteEmoteFromPackStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteEmoteFromPackStmt: %w", cerr)
		}
	}
	if q.deleteEmotePackStmt != nil {
		if cerr := q.deleteEmotePackStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteEmotePackStmt: %w", cerr)
		}
	}
	if q.deleteGuildStmt != nil {
		if cerr := q.deleteGuildStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteGuildStmt: %w", cerr)
		}
	}
	if q.deleteInviteStmt != nil {
		if cerr := q.deleteInviteStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteInviteStmt: %w", cerr)
		}
	}
	if q.deleteMessageStmt != nil {
		if cerr := q.deleteMessageStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteMessageStmt: %w", cerr)
		}
	}
	if q.deleteRoleStmt != nil {
		if cerr := q.deleteRoleStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteRoleStmt: %w", cerr)
		}
	}
	if q.dequipEmotePackStmt != nil {
		if cerr := q.dequipEmotePackStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing dequipEmotePackStmt: %w", cerr)
		}
	}
	if q.emailExistsStmt != nil {
		if cerr := q.emailExistsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing emailExistsStmt: %w", cerr)
		}
	}
	if q.expireSessionsStmt != nil {
		if cerr := q.expireSessionsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing expireSessionsStmt: %w", cerr)
		}
	}
	if q.getAvatarStmt != nil {
		if cerr := q.getAvatarStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAvatarStmt: %w", cerr)
		}
	}
	if q.getChannelPositionStmt != nil {
		if cerr := q.getChannelPositionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getChannelPositionStmt: %w", cerr)
		}
	}
	if q.getChannelsStmt != nil {
		if cerr := q.getChannelsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getChannelsStmt: %w", cerr)
		}
	}
	if q.getEmotePackEmotesStmt != nil {
		if cerr := q.getEmotePackEmotesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getEmotePackEmotesStmt: %w", cerr)
		}
	}
	if q.getEmotePacksStmt != nil {
		if cerr := q.getEmotePacksStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getEmotePacksStmt: %w", cerr)
		}
	}
	if q.getFileByHashStmt != nil {
		if cerr := q.getFileByHashStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getFileByHashStmt: %w", cerr)
		}
	}
	if q.getGuildDataStmt != nil {
		if cerr := q.getGuildDataStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getGuildDataStmt: %w", cerr)
		}
	}
	if q.getGuildListStmt != nil {
		if cerr := q.getGuildListStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getGuildListStmt: %w", cerr)
		}
	}
	if q.getGuildListPositionStmt != nil {
		if cerr := q.getGuildListPositionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getGuildListPositionStmt: %w", cerr)
		}
	}
	if q.getGuildMembersStmt != nil {
		if cerr := q.getGuildMembersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getGuildMembersStmt: %w", cerr)
		}
	}
	if q.getGuildOwnerStmt != nil {
		if cerr := q.getGuildOwnerStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getGuildOwnerStmt: %w", cerr)
		}
	}
	if q.getGuildPictureStmt != nil {
		if cerr := q.getGuildPictureStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getGuildPictureStmt: %w", cerr)
		}
	}
	if q.getLastGuildPositionInListStmt != nil {
		if cerr := q.getLastGuildPositionInListStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getLastGuildPositionInListStmt: %w", cerr)
		}
	}
	if q.getLocalUserIDStmt != nil {
		if cerr := q.getLocalUserIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getLocalUserIDStmt: %w", cerr)
		}
	}
	if q.getMessageStmt != nil {
		if cerr := q.getMessageStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getMessageStmt: %w", cerr)
		}
	}
	if q.getMessageAuthorStmt != nil {
		if cerr := q.getMessageAuthorStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getMessageAuthorStmt: %w", cerr)
		}
	}
	if q.getMessageDateStmt != nil {
		if cerr := q.getMessageDateStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getMessageDateStmt: %w", cerr)
		}
	}
	if q.getMessagesStmt != nil {
		if cerr := q.getMessagesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getMessagesStmt: %w", cerr)
		}
	}
	if q.getNonceInfoStmt != nil {
		if cerr := q.getNonceInfoStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getNonceInfoStmt: %w", cerr)
		}
	}
	if q.getPackOwnerStmt != nil {
		if cerr := q.getPackOwnerStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPackOwnerStmt: %w", cerr)
		}
	}
	if q.getPermissionsStmt != nil {
		if cerr := q.getPermissionsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPermissionsStmt: %w", cerr)
		}
	}
	if q.getPermissionsWithoutChannelStmt != nil {
		if cerr := q.getPermissionsWithoutChannelStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPermissionsWithoutChannelStmt: %w", cerr)
		}
	}
	if q.getRolePositionStmt != nil {
		if cerr := q.getRolePositionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getRolePositionStmt: %w", cerr)
		}
	}
	if q.getRolesForGuildStmt != nil {
		if cerr := q.getRolesForGuildStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getRolesForGuildStmt: %w", cerr)
		}
	}
	if q.getUserStmt != nil {
		if cerr := q.getUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserStmt: %w", cerr)
		}
	}
	if q.getUserByEmailStmt != nil {
		if cerr := q.getUserByEmailStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserByEmailStmt: %w", cerr)
		}
	}
	if q.getUserMetadataStmt != nil {
		if cerr := q.getUserMetadataStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserMetadataStmt: %w", cerr)
		}
	}
	if q.guildWithIDExistsStmt != nil {
		if cerr := q.guildWithIDExistsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing guildWithIDExistsStmt: %w", cerr)
		}
	}
	if q.guildsForUserStmt != nil {
		if cerr := q.guildsForUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing guildsForUserStmt: %w", cerr)
		}
	}
	if q.guildsForUserWithDataStmt != nil {
		if cerr := q.guildsForUserWithDataStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing guildsForUserWithDataStmt: %w", cerr)
		}
	}
	if q.incrementInviteStmt != nil {
		if cerr := q.incrementInviteStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing incrementInviteStmt: %w", cerr)
		}
	}
	if q.isIPWhitelistedStmt != nil {
		if cerr := q.isIPWhitelistedStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing isIPWhitelistedStmt: %w", cerr)
		}
	}
	if q.isUserWhitelistedStmt != nil {
		if cerr := q.isUserWhitelistedStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing isUserWhitelistedStmt: %w", cerr)
		}
	}
	if q.messageWithIDExistsStmt != nil {
		if cerr := q.messageWithIDExistsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing messageWithIDExistsStmt: %w", cerr)
		}
	}
	if q.moveChannelStmt != nil {
		if cerr := q.moveChannelStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing moveChannelStmt: %w", cerr)
		}
	}
	if q.moveGuildStmt != nil {
		if cerr := q.moveGuildStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing moveGuildStmt: %w", cerr)
		}
	}
	if q.moveRoleStmt != nil {
		if cerr := q.moveRoleStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing moveRoleStmt: %w", cerr)
		}
	}
	if q.numChannelsWithIDStmt != nil {
		if cerr := q.numChannelsWithIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing numChannelsWithIDStmt: %w", cerr)
		}
	}
	if q.openInvitesStmt != nil {
		if cerr := q.openInvitesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing openInvitesStmt: %w", cerr)
		}
	}
	if q.removeGuildFromListStmt != nil {
		if cerr := q.removeGuildFromListStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing removeGuildFromListStmt: %w", cerr)
		}
	}
	if q.removeUserFromGuildStmt != nil {
		if cerr := q.removeUserFromGuildStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing removeUserFromGuildStmt: %w", cerr)
		}
	}
	if q.removeUserFromRoleStmt != nil {
		if cerr := q.removeUserFromRoleStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing removeUserFromRoleStmt: %w", cerr)
		}
	}
	if q.resolveGuildIDStmt != nil {
		if cerr := q.resolveGuildIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing resolveGuildIDStmt: %w", cerr)
		}
	}
	if q.rolesForUserStmt != nil {
		if cerr := q.rolesForUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing rolesForUserStmt: %w", cerr)
		}
	}
	if q.sessionToUserIDStmt != nil {
		if cerr := q.sessionToUserIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing sessionToUserIDStmt: %w", cerr)
		}
	}
	if q.setGuildNameStmt != nil {
		if cerr := q.setGuildNameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing setGuildNameStmt: %w", cerr)
		}
	}
	if q.setGuildPictureStmt != nil {
		if cerr := q.setGuildPictureStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing setGuildPictureStmt: %w", cerr)
		}
	}
	if q.setPermissionsStmt != nil {
		if cerr := q.setPermissionsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing setPermissionsStmt: %w", cerr)
		}
	}
	if q.setStatusStmt != nil {
		if cerr := q.setStatusStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing setStatusStmt: %w", cerr)
		}
	}
	if q.updateAvatarStmt != nil {
		if cerr := q.updateAvatarStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateAvatarStmt: %w", cerr)
		}
	}
	if q.updateChannelNameStmt != nil {
		if cerr := q.updateChannelNameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateChannelNameStmt: %w", cerr)
		}
	}
	if q.updateMessageActionsStmt != nil {
		if cerr := q.updateMessageActionsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateMessageActionsStmt: %w", cerr)
		}
	}
	if q.updateMessageContentStmt != nil {
		if cerr := q.updateMessageContentStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateMessageContentStmt: %w", cerr)
		}
	}
	if q.updateMessageEmbedsStmt != nil {
		if cerr := q.updateMessageEmbedsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateMessageEmbedsStmt: %w", cerr)
		}
	}
	if q.updateMessageOverridesStmt != nil {
		if cerr := q.updateMessageOverridesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateMessageOverridesStmt: %w", cerr)
		}
	}
	if q.updateUsernameStmt != nil {
		if cerr := q.updateUsernameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateUsernameStmt: %w", cerr)
		}
	}
	if q.userInGuildStmt != nil {
		if cerr := q.userInGuildStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing userInGuildStmt: %w", cerr)
		}
	}
	if q.userIsLocalStmt != nil {
		if cerr := q.userIsLocalStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing userIsLocalStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                               DBTX
	tx                               *sql.Tx
	acquireEmotePackStmt             *sql.Stmt
	addEmoteToPackStmt               *sql.Stmt
	addFileHashStmt                  *sql.Stmt
	addForeignUserStmt               *sql.Stmt
	addLocalUserStmt                 *sql.Stmt
	addMessageStmt                   *sql.Stmt
	addNonceStmt                     *sql.Stmt
	addProfileStmt                   *sql.Stmt
	addSessionStmt                   *sql.Stmt
	addToGuildListStmt               *sql.Stmt
	addUserStmt                      *sql.Stmt
	addUserToGuildStmt               *sql.Stmt
	addUserToRoleStmt                *sql.Stmt
	createChannelStmt                *sql.Stmt
	createEmotePackStmt              *sql.Stmt
	createGuildStmt                  *sql.Stmt
	createGuildInviteStmt            *sql.Stmt
	createRoleStmt                   *sql.Stmt
	deleteChannelStmt                *sql.Stmt
	deleteEmoteFromPackStmt          *sql.Stmt
	deleteEmotePackStmt              *sql.Stmt
	deleteGuildStmt                  *sql.Stmt
	deleteInviteStmt                 *sql.Stmt
	deleteMessageStmt                *sql.Stmt
	deleteRoleStmt                   *sql.Stmt
	dequipEmotePackStmt              *sql.Stmt
	emailExistsStmt                  *sql.Stmt
	expireSessionsStmt               *sql.Stmt
	getAvatarStmt                    *sql.Stmt
	getChannelPositionStmt           *sql.Stmt
	getChannelsStmt                  *sql.Stmt
	getEmotePackEmotesStmt           *sql.Stmt
	getEmotePacksStmt                *sql.Stmt
	getFileByHashStmt                *sql.Stmt
	getGuildDataStmt                 *sql.Stmt
	getGuildListStmt                 *sql.Stmt
	getGuildListPositionStmt         *sql.Stmt
	getGuildMembersStmt              *sql.Stmt
	getGuildOwnerStmt                *sql.Stmt
	getGuildPictureStmt              *sql.Stmt
	getLastGuildPositionInListStmt   *sql.Stmt
	getLocalUserIDStmt               *sql.Stmt
	getMessageStmt                   *sql.Stmt
	getMessageAuthorStmt             *sql.Stmt
	getMessageDateStmt               *sql.Stmt
	getMessagesStmt                  *sql.Stmt
	getNonceInfoStmt                 *sql.Stmt
	getPackOwnerStmt                 *sql.Stmt
	getPermissionsStmt               *sql.Stmt
	getPermissionsWithoutChannelStmt *sql.Stmt
	getRolePositionStmt              *sql.Stmt
	getRolesForGuildStmt             *sql.Stmt
	getUserStmt                      *sql.Stmt
	getUserByEmailStmt               *sql.Stmt
	getUserMetadataStmt              *sql.Stmt
	guildWithIDExistsStmt            *sql.Stmt
	guildsForUserStmt                *sql.Stmt
	guildsForUserWithDataStmt        *sql.Stmt
	incrementInviteStmt              *sql.Stmt
	isIPWhitelistedStmt              *sql.Stmt
	isUserWhitelistedStmt            *sql.Stmt
	messageWithIDExistsStmt          *sql.Stmt
	moveChannelStmt                  *sql.Stmt
	moveGuildStmt                    *sql.Stmt
	moveRoleStmt                     *sql.Stmt
	numChannelsWithIDStmt            *sql.Stmt
	openInvitesStmt                  *sql.Stmt
	removeGuildFromListStmt          *sql.Stmt
	removeUserFromGuildStmt          *sql.Stmt
	removeUserFromRoleStmt           *sql.Stmt
	resolveGuildIDStmt               *sql.Stmt
	rolesForUserStmt                 *sql.Stmt
	sessionToUserIDStmt              *sql.Stmt
	setGuildNameStmt                 *sql.Stmt
	setGuildPictureStmt              *sql.Stmt
	setPermissionsStmt               *sql.Stmt
	setStatusStmt                    *sql.Stmt
	updateAvatarStmt                 *sql.Stmt
	updateChannelNameStmt            *sql.Stmt
	updateMessageActionsStmt         *sql.Stmt
	updateMessageContentStmt         *sql.Stmt
	updateMessageEmbedsStmt          *sql.Stmt
	updateMessageOverridesStmt       *sql.Stmt
	updateUsernameStmt               *sql.Stmt
	userInGuildStmt                  *sql.Stmt
	userIsLocalStmt                  *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                               tx,
		tx:                               tx,
		acquireEmotePackStmt:             q.acquireEmotePackStmt,
		addEmoteToPackStmt:               q.addEmoteToPackStmt,
		addFileHashStmt:                  q.addFileHashStmt,
		addForeignUserStmt:               q.addForeignUserStmt,
		addLocalUserStmt:                 q.addLocalUserStmt,
		addMessageStmt:                   q.addMessageStmt,
		addNonceStmt:                     q.addNonceStmt,
		addProfileStmt:                   q.addProfileStmt,
		addSessionStmt:                   q.addSessionStmt,
		addToGuildListStmt:               q.addToGuildListStmt,
		addUserStmt:                      q.addUserStmt,
		addUserToGuildStmt:               q.addUserToGuildStmt,
		addUserToRoleStmt:                q.addUserToRoleStmt,
		createChannelStmt:                q.createChannelStmt,
		createEmotePackStmt:              q.createEmotePackStmt,
		createGuildStmt:                  q.createGuildStmt,
		createGuildInviteStmt:            q.createGuildInviteStmt,
		createRoleStmt:                   q.createRoleStmt,
		deleteChannelStmt:                q.deleteChannelStmt,
		deleteEmoteFromPackStmt:          q.deleteEmoteFromPackStmt,
		deleteEmotePackStmt:              q.deleteEmotePackStmt,
		deleteGuildStmt:                  q.deleteGuildStmt,
		deleteInviteStmt:                 q.deleteInviteStmt,
		deleteMessageStmt:                q.deleteMessageStmt,
		deleteRoleStmt:                   q.deleteRoleStmt,
		dequipEmotePackStmt:              q.dequipEmotePackStmt,
		emailExistsStmt:                  q.emailExistsStmt,
		expireSessionsStmt:               q.expireSessionsStmt,
		getAvatarStmt:                    q.getAvatarStmt,
		getChannelPositionStmt:           q.getChannelPositionStmt,
		getChannelsStmt:                  q.getChannelsStmt,
		getEmotePackEmotesStmt:           q.getEmotePackEmotesStmt,
		getEmotePacksStmt:                q.getEmotePacksStmt,
		getFileByHashStmt:                q.getFileByHashStmt,
		getGuildDataStmt:                 q.getGuildDataStmt,
		getGuildListStmt:                 q.getGuildListStmt,
		getGuildListPositionStmt:         q.getGuildListPositionStmt,
		getGuildMembersStmt:              q.getGuildMembersStmt,
		getGuildOwnerStmt:                q.getGuildOwnerStmt,
		getGuildPictureStmt:              q.getGuildPictureStmt,
		getLastGuildPositionInListStmt:   q.getLastGuildPositionInListStmt,
		getLocalUserIDStmt:               q.getLocalUserIDStmt,
		getMessageStmt:                   q.getMessageStmt,
		getMessageAuthorStmt:             q.getMessageAuthorStmt,
		getMessageDateStmt:               q.getMessageDateStmt,
		getMessagesStmt:                  q.getMessagesStmt,
		getNonceInfoStmt:                 q.getNonceInfoStmt,
		getPackOwnerStmt:                 q.getPackOwnerStmt,
		getPermissionsStmt:               q.getPermissionsStmt,
		getPermissionsWithoutChannelStmt: q.getPermissionsWithoutChannelStmt,
		getRolePositionStmt:              q.getRolePositionStmt,
		getRolesForGuildStmt:             q.getRolesForGuildStmt,
		getUserStmt:                      q.getUserStmt,
		getUserByEmailStmt:               q.getUserByEmailStmt,
		getUserMetadataStmt:              q.getUserMetadataStmt,
		guildWithIDExistsStmt:            q.guildWithIDExistsStmt,
		guildsForUserStmt:                q.guildsForUserStmt,
		guildsForUserWithDataStmt:        q.guildsForUserWithDataStmt,
		incrementInviteStmt:              q.incrementInviteStmt,
		isIPWhitelistedStmt:              q.isIPWhitelistedStmt,
		isUserWhitelistedStmt:            q.isUserWhitelistedStmt,
		messageWithIDExistsStmt:          q.messageWithIDExistsStmt,
		moveChannelStmt:                  q.moveChannelStmt,
		moveGuildStmt:                    q.moveGuildStmt,
		moveRoleStmt:                     q.moveRoleStmt,
		numChannelsWithIDStmt:            q.numChannelsWithIDStmt,
		openInvitesStmt:                  q.openInvitesStmt,
		removeGuildFromListStmt:          q.removeGuildFromListStmt,
		removeUserFromGuildStmt:          q.removeUserFromGuildStmt,
		removeUserFromRoleStmt:           q.removeUserFromRoleStmt,
		resolveGuildIDStmt:               q.resolveGuildIDStmt,
		rolesForUserStmt:                 q.rolesForUserStmt,
		sessionToUserIDStmt:              q.sessionToUserIDStmt,
		setGuildNameStmt:                 q.setGuildNameStmt,
		setGuildPictureStmt:              q.setGuildPictureStmt,
		setPermissionsStmt:               q.setPermissionsStmt,
		setStatusStmt:                    q.setStatusStmt,
		updateAvatarStmt:                 q.updateAvatarStmt,
		updateChannelNameStmt:            q.updateChannelNameStmt,
		updateMessageActionsStmt:         q.updateMessageActionsStmt,
		updateMessageContentStmt:         q.updateMessageContentStmt,
		updateMessageEmbedsStmt:          q.updateMessageEmbedsStmt,
		updateMessageOverridesStmt:       q.updateMessageOverridesStmt,
		updateUsernameStmt:               q.updateUsernameStmt,
		userInGuildStmt:                  q.userInGuildStmt,
		userIsLocalStmt:                  q.userIsLocalStmt,
	}
}
