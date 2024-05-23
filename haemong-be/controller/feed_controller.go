package controller

import (
	"github.com/gofiber/fiber/v2"
	"haemong-be/dto/feed/request"
	"haemong-be/dto/feed/response"
	"haemong-be/service"
)

type FeedController struct {
	feedService *service.FeedService
}

func NewFeedController(s *service.FeedService) *FeedController {
	return &FeedController{
		feedService: s,
	}
}

func (r *FeedController) SaveFeed(c *fiber.Ctx) error {
	dto := new(request.CreateFeedReqDTO)
	_ = c.BodyParser(dto)
	err := r.feedService.SaveFeed(dto.UserId, dto.Title, dto.Content)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(201).SendString("글 게시 성공")
}

func (r *FeedController) ReadAllFeeds(c *fiber.Ctx) error {
	feeds, err := r.feedService.ReadAllFeeds()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	if len(feeds) == 0 {
		return c.Status(200).JSON([]response.ReadFeedResDTO{})
	}
	return c.Status(200).JSON(feeds)
}
