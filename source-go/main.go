package main

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/arienmalec/alexa-go"
	"github.com/aws/aws-lambda-go/lambda"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	log.Printf("Main function started")
	ctx := context.Background()
	connectionUri := "mongodb+srv://" + os.Getenv("iamKey") + ":" + os.Getenv("iamSecret") + "@alexatestskill.jufzbcg.mongodb.net/?authSource=%24external&authMechanism=MONGODB-AWS&retryWrites=true&w=majority&appName=AlexaTestSkill"
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionUri))

	if err != nil {
		log.Printf("Error connecting to DB: " + err.Error())
		panic(err)
	}

	defer client.Disconnect(ctx)

	connection := Connection{
		collection: client.Database("RecetardasPiolas").Collection("recipes"),
	}

	lambda.Start(connection.IntentDispatcher)
}

func (connection Connection) IntentDispatcher(ctx context.Context, request alexa.Request) (alexa.Response, error) {
	log.Printf("Intent Dispatcher started")
	var response alexa.Response
	switch request.Body.Intent.Name {
	case "GetRecipeIngredients":
		var recipe Recipe
		recipeName := request.Body.Intent.Slots["recipe"].Value
		if recipeName == "" {
			return alexa.NewSimpleResponse("Sin Nombre", "El nombre de la receta esta vacio"), errors.New("Nombre de receta vacio")
		}
		filter := bson.M{
			"title": bson.M{"$regex": recipeName, "$options": "i"},
		}
		if err := connection.collection.FindOne(ctx, filter).Decode(&recipe); err != nil {
			return alexa.NewSimpleResponse("Error", "Error al buscar la receta"), err
		}
		message := "Por supuesto, aqui tienes una lista de ingredientes para preparar " + recipe.Title + ". " + recipe.Ingredients
		response = alexa.NewSimpleResponse("Ingredientes", message)
	case "GetRandomRecipe":
		pipeline := []bson.D{
			{{Key: "$sample", Value: bson.D{{Key: "size", Value: 1}}}},
		}
		cursor, err := connection.collection.Aggregate(ctx, pipeline)
		if err != nil {
			return alexa.NewSimpleResponse("Error", "Error al buscar la receta"), err
		}
		defer cursor.Close(ctx)
		var result bson.M
		if !cursor.Next(ctx) {
			return alexa.NewSimpleResponse("Error", "Base de datos vacia"), err
		}
		if err := cursor.Decode(&result); err != nil {
			return alexa.NewSimpleResponse("Error", "Error al decodificar la receta"), err
		}
		marshaledBytes, err := bson.Marshal(result)
		if err != nil {
			return alexa.NewSimpleResponse("Error", "Error transformando la receta"), err
		}
		var recipe Recipe
		if err := bson.Unmarshal(marshaledBytes, &recipe); err != nil {
			return alexa.NewSimpleResponse("Error", "Error parseando la receta"), err
		}
		message := "Te sugiero que pruebes " + recipe.Title
		response = alexa.NewSimpleResponse("Receta", message)
	case "GetRecipeFromIngredient":
		ingredient := request.Body.Intent.Slots["ingredient"].Value
		filter := bson.M{
			"ingredients": bson.M{"$regex": ingredient, "$options": "i"},
		}

		cursor, err := connection.collection.Find(ctx, filter)
		if err != nil {
			return alexa.NewSimpleResponse("Error", "Ocurrio un error"), err
		}
		defer cursor.Close(ctx)

		limit := 10
		count := 0
		var recipeTitles string
		for cursor.Next(ctx) {
			var recipe Recipe
			if err := cursor.Decode(&recipe); err != nil {
				return alexa.NewSimpleResponse("Error", "Ocurrio un error al decodificar una receta"), err
			}
			if count > 0 {
				recipeTitles += ", "
			}
			recipeTitles += recipe.Title
			count++
			if count >= limit {
				break
			}
		}

		message := "Con " + ingredient + " puedes hacer las siguientes recetas: " + recipeTitles
		response = alexa.NewSimpleResponse("Recetas", message)
	default:
		response = alexa.NewSimpleResponse("Solicitud desconocida", "La solicitud no es conocida")
	}
	log.Printf("Intent dispatcher success!")
	return response, nil
}

// Stores a handle to the collection being used by the Lambda function
type Connection struct {
	collection *mongo.Collection
}

// A data structure representation of the collection schema
type Recipe struct {
	ID          primitive.ObjectID `bson:"_id"`
	Title       string             `bson:"title"`
	Url         string             `bson:"url"`
	Ingredients string             `bson:"ingredients"`
	Steps       string             `bson:"steps"`
	Uuid        string             `bson:"uuid"`
}
