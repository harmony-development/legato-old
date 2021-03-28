// Code generated by entc, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/harmony-development/legato/server/db/ent/entgen"
)

// The ChannelFunc type is an adapter to allow the use of ordinary
// function as Channel mutator.
type ChannelFunc func(context.Context, *entgen.ChannelMutation) (entgen.Value, error)

// Mutate calls f(ctx, m).
func (f ChannelFunc) Mutate(ctx context.Context, m entgen.Mutation) (entgen.Value, error) {
	mv, ok := m.(*entgen.ChannelMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *entgen.ChannelMutation", m)
	}
	return f(ctx, mv)
}

// The ForeignUserFunc type is an adapter to allow the use of ordinary
// function as ForeignUser mutator.
type ForeignUserFunc func(context.Context, *entgen.ForeignUserMutation) (entgen.Value, error)

// Mutate calls f(ctx, m).
func (f ForeignUserFunc) Mutate(ctx context.Context, m entgen.Mutation) (entgen.Value, error) {
	mv, ok := m.(*entgen.ForeignUserMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *entgen.ForeignUserMutation", m)
	}
	return f(ctx, mv)
}

// The GuildFunc type is an adapter to allow the use of ordinary
// function as Guild mutator.
type GuildFunc func(context.Context, *entgen.GuildMutation) (entgen.Value, error)

// Mutate calls f(ctx, m).
func (f GuildFunc) Mutate(ctx context.Context, m entgen.Mutation) (entgen.Value, error) {
	mv, ok := m.(*entgen.GuildMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *entgen.GuildMutation", m)
	}
	return f(ctx, mv)
}

// The InviteFunc type is an adapter to allow the use of ordinary
// function as Invite mutator.
type InviteFunc func(context.Context, *entgen.InviteMutation) (entgen.Value, error)

// Mutate calls f(ctx, m).
func (f InviteFunc) Mutate(ctx context.Context, m entgen.Mutation) (entgen.Value, error) {
	mv, ok := m.(*entgen.InviteMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *entgen.InviteMutation", m)
	}
	return f(ctx, mv)
}

// The LocalUserFunc type is an adapter to allow the use of ordinary
// function as LocalUser mutator.
type LocalUserFunc func(context.Context, *entgen.LocalUserMutation) (entgen.Value, error)

// Mutate calls f(ctx, m).
func (f LocalUserFunc) Mutate(ctx context.Context, m entgen.Mutation) (entgen.Value, error) {
	mv, ok := m.(*entgen.LocalUserMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *entgen.LocalUserMutation", m)
	}
	return f(ctx, mv)
}

// The MessageFunc type is an adapter to allow the use of ordinary
// function as Message mutator.
type MessageFunc func(context.Context, *entgen.MessageMutation) (entgen.Value, error)

// Mutate calls f(ctx, m).
func (f MessageFunc) Mutate(ctx context.Context, m entgen.Mutation) (entgen.Value, error) {
	mv, ok := m.(*entgen.MessageMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *entgen.MessageMutation", m)
	}
	return f(ctx, mv)
}

// The OverrideFunc type is an adapter to allow the use of ordinary
// function as Override mutator.
type OverrideFunc func(context.Context, *entgen.OverrideMutation) (entgen.Value, error)

// Mutate calls f(ctx, m).
func (f OverrideFunc) Mutate(ctx context.Context, m entgen.Mutation) (entgen.Value, error) {
	mv, ok := m.(*entgen.OverrideMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *entgen.OverrideMutation", m)
	}
	return f(ctx, mv)
}

// The ProfileFunc type is an adapter to allow the use of ordinary
// function as Profile mutator.
type ProfileFunc func(context.Context, *entgen.ProfileMutation) (entgen.Value, error)

// Mutate calls f(ctx, m).
func (f ProfileFunc) Mutate(ctx context.Context, m entgen.Mutation) (entgen.Value, error) {
	mv, ok := m.(*entgen.ProfileMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *entgen.ProfileMutation", m)
	}
	return f(ctx, mv)
}

// The SessionFunc type is an adapter to allow the use of ordinary
// function as Session mutator.
type SessionFunc func(context.Context, *entgen.SessionMutation) (entgen.Value, error)

// Mutate calls f(ctx, m).
func (f SessionFunc) Mutate(ctx context.Context, m entgen.Mutation) (entgen.Value, error) {
	mv, ok := m.(*entgen.SessionMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *entgen.SessionMutation", m)
	}
	return f(ctx, mv)
}

// The UserFunc type is an adapter to allow the use of ordinary
// function as User mutator.
type UserFunc func(context.Context, *entgen.UserMutation) (entgen.Value, error)

// Mutate calls f(ctx, m).
func (f UserFunc) Mutate(ctx context.Context, m entgen.Mutation) (entgen.Value, error) {
	mv, ok := m.(*entgen.UserMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *entgen.UserMutation", m)
	}
	return f(ctx, mv)
}

// Condition is a hook condition function.
type Condition func(context.Context, entgen.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m entgen.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m entgen.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m entgen.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op entgen.Op) Condition {
	return func(_ context.Context, m entgen.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m entgen.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m entgen.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m entgen.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
//
func If(hk entgen.Hook, cond Condition) entgen.Hook {
	return func(next entgen.Mutator) entgen.Mutator {
		return entgen.MutateFunc(func(ctx context.Context, m entgen.Mutation) (entgen.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, entgen.Delete|entgen.Create)
//
func On(hk entgen.Hook, op entgen.Op) entgen.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, entgen.Update|entgen.UpdateOne)
//
func Unless(hk entgen.Hook, op entgen.Op) entgen.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) entgen.Hook {
	return func(entgen.Mutator) entgen.Mutator {
		return entgen.MutateFunc(func(context.Context, entgen.Mutation) (entgen.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []entgen.Hook {
//		return []entgen.Hook{
//			Reject(entgen.Delete|entgen.Update),
//		}
//	}
//
func Reject(op entgen.Op) entgen.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []entgen.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...entgen.Hook) Chain {
	return Chain{append([]entgen.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() entgen.Hook {
	return func(mutator entgen.Mutator) entgen.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...entgen.Hook) Chain {
	newHooks := make([]entgen.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
