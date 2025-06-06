package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NOTE: This is a demo project. The code is not production ready and omits best practices for clarity and brevity.
// For example, error handling, input validation, and security are minimal or missing.

var mongoClient *mongo.Client
var mongoDB *mongo.Database

// Handler for listing all collections (datasets) in a given database.
// NOTE: In production, you should validate dbName and sanitize user input.
func listDatasetsHandler(c *gin.Context) {
	dbName := c.Param("db")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	db := mongoClient.Database(dbName)
	collections, err := db.ListCollectionNames(ctx, struct{}{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"datasets": collections})
}

// Handler for fetching all documents from a collection (dataset) in a given database.
// NOTE: In production, you should validate dbName and collName, and add pagination and security.
func getDatasetHandler(c *gin.Context) {
	dbName := c.Param("db")
	collName := c.Param("dataset")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db := mongoClient.Database(dbName)
	coll := db.Collection(collName)
	cur, err := coll.Find(ctx, struct{}{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cur.Close(ctx)
	var results []map[string]interface{}
	for cur.Next(ctx) {
		var doc map[string]interface{}
		if err := cur.Decode(&doc); err == nil {
			results = append(results, doc)
		}
	}
	c.JSON(http.StatusOK, results)
}

// Helper to sanitize names for MongoDB (replace spaces and dashes with underscores)
// NOTE: In production, use a stricter sanitizer and validate input.
func sanitizeName(name string) string {
	name = strings.ReplaceAll(name, " ", "_")
	name = strings.ReplaceAll(name, "-", "_")
	return name
}

func main() {
	_ = godotenv.Load()
	mongoURI := os.Getenv("MONGO_URI")
	dbName := os.Getenv("MONGO_DB")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	mongoClient = client
	mongoDB = client.Database(dbName)

	r := gin.Default()

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// New routes reflecting db and dataset naming
	// NOTE: These routes do not validate input and are for demonstration only.
	r.GET("/db/:db/datasets", listDatasetsHandler)
	r.GET("/db/:db/datasets/:dataset", getDatasetHandler)

	r.Run(":" + port)
}
