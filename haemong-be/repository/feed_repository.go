package repository

import (
	"database/sql"
	"errors"
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

func (r *FeedRepository) FindFeedByFeedId(feedId string) (FeedEntity, error) {
	e := new(FeedEntity)
	err := r.db.QueryRow("select * from tbl_feed where feedId = ?", feedId).Scan(
		&e.FeedId, &e.UserId, &e.Title, &e.Content, &e.LikeCount, &e.DislikeCount)
	if err != nil {
		return FeedEntity{}, errors.New("존재하지 않는 피드입니다")
	}
	return *e, nil
}

func (r *FeedRepository) AddLike(feedId string) error {
	feed, err := r.FindFeedByFeedId(feedId)
	if err != nil {
		return err
	}
	_, err = r.db.Exec("update tbl_feed set likeCount = ? where feedId = ?", feed.LikeCount+1, feedId)
	if err != nil {
		return errors.New("좋아요 추가에 실패했습니다")
	}
	return nil
}

func (r *FeedRepository) RemoveLike(feedId string) error {
	feed, err := r.FindFeedByFeedId(feedId)
	if err != nil {
		return err
	}
	_, err = r.db.Exec("update tbl_feed set likeCount = ? where feedId = ?", feed.LikeCount-1, feedId)
	if err != nil {
		return errors.New("좋아요 삭제에 실패했습니다")
	}
	return nil
}

func (r *FeedRepository) AddDislike(feedId string) error {
	feed, err := r.FindFeedByFeedId(feedId)
	if err != nil {
		return err
	}
	_, err = r.db.Exec("update tbl_feed set dislikeCount = ? where feedId = ?", feed.DislikeCount+1, feedId)
	if err != nil {
		return errors.New("싫어요 추가에 실패했습니다")
	}
	return nil
}

func (r *FeedRepository) RemoveDislike(feedId string) error {
	feed, err := r.FindFeedByFeedId(feedId)
	if err != nil {
		return err
	}
	_, err = r.db.Exec("update tbl_feed set likeCount = ? where feedId = ?", feed.DislikeCount-1, feedId)
	if err != nil {
		return errors.New("싫어요 삭제에 실패했습니다")
	}
	return nil
}

func (r *FeedRepository) ReedFeedsOrderPopularity() (list []FeedEntity, err error) {
	rows, err := r.db.Query("SELECT * FROM tbl_feed order by likeCount desc limit 10")
	if err != nil {
		return []FeedEntity{}, errors.New("인기순 피드 불러오기 실패")
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
