package movie

import (
	"context"
	"github.com/rasyidmm/movie-xsis/src/usecase/payload"
)

func (uc *MovieUsecase) GetMovie(ctx context.Context, id int64) (*payload.MoviePayload, error) {
	movieEntity, err := uc.repo.GetMovie(ctx, id)
	if err != nil {
		return nil, err
	}
	return &payload.MoviePayload{
		Id:          movieEntity.Id,
		Title:       movieEntity.Title,
		Description: movieEntity.Description,
		Rating:      movieEntity.Rating,
		Image:       movieEntity.Image,
		CreatedAt:   movieEntity.CreatedAt,
		UpdatedAt:   movieEntity.UpdatedAt,
	}, nil

}
