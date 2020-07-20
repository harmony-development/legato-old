package middleware

import (
	"context"
	"time"

	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

type visitor struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

func (m Middlewares) RateCleanup() {
	for {
		time.Sleep(3 * time.Minute)
		m.RateLock.Lock()
		for _, path := range m.RateLimits {
			for ip, v := range path {
				if time.Since(v.lastSeen) > 3*time.Minute {
					delete(path, ip)
				}
			}
		}
		m.RateLock.Unlock()
	}
}

func (m Middlewares) RateLimitInterceptor(c context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	ctx := c.(HarmonyContext)
	p, _ := peer.FromContext(c)
	l, exists := rpcConfigs[info.FullMethod]
	if exists {
		ctx.Limiter = m.GetVisitor(info.FullMethod, p.Addr.String(), l.RateLimit.Duration, l.RateLimit.Burst)
	}

	return handler(c, req)
}

func (m *Middlewares) GetVisitor(path, ip string, duration time.Duration, burst int) *rate.Limiter {
	m.RateLock.Lock()
	defer m.RateLock.Unlock()
	if _, exists := m.RateLimits[path]; !exists {
		m.RateLimits[path] = make(map[string]visitor)
	}
	v, exists := m.RateLimits[path][ip]
	if !exists {
		limiter := rate.NewLimiter(rate.Every(duration), burst)
		m.RateLimits[path][ip] = visitor{
			limiter:  limiter,
			lastSeen: time.Now(),
		}
		return limiter
	}
	return v.limiter
}

/* // RateLimit enforces a rate limit on clients
func (m *Middlewares) RateLimit(duration time.Duration, burst int) {
	return func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.(HarmonyContext)
			ctx.Limiter = m.GetVisitor(ctx.Path(), ctx.RealIP(), duration, burst)
			return handler(ctx)
		}
	}
}

// GetVisitor gets the rate limiter for a visitor
func (m *Middlewares) GetVisitor(path string, ip string, duration time.Duration, burst int) *rate.Limiter {
	m.RateLock.Lock()
	defer m.RateLock.Unlock()
	if _, exists := m.RateLimits[path]; !exists {
		m.RateLimits[path] = make(map[string]*visitor)
	}
	if v, exists := m.RateLimits[path][ip]; !exists {
		limiter := rate.NewLimiter(rate.Every(duration), burst)
		m.RateLimits[path][ip] = &visitor{
			limiter:  limiter,
			lastSeen: time.Now(),
		}
		return limiter
	} else {
		return v.limiter
	}
} */
