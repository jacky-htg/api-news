package repositories

import (
	"database/sql"

	"github.com/jacky-htg/api-news/libraries"
	"github.com/jacky-htg/api-news/models"
)

func UserGetByEmail(paramEmail string) (models.User, error) {
	return getUserRow(db.Query("SELECT `id`, `name`, `email`, `password`, `group_id`, `is_active`, `phone_number`, `photo`, `biography`, `birthdate`, `gender`, `created_at`, `updated_at`  FROM `users` WHERE `email`=?", paramEmail))
}

func UserGet(paramId uint) (models.User, error) {
	return getUserRow(db.Query("SELECT `id`, `name`, `email`, `password`, `group_id`, `is_active`, `phone_number`, `photo`, `biography`, `birthdate`, `gender`, `created_at`, `updated_at`  FROM `users` WHERE `id`=?", paramId))
}

func getUserRow(rows *sql.Rows, err error) (models.User, error) {
	var user models.User
	libraries.CheckError(err)

	if err != nil {
		return models.User{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var userNull models.UserNull
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Group.ID, &user.IsActive, &userNull.PhoneNumber, &userNull.Photo, &userNull.Biography, &userNull.Birthdate, &user.Gender, &user.CreatedAt, &user.UpdatedAt)
		libraries.CheckError(err)

		if err != nil {
			return models.User{}, err
		}

		user.PhoneNumber = userNull.PhoneNumber.String
		user.Photo = userNull.Photo.String
		user.Biography = userNull.Biography.String
		user.Birthdate = userNull.Birthdate.Time
	}

	err = rows.Err()
	libraries.CheckError(err)

	if err != nil {
		return models.User{}, err
	}

	user.Group, err = GroupGet(user.Group.ID)
	if err != nil {
		return models.User{}, err
	}

	return user, err
}
