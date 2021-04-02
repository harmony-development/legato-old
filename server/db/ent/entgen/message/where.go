// Code generated by entc, DO NOT EDIT.

package message

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/harmony-development/legato/server/db/ent/entgen/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
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
func IDNotIn(ids ...uint64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
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
func IDGT(id uint64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Createdat applies equality check predicate on the "createdat" field. It's identical to CreatedatEQ.
func Createdat(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedat), v))
	})
}

// Editedat applies equality check predicate on the "editedat" field. It's identical to EditedatEQ.
func Editedat(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEditedat), v))
	})
}

// Overrides applies equality check predicate on the "overrides" field. It's identical to OverridesEQ.
func Overrides(v []byte) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOverrides), v))
	})
}

// CreatedatEQ applies the EQ predicate on the "createdat" field.
func CreatedatEQ(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedat), v))
	})
}

// CreatedatNEQ applies the NEQ predicate on the "createdat" field.
func CreatedatNEQ(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedat), v))
	})
}

// CreatedatIn applies the In predicate on the "createdat" field.
func CreatedatIn(vs ...time.Time) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedat), v...))
	})
}

// CreatedatNotIn applies the NotIn predicate on the "createdat" field.
func CreatedatNotIn(vs ...time.Time) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedat), v...))
	})
}

// CreatedatGT applies the GT predicate on the "createdat" field.
func CreatedatGT(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedat), v))
	})
}

// CreatedatGTE applies the GTE predicate on the "createdat" field.
func CreatedatGTE(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedat), v))
	})
}

// CreatedatLT applies the LT predicate on the "createdat" field.
func CreatedatLT(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedat), v))
	})
}

// CreatedatLTE applies the LTE predicate on the "createdat" field.
func CreatedatLTE(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedat), v))
	})
}

// EditedatEQ applies the EQ predicate on the "editedat" field.
func EditedatEQ(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEditedat), v))
	})
}

// EditedatNEQ applies the NEQ predicate on the "editedat" field.
func EditedatNEQ(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEditedat), v))
	})
}

// EditedatIn applies the In predicate on the "editedat" field.
func EditedatIn(vs ...time.Time) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEditedat), v...))
	})
}

// EditedatNotIn applies the NotIn predicate on the "editedat" field.
func EditedatNotIn(vs ...time.Time) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEditedat), v...))
	})
}

// EditedatGT applies the GT predicate on the "editedat" field.
func EditedatGT(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEditedat), v))
	})
}

// EditedatGTE applies the GTE predicate on the "editedat" field.
func EditedatGTE(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEditedat), v))
	})
}

// EditedatLT applies the LT predicate on the "editedat" field.
func EditedatLT(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEditedat), v))
	})
}

// EditedatLTE applies the LTE predicate on the "editedat" field.
func EditedatLTE(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEditedat), v))
	})
}

// EditedatIsNil applies the IsNil predicate on the "editedat" field.
func EditedatIsNil() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldEditedat)))
	})
}

// EditedatNotNil applies the NotNil predicate on the "editedat" field.
func EditedatNotNil() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldEditedat)))
	})
}

// ActionsIsNil applies the IsNil predicate on the "actions" field.
func ActionsIsNil() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldActions)))
	})
}

// ActionsNotNil applies the NotNil predicate on the "actions" field.
func ActionsNotNil() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldActions)))
	})
}

// MetadataIsNil applies the IsNil predicate on the "metadata" field.
func MetadataIsNil() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldMetadata)))
	})
}

// MetadataNotNil applies the NotNil predicate on the "metadata" field.
func MetadataNotNil() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldMetadata)))
	})
}

// OverridesEQ applies the EQ predicate on the "overrides" field.
func OverridesEQ(v []byte) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOverrides), v))
	})
}

// OverridesNEQ applies the NEQ predicate on the "overrides" field.
func OverridesNEQ(v []byte) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOverrides), v))
	})
}

// OverridesIn applies the In predicate on the "overrides" field.
func OverridesIn(vs ...[]byte) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOverrides), v...))
	})
}

// OverridesNotIn applies the NotIn predicate on the "overrides" field.
func OverridesNotIn(vs ...[]byte) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOverrides), v...))
	})
}

// OverridesGT applies the GT predicate on the "overrides" field.
func OverridesGT(v []byte) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOverrides), v))
	})
}

// OverridesGTE applies the GTE predicate on the "overrides" field.
func OverridesGTE(v []byte) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOverrides), v))
	})
}

// OverridesLT applies the LT predicate on the "overrides" field.
func OverridesLT(v []byte) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOverrides), v))
	})
}

// OverridesLTE applies the LTE predicate on the "overrides" field.
func OverridesLTE(v []byte) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOverrides), v))
	})
}

// OverridesIsNil applies the IsNil predicate on the "overrides" field.
func OverridesIsNil() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldOverrides)))
	})
}

// OverridesNotNil applies the NotNil predicate on the "overrides" field.
func OverridesNotNil() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldOverrides)))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
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

// HasChannel applies the HasEdge predicate on the "channel" edge.
func HasChannel() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ChannelTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ChannelTable, ChannelColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChannelWith applies the HasEdge predicate on the "channel" edge with a given conditions (other predicates).
func HasChannelWith(preds ...predicate.Channel) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ChannelInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ChannelTable, ChannelColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOverride applies the HasEdge predicate on the "override" edge.
func HasOverride() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OverrideTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, OverrideTable, OverrideColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOverrideWith applies the HasEdge predicate on the "override" edge with a given conditions (other predicates).
func HasOverrideWith(preds ...predicate.Override) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OverrideInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, OverrideTable, OverrideColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.Message) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasReplies applies the HasEdge predicate on the "replies" edge.
func HasReplies() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RepliesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RepliesTable, RepliesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRepliesWith applies the HasEdge predicate on the "replies" edge with a given conditions (other predicates).
func HasRepliesWith(preds ...predicate.Message) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RepliesTable, RepliesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTextmessage applies the HasEdge predicate on the "textmessage" edge.
func HasTextmessage() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TextmessageTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, TextmessageTable, TextmessageColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTextmessageWith applies the HasEdge predicate on the "textmessage" edge with a given conditions (other predicates).
func HasTextmessageWith(preds ...predicate.TextMessage) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TextmessageInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, TextmessageTable, TextmessageColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFilemessage applies the HasEdge predicate on the "filemessage" edge.
func HasFilemessage() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FilemessageTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, FilemessageTable, FilemessageColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFilemessageWith applies the HasEdge predicate on the "filemessage" edge with a given conditions (other predicates).
func HasFilemessageWith(preds ...predicate.FileMessage) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FilemessageInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, FilemessageTable, FilemessageColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEmbedmessage applies the HasEdge predicate on the "embedmessage" edge.
func HasEmbedmessage() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EmbedmessageTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, EmbedmessageTable, EmbedmessageColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEmbedmessageWith applies the HasEdge predicate on the "embedmessage" edge with a given conditions (other predicates).
func HasEmbedmessageWith(preds ...predicate.EmbedMessage) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EmbedmessageInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, EmbedmessageTable, EmbedmessageColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Message) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Message) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
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
func Not(p predicate.Message) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		p(s.Not())
	})
}
