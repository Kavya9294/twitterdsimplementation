/* package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func main() {
	data := `{"user_id":"a1b2c3","username":"nikola"}`
	header := `{"alg": "HS256",
		"typ": "JWT"}`
	uEnc := base64.URLEncoding.EncodeToString([]byte(data))
	hEnc := base64.URLEncoding.EncodeToString([]byte(header))
	fmt.Println(uEnc)
	fmt.Println(hEnc)
	msg := uEnc + "." + hEnc
	fmt.Println(msg)
	sKey := "hellofromtheotherside"
	uDec, _ := base64.URLEncoding.DecodeString(uEnc)
	hDec, _ := base64.URLEncoding.DecodeString(hEnc)
	fmt.Println(string(uDec))
	fmt.Println(string(hDec))
	key := []byte(sKey)
	fmt.Println(string(key))
	h := hmac.New(sha256.New, key)
	fmt.Println(h)
	h.Write([]byte(msg))
	base := base64.URLEncoding.EncodeToString(h.Sum(nil))
	fmt.Println(string(base))
	final := msg + "," + base
	fmt.Println(final)

}
*/

// adding auth.go for JWT and custom claims which has an expiration time and which will enable required authentication paradigms
package main

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// jwtCustomClaims are custom claims extending default ones.
type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

func login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "jon" && password == "shhh!" {

		// Set custom claims
		claims := &jwtCustomClaims{
			"Jon Snow",
			true,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			},
		}

		// Create token with claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, echo.Map{
			"token": t,
		})
	}

	return echo.ErrUnauthorized
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	name := claims.Name
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Login route
	e.POST("/login", login)

	// Unauthenticated route
	e.GET("/", accessible)

	// Restricted group
	r := e.Group("/restricted")

	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &jwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	r.Use(middleware.JWTWithConfig(config))
	r.GET("", restricted)

	e.Logger.Fatal(e.Start(":9091"))
}
