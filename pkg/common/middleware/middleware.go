package middleware

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/sssash18/Digicart/pkg/common/models"
)

func Authenticate(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := strings.Split((r.Header.Get("Authorization")), " ")[1]
		if tokenString == "" {
			http.Error(w, "Unauthorized access", http.StatusUnauthorized)
			return
		}
		claims := &models.AuthClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				resp, _ := json.Marshal(&models.Response{
					Status: "error",
					Err:    err.Error(),
				})
				w.Write(resp)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			resp, _ := json.Marshal(&models.Response{
				Status: "error",
				Err:    err.Error(),
			})
			w.Write(resp)
			return
		}
		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			resp, _ := json.Marshal(&models.Response{
				Status: "error",
				Err:    err.Error(),
			})
			w.Write(resp)
			return
		}
		r.Header.Set("email", claims.Data.Email)
		r.Header.Set("phone", claims.Data.Phone)
		r.Header.Set("firstname", claims.Data.FirstName)
		r.Header.Set("lastname", claims.Data.LastName)
		r.Header.Set("userID", claims.Data.UserID)
		handler.ServeHTTP(w, r)
	})

}
