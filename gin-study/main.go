package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)
func main(){
    r := gin.Default()
    r.Use(gin.Logger())
    r.Use(gin.Recovery())
    // 作用于单个路由
    // r.GET("/benchmark", MyBenchLogger(), benchEndpoint)
    // r.GET("/",func(c *gin.Context){
    //     c.String(200,"Hello,Geektutu")
    // })
    // 作用于某个组
    // authorized := r.Group("/")
    // authorized.Use(AuthRequired())
    r.GET("/",func(c *gin.Context){
        c.String(http.StatusOK,"Who are you?")
    })
    r.GET("/user/:name",func(c *gin.Context){
        name := c.Param("name")
        c.String(http.StatusOK,"Hello %s", name)
    })
    // 获取    query参数
    r.GET("/users",func(c *gin.Context){
        name := c.Query("name")
        role := c.DefaultQuery("role","teacher")
        c.String(http.StatusOK,"%s is a %s",name,role)
    })
    r.POST("/form",func(c *gin.Context){
        username := c.PostForm("username")
        password := c.DefaultPostForm("password","000000") // 可设置默认值
        c.JSON(http.StatusOK,gin.H{
            "username":username,
            "password":password,
        })
    })
    r.POST("/posts",func(c *gin.Context){
        id := c.Query("id")
        page := c.DefaultQuery("page", "0")
        username := c.PostForm("username")
        password := c.DefaultPostForm("username","000000") // 可设置默认值
        c.JSON(http.StatusOK,gin.H{
            "id": id,
            "page":page,
            "username":username,
            "password":password,
        })
    })
    // 重定向
    r.GET("/redirect", func(c *gin.Context){
        c.Redirect(http.StatusMovedPermanently,"/index")
    })
    r.GET("/goindex", func(c *gin.Context){
        c.Request.URL.Path = "/"
        r.HandleContext(c)
    })

    // group routes 分组路由
    defaultHandler := func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "path": c.FullPath(),
        })
    }
    // group: v1
    v1 := r.Group("/v1")
    {
        v1.GET("/posts", defaultHandler)
        v1.GET("/series", defaultHandler)
    }
    // group: v2
    v2 := r.Group("/v2")
    {
        v2.GET("/posts", defaultHandler)
        v2.GET("/series", defaultHandler)
    }
    r.POST("/upload1",func(c *gin.Context){
        file,_ := c.FormFile("./main.go")
        c.String(http.StatusOK,"%s uploaded!",file.Filename)
    })
    r.POST("/upload2",func(c *gin.Context){
        form,_ := c.MultipartForm()
        files := form.File["upload[]"]
        for _,file := range files {
            log.Println(file.Filename)
        }
        c.String(http.StatusOK,"%d files uploaded!", len(files))
    })
    r.Run()
}

