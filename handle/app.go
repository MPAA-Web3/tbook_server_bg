package handle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"tbook_server_bg/daos"
	"tbook_server_bg/request"
	"tbook_server_bg/response"
)

func GetUser(c *gin.Context) {
	// 获取入参
	page := c.Query("page")         // 获取页码
	pageSize := c.Query("pageSize") // 获取每页大小

	// 从数据库获取用户数据
	users, total, err := daos.GetUsers(page, pageSize)
	if err != nil {
		c.JSON(500, gin.H{"code": 1, "msg": "获取用户数据失败"})
		return
	}

	// 格式化用户数据
	var userResponses []UserResponse
	for _, user := range users {
		// 查询邀请者信息
		invitee, err := daos.GetInviteeByUserID(user.UserID) // 假设存在此函数
		if err != nil {
			invitee = "" // 如果查询失败，设置为空
		}

		userResponses = append(userResponses, UserResponse{
			UserID:     user.UserID,
			Address:    user.Address,
			Balance:    user.Balance,
			CreateTime: user.CreatedAt.Format("2006-01-02T15:04:05Z"), // 格式化时间
			Invitee:    invitee,                                       // 添加邀请者信息
		})
	}

	response := map[string]interface{}{
		"code": 0,
		"data": map[string]interface{}{
			"list":     userResponses, // 使用格式化后的用户数据
			"page":     page,          // 使用入参
			"pageSize": pageSize,      // 使用入参
			"total":    total,         // 使用从数据库获取的总数
		},
		"msg": "success", // 返回消息
	}

	c.JSON(200, response) // 返回 JSON 响应
}

func GetTask(c *gin.Context) {
	// 获取入参
	page := c.Query("page")         // 获取页码
	pageSize := c.Query("pageSize") // 获取每页大小

	// 从数据库获取用户数据
	users, total, err := daos.GetUsers(page, pageSize)
	if err != nil {
		c.JSON(500, gin.H{"code": 1, "msg": "获取用户数据失败"})
		return
	}

	var userTasksResponses []response.UserTasks

	// 遍历用户，获取每个用户的任务
	for _, user := range users {
		// 调用 GetTasksByUserID 函数获取该用户的任务
		tasks, err := daos.GetTasksByUserID(user.UserID)
		if err != nil {
			tasks = []response.Task{} // 如果查询失败，设置为空数组
		}

		userTasksResponses = append(userTasksResponses, response.UserTasks{
			UserID: user.UserID,
			Task:   tasks, // 添加用户的任务数据
		})
	}

	response := map[string]interface{}{
		"code": 0,
		"data": map[string]interface{}{
			"list":     userTasksResponses, // 使用格式化后的用户任务数据
			"page":     page,               // 使用入参
			"pageSize": pageSize,           // 使用入参
			"total":    total,              // 使用从数据库获取的总数
		},
		"msg": "success", // 返回消息
	}
	c.JSON(200, response) // 返回 JSON 响应
}

func GetRewards(c *gin.Context) {
	// 获取入参
	page := c.Query("page")
	pageSize := c.Query("pageSize")

	// 从 physicalPrizes 数据库获取数据
	physicalPrizes, total, err := daos.GetPhysicalPrizes(page, pageSize)
	if err != nil {
		c.JSON(500, gin.H{"code": 1, "msg": "获取奖品数据失败"})
		return
	}

	// 创建返回结构体的切片
	var physicalPrizesResponses []response.Rewards

	for _, phy := range physicalPrizes {
		// 获取用户的奖品数据
		prizes, err := daos.GetPrizesByPrizeId(phy.PrizeId)
		if err != nil {
			// 如果获取奖品失败，跳过此用户，记录日志
			c.Error(err) // 记录错误日志，方便排查
			continue
		}

		// 遍历奖品，转换为响应结构体
		for _, prize := range prizes {
			physicalPrizesResponses = append(physicalPrizesResponses, response.Rewards{
				UserID: phy.UserID,
				Name:   prize.Name,
				Type:   prize.Type,
				Value:  prize.Value,
				Game:   "大转盘",
				Time:   phy.WinTime.Format("2006-01-02T15:04:05Z"), // 假设 `CreatedAt` 表示获取时间
			})
		}
	}

	// 创建返回的数据结构
	response := gin.H{
		"code": 0,
		"data": gin.H{
			"list":     physicalPrizesResponses, // 格式化的用户任务数据
			"page":     page,                    // 入参的页码
			"pageSize": pageSize,                // 入参的每页大小
			"total":    total,                   // 从数据库获取的总数
		},
		"msg": "success", // 返回消息
	}

	// 返回 JSON 响应
	c.JSON(200, response)
}

