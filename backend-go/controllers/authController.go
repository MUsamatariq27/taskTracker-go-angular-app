package controllers

import (
	//"fmt"
	"net/http"

	//"strings"
	"time"

	"github.com/MUsamaT/task-tracker/database"
	"github.com/MUsamaT/task-tracker/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("secret_key_goes_here")

func RegisterUser(ctx *gin.Context) {

	var input models.User
	err := ctx.ShouldBind(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Input!",
		})
		return
	}

	/*query := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		password TEXT NOT NULL
		);`

	_, err = database.DB.Exec(query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error ensuring users table exists!",
		})
		return
	}*/

	hashed, erro := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if erro != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error Hashing Password!",
		})
		return
	}

	insert := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3)`
	_, err = database.DB.Exec(insert, input.Name, input.Email, string(hashed))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email already Exists!"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User registered succesfully"})

}

func Login(c *gin.Context) {
	var input models.User

	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
		return
	}

	var user models.User
	err = database.DB.Get(&user, "SELECT * FROM users WHERE email=$1", input.Email)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found!"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Credentials!"})
		return
	}

	//generating token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Token!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})

}
