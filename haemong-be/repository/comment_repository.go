package repository

import (
	"database/sql"
	"errors"
)

type CommentRepository struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) *CommentRepository {
	return &CommentRepository{
		db: db,
	}
}

type CommentEntity struct {
	CommentId string `json:"commentId"`
	FeedId    string `json:"feedId"`
	Comment   string `json:"comment"`
	UserId    string `json:"userId"`
	Date      string `json:"date"`
}

func (r *CommentRepository) ReadCommentsOnTheFeed(feedId string) (entityList []CommentEntity, err error) {
	rows, err := r.db.Query("SELECT * FROM tbl_comment where feedId = ?", feedId)
	if err != nil {
		return entityList, errors.New("댓글 불러오기 실패")
	}
	commentEntity := CommentEntity{}
	for rows.Next() {
		err = rows.Scan(&commentEntity.CommentId,
			&commentEntity.FeedId,
			&commentEntity.Comment,
			&commentEntity.UserId,
			&commentEntity.Date)

		entityList = append(entityList, commentEntity)
	}
	return entityList, nil
}
