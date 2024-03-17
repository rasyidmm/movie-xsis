package movie

import (
	"context"
	"github.com/rasyidmm/movie-xsis/src/adapter/repository/entity"
	"github.com/rasyidmm/movie-xsis/src/usecase/payload"
)

func (uc *MovieUsecase) CreateMovie(ctx context.Context, input payload.CreateMoviePayload) error {
	movieEntity := entity.MovieEntity{
		Title:       input.Title,
		Description: input.Description,
		Rating:      input.Rating,
		Image:       input.Image,
	}
	err := uc.repo.CreateMovie(ctx, &movieEntity)
	if err != nil {
		return err
	}
	return nil
}
