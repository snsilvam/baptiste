package handlers

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
)

// Handler for our logout.
func HandlerLogOut(c *gin.Context) {
	logoutUrl, err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + "/v2/logout")
	if err != nil {
		fmt.Println("error al intentar conseguir la url del logout: ", err)
		c.JSON(500, MessageError{
			Message: err.Error(),
			Url:     "/logout",
		})
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}

	returnTo, err := url.Parse(scheme + "://" + c.Request.Host)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	parameters := url.Values{}
	parameters.Add("returnTo", returnTo.String())
	parameters.Add("client_id", os.Getenv("AUTH0_CLIENT_ID"))
	logoutUrl.RawQuery = parameters.Encode()

	c.Redirect(http.StatusTemporaryRedirect, logoutUrl.String())
}
