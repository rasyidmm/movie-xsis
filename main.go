package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "github.com/rasyidmm/movie-xsis/docs"
	"github.com/rasyidmm/movie-xsis/internal/config"
	"github.com/rasyidmm/movie-xsis/src/infrastructure/router"
	serviceMovie "github.com/rasyidmm/movie-xsis/src/infrastructure/services"
	container "github.com/rasyidmm/movie-xsis/src/shared/di"
	"github.com/rasyidmm/movie-xsis/src/usecase/movie"
)

// @title           movie-xsis
// @version         1.0
// @description     This is a backend service for movie-xsis
// @termsOfService  http://swagger.io/terms/

// @contact.name   rasyid
// @contact.email  rasyidmaulidmajid@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host
// @BasePath

func main() {
	app := fiber.New(
		fiber.Config{
			AppName: config.GetConfig().Server.Rest.Host,
		})
	app.Get("/swagger/*", swagger.HandlerDefault)
	v := app.Group("/v1")
	ctn := container.NewContainer()

	Apply(v, ctn)
	err := app.Listen(":" + string(config.GetConfig().Server.Rest.Port))
	if err != nil {
		panic(err)
	}
}

func Apply(f fiber.Router, ctn *container.Container) {
	router.NewMovie(f, serviceMovie.NewMovieService(ctn.Get("movie").(*movie.MovieUsecase)))
}
