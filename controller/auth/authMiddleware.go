package auth

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	model "github.com/goplay/model"
	auth "github.com/goplay/model/auth"
	"github.com/goplay/service/logger"
)

const (
	privKeyPath = "app.rsa"
	pubKeyPath  = "app.rsa.pub"
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

// read the key files before starting http handlers
func init() {
	signBytes, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		logger.Error(err)
	}
	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		logger.Error(err)
	}
	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		logger.Error(err)
	}
	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		logger.Error(err)
	}
	logger.Infof("Successfully parsed Cert and Key files")
}

//Authenticate to handle User authentication
func Authenticate(c *gin.Context) {

	//TODO: Add to validate creds passed by User

	// Create the Claims
	claims := auth.Claims{}
	claims.Data.AuthUser = auth.AuthUser{}
	expTime := time.Now().Add(time.Second * 10)
	claims.StandardClaims = jwt.StandardClaims{ExpiresAt: expTime.Unix(), Issuer: "root"}

	// Create a new jwt token, specifying signing method and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(signKey)
	if err != nil {
		logger.Error(err)
		processErr(c, err, http.StatusInternalServerError)
		return
	}
	jwtResponse := auth.JwtTokenResponse{JWTToken: tokenString, Status: model.SuccessStatus, Message: "Issued JWT Token"}
	c.JSON(http.StatusOK, jwtResponse)
}

//RefreshJWTToken refresh JWT token if user is authenticated
func RefreshJWTToken(c *gin.Context) {
	accessToken := c.Request.Header.Get("access-token")
	claims, err := getClaims(accessToken)
	if err != nil {
		processErr(c, err, http.StatusUnauthorized)
		return
	}
	expTime := time.Now().Add(time.Second * 10)
	claims.StandardClaims = jwt.StandardClaims{ExpiresAt: expTime.Unix(), Issuer: "root"}

	// Create a new jwt token, specifying signing method and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(signKey)
	if err != nil {
		logger.Error(err)
		processErr(c, err, http.StatusInternalServerError)
		return
	}
	jwtResponse := auth.JwtTokenResponse{JWTToken: tokenString, Status: model.SuccessStatus, Message: "Refreshed JWT Token"}
	c.JSON(http.StatusOK, jwtResponse)
}

//AuthorizationMiddleware to handle Authentication
func AuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.Request.Header.Get("access-token")
		_, err := getClaims(accessToken)
		if err != nil {
			processErr(c, err, http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}

func getClaims(accessToken string) (*auth.Claims, error) {
	err := errors.New("Error parsing JWT token")
	if accessToken == "" {
		err = fmt.Errorf("Empty JWT token")
		logger.Errorf("%v", err)
		return nil, err
	}
	// Parse the token from the request header
	token, err := jwt.ParseWithClaims(accessToken, &auth.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			err = fmt.Errorf("Unexpected signing method")
			logger.Errorf("%v", err)
			return nil, err
		}
		return verifyKey, nil
	})

	if err != nil || !token.Valid {
		err = fmt.Errorf("Error validating JWT token; Error: [%v]", err)
		logger.Errorf("%v", err)
		return nil, err
	}

	// Return claim
	if claims, ok := token.Claims.(*auth.Claims); ok {
		logger.Infof("Successfully parsed JWT token")
		return claims, nil
	}

	return nil, err
}

// RunAsRootUser sets the 'root' in the context.
func RunAsRootUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("user", "root")
		c.Next()
	}
}

func processErr(c *gin.Context, err error, errCode int) {
	response := model.GenericResponse{Status: model.ErrorStatus, Message: err.Error()}
	c.JSON(errCode, response)
	c.Abort()
	return
}
