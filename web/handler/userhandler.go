package handler

import (
	"database/sql"
	"net/http"

	"gitlab.com/parkby/parkby-service/db"
	"gitlab.com/parkby/parkby-service/db/finder"
	"gitlab.com/parkby/parkby-service/web/dto"
	"gitlab.com/parkby/parkby-service/web/mapper"

	"github.com/gin-gonic/gin"
)

// InitUserEndpoint initialize user endpoints
func InitUserEndpoint(router *gin.Engine) {
	router.POST("/users/register", register)
	router.POST("/users/login", login)
}

func register(c *gin.Context) {
	var userRegister dto.UserRegisterDTO
	c.BindJSON(&userRegister)

	user := finder.UserFinder.FindWithEmail(userRegister.Email)

	if user.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Email already exist"})
	} else {

		newUser := mapper.UserMapper.ToModel(userRegister)

		db.TransactionInvoker.Invoke(func(tx *sql.Tx) error {
			err := newUser.Save(tx)
			return err
		})

		c.JSON(http.StatusOK, gin.H{"message": "Email is OK"})
	}
}

func login(c *gin.Context) {
	var userLogin dto.UserLoginDTO
	c.BindJSON(&userLogin)

	user := finder.UserFinder.FindWithEmail(userLogin.Email)

	if user.ID == 0 || !user.EqualPassword(userLogin.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Invalid"})
	} else {
		c.JSON(http.StatusOK, gin.H{"token": "!@#$%"})
	}

}
