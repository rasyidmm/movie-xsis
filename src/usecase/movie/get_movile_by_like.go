package movie

import (
	"context"
	"github.com/rasyidmm/movie-xsis/src/adapter/repository/entity"
)

func (uc *MovieUsecase) GetMovieByLike(ctx context.Context, title string) (*[]entity.MovieEntity, error) {
	result, err := uc.repo.GetMovieByLike(ctx, title)
	if err != nil {
		return nil, err
	}
	return result, nil

}
