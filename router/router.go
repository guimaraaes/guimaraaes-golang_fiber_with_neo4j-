package router

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/guimaraaes/golang_fiber_with_neo4j/handler"
)

func Routes(app *fiber.App) {
	//cors
	app.Use(cors.New())

	//swagger
	app.Use("/swagger", swagger.Handler)

	//hello world
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	//movie
	app.Get("/movie", handler.GetMovie)
	app.Get("/movie/:title/:released", handler.GetMovieId)
	app.Post("/movie", handler.PostMovie)
	app.Put("/movie/:title/:released", handler.PutMovie)
	app.Delete("/movie/:title/:released", handler.DeleteMovie)

	//person
	app.Get("/person", handler.GetPerson)
	app.Get("/person/:name/:born", handler.GetPersonId)
	app.Post("/person", handler.PostPerson)
	app.Post("/person/with_relationship", handler.PostPersonWithRelationship)
	app.Post("/person/GETwith_relationship", handler.GetPersonRel)
	app.Put("/person/:name/:born", handler.PutPerson)
	app.Delete("/person/:name/:born", handler.DeletePerson)

	//algorithms
	app.Get("/algo_centrality/:node/:relationship", handler.Centrality)
	app.Get("/algo_community/:node/:relationship", handler.Community)
	app.Get("/algo_path/:node/:relationship", handler.Path)
	app.Get("/algo_pagerank/:node/:relationship", handler.PageRank)

}
