package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"

	"baptiste.com/authenticator"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Handler for our login.
func HandlerLogin(auth *authenticator.Authenticator) gin.HandlerFunc {
	return func(c *gin.Context) {
		state, err := generateRandomState()
		if err != nil {
			fmt.Println("error en la funcion generateRandomState de BAPTISTE: ", err)
			c.JSON(500, MessageError{
				Message: err.Error(),
				Url:     "/login",
			})
			//c.String(http.StatusInternalServerError, err.Error())
			return
		}

		// Save the state inside the session.
		session := sessions.Default(c)
		session.Set("state", state)
		if err := session.Save(); err != nil {
			fmt.Println("error en la funcion session.Save(): ", err)
			c.JSON(500, MessageError{
				Message: err.Error(),
				Url:     "/login",
			})
			//c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.Redirect(http.StatusTemporaryRedirect, auth.AuthCodeURL(state))
	}
}

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	state := base64.StdEncoding.EncodeToString(b)

	return state, nil
}
