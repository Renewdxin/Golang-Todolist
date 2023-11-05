# Golang-Todolist
最近跟着qimi老师做了一个ToDoList，但由于前端出了点问题，所以都是用postman进行测试

[原项目地址](git@github.com:Q1mi/bubble.git)

# 部分功能展示

删除代办

![Untitled](https://asdxz.oss-cn-beijing.aliyuncs.com/img/202311052235301.png)

查找代办

![Untitled](https://asdxz.oss-cn-beijing.aliyuncs.com/img/202311052235592.png)

下面给出思路

# 思路

其实这是一个很简单的增删改查的实现，要做的就三点：设计路由、设计todo事项的组成、设计增删改查以及返回参数。同时我也从中学到了一点：遇事不决，先写注释！

## 设计路由

设计路由针对这个而言很简单，因为总共四大项CRUD，查询函数里面再分单个和整体，所以就以实际操作命名，然后删除和查询单个后面再添加对应的id，类似于

```go
v1Group.GET("/todo/:id", func(c *gin.Context) {
		var todo Todo
		id := c.Param("id") // 从请求的URL参数中获取ID

		if err := DB.First(&todo, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "success",
				"data": todo,
			})
		}
})
```

## 设计todo事项的组成

这个最简单的三个点就是id，内容，状态（是否删除），所以做出以下设计，后续还可以进行增加内容，例如重复次数，提醒时间点，截止时间，等等

```go
type Todo struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}
```

## 设计增删改查以及返回参数

直接采用gorm框架，不过gorm有一定的改变，例如直接关闭数据库，不用手动实现，其余都在代码当中，这里贴一个初始化函数

```go
func initMySQL() (err error) {
	dsn := "root:pwd0@tcp(127.0.0.1:3306)/database?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}
```



## 反思

这是一个很简单的小demo，只需要一两个小时就可以搞定，但是还是锻炼了我很多，初学总是记不住很多东西，这么一个简单的东西可以很有效的让我把基础再度巩固
