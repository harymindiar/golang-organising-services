package model

type News struct {
	Id      int
	Content string
}

func GetNews() ([]*News, error) {
	rows, err := db.Query("SELECT * FROM news")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	news := make([]*News, 0)
	for rows.Next() {
		rowNews := new(News)
		err := rows.Scan(&rowNews.Id, &rowNews.Content)
		if err != nil {
			return nil, err
		}
		news = append(news, rowNews)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return news, nil
}
