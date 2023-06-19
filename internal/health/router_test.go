package health

import (
	"github.com/ifaniqbal/go-base-project/test/mocks/pkg/httpserver"
	"testing"

	"github.com/ifaniqbal/go-base-project/pkg/httpserver"
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
