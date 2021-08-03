package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scorelist/dao"
	"scorelist/models"
)

func CreateScore(c *gin.Context) {
	var score models.Score
	c.BindJSON(&score)
	err:=models.CreateAScore(&score)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}else{
		//c.JSON(http.StatusOK, score)
		c.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg": "success",
			"data": score,
		})
	}
}

func GetScorelist(c *gin.Context) {
	// 查询todo这个表里的所有数据
	todoList, err := models.GetAllScore()
	if err!= nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}else {
		c.JSON(http.StatusOK, todoList)
	}
}

func GetAScore(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	var score models.Score
	score2 := dao.DB.Where("id=?", id).First(score)
	c.BindJSON(&score2)
	c.JSON(http.StatusOK, gin.H{
		"code": 2000,
		"msg": "success",
		"data": score2,
	})

}
func GetAScore3(c *gin.Context) {
	id := c.Query("id")
	//c.JSON(http.StatusOK,gin.H{
	//	"code": 2000,
	//	"msg": "success",
	//	"id":id,
	//})

	todo := models.GetAScore2(id)
	//if err != nil {
	//	c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	//	return
	//}
	c.BindJSON(&todo)
	c.JSON(http.StatusOK, gin.H{
		"code": 2000,
		"msg": "success",
		"data": todo,
	})

}
func UpdateAScore(c *gin.Context) {
	id := c.Query("id")
	//c.JSON(http.StatusOK,gin.H{
	//	"code": 2000,
	//	"msg": "success",
	//	"id":id,
	//})
	todo, err := models.GetAScore(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.BindJSON(&todo)
	if err = models.UpdateAScore(todo); err!= nil{
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}else{
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteAScore(c *gin.Context) {
	id := c.Query("id")

	if err := models.DeleteAScore(id);err!=nil{
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}else{
		c.JSON(http.StatusOK, gin.H{id:"deleted"})
	}
}

func IsLike(c *gin.Context){
	//chat_id := c.PostForm("chat_id")
	//auth_code := c.PostForm("auth_code")
	name := c.PostForm("name")
	err := models.UpdateMsg(name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
			"data": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg": "请求成功",
		"data": nil,
	})
	return
}
