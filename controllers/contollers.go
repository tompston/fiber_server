package controllers

import (
	"context"
	"fiber_server/settings/database"
	"fiber_server/utils/response"
	"log"

	"github.com/gofiber/fiber/v2"
)

// define the function that runs the ent migrations
func EntMigrate(c *fiber.Ctx) error {

	client, err := database.GetDbConnEnt()

	if err != nil {
		log.Fatal(err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("Failed creating schema resources: %v", err)
	}

	defer client.Close()

	return response.ResponseSuccess(c, nil, "Migrations run successfully!")
}
