package authenticator

import (
	"errors"
	"github.com/ifaniqbal/go-base-project/internal/utils"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

type JwtAuthenticator struct {
	secret  string
	isValid ValidatorFunc
}

type ValidatorFunc func(payload any) bool

const authenticationPayloadKey = "authentication-payload"

// Authenticate creates and returns a JwtAuthenticator middleware that validates
// JWT tokens with the provided secret.
func (a JwtAuthenticator) Authenticate(secret string) utils.HandlerFunc {
	jwtAuthenticator := JwtAuthenticator{secret: secret}
	return jwtAuthenticator.handle
}

// Authorize creates and returns a JwtAuthenticator middleware that validates
// JWT tokens with the provided secret and custom validation function.
func (a JwtAuthenticator) Authorize(secret string, isValid ValidatorFunc) utils.HandlerFunc {
	jwtAuthenticator := JwtAuthenticator{
		secret:  secret,
		isValid: isValid,
	}
	return jwtAuthenticator.handle
}

// SendUnauthorizedResponse sends an unauthorized response
// (HTTP status code 401) to the client.
func (a JwtAuthenticator) SendUnauthorizedResponse(ctx utils.Context) {
	ctx.Status(http.StatusUnauthorized)
}

// GetPayloadFromContext returns the authentication payload stored in
// the provided http server context.
func (a JwtAuthenticator) GetPayloadFromContext(ctx utils.Context) (any, bool) {
	return ctx.Get(authenticationPayloadKey)
}

// handle validates JWT tokens and sets the authentication payload in
// the request context. If the token is invalid or the payload fails validation,
// an unauthorized response is sent and the request is aborted.
func (a JwtAuthenticator) handle(ctx utils.Context) {
	bearerToken := strings.Replace(ctx.GetHeader("Authorization"), "Bearer ", "", -1)
	jwtToken, err := a.parseToken(bearerToken)
	payload, claimsOk := jwtToken.Claims.(jwt.MapClaims)
	if err != nil || !claimsOk {
		a.SendUnauthorizedResponse(ctx)
		ctx.Abort()
		return
	}

	ctx.Set(authenticationPayloadKey, payload)

	if a.isValid != nil && !a.isValid(payload) {
		a.SendUnauthorizedResponse(ctx)
		ctx.Abort()
		return
	}

	ctx.Next()
}

// parseToken uses the jwt.Parse function to parse the bearerToken string and
// returns a pointer to a jwt.Token type and an error.
func (a JwtAuthenticator) parseToken(bearerToken string) (*jwt.Token, error) {
	return jwt.Parse(bearerToken, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(a.secret), nil
	})
}
