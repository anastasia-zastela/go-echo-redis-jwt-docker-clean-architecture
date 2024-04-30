package middleware

import (
	"github.com/anastasia-zastela/api-mc/config"
	"github.com/anastasia-zastela/api-mc/internal/auth"
	"github.com/anastasia-zastela/api-mc/internal/session"
	"github.com/anastasia-zastela/api-mc/pkg/logger"
)

// Middleware manager
type MiddlewareManager struct {
	sessUC  session.UCSession
	authUC  auth.UseCase
	cfg     *config.Config
	origins []string
	logger  logger.Logger
}

// Middleware manager constructor
func NewMiddlewareManager(sessUC session.UCSession, authUC auth.UseCase, cfg *config.Config, origins []string, logger logger.Logger) *MiddlewareManager {
	return &MiddlewareManager{sessUC: sessUC, authUC: authUC, cfg: cfg, origins: origins, logger: logger}
}
