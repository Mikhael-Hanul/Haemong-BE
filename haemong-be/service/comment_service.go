package service

import (
	"errors"
	"haemong-be/dto/comment/request"
	"haemong-be/dto/comment/response"
	"haemong-be/repository"
	"time"
)

type CommentService struct {
	commentRepo *repository.CommentRepository
}

func NewCommentService(commentRepo *repository.CommentRepository) *CommentService {
	return &CommentService{
		commentRepo: commentRepo,
	}
}

func (r *CommentService) ReadCommentsOnTheFeed(feedId string) (list []response.ReadCommentResDTO, err error) {
	entities, err := r.commentRepo.ReadCommentsOnTheFeed(feedId)
	if err != nil {
		return nil, err
	}
	if len(entities) == 0 {
		return []response.ReadCommentResDTO{}, nil
	}
	for _, v := range entities {
		a := response.ReadCommentResDTO{}
		a.CommentId = v.CommentId
		a.Comment = v.Comment
		a.Date = v.Date
		a.UserId = v.UserId
		list = append(list, a)
	}
	return list, nil
}

func (r *CommentService) CreateComment(dto request.CreateCommentReqDTO) error {
	err := r.commentRepo.CreateComment(dto.FeedId, dto.Comment, dto.UserId, time.Now().Format(time.DateTime))
	if err != nil {
		return errors.New("댓글 작성 실패")
	}
	return nil
}
