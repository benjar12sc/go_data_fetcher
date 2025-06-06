package main

import (
	"context"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/yaml.v3"
)

type MongoConfig struct {
	URI      string `yaml:"uri"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func main() {
	folderPath := flag.String("folder_path", "", "Path to folder containing CSV files to load")
	configPath := flag.String("config", "mongo_config.yaml", "Path to MongoDB config YAML file")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s --folder_path <csv_folder> [--config <mongo_config.yaml>]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	if *folderPath == "" {
		flag.Usage()
		os.Exit(1)
	}

	cfg, err := loadMongoConfig(*configPath)
	if err != nil {
		log.Fatalf("Failed to load Mongo config: %v", err)
	}

	dbName := sanitizeName(filepath.Base(filepath.Clean(*folderPath)))

	clientOpts := options.Client().ApplyURI(cfg.URI)
	if cfg.Username != "" && cfg.Password != "" {
		clientOpts.SetAuth(options.Credential{
			Username: cfg.Username,
			Password: cfg.Password,
		})
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(ctx)
	db := client.Database(dbName)

	files, err := os.ReadDir(*folderPath)
	if err != nil {
		log.Fatalf("Failed to read folder: %v", err)
	}

	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".csv") {
			continue
		}
		collName := sanitizeName(strings.TrimSuffix(file.Name(), ".csv"))
		csvPath := filepath.Join(*folderPath, file.Name())
		if err := importCSVToMongo(ctx, db, collName, csvPath); err != nil {
			log.Printf("Failed to import %s: %v", file.Name(), err)
		} else {
			fmt.Printf("Imported %s into %s.%s\n", file.Name(), dbName, collName)
		}
	}
}

func loadMongoConfig(path string) (*MongoConfig, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var cfg MongoConfig
	dec := yaml.NewDecoder(f)
	if err := dec.Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func sanitizeName(name string) string {
	name = strings.ReplaceAll(name, " ", "_")
	name = strings.ReplaceAll(name, "-", "_")
	name = regexp.MustCompile(`[()]+`).ReplaceAllString(name, "")
	return name
}

func importCSVToMongo(ctx context.Context, db *mongo.Database, collName, csvPath string) error {
	f, err := os.Open(csvPath)
	if err != nil {
		return err
	}
	defer f.Close()
	reader := csv.NewReader(f)
	headers, err := reader.Read()
	if err != nil {
		return err
	}
	var docs []interface{}
	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return err
		}
		doc := make(map[string]interface{})
		for i, h := range headers {
			if i < len(record) {
				doc[h] = record[i]
			}
		}
		docs = append(docs, doc)
	}
	if len(docs) > 0 {
		_, err = db.Collection(collName).InsertMany(ctx, docs)
	}
	return err
}
