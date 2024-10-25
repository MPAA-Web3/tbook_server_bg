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
	r.Use(handle.Core())
	route(r)
	err := r.Run(":" + configs.Config().Port)
	if err != nil {
		return
	}
}

func route(r *gin.Engine) {
	// 不鉴权接口
	public := r.Group("/api/v1")
	{
		public.GET("/ping", handle.GetPing)               // 不鉴权的测试接口 ✅
		public.GET("/getUser", handle.GetUser)            // 获取用户
		public.GET("/getTask", handle.GetTask)            // 获取任务
		public.GET("/getRewards", handle.GetRewards)      // 获取奖品列表
		public.GET("/getInvitee", handle.Invitee)         // 获取邀请列表invitee ❌
		public.GET("/getCard", handle.Card)               // 获取卡片Cards抽奖次数 积分换抽奖次数
		public.GET("/getSwap", handle.GetSwap)            // 获取Swap ton换积分
		public.GET("/getPrizeList", handle.GetPrizeList)  // 获取奖品配置列表✅
		public.POST("/setPrizeList", handle.SetPrizeList) // 新增或修改奖品配置✅
		public.GET("/getTaskType", handle.GetTaskList)    // 获取任务列表✅
		public.POST("/setTaskType", handle.SetTaskList)   // 配置任务列表✅
		public.GET("/getImgList", handle.GetImgList)      // 获取图片列表✅
		public.POST("/setImgList", handle.SetImgList)     // 配置图片列表✅
		public.GET("/getTaskList", handle.GetTasks)
		public.POST("/setTaskList", handle.SetTasks)
	}
}
