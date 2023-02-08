package api

import (
	"errors"
	"testing"

	mocks2 "bitbucket.org/ifan-moladin/base-project/mocks/modules/api"
	mocks "bitbucket.org/ifan-moladin/base-project/mocks/utils/httpserver"
	"bitbucket.org/ifan-moladin/base-project/utils/httpserver"
)

func TestApi_Start(t *testing.T) {
	type fields struct {
		httpServer httpserver.HttpServer
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
