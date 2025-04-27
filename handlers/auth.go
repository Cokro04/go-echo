package handlers

import (
	"net/http"
	"simple-store-app/config"
	"simple-store-app/models"
	"simple-store-app/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword

	// Insert into database
	_, err = config.DB.Exec("INSERT INTO users (username, password) VALUES (?, ?)", u.Username, u.Password)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, u)
}

func Login(c echo.Context) error {
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	// Fetch user from database
	row := config.DB.QueryRow("SELECT password FROM users WHERE username = ?", u.Username)
	var hashedPassword string
	if err := row.Scan(&hashedPassword); err != nil {
		return echo.ErrUnauthorized
	}

	// Check password
	if !utils.CheckPasswordHash(u.Password, hashedPassword) {
		return echo.ErrUnauthorized
	}

	// Generate JWT
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = u.Username

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}
