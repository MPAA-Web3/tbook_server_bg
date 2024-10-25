package daos

import (
	"strconv"
	"tbook_server_bg/models"
	"tbook_server_bg/response"
)

// 假设 db 是您的数据库连接

func GetUsers(page string, pageSize string) ([]models.User, int, error) {
	// 将字符串转换为整数
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return nil, 0, err
	}
	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		return nil, 0, err
	}

	var users []models.User
	var total int64

	// 查询总数
	if err := DB.Model(&models.User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询用户数据
	if err := DB.Offset((pageInt - 1) * pageSizeInt).Limit(pageSizeInt).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, int(total), nil
}

// GetInviteeByUserID 根据用户ID获取一级邀请者信息
func GetInviteeByUserID(userID string) (string, error) {
	var invitation models.Invitation
	if err := DB.Where("invitee_user_id = ? AND level = ?", userID, 1).First(&invitation).Error; err != nil {
		return "", err // 如果没有找到邀请记录，返回错误
	}
	return invitation.InviterID, nil // 返回邀请者ID
}

// GetTasksByUserID retrieves the list of tasks for a given user by their UserID.
func GetTasksByUserID(userID string) ([]response.Task, error) {
	// 从数据库获取用户
	var user models.User
	if err := DB.Where("user_id = ?", userID).First(&user).Error; err != nil {
		return nil, err // 如果查询失败，返回错误
	}

	// 构建任务列表
	tasks := []response.Task{
		{
			Name:   "Joined Discord",   // 任务名称
			Status: user.JoinedDiscord, // 任务是否已完成
		},
		{
			Name:   "Joined X",
			Status: user.JoinedX,
		},
		{
			Name:   "Joined Telegram",
			Status: user.JoinedTelegram,
		},
		{
			Name:   "Joined Telegram Group",
			Status: user.JoinedTelegramGroup,
		},
	}

	return tasks, nil // 返回任务列表
}
