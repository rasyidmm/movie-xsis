package movie

import (
	"context"
	"github.com/pkg/errors"
	"github.com/rasyidmm/movie-xsis/src/adapter/repository/entity"
	"github.com/rasyidmm/movie-xsis/src/shared/mock"
	"github.com/rasyidmm/movie-xsis/src/usecase/payload"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMovieUsecase_UpdateMovie(t *testing.T) {

	repo := &mock.MovieRepository{}
	Convey("TestMovieUsecase_UpdateMovie", t, func() {
		Convey("Positive Case", func() {
			repo.On("UpdateMovie", context.Background(), &entity.MovieEntity{}).Return(nil).Once()
			uc := NewMovieUsecase(repo)
			err := uc.UpdateMovie(context.Background(), payload.UpdateMoviePayload{})
			So(err, ShouldBeNil)
		})
		Convey("Negative Case", func() {
			repo.On("UpdateMovie", context.Background(), &entity.MovieEntity{}).Return(errors.New("error")).Once()
			uc := NewMovieUsecase(repo)
			err := uc.UpdateMovie(context.Background(), payload.UpdateMoviePayload{})
			So(err, ShouldNotBeNil)
		})
	})
}
