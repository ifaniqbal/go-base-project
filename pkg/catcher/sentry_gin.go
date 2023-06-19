package catcher

import (
	"github.com/getsentry/sentry-go"
	"github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/ifaniqbal/go-base-project/pkg/environment"
	"github.com/ifaniqbal/go-base-project/pkg/httpserver"
)

// SentryGin is an implementation of a Sentry error catcher.Catcher
// for the Gin framework.
type SentryGin struct {
	env environment.Environment
	hub Hub
}

type Hub interface {
	CaptureMessage(message string) *sentry.EventID
	CaptureException(exception error) *sentry.EventID
}

// NewSentryGin creates a new SentryGin instance.
func NewSentryGin(
	env environment.Environment,
	hub *sentry.Hub,
) *SentryGin {
	return &SentryGin{
		env: env,
		hub: hub,
	}
}

// Default creates a default SentryGin instance.
func Default() *SentryGin {
	return NewSentryGin(environment.Default(), nil)
}

// Init initializes the Sentry client with the provided environment and hub.
// The trace sample rate is set to 1.0 for non-production environments, and 0.2
// for production environments. The Sentry client is initialized with the DSN
// and environment from the environment variable, and the current hub.
// The function returns an error if the Sentry client fails to initialize.
func (s *SentryGin) Init() error {
	traceSampleRate := 1.0
	sentryEnvironment := s.env.Get("SENTRY_ENVIRONMENT")
	if sentryEnvironment == "production" {
		traceSampleRate = 0.2
	}
	s.hub = sentry.CurrentHub()
	return sentry.Init(sentry.ClientOptions{
		Dsn:              s.env.Get("SENTRY_DSN"),
		Environment:      sentryEnvironment,
		TracesSampleRate: traceSampleRate,
	})
}

// WithContext adds Sentry context to the Gin context by using
// sentrygin.New function. It returns a new instance of SentryGin
// with the updated hub extracted from the Gin context.
// If the passed context is not of type *gin.Context,
// the original instance of SentryGin is returned.
func (s *SentryGin) WithContext(ctx httpserver.Context) Catcher {
	ginCtx, ok := ctx.(*gin.Context)
	if !ok {
		return s
	}
	options := sentrygin.Options{
		Repanic: true,
	}
	sentrygin.New(options)(ginCtx)
	return NewSentryGin(s.env, sentrygin.GetHubFromContext(ginCtx))
}

// Catch captures an error in Sentry.
func (s *SentryGin) Catch(err error) {
	s.hub.CaptureException(err)
}

// CaptureMessage captures a message in Sentry.
func (s *SentryGin) CaptureMessage(msg string) {
	s.hub.CaptureMessage(msg)
}
