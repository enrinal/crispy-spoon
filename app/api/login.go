package api

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Username string
	Role     string
	Password string
}

var users = map[string]User{
	"rogu": {
		Username: "rogu",
		Role:     "admin",
		Password: "123456",
	},
	"dodi": {
		Username: "dodi",
		Role:     "user",
		Password: "123456",
	},
	"joko": {
		Username: "joko",
		Role:     "user",
		Password: "123456",
	},
}

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

var JwtKey = []byte("secret")

func (api *API) Login(c *gin.Context) {
	var creds Credentials

	// decode request body into struct and check for errors
	err := c.ShouldBindJSON(&creds)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: err.Error(),
		})
	}

	// check if the user exists
	user, ok := users[creds.Username]
	if !ok {
		c.JSON(http.StatusUnauthorized, ErrorResponse{
			Error: "invalid credentials",
		})
		return
	}

	// check password is correct
	if user.Password != creds.Password {
		c.JSON(http.StatusUnauthorized, ErrorResponse{
			Error: "invalid credentials",
		})
		return
	}

	// set expiration time
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Username: creds.Username,
		Role:     user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Subject:   "login",
		},
	}

	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sign token and handle error
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}
