package controller

import (
	"github.com/gofiber/fiber/v2"
	"haemong-be/service"
)

type CommentController struct {
	commentService *service.CommentService
}

func NewCommentController(s *service.CommentService) *CommentController {
	return &CommentController{
		commentService: s,
	}
}

func (r *CommentController) ReadCommentsOnTheFeed(c *fiber.Ctx) error {
	feedId := c.Params("feedId")
	comments, err := r.commentService.ReadCommentsOnTheFeed(feedId)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(200).JSON(comments)
}
