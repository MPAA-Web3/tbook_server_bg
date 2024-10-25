package daos

import (
	"gorm.io/gorm"
	"strconv"
	"tbook_server_bg/models"
)

func GetImgList(page string, pageSize string) ([]models.Image, int, error) {
	// 将字符串转换为整数
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return nil, 0, err
	}
	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		return nil, 0, err
	}

	var image []models.Image
	var total int64

	// 查询总数
	if err := DB.Model(&models.Image{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询用户数据
	if err := DB.Offset((pageInt - 1) * pageSizeInt).Limit(pageSizeInt).Find(&image).Error; err != nil {
		return nil, 0, err
	}

	return image, int(total), nil
}

// 添加image
func CreateOrUpdateImage(key string, url string) error {
	var existingImage models.Image

	// 查找是否已有相同 Key 的记录
	err := DB.Where("key = ?", key).First(&existingImage).Error
	if err != nil {
		// 如果没有找到匹配项 (ErrRecordNotFound)，创建新的记录
		if err == gorm.ErrRecordNotFound {
			newImage := models.Image{
				Key: key,
				Url: url,
			}
			if err := DB.Create(&newImage).Error; err != nil {
				return err
			}
		} else {
			// 其他数据库错误
			return err
		}
	} else {
		// 如果找到匹配项，更新记录的 URL
		existingImage.Url = url
		if err := DB.Save(&existingImage).Error; err != nil {
			return err
		}
	}

	return nil
}
