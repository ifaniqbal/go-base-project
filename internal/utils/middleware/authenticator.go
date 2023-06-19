package middleware

import (
	"github.com/ifaniqbal/go-base-project/internal/utils"
	"github.com/ifaniqbal/go-base-project/pkg/ginwrapper/middleware/authenticator"
)

type Authenticator interface {
	Authenticate(secret string) utils.HandlerFunc
	Authorize(secret string, isValid authenticator.ValidatorFunc) utils.HandlerFunc
	SendUnauthorizedResponse(ctx utils.Context)
	GetPayloadFromContext(ctx utils.Context) (any, bool)
}