func Invitee(c *gin.Context) {
	// 获取入参
	page := c.Query("page")         // 获取页码
	pageSize := c.Query("pageSize") // 获取每页大小

	// 从数据库获取用户数据
	users, total, err := daos.GetUsers(page, pageSize)
	if err != nil {
		c.JSON(500, gin.H{"code": 1, "msg": "获取用户数据失败"})
		return
	}
	var inviteeResponses []response.InviteeTasks
	// 遍历用户，获取每个用户的任务
	for _, user := range users {
		//// 调用 GetTasksByUserID 函数获取该用户的任务
		//tasks, err := daos.GetTasksByUserID(user.UserID)
		//if err != nil {
		//	tasks = []response.Task{} // 如果查询失败，设置为空数组
		//}

		inviteeResponses = append(inviteeResponses, response.InviteeTasks{
			UserID:   user.UserID,
			Num:      1,
			Wallet:   "1",
			Twitter:  "1",
			Telegram: "1",
		})
	}

	// 创建返回的数据结构
	response := gin.H{
		"code": 0,
		"data": gin.H{
			"list":     inviteeResponses, // 格式化的用户任务数据
			"page":     page,             // 入参的页码
			"pageSize": pageSize,         // 入参的每页大小
			"total":    total,            // 从数据库获取的总数
		},
		"msg": "success", // 返回消息
	}
	// 返回 JSON 响应
	c.JSON(200, response)
}

func Card(c *gin.Context) {
	// 获取入参
	page := c.Query("page")         // 获取页码
	pageSize := c.Query("pageSize") // 获取每页大小

	// 从数据库获取用户数据
	purchaseRecord, total, err := daos.GetPurchaseRecord(page, pageSize)
	if err != nil {
		c.JSON(500, gin.H{"code": 1, "msg": "获取用户数据失败"})
		return
	}

	// 创建一个用于存储响应的切片
	var purchaseResponses []response.Purchase
	for _, purchase := range purchaseRecord {
		// 将 models.PurchaseRecord 转换为 response.Purchase
		purchaseResponses = append(purchaseResponses, response.Purchase{
			UserID: purchase.UserID,                                      // 从数据库字段获取数据
			Time:   purchase.PurchaseTime.Format("2006-01-02T15:04:05Z"), // 格式化时间
			Points: purchase.PointsSpent * float64(purchase.CardCount),   // 计算积分
			Level:  purchase.Type,                                        // 等级
		})
	}

	// 创建返回的数据结构
	response := gin.H{
		"code": 0,
		"data": gin.H{
			"list":     purchaseResponses, // 使用格式化的用户任务数据
			"page":     page,              // 入参的页码
			"pageSize": pageSize,          // 入参的每页大小
			"total":    total,             // 从数据库获取的总数
		},
		"msg": "success", // 返回消息
	}
	// 返回 JSON 响应
	c.JSON(200, response)
}

func GetSwap(c *gin.Context) {
	// 获取入参
	page := c.Query("page")         // 获取页码
	pageSize := c.Query("pageSize") // 获取每页大小

	// 从数据库获取用户数据
	order, total, err := daos.GetOrder(page, pageSize)
	if err != nil {
		c.JSON(500, gin.H{"code": 1, "msg": "获取用户数据失败"})
		return
	}

	// 创建一个用于存储响应的切片
	var orderResponses []response.Order
	for _, or := range order {
		// 将 models.Order 转换为 response.Order
		orderResponses = append(orderResponses, response.Order{
			UserID: or.UserID,                                   // 用户 ID
			Time:   or.CreatedAt.Format("2006-01-02T15:04:05Z"), // 格式化时间
			Ton:    or.Amount,                                   // TON 金额
			Value:  or.Amount,                                   // 订单值
		})
	}

	// 创建返回的数据结构
	response := gin.H{
		"code": 0,
		"data": gin.H{
			"list":     orderResponses, // 使用格式化的订单数据
			"page":     page,           // 入参的页码
			"pageSize": pageSize,       // 入参的每页大小
			"total":    total,          // 从数据库获取的总数
		},
		"msg": "success", // 返回消息
	}
	// 返回 JSON 响应
	c.JSON(200, response)
}

