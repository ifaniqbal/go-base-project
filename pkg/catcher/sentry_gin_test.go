package catcher

import (
	"errors"
	"github.com/ifaniqbal/go-base-project/test/mocks/pkg/environment"
	mocks3 "github.com/ifaniqbal/go-base-project/test/mocks/pkg/httpserver"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/ifaniqbal/go-base-project/pkg/environment"
	"github.com/ifaniqbal/go-base-project/pkg/httpserver"
	"github.com/stretchr/testify/assert"
)

func TestSentryGinCatcher_Init(t *testing.T) {
	mockEnv := mocks.NewEnvironment(t)
	mockEnv.EXPECT().Get("SENTRY_ENVIRONMENT").Return("production").Once()
	mockEnv.EXPECT().Get("SENTRY_DSN").Return("").Once()

	s := NewSentryGin(mockEnv, nil)
	err := s.Init()
	assert.Nil(t, err)
	assert.Equal(t, sentry.CurrentHub(), s.hub)
}

func TestSentryGinCatcher_WithContext(t *testing.T) {
	type fields struct {
		env environment.Environment
	}
	type args struct {
		ctx httpserver.Context
	}
	type testCase struct {
		name   string
		fields fields
		args   args
		want   Catcher
	}
	var tests []testCase
	mockEnv := mocks.NewEnvironment(t)
	var f = fields{
		env: mockEnv,
	}
	mockCtx := mocks3.NewContext(t)
	var a = args{ctx: mockCtx}

	var name = "cast failed"
	tests = append(tests, testCase{
		name:   name,
		fields: f,
		args:   a,
		want: &SentryGin{
			env: mockEnv,
		},
	})

	name = "success"
	w := httptest.NewRecorder()
	testCtx, _ := gin.CreateTestContext(w)
	req, _ := http.NewRequest("GET", "/", nil)
	testCtx.Request = req
	var ctx httpserver.Context
	ctx = testCtx
	ginCtx, _ := ctx.(*gin.Context)
	options := sentrygin.Options{
		Repanic: true,
	}
	sentrygin.New(options)(ginCtx)
	tests = append(tests, testCase{
		name:   name,
		fields: f,
		args:   args{ctx: ctx},
		want:   NewSentryGin(mockEnv, sentrygin.GetHubFromContext(ginCtx)),
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SentryGin{
				env: tt.fields.env,
			}
			assert.Equalf(t, tt.want, s.WithContext(tt.args.ctx), "WithContext(%v)", tt.args.ctx)
		})
	}
}

func TestSentryGinCatcher_Catch(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		defer func() {
			r := recover()
			assert.NotNil(t, r)
		}()
		s := Default()
		s.Catch(errors.New("e"))
	})
}

func TestSentryGinCatcher_CaptureMessage(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		defer func() {
			r := recover()
			assert.NotNil(t, r)
		}()
		s := Default()
		s.CaptureMessage("m")
	})
}
