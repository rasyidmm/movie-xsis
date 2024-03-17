package movie

import (
	"context"
	"github.com/pkg/errors"
	"github.com/rasyidmm/movie-xsis/src/adapter/repository/entity"
	"github.com/rasyidmm/movie-xsis/src/shared/mock"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMovieUsecase_GetMovie(t *testing.T) {

	repo := &mock.MovieRepository{}
	Convey("TestMovieUsecase_GetMovie", t, func() {
		Convey("Positive Case", func() {
			repo.On("GetMovie", context.Background(), int64(2)).Return(&entity.MovieEntity{}, nil).Once()
			uc := NewMovieUsecase(repo)
			res, err := uc.GetMovie(context.Background(), int64(2))
			So(err, ShouldBeNil)
			So(res, ShouldNotBeNil)
		})
		Convey("Negative Case", func() {
			repo.On("GetMovie", context.Background(), int64(2)).Return(nil, errors.New("error")).Once()
			uc := NewMovieUsecase(repo)
			res, err := uc.GetMovie(context.Background(), int64(2))
			So(err, ShouldNotBeNil)
			So(res, ShouldBeNil)
		})
	})
}
