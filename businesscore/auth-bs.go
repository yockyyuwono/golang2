package businesscore

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func ExampleParseWithClaims_customClaimsType() {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJleHAiOjE1MDAwLCJpc3MiOiJ0ZXN0In0.HE7fK0xOQwFEr4WDgRWj4teRPZ6i3GLwD5YCm6Pwu_c"

	type MyCustomClaims struct {
		Foo string `json:"foo"`
		jwt.StandardClaims
	}

	// sample token is expired.  override time so it parses as valid
	at(time.Unix(0, 0), func() {
		token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("AllYourBase"), nil
		})

		if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
			fmt.Printf("%v %v", claims.Foo, claims.StandardClaims.ExpiresAt)
		} else {
			fmt.Println(err)
		}
	})

	// Output: bar 15000

}
func at(t time.Time, f func()) {
	jwt.TimeFunc = func() time.Time {
		return t
	}
	f()
	jwt.TimeFunc = time.Now
}
