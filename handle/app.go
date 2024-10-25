package handle

import (
	"github.com/gin-gonic/gin"
	"log"
	"tbook_server_bg/daos"
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
			Name: or.Key,
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
func SetPrizeList(c *gin.Context) {
	key := c.Query("key") // 获取页码
	url := c.Query("url") // 获取每页大小

	err := daos.CreateOrUpdateImage(key, url)
	if err != nil {
		log.Println("Error creating image:", err)
		c.JSON(500, gin.H{"code": 1, "msg": "Error creating image:"})
	}

	c.JSON(200, "true")
}

// GetTaskList
func GetTaskList(c *gin.Context) {
	page := c.Query("page")         // 获取页码
	pageSize := c.Query("pageSize") // 获取每页大小

}

// SetTaskList
func SetTaskList(c *gin.Context) {

}
