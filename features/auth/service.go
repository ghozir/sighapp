package auth

import (
	"time"

	env "github.com/ghozir/sighapp/config"
	authcontract "github.com/ghozir/sighapp/contracts"
	"github.com/ghozir/sighapp/entities"
	"github.com/ghozir/sighapp/features/auth/dto"
	"github.com/ghozir/sighapp/features/auth/presenter"
	"github.com/ghozir/sighapp/utils"
	"github.com/ghozir/sighapp/utils/exception"
	"github.com/ghozir/sighapp/utils/logger"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	LoginUser(c *fiber.Ctx, req dto.LoginRequest) (*presenter.LoginResult, error)
}

type authService struct {
	repo authcontract.Repository
}

func NewAuthService(repo authcontract.Repository) Service {
	return &authService{repo: repo}
}

func (s *authService) LoginUser(c *fiber.Ctx, req dto.LoginRequest) (*presenter.LoginResult, error) {
	user, err := s.repo.FindUserByEmail(req.Email)
	if err != nil {
		logger.Error("User not found", err)
		return nil, exception.NotFound("user not found")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		logger.Error("Wrong Password", err)
		return nil, exception.Unauthorized("wrong password")
	}

	token, jti, err := utils.GenerateJWT(user.ID.Hex())
	if err != nil {
		logger.Error("Failed to generate token", err)
		return nil, exception.InternalServerError("failed to generate token")
	}

	session := entities.Session{
		UserID:    user.ID.Hex(),
		JTI:       jti,
		UserAgent: c.Get("User-Agent"),
		IP:        c.IP(),
		ExpiresAt: time.Now().Add(env.Config.JWTExpiresIn),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err = s.repo.InsertOneSession(session)
	if err != nil {
		logger.Error("Failed to insert session", err)
		return nil, exception.InternalServerError("failed to insert session")
	}

	return &presenter.LoginResult{
		Token: token,
	}, nil
}
