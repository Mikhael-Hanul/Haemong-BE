package repository

import (
	"database/sql"
	"errors"
)

type FeedRepository struct {
	db *sql.DB
}

type FeedEntity struct {
	feedId   string
	userId   string
	userName string
	title    string
	content  string
}

func NewFeedRepository(db *sql.DB) *FeedRepository {
	return &FeedRepository{
		db: db,
	}
}

func (r *FeedRepository) SaveFeed(feedId, userId, userName, title, content string) error {
	_, err := r.db.Exec("insert into tbl_feed value (?, ?, ?, ?, ?)", feedId, userId, userName, title, content)
	if err != nil {
		return errors.New("피드 등록에 실패함 : " + err.Error())
	}
	return nil
}
