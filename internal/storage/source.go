package storage

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/samber/lo"
	"news-feed-bot/internal/model"
	"time"
)

type SourcePostgresStorage struct {
	db *sqlx.DB
}

func (s *SourcePostgresStorage) Sources(ctx context.Context) ([]model.Source, error) {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var sources []dbSource
	if err := conn.SelectContext(ctx, &sources, `SELECT * FROM sources`); err != nil {
		return nil, err
	}

	return lo.Map(sources, func(source dbSource, _ int) model.Source { return model.Source(source) }), nil
}

func (s *SourcePostgresStorage) SourceByID(ctx context.Context, id int64) (*model.Source, error) {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
}

func (s *SourcePostgresStorage) Add(ctx context.Context, source model.Source) (int64, error) {

}

func (s *SourcePostgresStorage) Delete(ctx context.Context, id int64) error {

}

type dbSource struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	FeedURL   string    `db:"feed_url"`
	CreatedAt time.Time `db:"created_at"`
}
