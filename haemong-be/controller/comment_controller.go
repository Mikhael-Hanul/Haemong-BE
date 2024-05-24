package controller

import (
	"github.com/gofiber/fiber/v2"
	"haemong-be/dto/comment/request"
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

func (r *CommentController) CreateComment(c *fiber.Ctx) error {
	dto := new(request.CreateCommentReqDTO)
	_ = c.BodyParser(dto)
	err := r.commentService.CreateComment(*dto)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(200).SendString("댓글 작성 성공")
}
