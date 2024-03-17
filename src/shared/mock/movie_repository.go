package mock

import (
	"context"
	"github.com/rasyidmm/movie-xsis/src/adapter/repository/entity"
	"github.com/stretchr/testify/mock"
)

type MovieRepository struct {
	mock.Mock
}

func (m *MovieRepository) CreateMovie(ctx context.Context, input *entity.MovieEntity) error {
	args := m.Called(ctx, input)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}
func (m *MovieRepository) UpdateMovie(ctx context.Context, input *entity.MovieEntity) error {
	args := m.Called(ctx, input)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}
func (m *MovieRepository) DeleteMovie(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}
func (m *MovieRepository) GetMovie(ctx context.Context, id int64) (*entity.MovieEntity, error) {
	args := m.Called(ctx, id)
	if args.Get(0) != nil {
		return args.Get(0).(*entity.MovieEntity), nil
	}
	return nil, args.Error(1)
}
func (m *MovieRepository) GetMovies(ctx context.Context) (*[]entity.MovieEntity, error) {
	args := m.Called(ctx)
	if args.Get(0) != nil {
		return args.Get(0).(*[]entity.MovieEntity), nil
	}
	return nil, args.Error(1)
}
