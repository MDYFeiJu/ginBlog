package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 添加类型

func AddCategory(c *gin.Context) {
	var cg model.Category
	c.ShouldBindJSON(&cg)
	code := model.CreateCategory(&cg)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    cg,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询单个分类下的文章
//func FindCategorys(c *gin.Context) {
//	arts := model.FindCategory()
//}

// 查询分类是否存在

// 查询分类列表

func GetCategorys(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	Categorys := model.GetCategory(pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    Categorys,
		"message": errmsg.GetErrMsg(code),
	})
}

//
// 编辑分类

func EditCategory(c *gin.Context) {
	var cg model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&cg)
	code := model.CheckCategory(cg.Name)
	if code == errmsg.SUCCESS {
		model.EditCategory(id, &cg)
	}
	if code == errmsg.ErrorCategoryNameUsed {
		c.Abort()
	}
}

// 删除分类

func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteCategory(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
