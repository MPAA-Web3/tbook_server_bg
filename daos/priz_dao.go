package daos

import (
	"tbook_server_bg/models"
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
