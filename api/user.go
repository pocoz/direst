package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/dieqnt/direst/models"
	"github.com/dieqnt/direst/storage"
)

func New(psql *storage.Engine) *User {
	return &User{psql: psql}
}

type User struct {
	psql *storage.Engine
}

func (r *User) UserGet(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	user, err := r.psql.UserGetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (r *User) UserGetList(c *gin.Context) {
	userList, err := r.psql.UserGetList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, userList)
}

func (r *User) UserCreate(c *gin.Context) {
	var requestBody models.User
	err := c.BindJSON(&requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	err = r.psql.UserInsert(&requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func (r *User) UserUpdate(c *gin.Context) {
	var requestBody models.User
	err := c.BindJSON(&requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	err = r.psql.UserUpdate(&requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (r *User) UserDelete(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	err = r.psql.UserDelete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}
