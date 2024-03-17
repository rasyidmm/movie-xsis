package movie

import (
	"context"
	"github.com/rasyidmm/movie-xsis/src/adapter/repository/entity"
	"github.com/rasyidmm/movie-xsis/src/usecase/payload"
)

func (uc *MovieUsecase) UpdateMovie(ctx context.Context, input payload.UpdateMoviePayload) error {
	movieEntity := entity.MovieEntity{
		Id:          input.Id,
		Title:       input.Title,
		Description: input.Description,
		Rating:      input.Rating,
		Image:       input.Image,
	}
	err := uc.repo.UpdateMovie(ctx, &movieEntity)
	if err != nil {
		return err
	}
	return nil

}
