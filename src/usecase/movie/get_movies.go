package movie

import (
	"context"
	"github.com/rasyidmm/movie-xsis/src/usecase/payload"
)

func (uc *MovieUsecase) GetMovies(ctx context.Context) (*[]payload.MoviePayload, error) {
	movieEntities, err := uc.repo.GetMovies(ctx)
	if err != nil {
		return nil, err
	}
	var moviePayloads []payload.MoviePayload
	for _, movieEntity := range *movieEntities {
		moviePayloads = append(moviePayloads, payload.MoviePayload{
			Id:          movieEntity.Id,
			Title:       movieEntity.Title,
			Description: movieEntity.Description,
			Rating:      movieEntity.Rating,
			Image:       movieEntity.Image,
			CreatedAt:   movieEntity.CreatedAt,
			UpdatedAt:   movieEntity.UpdatedAt,
		})
	}
	return &moviePayloads, nil

}
