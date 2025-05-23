package problem

import (
	"fmt"
	"time"

	"github.com/ghozir/sighapp/entities"
	problemdto "github.com/ghozir/sighapp/features/problem/dto"
	problempresenter "github.com/ghozir/sighapp/features/problem/presenter"
	"github.com/ghozir/sighapp/utils/logger"
	"github.com/gofiber/fiber/v2"
)

type Service interface {
	InsertProblem(c *fiber.Ctx, req problemdto.InsertProblemRequest) (*problempresenter.InsertProblemResult, error)
}

type problemService struct {
	repo Repository
}

func NewProblemService(repo Repository) Service {
	return &problemService{repo: repo}
}

func (s *problemService) InsertProblem(c *fiber.Ctx, req problemdto.InsertProblemRequest) (*problempresenter.InsertProblemResult, error) {
	var createdBy string
	if user, ok := c.Locals("user").(entities.SafeUser); ok {
		createdBy = user.ID.Hex()
	} else {
		createdBy = fmt.Sprintf("anonim-%d", time.Now().Unix())
	}

	problem := entities.Problem{
		Content:    req.Content,
		CreatedAt:  time.Now(),
		CreatedBy:  createdBy,
		Categories: []string{}, // diisi worker nanti
	}

	data, err := s.repo.InsertOneProblem(problem)
	if err != nil {
		logger.Error("Failed to insert problem", err)
		return nil, err
	}

	return &problempresenter.InsertProblemResult{
		ContentId: data.ID.Hex(),
	}, nil
}
