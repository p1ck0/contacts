package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/p1ck0/contacts/contact"
	"github.com/p1ck0/contacts/database"
)

func init() {
	err := database.CreateDB()
	if err != nil {
		log.Panicln(err)
	}
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/getContacts", contact.GetContacts)
	app.Post("/api/v1/newContact", contact.NewContact)
	app.Post("/api/v1/editContact/:id", contact.EditContact)
	app.Delete("/api/v1/delContact/:id", contact.DeleteContact)
}

func main() {
	app := fiber.New()
	app.Use(cors.New())

	setupRoutes(app)

	log.Fatal(app.Listen(":8888"))
}
