package main

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/gin-gonic/gin"
	"tbook_server_bg/configs"
	"tbook_server_bg/daos"
	"tbook_server_bg/handle"
)

// 初始化日志
func initLogger() {
	err := logs.SetLogger(logs.AdapterFile, `{"filename":"logs/mpaa.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
	if err != nil {
		return
	}
}

func main() {
	initLogger() // 初始化日志
	configs.Config()
	configs.ParseConfig("./configs/config.yaml") // 加载 configs 目录中的配置文件
	daos.InitMysql()
	configs.NewRedis()
	r := gin.Default()
	route(r)
	r.Use(handle.Core())
	err := r.Run(":" + configs.Config().Port)
	if err != nil {
		return
	}
}

func route(r *gin.Engine) {
	// 不鉴权接口
	public := r.Group("/api/v1")
	{
		public.GET("/ping", handle.GetPing) // 不鉴权的测试接口 ✅

	}

}
