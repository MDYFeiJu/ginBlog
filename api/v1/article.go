package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 添加文章
func CreateArt(c *gin.Context) {
	var art model.Article
	c.ShouldBindJSON(&art)
	code := model.CreateArt(&art)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    art,
		"message": errmsg.GetErrMsg(code),
	})

}

// 查询单个文章

// 查询文章列表
func FindOneArts(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	arts := model.GetArts(pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    arts,
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑文章
func EditArt(c *gin.Context) {
	// todo
}

// 删除文章
func DeleteArt(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteArt(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
