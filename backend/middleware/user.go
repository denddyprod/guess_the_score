package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"guess_the_score/backend/models"
	"guess_the_score/backend/utils/constants"
	"guess_the_score/backend/views"
	"log"
	"net/http"
	"strings"
)

type RequireUser struct {
	UserService        models.UserService
}

func (self *RequireUser) Required(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := extractToken(r)
		if err != nil {
			log.Println("Token not provided or malformed")
			log.Println(constants.ErrNoToken.Error())
			data := map[string]string{"success": "false", "errorMsg": constants.ErrNoToken.Error()}
			views.SendResponse(w, data, http.StatusForbidden)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			//Make sure that the token method conform to "SigningMethodHMAC"
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("shhhhhh"), nil
		})
		if err != nil {
			log.Println(constants.ErrInternalServer.Error())
			data := map[string]string{"success": "false", "errorMsg": constants.ErrInternalServer.Error()}
			views.SendResponse(w, data, http.StatusForbidden)
			return
		}

		var username string
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			username, ok =  claims["username"].(string)
			if !ok {
				log.Println(constants.ErrClaimsExtract.Error())
				data := map[string]string{"success": "false", "errorMsg": constants.ErrClaimsExtract.Error()}
				views.SendResponse(w, data, http.StatusForbidden)
				return
			}
		} else {
			log.Println("Expired token. Please try again")
			data := map[string]string{"success": "false", "errorMsg": constants.ErrUserMustLogin.Error()}
			views.SendResponse(w, data, http.StatusForbidden)
			return
		}

		log.Println(username)

		user, err := self.UserService.FindByUsername(username)
		if err != nil {
			log.Println(constants.ErrNotFound.Error(), username)
			data := map[string]string{"success": "false", "errorMsg": constants.ErrNotFound.Error()}
			views.SendResponse(w, data, http.StatusForbidden)
			return
		}

		ctx := context.WithValue(r.Context(), "user", user)
		r = r.WithContext(ctx)

		next(w, r)
	})
}

func extractToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	authHeaderContent := strings.Split(authHeader, " ")
	if len(authHeaderContent) != 2 {
		return "", errors.New("Token not provided or malformed")
	}
	return authHeaderContent[1], nil
}
