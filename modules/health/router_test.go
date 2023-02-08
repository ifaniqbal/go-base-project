package health

import (
	"testing"

	mocks "bitbucket.org/ifan-moladin/base-project/mocks/utils/httpserver"
	"bitbucket.org/ifan-moladin/base-project/utils/httpserver"
	"github.com/stretchr/testify/mock"
)

func TestRouter_Route(t *testing.T) {
	type args struct {
		route httpserver.RouteHandler
	}
	mockRoute := mocks.NewRouteHandler(t)
	mockRoute.EXPECT().GET("/", mock.AnythingOfType("httpserver.HandlerFunc")).Return(mockRoute).Once()
	tests := []struct {
		name string
		args args
	}{
		{
			name: "route",
			args: args{route: mockRoute},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRouter(nil, nil)
			r.Route(tt.args.route)
		})
	}
}
