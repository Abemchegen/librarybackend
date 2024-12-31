package infrastracture

import (
	"fmt"
	"librarybackend/config"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		env, err := config.LoadEnv()
		if err != nil {
			fmt.Print("Error in env.load")
		}

		var JwtSecret = []byte(env.AccessTokenSecret)

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
			c.JSON(401, gin.H{"error": "Invalid authorization header"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(authParts[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return JwtSecret, nil
		})

		if err != nil {
			fmt.Printf("Error parsing token: %v\n", err)
			c.JSON(401, gin.H{"error": "Invalid JWT"})
			c.Abort()
			return
		}

		if !token.Valid {
			fmt.Println("Token is not valid")
			c.JSON(401, gin.H{"error": "Invalid JWT"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(401, gin.H{"error": "Invalid JWT"})
			c.Abort()
			return
		}

		// c.Set("user", claims)
		c.Set("user_id", claims["user_id"])
		c.Set("name", claims["name"])
		c.Set("role", claims["role"])
		fmt.Println(claims)

		// c.Set("Roles", claims["role"])
		c.Next()
	}
}

func RoleMiddleware(Role_ string) gin.HandlerFunc {
	return func(c *gin.Context) {

		role := c.GetString("role")
		if role == "" {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Role not found in claims"})
			c.Abort()
			return
		}

		if role != Role_ {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func EligibilityMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Retrieve the is_activated claim as a boolean
		isActivated, exists := c.Get("is_activated")
		if !exists {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "is_activated not found in claims"})
			c.Abort()
			return
		}

		// Convert to bool if necessary
		if isActivatedBool, ok := isActivated.(bool); ok {
			if !isActivatedBool {
				c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Please activate your account to access this service."})
				c.Abort()
				return
			}
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "is_activated claim is not a valid boolean"})
			c.Abort()
			return
		}

		c.Next()
	}
}
