package authenticator

import (
	"github.com/ifaniqbal/go-base-project/pkg/httpserver"
)

type Authenticator interface {
	Authenticate(secret string) httpserver.HandlerFunc
	Authorize(secret string, isValid ValidatorFunc) httpserver.HandlerFunc
	SendUnauthorizedResponse(ctx httpserver.Context)
	GetPayloadFromContext(ctx httpserver.Context) (any, bool)
}
