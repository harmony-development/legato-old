package middleware

import (
	"net/http"
	"time"

	"github.com/harmony-development/hrpc/server"
	"github.com/harmony-development/legato/server/http/responses"
	"github.com/labstack/echo/v4"
	"golang.org/x/time/rate"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
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

func (m Middlewares) RateLimitInterceptor(meth *descriptorpb.MethodDescriptorProto, serv *descriptorpb.ServiceDescriptorProto, d *descriptorpb.FileDescriptorProto, h server.Handler) server.Handler {
	return func(c echo.Context, req protoreflect.ProtoMessage) (protoreflect.ProtoMessage, error) {
		ctx := c.(HarmonyContext)
		l, exists := rpcConfigs[meth.GetName()]
		if exists {
			if !m.GetVisitor(meth.GetName(), ctx.RealIP(), l.RateLimit.Duration, l.RateLimit.Burst).Allow() {
				return nil, echo.NewHTTPError(http.StatusTooManyRequests, responses.TooManyRequests)
			}
		}
		return h(ctx, req)
	}
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
