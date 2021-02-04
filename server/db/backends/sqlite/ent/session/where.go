// Code generated by entc, DO NOT EDIT.

package session

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/harmony-development/legato/server/db/backends/sqlite/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// SessionID applies equality check predicate on the "sessionID" field. It's identical to SessionIDEQ.
func SessionID(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSessionID), v))
	})
}

// Expires applies equality check predicate on the "expires" field. It's identical to ExpiresEQ.
func Expires(v time.Time) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExpires), v))
	})
}

// SessionIDEQ applies the EQ predicate on the "sessionID" field.
func SessionIDEQ(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSessionID), v))
	})
}

// SessionIDNEQ applies the NEQ predicate on the "sessionID" field.
func SessionIDNEQ(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSessionID), v))
	})
}

// SessionIDIn applies the In predicate on the "sessionID" field.
func SessionIDIn(vs ...string) predicate.Session {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Session(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSessionID), v...))
	})
}

// SessionIDNotIn applies the NotIn predicate on the "sessionID" field.
func SessionIDNotIn(vs ...string) predicate.Session {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Session(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSessionID), v...))
	})
}

// SessionIDGT applies the GT predicate on the "sessionID" field.
func SessionIDGT(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSessionID), v))
	})
}

// SessionIDGTE applies the GTE predicate on the "sessionID" field.
func SessionIDGTE(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSessionID), v))
	})
}

// SessionIDLT applies the LT predicate on the "sessionID" field.
func SessionIDLT(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSessionID), v))
	})
}

// SessionIDLTE applies the LTE predicate on the "sessionID" field.
func SessionIDLTE(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSessionID), v))
	})
}

// SessionIDContains applies the Contains predicate on the "sessionID" field.
func SessionIDContains(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSessionID), v))
	})
}

// SessionIDHasPrefix applies the HasPrefix predicate on the "sessionID" field.
func SessionIDHasPrefix(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSessionID), v))
	})
}

// SessionIDHasSuffix applies the HasSuffix predicate on the "sessionID" field.
func SessionIDHasSuffix(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSessionID), v))
	})
}

// SessionIDEqualFold applies the EqualFold predicate on the "sessionID" field.
func SessionIDEqualFold(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSessionID), v))
	})
}

// SessionIDContainsFold applies the ContainsFold predicate on the "sessionID" field.
func SessionIDContainsFold(v string) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSessionID), v))
	})
}

// ExpiresEQ applies the EQ predicate on the "expires" field.
func ExpiresEQ(v time.Time) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExpires), v))
	})
}

// ExpiresNEQ applies the NEQ predicate on the "expires" field.
func ExpiresNEQ(v time.Time) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExpires), v))
	})
}

// ExpiresIn applies the In predicate on the "expires" field.
func ExpiresIn(vs ...time.Time) predicate.Session {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Session(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldExpires), v...))
	})
}

// ExpiresNotIn applies the NotIn predicate on the "expires" field.
func ExpiresNotIn(vs ...time.Time) predicate.Session {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Session(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldExpires), v...))
	})
}

// ExpiresGT applies the GT predicate on the "expires" field.
func ExpiresGT(v time.Time) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldExpires), v))
	})
}

// ExpiresGTE applies the GTE predicate on the "expires" field.
func ExpiresGTE(v time.Time) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldExpires), v))
	})
}

// ExpiresLT applies the LT predicate on the "expires" field.
func ExpiresLT(v time.Time) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldExpires), v))
	})
}

// ExpiresLTE applies the LTE predicate on the "expires" field.
func ExpiresLTE(v time.Time) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldExpires), v))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.LocalUser) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Session) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Session) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Session) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		p(s.Not())
	})
}
