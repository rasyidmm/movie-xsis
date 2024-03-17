package di

import (
	"github.com/rasyidmm/movie-xsis/src/adapter/db"
	repoMovie "github.com/rasyidmm/movie-xsis/src/adapter/repository/movie"
	"github.com/rasyidmm/movie-xsis/src/usecase/movie"
	"github.com/sarulabs/di/v2"
)

type Container struct {
	ctn di.Container
}

func NewContainer() *Container {
	builder, _ := di.NewBuilder()
	_ = builder.Add([]di.Def{
		{Name: "movie", Build: movieUsecase},
	}...)
	return &Container{
		ctn: builder.Build(),
	}
}

func (c *Container) Get(name string) interface{} {
	return c.ctn.Get(name)

}
func movieUsecase(_ di.Container) (interface{}, error) {
	repo := repoMovie.NewMovieDataHandler(db.DB)
	return movie.NewMovieUsecase(repo), nil
}
