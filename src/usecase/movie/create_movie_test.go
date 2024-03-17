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

func TestMovieUsecase_CreateMovie(t *testing.T) {
	repo := &mock.MovieRepository{}
	Convey("TestMovieUsecase_CreateMovie", t, func() {
		Convey("Positive Case", func() {
			repo.On("CreateMovie", context.Background(), &entity.MovieEntity{}).Return(nil).Once()
			uc := NewMovieUsecase(repo)
			err := uc.CreateMovie(context.Background(), payload.CreateMoviePayload{})
			So(err, ShouldBeNil)
		})
		Convey("Negative Case", func() {
			repo.On("CreateMovie", context.Background(), &entity.MovieEntity{}).Return(errors.New("error")).Once()
			uc := NewMovieUsecase(repo)
			err := uc.CreateMovie(context.Background(), payload.CreateMoviePayload{})
			So(err, ShouldNotBeNil)
		})

	})
}
