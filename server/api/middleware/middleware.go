package middleware

import "context"

// HarmonyContext contains a custom context for passing data from middleware to handlers
type HarmonyContext struct {
	context.Context
	UserID uint64
}
