package repositories

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jacky-htg/api-news/libraries"
)

func AuthCheck(paramEmail string, controller string, method string) (bool, error) {
	var isAuth = false
	var err error

	rows, err := db.Query(""+
		"SELECT `users`.`id` "+
		"FROM `users` "+
		"JOIN `groups` ON (`users`.`group_id` = `groups`.`id`) "+
		"JOIN `access_groups` ON (`groups`.id = `access_groups`.`group_id`) "+
		"JOIN `access` ON (`access_groups`.`access_id` = `access`.`id`) "+
		"WHERE `users`.`email` = ? AND `users`.`is_active` = 1 "+
		"AND (`access`.`name` = 'root' OR `access`.`name` = ? OR `access`.`name` = ? ) "+
		"", paramEmail, controller, controller+"."+method)

	libraries.CheckError(err)

	if err != nil {
		return false, err
	}

	for rows.Next() {
		var id uint
		err = rows.Scan(&id)
		libraries.CheckError(err)

		if err != nil {
			return false, err
		} else {
			if id > 0 {
				isAuth = true
			}
		}
	}

	return isAuth, err
}
