[*back to doc*](https://github.com/malw2020/learn/tree/master/doc#contents)<br>

### Gin

Gin is a HTTP web framework written in Go (Golang). It features a Martini-like API with much better performance -- up to 40 times faster. If you need smashing performance, get yourself some Gin.

from: [*https://github.com/gin-gonic/gin*](https://github.com/gin-gonic/gin)



[↑ top](#gin)
<br><br>


## Quick-start

	package main
	
	import (
		"fmt"
		"net/http"
	
		"gopkg.in/gin-gonic/gin.v1"
	)

	func main() {
		gin.SetMode(gin.DebugMode) //全局设置环境，此为开发环境，线上环境为gin.ReleaseMode
		router := gin.Default()    //获得路由实例
	
		//添加中间件
		router.Use(Middleware)
	
		//注册接口
		router.GET("/simple/get", GetHandler)
		router.POST("/simple/post", PostHandler)
		router.PUT("/simple/put", PutHandler)
		router.DELETE("/simple/delete", DeleteHandler)
	
		//监听端口
		http.ListenAndServe(":8005", router)
	
	}
	
	func Middleware(c *gin.Context) {
		fmt.Println("this is a middleware!")
	}
	
	func GetHandler(c *gin.Context) {
		name, exist := c.GetQuery("name")
		if !exist {
			name = "the name key is not exist!"
		}
	
		age, exist := c.GetQuery("age")
		if !exist {
			age = "the age key is not exist!"
		}
	
		c.Data(http.StatusOK, "text/plain", []byte(fmt.Sprintf("get success! name:%s, age:%s\n", name, age)))
		return
	}
	
	func PostHandler(c *gin.Context) {
	
		type JsonHolder struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
		}
		var json JsonHolder
	
		err := c.BindJSON(&json)
		if err != nil {
			holder := JsonHolder{Id: 2, Name: "ttt"}
			//若返回json数据，可以直接使用gin封装好的JSON方法
			c.JSON(http.StatusOK, holder)
		} else {
			c.JSON(http.StatusOK, JsonHolder{Id: json.Id, Name: json.Name})
		}
	
		return
	}
	
	func PutHandler(c *gin.Context) {
		c.Data(http.StatusOK, "text/plain", []byte("put success!\n"))
		return
	}
	
	func DeleteHandler(c *gin.Context) {
		c.Data(http.StatusOK, "text/plain", []byte("delete success!\n"))
		return
	}

	
[↑ top](#gin)
<br><br>

## Installation
	go get gopkg.in/gin-gonic/gin.v1

[↑ top](#gin)
<br><br>








