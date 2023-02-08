package authenticator

import (
	"net/http"
	"reflect"
	"testing"

	mocks "bitbucket.org/ifan-moladin/base-project/mocks/utils/httpserver"
	"bitbucket.org/ifan-moladin/base-project/utils/httpserver"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestJwtAuthenticator_handle(t *testing.T) {
	type fields struct {
		secret  string
		isValid ValidatorFunc
	}
	type args struct {
		ctx httpserver.Context
	}
	type testCase struct {
		name   string
		fields fields
		args   args
	}
	var tests []testCase

	var name = "error"
	var f = fields{
		secret: "secret",
	}
	mockCtx := mocks.NewContext(t)
	var a = args{ctx: mockCtx}
	var invalidBearerToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.5yDeOpW5CSQbX91bDg_M8taycvXTEvq6ntTkcSuJRfs"
	mockCtx.EXPECT().GetHeader("Authorization").Return("Bearer " + invalidBearerToken).Once()
	mockCtx.EXPECT().Status(http.StatusUnauthorized).Once()
	mockCtx.EXPECT().Abort().Once()
	tests = append(tests, testCase{
		name:   name,
		fields: f,
		args:   a,
	})

	name = "success"
	var validBearerToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.XbPfbIHMI6arZ3Y922BhjWgQzWXcXNrz0ogtVhfEd2o"
	mockCtx.EXPECT().GetHeader("Authorization").Return("Bearer " + validBearerToken).Once()
	var jwtToken, _ = (JwtAuthenticator{
		secret: f.secret,
	}).parseToken(validBearerToken)
	mockCtx.EXPECT().Set(authenticationPayloadKey, jwtToken.Claims.(jwt.MapClaims)).Once()
	mockCtx.EXPECT().Next().Once()
	tests = append(tests, testCase{
		name:   name,
		fields: f,
		args:   a,
	})

	name = "is not valid"
	f.isValid = func(payload any) bool {
		return false
	}
	mockCtx.EXPECT().GetHeader("Authorization").Return("Bearer " + validBearerToken).Once()
	payload := jwtToken.Claims.(jwt.MapClaims)
	mockCtx.EXPECT().Set(authenticationPayloadKey, payload).Once()
	mockCtx.EXPECT().Status(http.StatusUnauthorized).Once()
	mockCtx.EXPECT().Abort().Once()
	tests = append(tests, testCase{
		name:   name,
		fields: f,
		args:   a,
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := JwtAuthenticator{
				secret:  tt.fields.secret,
				isValid: tt.fields.isValid,
			}
			a.handle(tt.args.ctx)
		})
	}
}

func TestJwtAuthenticator_GetPayloadFromContext(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		a := JwtAuthenticator{}
		mockCtx := mocks.NewContext(t)
		mockCtx.EXPECT().Get(authenticationPayloadKey).Return(nil, false).Once()
		got, got1 := a.GetPayloadFromContext(mockCtx)
		if !reflect.DeepEqual(got, nil) {
			t.Errorf("GetPayloadFromContext() got = %v, want %v", got, nil)
		}
		if got1 != false {
			t.Errorf("GetPayloadFromContext() got1 = %v, want %v", got1, false)
		}
	})
}

func TestJwtAuthenticator_Authenticate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		a := JwtAuthenticator{}
		got := a.Authenticate("secret")
		assert.NotNil(t, got)
	})
}

func TestJwtAuthenticator_Authorize(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		a := JwtAuthenticator{}
		got := a.Authorize("secret", func(payload any) bool {
			return false
		})
		assert.NotNil(t, got)
	})
}
