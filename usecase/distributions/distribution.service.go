package distributions

import (
	"context"

	logger "github.com/hthl85/aws-lambda-logger"
	"github.com/hthl85/aws-vanguard-etf-ca-norm-dividends/entities"
)

// Service sector
type Service struct {
	repo Repo
	log  logger.ContextLog
}

// NewService create new service
func NewService(repo Repo, log logger.ContextLog) *Service {
	return &Service{
		repo: repo,
		log:  log,
	}
}

// FindFundDistributions finds all fund distributions
func (s *Service) FindFundDistributions(ctx context.Context) ([]*entities.FundDistribution, error) {
	s.log.Info(ctx, "find all fund distributions")
	return s.repo.FindFundDistributions(ctx)
}
