// Code generated by entc, DO NOT EDIT.

package entgen

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/harmony-development/legato/server/db/ent/entgen/role"
)

// Role is the model entity for the Role schema.
type Role struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Color holds the value of the "color" field.
	Color int `json:"color,omitempty"`
	// Hoist holds the value of the "hoist" field.
	Hoist bool `json:"hoist,omitempty"`
	// Pingable holds the value of the "pingable" field.
	Pingable bool `json:"pingable,omitempty"`
	// Position holds the value of the "position" field.
	Position string `json:"position,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RoleQuery when eager-loading is set.
	Edges      RoleEdges `json:"edges"`
	guild_role *uint64
}

// RoleEdges holds the relations/edges for other nodes in the graph.
type RoleEdges struct {
	// Members holds the value of the members edge.
	Members []*User `json:"members,omitempty"`
	// Permission holds the value of the permission edge.
	Permission []*Permission `json:"permission,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// MembersOrErr returns the Members value or an error if the edge
// was not loaded in eager-loading.
func (e RoleEdges) MembersOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.Members, nil
	}
	return nil, &NotLoadedError{edge: "members"}
}

// PermissionOrErr returns the Permission value or an error if the edge
// was not loaded in eager-loading.
func (e RoleEdges) PermissionOrErr() ([]*Permission, error) {
	if e.loadedTypes[1] {
		return e.Permission, nil
	}
	return nil, &NotLoadedError{edge: "permission"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Role) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case role.FieldHoist, role.FieldPingable:
			values[i] = &sql.NullBool{}
		case role.FieldID, role.FieldColor:
			values[i] = &sql.NullInt64{}
		case role.FieldName, role.FieldPosition:
			values[i] = &sql.NullString{}
		case role.ForeignKeys[0]: // guild_role
			values[i] = &sql.NullInt64{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Role", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Role fields.
func (r *Role) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case role.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = uint64(value.Int64)
		case role.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				r.Name = value.String
			}
		case role.FieldColor:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field color", values[i])
			} else if value.Valid {
				r.Color = int(value.Int64)
			}
		case role.FieldHoist:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field hoist", values[i])
			} else if value.Valid {
				r.Hoist = value.Bool
			}
		case role.FieldPingable:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field pingable", values[i])
			} else if value.Valid {
				r.Pingable = value.Bool
			}
		case role.FieldPosition:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field position", values[i])
			} else if value.Valid {
				r.Position = value.String
			}
		case role.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field guild_role", value)
			} else if value.Valid {
				r.guild_role = new(uint64)
				*r.guild_role = uint64(value.Int64)
			}
		}
	}
	return nil
}

// QueryMembers queries the "members" edge of the Role entity.
func (r *Role) QueryMembers() *UserQuery {
	return (&RoleClient{config: r.config}).QueryMembers(r)
}

// QueryPermission queries the "permission" edge of the Role entity.
func (r *Role) QueryPermission() *PermissionQuery {
	return (&RoleClient{config: r.config}).QueryPermission(r)
}

// Update returns a builder for updating this Role.
// Note that you need to call Role.Unwrap() before calling this method if this Role
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Role) Update() *RoleUpdateOne {
	return (&RoleClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Role entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Role) Unwrap() *Role {
	tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("entgen: Role is not a transactional entity")
	}
	r.config.driver = tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Role) String() string {
	var builder strings.Builder
	builder.WriteString("Role(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteString(", name=")
	builder.WriteString(r.Name)
	builder.WriteString(", color=")
	builder.WriteString(fmt.Sprintf("%v", r.Color))
	builder.WriteString(", hoist=")
	builder.WriteString(fmt.Sprintf("%v", r.Hoist))
	builder.WriteString(", pingable=")
	builder.WriteString(fmt.Sprintf("%v", r.Pingable))
	builder.WriteString(", position=")
	builder.WriteString(r.Position)
	builder.WriteByte(')')
	return builder.String()
}

// Roles is a parsable slice of Role.
type Roles []*Role

func (r Roles) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}