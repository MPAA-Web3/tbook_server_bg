package daos

import (
	"strconv"
	"tbook_server_bg/models"
)

func GetPurchaseRecord(page string, pageSize string) ([]models.PurchaseRecord, int, error) {
	// 将字符串转换为整数
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return nil, 0, err
	}
	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		return nil, 0, err
	}

	var purchaseRecord []models.PurchaseRecord
	var total int64

	// 查询总数
	if err := DB.Model(&models.PurchaseRecord{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询用户数据
	if err := DB.Offset((pageInt - 1) * pageSizeInt).Limit(pageSizeInt).Find(&purchaseRecord).Error; err != nil {
		return nil, 0, err
	}

	return purchaseRecord, int(total), nil
}
