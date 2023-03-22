package helper

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"app/pkg/config"
)

type token struct {
	Platform Platform `json:"platform" bson:"platform"`
	UserType UserType `json:"user_type" bson:"user_type"`
}

type Platform struct {
	Web uint8 `json:"web" bson:"web"`
	App uint8 `json:"app" bson:"app"`
}

type UserType struct {
	Portal uint8 `json:"portal" bson:"portal"`
	User   uint8 `json:"user" bson:"user"`
}

type PayloadClaims struct {
	ID       uint64 `json:"id" bson:"id"`
	RoleID   int    `json:"role_id" bson:"role_id"`
	Platform uint8  `json:"platform" bson:"platform"`
	UserType uint8  `json:"user_type" bson:"user_type"`
	jwt.RegisteredClaims
}

type Payload struct {
	ID        uint64 `json:"id" bson:"id"`
	RoleID    int    `json:"role_id" bson:"role_id"`
	Platform  uint8  `json:"platform" bson:"platform"`
	UserType  uint8  `json:"user_type" bson:"user_type"`
	ExpiresAt int64  `json:"exp" bson:"exp"`
}

var platform = Platform{
	Web: 1,
	App: 2,
}

var userType = UserType{
	Portal: 1,
	User:   2,
}

var Token = token{
	Platform: platform,
	UserType: userType,
}

/* -------------------------------------------------------------------------- */
/*                                Create Token                                */
/* -------------------------------------------------------------------------- */
// payloadClaims := helper.PayloadClaims{
// 	ID:   1,
// 	Role: "admin",
// }
// token, err := helper.CreateToken(payloadClaims)
// if err != nil {
// 	fmt.Println("t error:", err)
// }
// fmt.Println("token:", token)
/* -------------------------------------------------------------------------- */
func (t *token) CreateToken(payloadClaims PayloadClaims) (string, error) {
	// Set the expiration time for the token
	expiresIn := time.Hour * 24
	expirationTime := time.Now().Add(expiresIn).Unix()

	// Convert the Unix timestamp to a time.Time value
	expTime := time.Unix(expirationTime, 0)

	// Set the claims for the token
	claims := &PayloadClaims{
		ID:       payloadClaims.ID,
		RoleID:   payloadClaims.RoleID,
		Platform: payloadClaims.Platform,
		UserType: payloadClaims.UserType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   strconv.Itoa(int(payloadClaims.ID)),
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
func (t *token) ParseToken(tokenString string) (*Payload, error) {
	token, err := jwt.ParseWithClaims(tokenString, &PayloadClaims{}, func(token *jwt.Token) (interface{}, error) {
		signingKey := []byte(config.AppConfig.JWT_SIGNING_KEY)
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*PayloadClaims)
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
	payload.ID = id
	payload.RoleID = claims.RoleID
	payload.Platform = claims.Platform
	payload.UserType = claims.UserType

	return payload, nil
}
