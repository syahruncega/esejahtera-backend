package middleware

import (
	"fmt"
	"kemiskinan/config"
	"kemiskinan/model"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(cntx *gin.Context) {
	var tokenString, err = cntx.Cookie("Authorization")
	if err != nil {
		cntx.AbortWithStatus(http.StatusUnauthorized)
	}

	// Parse takes the token string and a function for looking up the key.
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("APP_SECRET_KEY")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			cntx.AbortWithStatus(http.StatusUnauthorized)
		}

		var user model.User
		var err = config.DB.Take(&user, claims["uId"]).Error

		if err != nil {
			cntx.AbortWithStatus(http.StatusUnauthorized)
		}

		cntx.Set("user", user)
	} else {
		cntx.AbortWithStatus(http.StatusUnauthorized)
	}

	cntx.Next()
}
