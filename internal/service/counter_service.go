package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/blaze-d83/go-GoTTH/internal/repository"
	"github.com/blaze-d83/go-GoTTH/pkg/logger"
)

type CounterService struct {
	db      *sql.DB
	logger  logger.LoggerStrategy
	queries *repository.Queries
}

func NewCounterService(db *sql.DB, logger logger.LoggerStrategy) *CounterService {
	return &CounterService{
		db:      db,
		logger:  logger,
		queries: repository.New(db),
	}

}

func (s *CounterService) GetCounter(ctx context.Context) (int64, error) {
	startTime := time.Now()
	counter, err := s.queries.GetCounter(ctx)
	duration := time.Since(startTime)
	if err != nil {
		s.logger.LogError(nil, err)
	} else {
		s.logger.LogEvent("Fetched counter values")
		s.logger.LogResponses(nil, 200, duration)
	}
	return counter, err
}

func (s *CounterService) IncrementCounter(ctx context.Context) error {
	err := s.queries.IncrementCounter(ctx)
	if err != nil {
		s.logger.LogError(nil, err)
		return err
	}
	s.logger.LogEvent("Counter Increment")
	return nil
}

func (s *CounterService) DecrementCounter(ctx context.Context) error {
	err := s.queries.DecrementCounter(ctx)
	if err != nil {
		s.logger.LogError(nil, err)
		return err
	}
	s.logger.LogEvent("Counter Decrement")
	return nil
}
