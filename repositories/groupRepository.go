package repositories

import (
	"github.com/jacky-htg/api-news/libraries"
	"github.com/jacky-htg/api-news/models"
)

func GroupGet(paramId uint) (models.Group, error) {

	var group models.Group
	rows, err := db.Query("SELECT `id`, `title`, `created_at`, `updated_at`  FROM `groups` WHERE `id`=?", paramId)
	libraries.CheckError(err)

	if err != nil {
		return models.Group{}, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&group.ID, &group.Title, &group.CreatedAt, &group.UpdatedAt)
		libraries.CheckError(err)

		if err != nil {
			return models.Group{}, err
		}
	}
	err = rows.Err()
	libraries.CheckError(err)

	if err != nil {
		return models.Group{}, err
	}

	return group, err
}
