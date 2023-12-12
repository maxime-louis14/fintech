package controller

import (
	"context"
	"fintech/database"
	"fintech/model"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"

	"time"
)

var controllerQuestion *mongo.Collection = database.GetCollection(database.DB, "Sondage")
var validate = validator.New()

func PostForm(c *fiber.Ctx) error {
	// Parse le corps de la requête dans une structure de modèle
	var response model.Response
	if err := c.BodyParser(&response); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Erreur de parsing du corps de la requête",
		})
	}

	// Valide la structure du modèle
	if err := validate.Struct(response); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Ajoute la date de création
	response.CreatedAt = time.Now()

	// Insère la réponse dans la base de données
	result, err := controllerQuestion.InsertOne(context.Background(), response)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erreur lors de l'insertion dans la base de données",
		})
	}

	// Retourne la réponse avec l'ID généré
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Réponse ajoutée avec succès",
		"ID":      result.InsertedID,
	})
}
