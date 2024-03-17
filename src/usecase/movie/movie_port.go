package movie

import (
	"context"
	"github.com/rasyidmm/movie-xsis/src/adapter/repository/movie"
	"github.com/rasyidmm/movie-xsis/src/usecase/payload"
)

type MovieUsecase struct {
	repo movie.MovieRepository
}

func NewMovieUsecase(r movie.MovieRepository) *MovieUsecase {
	return &MovieUsecase{
		repo: r,
	}
}

type MoviePort interface {
	CreateMovie(ctx context.Context, input payload.CreateMoviePayload) error
	UpdateMovie(ctx context.Context, input payload.UpdateMoviePayload) error
	DeleteMovie(ctx context.Context, id int64) error
	GetMovie(ctx context.Context, id int64) (*payload.MoviePayload, error)
	GetMovies(ctx context.Context) (*[]payload.MoviePayload, error)
}