func GetImgList(c *gin.Context) {
	page := c.Query("page")         // 获取页码
	pageSize := c.Query("pageSize") // 获取每页大小

	// 从数据库获取用户数据
	imageList, total, err := daos.GetImgList(page, pageSize)
	if err != nil {
		c.JSON(500, gin.H{"code": 1, "msg": "获取用户数据失败"})
		return
	}

	// 创建一个用于存储响应的切片
	var imgListResponses []response.ImgList
	for _, or := range imageList {

		imgListResponses = append(imgListResponses, response.ImgList{
			Name: or.Name,
			Url:  or.Url,
		})
	}

	// 创建返回的数据结构
	response := gin.H{
		"code": 0,
		"data": gin.H{
			"list":     imgListResponses, // 使用格式化的订单数据
			"page":     page,             // 入参的页码
			"pageSize": pageSize,         // 入参的每页大小
			"total":    total,            // 从数据库获取的总数
		},
		"msg": "success", // 返回消息
	}
	// 返回 JSON 响应
	c.JSON(200, response)
}

// SetPrizeList
func SetImgList(c *gin.Context) {

	var setImgList request.SetImgList
	if err := c.ShouldBindJSON(&setImgList); err != nil {
		fmt.Println(err.Error())
		c.JSON(500, gin.H{"code": 1, "msg": "Error req:"})
		return
	}

	err := daos.CreateOrUpdateImage(setImgList.Name, setImgList.Url)
	if err != nil {
		log.Println("Error creating image:", err)
		c.JSON(500, gin.H{"code": 1, "msg": "Error creating image:"})
	}

	// 创建返回的简单成功结构
	response := gin.H{
		"code": 0,
		"msg":  "success",
	}

	c.JSON(200, response)

}

// GetTaskList
func GetTaskList(c *gin.Context) {
	// 从数据库获取用户数据
	taskList, err := daos.GetTaskList()
	if err != nil {
		c.JSON(500, gin.H{"code": 1, "msg": "获取Prize数据失败"})
		return
	}
	// 创建一个用于存储响应的任务列表
	taskListResponses := []response.TaskList{
		{
			Name:  "TelegramChannel",
			Value: taskList.TelegramChannelAmount,
		},
		{
			Name:  "TelegramGroup",
			Value: taskList.TelegramGroupAmount,
		},
		{
			Name:  "Twitter",
			Value: taskList.TwitterAmount,
		},
		{
			Name:  "Discord",
			Value: taskList.DiscordAmount,
		},
	}

	// 创建返回的数据结构
	response := gin.H{
		"code": 0,
		"data": gin.H{
			"list": taskListResponses, // 使用格式化的订单数据
		},
		"msg": "success", // 返回消息
	}
	// 返回 JSON 响应
	c.JSON(200, response)

}

// SetTaskList
func SetTaskList(c *gin.Context) {
	var setTaskList request.SetTaskList
	if err := c.ShouldBindJSON(&setTaskList); err != nil {
		fmt.Println(err.Error())
		c.JSON(500, gin.H{"code": 1, "msg": "Error req:"})
		return
	}
	// 将 Value 从 string 转换为 int64
	valueInt64, err := strconv.ParseInt(string(setTaskList.Value), 10, 64)
	if err != nil {
		log.Println("Error parsing Value to int64:", err)
		return // 解析失败时返回，避免后续的数据库操作
	}
	err = daos.CreateOrUpdateTaskList(setTaskList.Name, valueInt64)
	if err != nil {
		log.Println("Error creating or updating prize:", err)
		c.JSON(500, gin.H{"code": 1, "msg": "Error creating or updating prize:"})
		return
	}
	// 创建返回的数据结构
	response := gin.H{
		"code": 0,
		"msg":  "success",
	}

	c.JSON(200, response)

}

