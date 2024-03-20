package movie

import (
	"context"
	"github.com/rasyidmm/movie-xsis/src/adapter/repository/entity"
)

type MovieRepository interface {
	CreateMovie(ctx context.Context, input *entity.MovieEntity) error
	UpdateMovie(ctx context.Context, input *entity.MovieEntity) error
	DeleteMovie(ctx context.Context, id int64) error
	GetMovie(ctx context.Context, id int64) (*entity.MovieEntity, error)
	GetMovies(ctx context.Context) (*[]entity.MovieEntity, error)
	GetMovieByLike(ctx context.Context, title string) (*[]entity.MovieEntity, error)
}
