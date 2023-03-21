package helper

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"app/pkg/config"
)

type UserClaims struct {
	ID   uint64 `json:"id"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

type Payload struct {
	ID        uint64 `json:"id"`
	Role      string `json:"role"`
	ExpiresAt int64  `json:"exp"`
}

/* -------------------------------------------------------------------------- */
/*                                Create Token                                */
/* -------------------------------------------------------------------------- */
// userClaims := helper.UserClaims{
// 	ID:   1,
// 	Role: "admin",
// }
// token, err := helper.CreateToken(userClaims)
// if err != nil {
// 	fmt.Println("t error:", err)
// }
// fmt.Println("token:", token)
/* -------------------------------------------------------------------------- */
func CreateToken(userClaims UserClaims) (string, error) {
	// Set the expiration time for the token
	expiresIn := time.Hour * 24
	expirationTime := time.Now().Add(expiresIn).Unix()

	// Convert the Unix timestamp to a time.Time value
	expTime := time.Unix(expirationTime, 0)

	// Set the claims for the token
	claims := &UserClaims{
		ID:   userClaims.ID,
		Role: userClaims.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   strconv.Itoa(int(userClaims.ID)),
		},
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the specified key
	signingKey := []byte(config.AppConfig.JWT_SIGNING_KEY)
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

/* -------------------------------------------------------------------------- */
/*                                 Parse Token                                */
/* -------------------------------------------------------------------------- */
// payload, err := helper.ParseToken(token)
// if err != nil {
// 	fmt.Println("p error:", err)
// }
// fmt.Println("payload:", payload)
// fmt.Println("time:", time.Unix(payload.ExpiresAt, 0))
/* -------------------------------------------------------------------------- */
func ParseToken(tokenString string) (*Payload, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		signingKey := []byte(config.AppConfig.JWT_SIGNING_KEY)
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	payload := &Payload{
		ExpiresAt: claims.ExpiresAt.Time.Unix(),
	}

	id, err := strconv.ParseUint(claims.Subject, 10, 32)
	if err != nil {
		return nil, err
	}
	payload.ID = uint64(id)
	payload.Role = claims.Role

	return payload, nil
}
