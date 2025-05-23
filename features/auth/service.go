package auth

import (
	"time"

	env "github.com/ghozir/sighapp/config"
	"github.com/ghozir/sighapp/entities"
	authdto "github.com/ghozir/sighapp/features/auth/dto"
	authpresenter "github.com/ghozir/sighapp/features/auth/presenter"
	"github.com/ghozir/sighapp/utils"
	"github.com/ghozir/sighapp/utils/exception"
	"github.com/ghozir/sighapp/utils/logger"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	LoginUser(c *fiber.Ctx, req authdto.LoginRequest) (*authpresenter.LoginResult, error)
}

type authService struct {
	repo Repository
}

func NewAuthService(repo Repository) Service {
	return &authService{repo: repo}
}

func (s *authService) LoginUser(c *fiber.Ctx, req authdto.LoginRequest) (*authpresenter.LoginResult, error) {
	user, err := s.repo.FindUserByEmail(req.Email)
	if err != nil {
		logger.Error("User not found", err)
		return nil, exception.NotFound("user not found")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		logger.Error("Wrong password", err)
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

	if _, err := s.repo.InsertOneSession(session); err != nil {
		logger.Error("Failed to insert session", err)
		return nil, exception.InternalServerError("failed to insert session")
	}

	return &authpresenter.LoginResult{Token: token}, nil
}
