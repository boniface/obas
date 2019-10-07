package middleware

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"gopkg.in/square/go-jose.v2/jwt"
	"net/http"
)

type LoginSession struct {
	SessionManager *scs.SessionManager
}

func (session LoginSession) RequireAuthenticatedUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(write http.ResponseWriter, request *http.Request) {
		if !VeryFyTheToken(session.SessionManager, request) {
			http.Redirect(write, request, "/login", 302)
			return
		}
		// Otherwise call the next handler in the chain.
		next.ServeHTTP(write, request)
	})
}

func VeryFyTheToken(manager *scs.SessionManager, request *http.Request) bool {
	email := manager.GetString(request.Context(), "userId")
	token := manager.GetString(request.Context(), "token")
	webToken, err := jwt.ParseSigned(token)
	if err != nil {
		fmt.Println("failed to parse JWT:%+v", err)
		return false
	}
	var claims map[string]interface{}
	// decode JWT token without verifying the signature g)
	_ = webToken.UnsafeClaimsWithoutVerification(&claims)
	manager.Put(request.Context(), "role", claims["role"])
	return email == claims["email"]
}
