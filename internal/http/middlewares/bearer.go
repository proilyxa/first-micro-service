package middlewares

import (
	"context"
	"education-project/internal/models"
	"education-project/internal/repositories"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type authUser struct{}

func BearerAuth(userRepo repositories.UserRepository) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			authToken, ok := strings.CutPrefix(authHeader, "Bearer ")
			if !ok {
				authFailed(w)
				return
			}

			user, err := userRepo.FindByAuthToken(r.Context(), authToken)
			if err != nil {
				authFailed(w)
				return
			}

			r = r.WithContext(context.WithValue(r.Context(), authUser{}, user))
			next.ServeHTTP(w, r)
		})
	}
}

func GetAuthUser(ctx context.Context) *models.User {
	return ctx.Value(authUser{}).(*models.User)
}

func authFailed(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	resp := make(map[string]string)
	resp["message"] = "Unauthorized"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}
