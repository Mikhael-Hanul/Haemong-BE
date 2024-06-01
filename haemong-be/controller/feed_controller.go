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

func (r *FeedController) AddLike(c *fiber.Ctx) error {
	feedId := c.Params("feedId")
	err := r.feedService.AddLike(feedId)
	if err != nil {
		return c.Status(500).JSON(map[string]string{"message": err.Error()})
	}
	return c.Status(200).JSON(map[string]string{"message": "좋아요 추가를 완료했습니다."})
}

func (r *FeedController) RemoveLike(c *fiber.Ctx) error {
	feedId := c.Params("feedId")
	err := r.feedService.RemoveLike(feedId)
	if err != nil {
		return c.Status(500).JSON(map[string]string{"message": err.Error()})
	}
	return c.Status(200).JSON(map[string]string{"message": "좋아요 삭제를 완료했습니다."})
}

func (r *FeedController) AddDislike(c *fiber.Ctx) error {
	feedId := c.Params("feedId")
	err := r.feedService.AddDislike(feedId)
	if err != nil {
		return c.Status(500).JSON(map[string]string{"message": err.Error()})
	}
	return c.Status(200).JSON(map[string]string{"message": "싫어요 추가를 완료했습니다."})
}

func (r *FeedController) RemoveDislike(c *fiber.Ctx) error {
	feedId := c.Params("feedId")
	err := r.feedService.RemoveDislike(feedId)
	if err != nil {
		return c.Status(500).JSON(map[string]string{"message": err.Error()})
	}
	return c.Status(200).JSON(map[string]string{"message": "싫어요 삭제를 완료했습니다."})
}
