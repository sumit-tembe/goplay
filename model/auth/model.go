package auth

import jwt "github.com/dgrijalva/jwt-go"

//AuthUser holds the authenticated user details from User Agent API
type AuthUser struct {
	UserID     string `json:"uid" valid:"required"`
	UserStatus string `json:"userstatus"`
	FirstName  string `json:"firstname"`
	MiddleName string `json:"middlename"`
	LastName   string `json:"lastname"`
}

// Claims ...
type Claims struct {
	Data struct {
		AuthUser AuthUser `json:"scope"`
	} `json:"data"`
	jwt.StandardClaims
}

//JwtTokenResponse ...
type JwtTokenResponse struct {
	JWTToken string `json:"jwt-token"`
	Status   string `json:"status"`
	Message  string `json:"message"`
}
