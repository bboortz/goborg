package appcontext

import (
	"context"

	"go.uber.org/zap"
)

type correlationIdType int

const (
	requestIdKey correlationIdType = iota
	sessionIdKey
	pkgNameKey
)

var ctx context.Context
var logger *zap.Logger

func init() {
	// create a new context
	ctx = NewCtx()

	// a fallback/root logger for events without context
	logger, _ = zap.NewDevelopment()
	defer logger.Sync() // flushes buffer, if any
	logger.Debug("logging init function finalized")
}

func Ctx() context.Context {
	return ctx
}

func NewCtx() context.Context {
	ctx := context.Background()
	return WithPkgName(ctx, "main")
}

// WithPkg returns a context with pkg name
func WithPkgName(ctx context.Context, pkgName string) context.Context {
	return context.WithValue(ctx, pkgNameKey, pkgName)
}

// WithRqId returns a context which knows its request ID
func WithRqId(ctx context.Context, requestId string) context.Context {
	return context.WithValue(ctx, requestIdKey, requestId)
}

// WithSessionId returns a context which knows its session ID
func WithSessionId(ctx context.Context, sessionId string) context.Context {
	return context.WithValue(ctx, sessionIdKey, sessionId)
}

// Logger returns a zap logger with as much context as possible
func Logger(ctx context.Context) *zap.Logger {
	newLogger := logger
	if ctx != nil {
		if ctxRqId, ok := ctx.Value(requestIdKey).(string); ok {
			newLogger = newLogger.With(zap.String("rqId", ctxRqId))
		}
		if ctxSessionId, ok := ctx.Value(sessionIdKey).(string); ok {
			newLogger = newLogger.With(zap.String("sessionId", ctxSessionId))
		}
		if ctxPkgName, ok := ctx.Value(pkgNameKey).(string); ok {
			newLogger = newLogger.With(zap.String("pkgName", ctxPkgName))
		}
	}
	return newLogger
}
