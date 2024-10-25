package daos

import (
	"gorm.io/gorm"
	"log"
	"strconv"
	"tbook_server_bg/models"
)

func GetTaskList() (*models.APP, error) {
	var app models.APP

	// 查询 APP 表中的第一条记录
	if err := DB.First(&app).Error; err != nil {
		return nil, err
	}

	return &app, nil
}

// CreateOrUpdateTaskList 根据名称查找并更新相应的字段
func CreateOrUpdateTaskList(name string, value int64) error {
	var app models.APP

	// 查找当前的 APP 记录
	if err := DB.First(&app).Error; err != nil {
		return err // 如果找不到记录，返回错误
	}

	// 根据名称更新相应的字段
	switch name {
	case "TelegramChannel":
		app.TelegramChannelAmount = value
	case "TelegramGroup":
		app.TelegramGroupAmount = value
	case "Twitter":
		app.TwitterAmount = value
	case "Discord":
		app.DiscordAmount = value
	default:
		return nil // 如果名称不匹配，返回 nil
	}

	// 更新 APP 记录
	if err := DB.Save(&app).Error; err != nil {
		log.Println("Failed to update APP:", err)
		return err
	}

	return nil
}

func GetTasks(page string, pageSize string) ([]models.Tasks, int, error) {
	// 将字符串转换为整数
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return nil, 0, err
	}
	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		return nil, 0, err
	}

	var prize []models.Tasks
	var total int64

	// 查询总数
	if err := DB.Model(&models.Tasks{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询用户数据
	if err := DB.Offset((pageInt - 1) * pageSizeInt).Limit(pageSizeInt).Find(&prize).Error; err != nil {
		return nil, 0, err
	}

	return prize, int(total), nil
}

func CreateOrUpdateSetTasks(id int, name, url, taskType, state string) error {
	var task models.Tasks

	// 如果 id 为 0，直接创建新记录
	if id == 8888 {
		newTask := models.Tasks{
			Name:  name,
			Url:   url,
			Type:  taskType,
			State: state,
		}
		if err := DB.Create(&newTask).Error; err != nil {
			log.Println("Error creating new task:", err)
			return err
		}
		return nil
	}

	// 否则，尝试查找现有记录
	if err := DB.Where("id = ?", id).First(&task).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// 未找到记录，创建新的 task
			newTask := models.Tasks{
				Name:  name,
				Url:   url,
				Type:  taskType,
				State: state,
			}
			if err := DB.Create(&newTask).Error; err != nil {
				log.Println("Error creating new task:", err)
				return err
			}
		} else {
			// 其他错误
			log.Println("Error finding task:", err)
			return err
		}
	} else {
		// 如果找到记录，更新 URL 和 Type
		task.Url = url
		task.Name = name
		task.Type = taskType
		task.State = state
		if err := DB.Save(&task).Error; err != nil {
			log.Println("Error updating task:", err)
			return err
		}
	}

	return nil
}
