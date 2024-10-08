package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostUserHandler(c *gin.Context) {
	/* var user models.UserInsert

	if err := c.BindJSON(&user); err != nil {
		fmt.Println("error al recibir el objeto(user) en el request", err)

		c.JSON(404, MessageError{
			Message: err.Error(),
			Url:     "/users",
		})
		return
	}

	err := repository.InsertUser(c, &user)
	if err != nil {
		c.JSON(500, MessageError{
			Message: err.Error(),
			Url:     "/users",
		})
		return
	} */

	c.IndentedJSON(http.StatusOK, "Usuario creado con exito.")
}

func GetUser(c *gin.Context) {
	/* idDoc := c.Param("id")
	user, err := repository.GetUser(c, idDoc)

	if err != nil {
		c.JSON(500, MessageError{
			Message: err.Error(),
			Url:     "/users/" + idDoc,
		})
		return
	} */
	user := "test"

	c.JSON(http.StatusOK, user)
}

func GetAllUsers(c *gin.Context) {
	/* users, err := repository.GetAllUsers(c)
	if err != nil {
		c.JSON(500, MessageError{
			Message: err.Error(),
			Url:     "/users/",
		})
		return
	} */
	users := "test"

	c.JSON(http.StatusOK, users)
}

func UpdateUser(c *gin.Context) {
	/* var user models.Users

	if err := c.BindJSON(&user); err != nil {
		fmt.Println("error al recibir el objeto(UserUpdated) en el request", err)

		c.JSON(404, MessageError{
			Message: err.Error(),
			Url:     "/users",
		})
		return
	}

	if err := repository.UpdateUser(c, &user); err != nil {
		c.JSON(500, MessageError{
			Message: err.Error(),
			Url:     "/users",
		})
		return
	} */

	c.IndentedJSON(http.StatusOK, "Usuario actualizado con exito.")
}
