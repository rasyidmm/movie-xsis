package movie

import "context"

func (uc *MovieUsecase) DeleteMovie(ctx context.Context, id int64) error {
	err := uc.repo.DeleteMovie(ctx, id)
	if err != nil {
		return err
	}
	return nil

}
