package server

import (
	"strconv"
	"url-shortener/model"
	"url-shortener/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func getAllRedirects(c *fiber.Ctx) error{
	golies, err := model.GetAllGolies()
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error getting all goly links" + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(golies)
}

func getGoly(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not parse id: " + err.Error(),
		})
	}

	goly, err := model.GetGoly(uint64(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not retrieve goly from db: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(goly)
}


func createGoly(c *fiber.Ctx) error{
	c.Accepts("application/json")

	var goly model.Goly
	err := c.BodyParser(&goly)
	if err!=nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message" : "error parsing JSON" + err.Error(),
		})
}
	if goly.Random{
		goly.Goly = utils.RandomURL(8)
	}
	err = model.CreateGoly(goly)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"could not create goly in db" + err.Error(),
	})
}
	return c.Status(fiber.StatusOK).JSON(goly)
}

func updateGoly(c *fiber.Ctx) error{
	c.Accepts("application/json")

	var goly model.Goly

	err := c.BodyParser(&goly)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not parse json" + err.Error(),
		})
	}	
	err = model.UpdateGoly(goly)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"could not update goly link in DB" + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(goly)
}

func deleteGoly(c *fiber.Ctx) error{
	id, err := strconv.ParseUint(c.Params(("id"),)
}

func SetupAndList(){

	router := fiber.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	
	router.Get("/goly", getAllRedirects)
	router.Get("/goly/:id", getGoly)
	router.Post("/goly", createGoly)
	router.Patch("/goly", updateGoly)
	router.Delete("/goly/:id", deleteGoly)

	router.Listen(":3000")
}