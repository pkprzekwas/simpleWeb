package jwt

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkprzekwas/simpleWeb/pkg/models"
	"github.com/pkprzekwas/simpleWeb/pkg/utils"
	"net/http"
	"os"
	"strings"
)

const (
	EmptyToken = ""
)

var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		noAuth := []string{"/api/user/new", "/api/user/login"}
		requestPath := r.URL.Path

		for _, path := range noAuth {
			if requestPath == path {
				next.ServeHTTP(w, r)
				return
			}
		}

		response := make(map[string]interface{})
		tokenHeader := r.Header.Get("Authorization")

		if tokenHeader == EmptyToken {
			response = utils.Message(false, "Missing auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			utils.Respond(w, response)
			return
		}

		tokenSplit := strings.Split(tokenHeader, " ")
		if len(tokenSplit) != 2 {
			response = utils.Message(false, "Invalid auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			utils.Respond(w, response)
			return
		}

		tokenPart := tokenSplit[1]
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})
		if err != nil || !token.Valid {
			response = utils.Message(false, "Malformed auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			utils.Respond(w, response)
			return
		}

		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
