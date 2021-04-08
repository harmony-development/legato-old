// Code generated by entc, DO NOT EDIT.

package guild

const (
	// Label holds the string label denoting the guild type in the database.
	Label = "guild"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldOwner holds the string denoting the owner field in the database.
	FieldOwner = "owner"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldPicture holds the string denoting the picture field in the database.
	FieldPicture = "picture"
	// FieldMetadata holds the string denoting the metadata field in the database.
	FieldMetadata = "metadata"
	// EdgeInvite holds the string denoting the invite edge name in mutations.
	EdgeInvite = "invite"
	// EdgeBans holds the string denoting the bans edge name in mutations.
	EdgeBans = "bans"
	// EdgeChannel holds the string denoting the channel edge name in mutations.
	EdgeChannel = "channel"
	// EdgeRole holds the string denoting the role edge name in mutations.
	EdgeRole = "role"
	// EdgePermissionNode holds the string denoting the permission_node edge name in mutations.
	EdgePermissionNode = "permission_node"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the guild in the database.
	Table = "guilds"
	// InviteTable is the table the holds the invite relation/edge.
	InviteTable = "invites"
	// InviteInverseTable is the table name for the Invite entity.
	// It exists in this package in order to avoid circular dependency with the "invite" package.
	InviteInverseTable = "invites"
	// InviteColumn is the table column denoting the invite relation/edge.
	InviteColumn = "guild_invite"
	// BansTable is the table the holds the bans relation/edge.
	BansTable = "users"
	// BansInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	BansInverseTable = "users"
	// BansColumn is the table column denoting the bans relation/edge.
	BansColumn = "guild_bans"
	// ChannelTable is the table the holds the channel relation/edge.
	ChannelTable = "channels"
	// ChannelInverseTable is the table name for the Channel entity.
	// It exists in this package in order to avoid circular dependency with the "channel" package.
	ChannelInverseTable = "channels"
	// ChannelColumn is the table column denoting the channel relation/edge.
	ChannelColumn = "guild_channel"
	// RoleTable is the table the holds the role relation/edge. The primary key declared below.
	RoleTable = "guild_role"
	// RoleInverseTable is the table name for the Role entity.
	// It exists in this package in order to avoid circular dependency with the "role" package.
	RoleInverseTable = "roles"
	// PermissionNodeTable is the table the holds the permission_node relation/edge.
	PermissionNodeTable = "permission_nodes"
	// PermissionNodeInverseTable is the table name for the PermissionNode entity.
	// It exists in this package in order to avoid circular dependency with the "permissionnode" package.
	PermissionNodeInverseTable = "permission_nodes"
	// PermissionNodeColumn is the table column denoting the permission_node relation/edge.
	PermissionNodeColumn = "guild_permission_node"
	// UserTable is the table the holds the user relation/edge. The primary key declared below.
	UserTable = "user_guild"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
)

// Columns holds all SQL columns for guild fields.
var Columns = []string{
	FieldID,
	FieldOwner,
	FieldName,
	FieldPicture,
	FieldMetadata,
}

var (
	// RolePrimaryKey and RoleColumn2 are the table columns denoting the
	// primary key for the role relation (M2M).
	RolePrimaryKey = []string{"guild_id", "role_id"}
	// UserPrimaryKey and UserColumn2 are the table columns denoting the
	// primary key for the user relation (M2M).
	UserPrimaryKey = []string{"user_id", "guild_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
