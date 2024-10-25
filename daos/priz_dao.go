package daos

import (
	"strconv"
	"tbook_server_bg/models"
	"time"
)

// GetPrizesByUserID retrieves the prizes associated with a specific user by their user ID
func GetPrizesByPrizeId(iD uint) ([]models.Prize, error) {
	var prizes []models.Prize

	// Assuming there is a relationship between User and Prize via a foreign key (user_id)
	err := DB.Where("id = ?", iD).Find(&prizes).Error
	if err != nil {
		return nil, err
	}

	return prizes, nil
}

func GetPrizeList(page string, pageSize string) ([]models.Prize, int, error) {
	// 将字符串转换为整数
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return nil, 0, err
	}
	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		return nil, 0, err
	}

	var prize []models.Prize
	var total int64

	// 查询总数
	if err := DB.Model(&models.Prize{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询用户数据
	if err := DB.Offset((pageInt - 1) * pageSizeInt).Limit(pageSizeInt).Find(&prize).Error; err != nil {
		return nil, 0, err
	}

	return prize, int(total), nil
}

// CreateOrUpdatePrize 创建或更新奖品信息
func CreateOrUpdatePrize(id string, name string, imageURL string, prizeType string, value string, probabilityStr string, quotaStr string) error {
	var existingPrize models.Prize

	// 查找是否已有相同 ID 的记录
	if err := DB.Where("id = ?", id).First(&existingPrize).Error; err != nil {
		// 如果没有找到匹配项，创建新的记录
		probability, err := strconv.ParseFloat(probabilityStr, 64)
		if err != nil {
			return err // 返回转换错误
		}

		quota, err := strconv.ParseInt(quotaStr, 10, 64)
		if err != nil {
			return err // 返回转换错误
		}

		newPrize := models.Prize{
			Name:              name,
			ImageURL:          imageURL,
			Type:              prizeType,
			Value:             value,
			Probability:       probability, // 默认概率
			IsTimeBased:       false,       // 默认不基于时间
			StartTime:         time.Now(),  // 默认开始时间
			EndTime:           time.Now(),  // 默认结束时间
			PlayMode:          "1",         // 默认玩法参数
			IsAutoDistributed: false,       // 默认不自动发放
			Quota:             quota,       // 默认抽中名额限制
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
		}

		if err := DB.Create(&newPrize).Error; err != nil {
			return err
		}
	} else {
		// 如果找到匹配项，更新记录
		probability, err := strconv.ParseFloat(probabilityStr, 64)
		if err != nil {
			return err // 返回转换错误
		}

		quota, err := strconv.ParseInt(quotaStr, 10, 64)
		if err != nil {
			return err // 返回转换错误
		}

		existingPrize.Name = name
		existingPrize.ImageURL = imageURL
		existingPrize.Type = prizeType
		existingPrize.Value = value
		existingPrize.Probability = probability // 更新概率
		existingPrize.Quota = quota             // 更新抽中名额限制
		existingPrize.IsTimeBased = false       // 默认不基于时间
		existingPrize.StartTime = time.Now()    // 默认开始时间
		existingPrize.EndTime = time.Now()      // 默认结束时间
		existingPrize.CreatedAt = time.Now()
		existingPrize.UpdatedAt = time.Now()
		existingPrize.PlayMode = "1"            // 默认玩法参数
		existingPrize.IsAutoDistributed = false // 默认不自动发放

		// 保存更新后的记录
		if err := DB.Save(&existingPrize).Error; err != nil {
			return err
		}
	}

	return nil
}
