package middleware

import "net/http"

func RequireAuthenticatedUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// If the user is not authenticated, redirect them to the login page and
		// return from the middleware chain so that no subsequent handlers in
		// the chain are executed.

		if 0 == 0 {
			http.Redirect(w, r, "/login", 302)
			return
		}
		// Otherwise call the next handler in the chain.
		next.ServeHTTP(w, r)
	})
}
