package controllers

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"fintech/database"
	"fintech/models"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

var questionnaireCollection *mongo.Collection = database.GetFormulaireCollection(database.DB)
var validateQuestionnaire = validator.New()

func PostFormulaire(c *fiber.Ctx) error {
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
		"ID":      result.InsertedIDs,
	})
}

func GetFormulaireCollection(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := questionnaireCollection
	filter := bson.M{}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).SendString("Erreur lors de la récupération des données.")
	}
	defer cursor.Close(ctx)

	var result []models.Questionnaire
	for cursor.Next(ctx) {
		var entry models.Questionnaire
		if err := cursor.Decode(&entry); err != nil {
			log.Println(err)
			return c.Status(http.StatusInternalServerError).SendString("Erreur lors du décodage des données.")
		}
		result = append(result, entry)
	}

	if err := cursor.Err(); err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).SendString("Erreur lors de la récupération des données.")
	}

	// Récupérer le format souhaité depuis le paramètre "format" dans l'URL
	format := c.Query("format")

	switch format {
	case "json":
		if err := saveJSON(result); err != nil {
			log.Println(err)
			return c.Status(http.StatusInternalServerError).SendString("Erreur lors de la sauvegarde des données au format JSON.")
		}
		return c.Download("datafile/data.json")

	case "csv":
		if err := saveCSV(result); err != nil {
			log.Println(err)
			return c.Status(http.StatusInternalServerError).SendString("Erreur lors de la sauvegarde des données au format CSV.")
		}
		return c.Download("datafile/data.csv")

	case "sqlite":
		if err := saveSQLite(result); err != nil {
			log.Println(err)
			return c.Status(http.StatusInternalServerError).SendString("Erreur lors de la sauvegarde des données au format SQLite.")
		}
		return c.Download("datafile/data.sqlite")

	default:
		return c.Status(http.StatusBadRequest).SendString("Format non pris en charge.")
	}
}

func saveSQLite(data []models.Questionnaire) error {
	err := os.Mkdir("datafile/", os.ModePerm)
	if err != nil && !os.IsExist(err) {
		return err
	}

	file, err := os.Create("datafile/data.sqlite")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(data)

}

func saveJSON(data interface{}) error {
	err := os.Mkdir("datafile/", os.ModePerm)
	if err != nil && !os.IsExist(err) {
		return err
	}

	file, err := os.Create("datafile/data.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(data)
}

func saveCSV(data []models.Questionnaire) error {
	err := os.Mkdir("datafile/", os.ModePerm)
	if err != nil && !os.IsExist(err) {
		return err
	}

	file, err := os.Create("datafile/data.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Écrivez l'en-tête CSV
	header := []string{"title", "email", "reponse", "question", "createdAt"}
	if err := writer.Write(header); err != nil {
		return err
	}

	// Écrivez les données CSV
	for _, entry := range data {
		record := []string{
			entry.Title,
			entry.Email,
			entry.Reponse,
			entry.Question,
			entry.CreatedAt.String(), // Convertir la date en chaîne
		}
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}

// respondWithFormat renvoie les données au client dans le format demandé
func respondWithFormat(c *fiber.Ctx, data interface{}) error {
	format := c.Query("format") // Obtenez le format demandé depuis la requête

	switch format {
	case "json":
		return c.JSON(data)
	case "csv":
		return respondCSV(c, data)
	case "xml":
		return respondXML(c, data)
	default:
		// Par défaut, renvoyer en JSON
		return c.JSON(data)
	}
}

// respondCSV renvoie les données au client au format CSV
func respondCSV(c *fiber.Ctx, data interface{}) error {
	c.Response().Header.Set("Content-Type", "text/csv")
	c.Response().Header.Set("Content-Disposition", "attachment; filename=data.csv")

	writer := csv.NewWriter(c.Response().BodyWriter())
	defer writer.Flush()

	// Écrivez les données CSV ici en utilisant writer.Write()

	return nil
}

// respondXML renvoie les données au client au format XML
func respondXML(c *fiber.Ctx, data interface{}) error {
	c.Response().Header.Set("Content-Type", "application/xml")
	c.Response().Header.Set("Content-Disposition", "attachment; filename=data.xml")

	encoder := xml.NewEncoder(c.Response().BodyWriter())
	defer encoder.Flush()

	// Encodez les données XML ici en utilisant encoder.Encode()

	return nil
}
