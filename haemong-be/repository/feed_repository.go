package repository

import (
	"database/sql"
	"errors"
	"fmt"
)

type FeedRepository struct {
	db *sql.DB
}

type FeedEntity struct {
	FeedId       string
	UserId       string
	Title        string
	Content      string
	LikeCount    int
	DislikeCount int
}

func NewFeedRepository(db *sql.DB) *FeedRepository {
	return &FeedRepository{
		db: db,
	}
}

func (r *FeedRepository) SaveFeed(feedId, userId, title, content string) error {
	_, err := r.db.Exec("insert into tbl_feed value (?, ?, ?, ?, ?, ?)", feedId, userId, title, content, 0, 0)
	if err != nil {
		return errors.New("피드 등록에 실패함 : " + err.Error())
	}
	return nil
}

func (r *FeedRepository) ReadAllFeeds() (list []FeedEntity, err error) {
	rows, err := r.db.Query("SELECT * FROM tbl_feed")
	if err != nil {
		fmt.Println(err)
		return []FeedEntity{}, errors.New("피드 불러오기 실패")
	}
	feedEntity := FeedEntity{}
	for rows.Next() {
		err = rows.Scan(&feedEntity.FeedId,
			&feedEntity.UserId,
			&feedEntity.Title,
			&feedEntity.Content,
			&feedEntity.LikeCount,
			&feedEntity.DislikeCount)

		list = append(list, feedEntity)
	}
	return list, nil
}
