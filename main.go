package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

var DB *gorm.DB

type Todo struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

func initMySQL() (err error) {
	dsn := "root:26221030@tcp(127.0.0.1:13306)/first?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

func main() {
	//创建链接数据库
	if err := initMySQL(); err != nil {
		panic(err)
	}
	DB.AutoMigrate(&Todo{})

	r := gin.Default()
	r.Static("/static", "static")
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	v1Group := r.Group("v1")
	{
		//添加C
		v1Group.POST("/todo", func(c *gin.Context) {
			var todo Todo
			c.BindJSON(&todo)
			if err := DB.Create(&todo).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{"code": 200,
					"msg":  "success",
					"data": todo})
			}
		})

		//查找R
		v1Group.GET("/todo", func(c *gin.Context) {
			var todoList []Todo
			if err := DB.Find(&todoList).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{"code": 200,
					"msg":  "success",
					"data": todoList})
			}

		})
		v1Group.GET("/todo/:id", func(c *gin.Context) {
		})

		//修改U
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			id, _ := c.Params.Get("id")
			var todo Todo
			DB.Where("id=?", id).First(&todo)
			if err := DB.Where("id=?", id).First(&todo).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "INVALID"})
				return
			}
			c.BindJSON(&todo)
			if err := DB.Save(&todo).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{"code": 200,
					"msg":  "success",
					"data": todo})
			}
		})
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusBadRequest, gin.H{"error": "INVALID"})
				return
			}
			if err := DB.Where("id=?", id).Delete(&Todo{}).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			} else {
				c.JSON(http.StatusOK, gin.H{"code": 200,
					"msg":  "DELETED",
					"data": nil})
			}
		})
	}
	//删除D
	r.Run()
}
