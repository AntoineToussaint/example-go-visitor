package infra

import (
	"context"
	"database/sql"
	"github.com/codefly-dev/go-grpc/base/pkg/business"
	"github.com/codefly-dev/go-grpc/base/pkg/gen"
	_ "github.com/lib/pq"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore(connStr string) (*PostgresStore, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return &PostgresStore{db: db}, nil
}

func (s *PostgresStore) Visit(ctx context.Context) (int, error) {
	today := time.Now().Format("2006-01-02")
	_, err := s.db.ExecContext(ctx, "INSERT INTO store (date, visits) VALUES ($1, 0) ON CONFLICT (date) DO NOTHING", today)
	if err != nil {
		return 0, err
	}

	var visitNumber int
	err = s.db.QueryRowContext(ctx, "UPDATE store SET visits = visits + 1 WHERE date = $1 RETURNING visits", today).Scan(&visitNumber)
	if err != nil {
		return 0, err
	}
	return visitNumber, nil
}

func (s *PostgresStore) GetStatistics(ctx context.Context) (*business.Statistics, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT date, visits FROM store")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var totalVisits int
	var visits []*gen.DailyVisitStatistics
	for rows.Next() {
		var visit gen.DailyVisitStatistics
		var date time.Time
		if err := rows.Scan(&date, &visit.Visits); err != nil {
			return nil, err
		}
		visit.Date = timestamppb.New(date)
		visits = append(visits, &visit)
		totalVisits += int(visit.Visits)
	}
	return &business.Statistics{
		Total:  totalVisits,
		Visits: visits,
	}, nil
}
