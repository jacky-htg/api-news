package repositories

import (
	"database/sql"
	"strconv"

	"github.com/jacky-htg/api-news/libraries"
	"github.com/jacky-htg/api-news/models"
)

func TopicGetList(param map[string]string) ([]models.Topic, error) {
	var query string
	query = "SELECT " +
		"`id`, `title`, `created_at`, `updated_at`" +
		" FROM `topics`"

	if len(param["search"]) > 0 && param["search"] != "" {
		query += " WHERE `title` LIKE '%" + param["search"] + "%'"
	}

	if len(param["sortby"]) > 0 && param["sortby"] != "" {
		query += " ORDER BY `" + param["sortby"] + "`"
		if len(param["order"]) > 0 && param["order"] != "" {
			query += " " + param["order"]
		} else {
			query += " DESC"
		}
	} else {
		if len(param["order"]) > 0 && param["order"] != "" {
			query += " ORDER BY `id` " + param["order"]
		} else {
			query += " ORDER BY `id` DESC"
		}
	}

	var offset int
	limit, _ := strconv.Atoi(param["limit"])
	page, _ := strconv.Atoi(param["page"])
	offset = page*limit - limit

	query += " LIMIT " + param["limit"] + " OFFSET " + strconv.Itoa(offset)

	return topicGetRow(
		db.Query(
			query,
		),
	)
}

func TopicGet(paramsID uint) (models.Topic, error) {
	rows, err := db.Query(
		"SELECT `id`, `title`, `created_at`, `updated_at` FROM `topics` WHERE `id`=?",
		paramsID,
	)
	libraries.CheckError(err)
	if err != nil {
		return models.Topic{}, err
	}

	defer rows.Close()

	return getTopicByRow(rows)
}

func TopicFindFirst() (models.Topic, error) {
	rows, err := db.Query("SELECT `id`, `title`, `created_at`, `updated_at` FROM `topics` ORDER BY id ASC LIMIT 1")
	libraries.CheckError(err)
	if err != nil {
		return models.Topic{}, err
	}

	defer rows.Close()

	return getTopicByRow(rows)
}

func TopicFindLast() (models.Topic, error) {
	rows, err := db.Query("SELECT `id`, `title`, `created_at`, `updated_at` FROM `topics` ORDER BY id DESC LIMIT 1")
	libraries.CheckError(err)
	if err != nil {
		return models.Topic{}, err
	}

	defer rows.Close()

	return getTopicByRow(rows)
}

func TopicIsExist(title string, id uint) (bool, error) {
	var o models.Topic
	var err error
	var isExist bool = false

	if id > 0 {
		o, err = topicIsExistById(title, id)
	} else {
		o, err = topicIsExist(title)
	}

	if err != nil {
		return isExist, err
	}

	if o.ID > 0 {
		isExist = true
	}

	return isExist, nil
}

func TopicStore(o models.Topic) (models.Topic, error) {
	stmt, err := db.Prepare("INSERT INTO topics (`title`) VALUES (?)")
	libraries.CheckError(err)
	if err != nil {
		return models.Topic{}, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(o.Title)
	libraries.CheckError(err)
	if err != nil {
		return models.Topic{}, err
	}

	id, err := res.LastInsertId()
	libraries.CheckError(err)
	if err != nil {
		return models.Topic{}, err
	}

	o.ID = uint(id)

	topic, err := TopicGet(o.ID)
	if err != nil {
		return models.Topic{}, err
	}

	return topic, nil
}

func TopicUpdate(o models.Topic) (models.Topic, error) {
	stmt, err := db.Prepare(
		"UPDATE topics" +
			" SET `title`=?, `updated_at`=NOW()" +
			" WHERE id=?",
	)
	libraries.CheckError(err)
	if err != nil {
		return models.Topic{}, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		o.Title,
		o.ID,
	)
	libraries.CheckError(err)
	if err != nil {
		return models.Topic{}, nil
	}

	o, err = TopicGet(o.ID)
	if err != nil {
		return models.Topic{}, nil
	}

	return o, nil
}

func TopicDestroy(id uint) (bool, error) {
	stmt, err := db.Prepare(
		"DELETE FROM topics WHERE `id`=?",
	)
	libraries.CheckError(err)
	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)
	libraries.CheckError(err)
	if err != nil {
		return false, err
	}

	stmt, err = db.Prepare(
		"DELETE FROM news_topics WHERE `topic_id`=?",
	)
	libraries.CheckError(err)
	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)
	libraries.CheckError(err)
	if err != nil {
		return false, err
	}

	return true, nil
}

func TopicGetByNewsId(news_id uint) ([]models.Topic, error) {
	var topics []models.Topic

	rows, err := db.Query("SELECT `topic_id` FROM `news_topics` WHERE `news_id`=?", news_id)

	libraries.CheckError(err)
	if err != nil {
		return []models.Topic{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var topic models.Topic

		err := rows.Scan(&topic.ID)
		libraries.CheckError(err)
		if err != nil {
			return []models.Topic{}, err
		}

		topic, err = TopicGet(topic.ID)
		libraries.CheckError(err)
		if err != nil {
			return []models.Topic{}, err
		}

		topics = append(topics, topic)
	}

	err = rows.Err()
	if err != nil {
		return []models.Topic{}, err
	}

	return topics, nil
}

func topicGetRow(rows *sql.Rows, err error) ([]models.Topic, error) {
	var topics []models.Topic
	libraries.CheckError(err)

	if err != nil {
		return []models.Topic{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var topic models.Topic

		err := rows.Scan(
			&topic.ID,
			&topic.Title,
			&topic.CreatedAt,
			&topic.UpdatedAt,
		)
		libraries.CheckError(err)
		if err != nil {
			return []models.Topic{}, err
		}

		topics = append(topics, topic)
	}

	err = rows.Err()
	if err != nil {
		return []models.Topic{}, err
	}

	return topics, nil
}

func topicIsExistById(title string, id uint) (models.Topic, error) {
	var o models.Topic
	rows, err := db.Query(
		"SELECT `id` FROM `topics` WHERE `title`=? AND `id`<>?",
		title,
		id,
	)
	libraries.CheckError(err)
	if err != nil {
		return models.Topic{}, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&o.ID)
		libraries.CheckError(err)
		if err != nil {
			return models.Topic{}, err
		}
	}

	err = rows.Err()
	libraries.CheckError(err)
	if err != nil {
		return models.Topic{}, err
	}

	return o, nil
}

func topicIsExist(title string) (models.Topic, error) {
	var o models.Topic

	rows, err := db.Query(
		"SELECT `id` FROM `topics` WHERE `title`=?",
		title,
	)
	libraries.CheckError(err)
	if err != nil {
		return models.Topic{}, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&o.ID)
		libraries.CheckError(err)
		if err != nil {
			return models.Topic{}, err
		}
	}

	err = rows.Err()
	libraries.CheckError(err)
	if err != nil {
		return models.Topic{}, err
	}

	return o, nil
}

func getTopicByRow(rows *sql.Rows) (models.Topic, error) {
	var o models.Topic

	for rows.Next() {
		err := rows.Scan(
			&o.ID,
			&o.Title,
			&o.CreatedAt,
			&o.UpdatedAt,
		)
		libraries.CheckError(err)
		if err != nil {
			return models.Topic{}, err
		}
	}

	err = rows.Err()
	libraries.CheckError(err)
	if err != nil {
		return models.Topic{}, err
	}

	return o, nil
}
