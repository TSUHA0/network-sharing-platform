package v1

import (
	"University-Information-Website/middleware"
	"University-Information-Website/model"
	"University-Information-Website/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddComment(c *gin.Context) {
	var data model.SingleComment
	CoursesId := c.Param("courses_id")
	LessonId := c.Param("lesson_id")
	if err := c.ShouldBindJSON(&data); err != nil {
		error := errmsg.SetErrorResponse(c.Request.Method, c.Request.URL.Path, http.StatusBadRequest,
			errmsg.GetErrMsg(errmsg.PARSEBODYFAIL))
		c.JSON(http.StatusBadRequest, error)
		return
	}

	code := model.InsertComment(&data, CoursesId, LessonId)
	if code != errmsg.SUCCESS {
		error := errmsg.SetErrorResponse(c.Request.Method, c.Request.URL.Path, code,
			errmsg.GetErrMsg(code))
		c.JSON(http.StatusBadRequest, error)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

func DeleteComment(c *gin.Context) {
	CoursesId := c.Param("courses_id")
	LessonId := c.Param("lesson_id")
	CommentId := c.Param("comment_id")
	UserID := c.Param("user_id")
	tokenHeader := c.Request.Header.Get("Authorization")
	userId, _, role, code := middleware.ParseToken(tokenHeader)

	if code != errmsg.SUCCESS {
		code = errmsg.ERROR_USER_DEL_ERROR
		error := errmsg.SetErrorResponse(c.Request.Method, c.Request.URL.Path, code,
			errmsg.GetErrMsg(code))
		c.JSON(http.StatusBadRequest, error)
		return
	}

	if role >= 2 && userId == UserID {
		code = model.DeleteComment(CoursesId, LessonId, CommentId)
	} else if role < 2 && userId != UserID {
		code = model.DeleteComment(CoursesId, LessonId, CommentId)
	} else {
		code = errmsg.ERROR_USER_NOT_RIGHT
		error := errmsg.SetErrorResponse(c.Request.Method, c.Request.URL.Path, code,
			errmsg.GetErrMsg(code))
		c.JSON(http.StatusBadRequest, error)
		return
	}

	if code != errmsg.SUCCESS {
		error := errmsg.SetErrorResponse(c.Request.Method, c.Request.URL.Path, code,
			errmsg.GetErrMsg(code))
		c.JSON(http.StatusBadRequest, error)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