// GetPrizeList
func GetPrizeList(c *gin.Context) {
	page := c.Query("page")         // 获取页码
	pageSize := c.Query("pageSize") // 获取每页大小

	// 从数据库获取用户数据
	prizeList, total, err := daos.GetPrizeList(page, pageSize)
	if err != nil {
		c.JSON(500, gin.H{"code": 1, "msg": "获取Prize数据失败"})
		return
	}

	// 创建一个用于存储响应的切片
	var prizeResponses []response.PrizeList
	for _, or := range prizeList {

		prizeResponses = append(prizeResponses, response.PrizeList{
			ID:          or.ID,          // 从 prizeList 中获取的奖品 ID，将根据实际数据替换
			Name:        or.Name,        // 从 prizeList 中获取的奖品名称，将根据实际数据替换
			ImageURL:    or.ImageURL,    // 从 prizeList 中获取的奖品图片 URL，将根据实际数据替换
			Type:        or.Type,        // 从 prizeList 中获取的奖品类型，将根据实际数据替换
			Probability: or.Probability, //中奖概率
			Value:       or.Value,       // 从 prizeList 中获取的奖品值，将根据实际数据替换
			QuotaStr:    strconv.FormatInt(or.Quota, 10),
		})
	}

	// 创建返回的数据结构
	response := gin.H{
		"code": 0,
		"data": gin.H{
			"list":     prizeResponses, // 使用格式化的订单数据
			"page":     page,           // 入参的页码
			"pageSize": pageSize,       // 入参的每页大小
			"total":    total,          // 从数据库获取的总数
		},
		"msg": "success", // 返回消息
	}
	// 返回 JSON 响应
	c.JSON(200, response)

}

// SetPrizeList
func SetPrizeList(c *gin.Context) {
	var setPrizeList request.SetPrizeList
	if err := c.ShouldBindJSON(&setPrizeList); err != nil {
		fmt.Println(err.Error())
		c.JSON(500, gin.H{"code": 1, "msg": "Error req:"})
		return
	}

	// 创建或更新奖品信息
	err := daos.CreateOrUpdatePrize(setPrizeList.ID.String(), setPrizeList.Name, setPrizeList.ImageURL, setPrizeList.PrizeType, setPrizeList.Value, setPrizeList.Probability.String(), setPrizeList.QuotaStr)
	if err != nil {
		log.Println("Error creating or updating prize:", err)
		c.JSON(500, gin.H{"code": 1, "msg": "Error creating or updating prize:"})
		return
	}
	// 创建返回的数据结构
	response := gin.H{
		"code": 0,
		"msg":  "success",
	}

	c.JSON(200, response)

}

// GetTasks
func GetTasks(c *gin.Context) {
	page := c.Query("page")         // 获取页码
	pageSize := c.Query("pageSize") // 获取每页大小

	// 从数据库获取用户数据
	getTasks, total, err := daos.GetTasks(page, pageSize)
	if err != nil {
		c.JSON(500, gin.H{"code": 1, "msg": "获取用户数据失败"})
		return
	}

	// 创建一个用于存储响应的切片
	var getTasksResponses []response.GetTasks
	for _, or := range getTasks {

		getTasksResponses = append(getTasksResponses, response.GetTasks{
			Name:  or.Name,
			Url:   or.Url,
			Type:  or.Type,
			Id:    int(or.ID),
			State: or.State,
		})
	}

	// 创建返回的数据结构
	response := gin.H{
		"code": 0,
		"data": gin.H{
			"list":     getTasksResponses, // 使用格式化的订单数据
			"page":     page,              // 入参的页码
			"pageSize": pageSize,          // 入参的每页大小
			"total":    total,             // 从数据库获取的总数
		},
		"msg": "success", // 返回消息
	}
	// 返回 JSON 响应
	c.JSON(200, response)
}

// SetTasks
func SetTasks(c *gin.Context) {
	var setTasks request.SetTasks
	if err := c.ShouldBindJSON(&setTasks); err != nil {
		fmt.Println(err.Error())
		c.JSON(500, gin.H{"code": 1, "msg": "Error req:"})
		return
	}
	err := daos.CreateOrUpdateSetTasks(setTasks.ID, setTasks.Name, setTasks.Url, setTasks.Type, setTasks.State)
	if err != nil {
		log.Println("Error creating image:", err)
		c.JSON(500, gin.H{"code": 1, "msg": "Error creating image:"})
	}

	// 创建返回的简单成功结构
	response := gin.H{
		"code": 0,
		"msg":  "success",
	}

	c.JSON(200, response)
}
