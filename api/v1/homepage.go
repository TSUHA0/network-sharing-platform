package v1

import (
	"University-Information-Website/model"
	"University-Information-Website/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

//
func GetHomepage(c *gin.Context) {
	data, code := model.Retrieve()
	if code != errmsg.SUCCESS {
		error := errmsg.SetErrorResponse(c.Request.Method, c.Request.URL.Path, code,
			errmsg.GetErrMsg(code))
		c.JSON(http.StatusBadRequest, error)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}
