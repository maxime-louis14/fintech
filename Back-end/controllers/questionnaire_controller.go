package controllers

import (
	"context"
	"fintech/database"
	"fintech/models"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

var questionnaireCollection *mongo.Collection = database.GetQuestionnaireCollection(database.DB)
var validateQuestionnaire = validator.New()

func PostQuestionnaires(c *fiber.Ctx) error {
	// Parse le corps de la requête dans une structure de modèle
	var responses []models.Questionnaire
	if err := c.BodyParser(&responses); err != nil {
		fmt.Println("Erreur de parsing du corps de la requête:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Erreur de parsing du corps de la requête",
		})
	}

	// Valide la structure du modèle pour chaque réponse
	for i := range responses {
		if err := validateQuestionnaire.Struct(responses[i]); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		// Ajoute la date de création
		responses[i].CreatedAt = time.Now()
	}

	// Créer des documents à partir des réponses
	var documents []interface{}
	for _, response := range responses {
		documents = append(documents, response)
	}

	// Insère les documents dans la collection "questionnaire"
	result, err := questionnaireCollection.InsertMany(context.Background(), documents)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erreur lors de l'insertion dans la base de données",
		})
	}

	// Retourne la réponse avec les IDs générés
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Réponses ajoutées avec succès",
		"ID":     result.InsertedIDs,
	})
}
