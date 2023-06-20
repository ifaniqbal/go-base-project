package api

import (
	"errors"
	"github.com/ifaniqbal/go-base-project/internal/utils"
	mocks2 "github.com/ifaniqbal/go-base-project/test/mocks/internal_/api"
	mocks "github.com/ifaniqbal/go-base-project/test/mocks/internal_/utils"
	"testing"
)

func TestApi_Start(t *testing.T) {
	type fields struct {
		httpServer utils.HttpServer
		routers    []Router
	}
	type testCase struct {
		name    string
		fields  fields
		wantErr bool
	}
	var tests []testCase
	var name string
	var f fields

	name = "httpServer.Run error"
	mockHttpServer := mocks.NewHttpServer(t)
	mockRouter := mocks2.NewRouter(t)
	f = fields{
		httpServer: mockHttpServer,
		routers:    []Router{mockRouter},
	}
	mockHttpServer.EXPECT().Group("/").Return(nil).Once()
	mockRouter.EXPECT().Route(nil).Once()
	mockHttpServer.EXPECT().Run().Return(errors.New("e")).Once()
	tests = append(tests, testCase{
		name:    name,
		fields:  f,
		wantErr: true,
	})

	name = "success"
	mockHttpServer.EXPECT().Group("/").Return(nil).Once()
	mockRouter.EXPECT().Route(nil).Once()
	mockHttpServer.EXPECT().Run().Return(nil).Once()
	tests = append(tests, testCase{
		name:    name,
		fields:  f,
		wantErr: false,
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := Api{
				httpServer: tt.fields.httpServer,
				routers:    tt.fields.routers,
			}
			if err := a.Start(); (err != nil) != tt.wantErr {
				t.Errorf("Start() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
