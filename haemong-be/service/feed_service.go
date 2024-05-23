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
	err := r.feedRepo.SaveFeed(uuid.New().String(), userId, title, content)
	if err != nil {
		return err
	}
	return nil
}

func (r *FeedService) ReadAllFeeds() (list []repository.FeedEntity, err error) {
	return r.feedRepo.ReadAllFeeds()
}
