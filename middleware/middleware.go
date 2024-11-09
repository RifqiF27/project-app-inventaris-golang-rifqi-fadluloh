package middleware_auth

import (
	"inventaris/service"
	"net/http"
)

func SessionMiddleware(service *service.UserService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("session_token")
			if err != nil || cookie.Value == "" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			userID, err := service.GetUserIDBySessionID(cookie.Value)
			if err != nil || userID == 0 {
				http.Error(w, "Invalid session", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
