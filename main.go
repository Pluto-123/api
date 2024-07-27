package main

import (
	g "api_project/global"
	"api_project/middleware"
	"api_project/routes"
	"flag"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	//wait-for-it 太麻烦了，直接延迟两秒启动等待mysql部署完成
	time.Sleep(time.Second * 2)

	configPath := flag.String("c", "config.yml", "配置文件路径")
	flag.Parse()

	// 根据命令行参数读取配置文件, 其他变量的初始化依赖于配置文件对象
	conf := g.ReadConfig(*configPath)

	// 初始化数据库
	db := g.InitDatabase(conf)

	r := gin.New()
	r.SetTrustedProxies([]string{"*"})
	// 使用gin自带的日志和恢复中间件
	r.Use(gin.Logger(), gin.Recovery())

	r.Use(middleware.CORS())
	r.Use(middleware.WithGormDB(db))

	routes.RegisterHandlers(r)
	r.Run(conf.Server.Port)
}
