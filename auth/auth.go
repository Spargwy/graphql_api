package auth

import (
	"context"
	"fmt"
	"gql_app/graph/model"
	"gql_app/graph/resolvers/storage"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
)

var viewerCtxKey = &contextKey{"viewer"}

type contextKey struct {
	name string
}

func Middleware(resolver storage.Psql) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header["Authorization"] == nil {
				next.ServeHTTP(w, r)
				return
			}
			token, err := validateToken(r.Header["Authorization"][0])

			if err != nil {
				w.WriteHeader(http.StatusForbidden)
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)

			if !ok || !token.Valid {
				w.WriteHeader(http.StatusForbidden)
				return
			}

			var viewer model.Viewer
			user, err := resolver.SelectUserByID(int(claims["user_id"].(float64)))

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			viewer.User = &user
			ctx := context.WithValue(r.Context(), viewerCtxKey, &viewer)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func ForContext(ctx context.Context) *model.Viewer {
	raw, _ := ctx.Value(viewerCtxKey).(*model.Viewer)
	return raw
}

func validateToken(authorizationHeader string) (token *jwt.Token, err error) {
	if authorizationHeader == "" {
		return nil, fmt.Errorf("token is empty")
	}

	token, err = jwt.Parse(authorizationHeader, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error in parsing")
		}
		return []byte(os.Getenv("SECRET_FOR_JWT")), nil
	})

	if err != nil {
		return
	}

	return
}
