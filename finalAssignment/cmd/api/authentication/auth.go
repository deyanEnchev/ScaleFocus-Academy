package authentication

import (
	"context"
	"crypto/subtle"
	"final/cmd/persistence"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Authentication(username, password string, c echo.Context) (bool, error) {
	repo := persistence.NewRepository(persistence.ConnectToDb())
	res, err := repo.GetUsers(context.Background())
	if err != nil {
		fmt.Println("Failed fetching users from DB.")
		return false, err
	}

	for _, row := range res {
		if subtle.ConstantTimeCompare([]byte(username), []byte(row.Username)) == 1 {
			if err := bcrypt.CompareHashAndPassword([]byte(row.Password), []byte(password)); err == nil {
				c.Set("UserID", int(row.ID))
				return true, nil
			}

			return false, c.String(http.StatusUnauthorized, "Wrong password.")
		}
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("Failed encrypting user's password: %v", err)
		return false, c.String(http.StatusInternalServerError, "Failed creating user.")
	}
	err = repo.CreateUserInDb(username, string(hash))
	if err != nil {
		fmt.Printf("Couldn't create user: %v", err)
		return true, c.String(http.StatusInternalServerError, "Failed creating user.")
	}

	res, err = repo.GetUsers(context.Background())
	if err != nil {
		fmt.Println("Failed fetching users from DB.")
		return true, err
	}

	c.Set("UserID", int(res[len(res)-1].ID))
	return true, c.String(http.StatusCreated, "New user created.")
}
