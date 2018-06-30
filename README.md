## vgo（we go）<br>

*vgo是一个轻量级的、简单、稳定、高效、灵活的Web应用程序框架。它旨在使入门快速轻松，快开发和干净实用的设计，并能扩展到复杂的应用程序。<br>  
*vgo采用httprouter路由开发，可以轻松实现数万至数十万个开放链接<br>  
*vgo提供建议，但不强制执行任何依赖或项目布局。开发人员需要选择他们想使用的工具和库。社提供了许多扩展功能，可以轻松添加新功能。<br>  
  
  
<br>   
<br> 

 
#### 适用场景：<br>

<ol>

    api服务应用<br>
    微服务应用<br>
    企业级web应用<br>
    自动化运维相关应用<br>  

</ol>

<br>  

#### 获得更多帮助：<br>
qq群：707362951<br>  
<br>


#### 功能特点: <br>
<ol>

    使用简单、基于httprouter强大路由层<br>
	基于强大的sqlx数据库操作，并且用户可以选择xorm或者gorm两种引擎<br>
	创建强大的redis连接池，用户只需要在配置中配置最大连接数，链接超时时间，链<br>接地址，最大的活动链接即可<br>
	天生支持前后端完全分离协作开发<br>
	支持ini格式解析数据，并且可以单独拿来即用<br>
	支持服务注册、服务发现<br>  

</ol>
<br>  
<br>  
<br>  

####  代码示例  <br>  

#####  目录结构  <br>
project<br>
---controller<br>
---------c.go<br>
---main<br>
---------main.go<br>
---router<br>
---------router.go
----------------------------------
* c.go
```
package controller

import (
   "fmt"
   "net/http"
   "github.com/julienschmidt/httprouter"
   "golang.org/x/net/context"
)


func T(w http.ResponseWriter, r *http.Request,c httprouter.Params) {
   _,cancel:=context.WithCancel(context.Background())
   defer cancel()
   fmt.Fprint(w, "Welcome!\n")
}
```
* main.go
```
package main

import(
   "vgo/vgo"
   _ "vgo/router"
)

func main(){
   vgo.Run(":8080")
}
```
* router.go
```
package router

import(
   "vgo/router"
   "project /controller"
)

func init(){
   router.Router.GET("/",controller.T)
}
```

