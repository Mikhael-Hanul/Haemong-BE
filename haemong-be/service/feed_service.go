package service

import (
	"github.com/google/uuid"
	"haemong-be/repository"
)

type FeedService struct {
	feedRepo *repository.FeedRepository
	userRepo *repository.UserRepository
}

func NewFeedService(feedRepo *repository.FeedRepository, userRepo *repository.UserRepository) *FeedService {
	return &FeedService{
		feedRepo: feedRepo,
		userRepo: userRepo,
	}
}

func (r *FeedService) SaveFeed(userId, title, content string) error {
	user, err := r.userRepo.FindUserById(userId)
	if err != nil {
		return err
	}
	err = r.feedRepo.SaveFeed(uuid.New().String(), userId, user.Name, title, content)
	if err != nil {
		return err
	}
	return nil
}
