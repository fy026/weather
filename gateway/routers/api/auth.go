package api

import (
	"net/http"

	"github.com/fy026/weather/gateway/pkg/e"
	"github.com/fy026/weather/gateway/pkg/util"
	"github.com/gin-gonic/gin"
)

const (
	USERNAME     string = "test01"
	PASSWORD     string = "123456"
	SERVICE_NAME string = "ts"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	isExist := checkAuth(username, password)
	if isExist {
		token, err := util.GenerateToken(username, password)
		if err != nil {
			code = e.ERROR_AUTH_TOKEN
		} else {
			data["token"] = token

			code = e.SUCCESS
		}

	} else {
		code = e.ERROR_AUTH
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func checkAuth(username, password string) bool {
	ok := false
	if username == USERNAME && password == PASSWORD {
		ok = true
	}
	return ok
}
