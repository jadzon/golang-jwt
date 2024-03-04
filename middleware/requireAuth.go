package middleware

import (
	"fmt"
	"jwt_najnowszy/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context, db models.Database) {
	tokenString, err := c.Cookie("Authcookerson")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, _ = token, err
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("TOKEN_SECRET_KEY")), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["expiresAt"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// for _, u := range data.Users {
		// 	if u.ID == claims["id"] {
		// 		theUser = u
		// 		break
		// 	}
		// }
		theUser, err := db.GetUserByID(int(claims["id"].(float64)))
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Set("user", theUser)
		c.Next()

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}
