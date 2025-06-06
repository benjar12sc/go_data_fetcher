# Go REST API for NRC Datasets

This project provides a REST API for accessing NRC datasets loaded into MongoDB. It is designed to work with the data structure created by the `database_populator` tool in this repository.

## Features
- Endpoints for each NRC dataset (one per MongoDB collection)
- Docker Compose for running the API and MongoDB together
- Reference to the NRC data source and project structure

## Example endpoints

- `GET /healthz` — Health check
- `GET /datasets` — List all available datasets (MongoDB collections)
- `GET /datasets/{name}` — Get all documents from a dataset (collection)

## Running locally

1. Start MongoDB (or use Docker Compose):
   ```
   docker-compose up --build
   ```
2. Run the API server:
   ```
   go run main.go
   ```

## Environment variables

- `MONGO_URI` — MongoDB connection string (default: `mongodb://root:example@mongo:27017`)
- `MONGO_DB` — Database name (default: `nrc`)
- `PORT` — API port (default: `8080`)

## Data loading

Use the `database_populator` tool from the main project to load CSVs into MongoDB before querying via the API.

## Data Source

All data is sourced from the U.S. Nuclear Regulatory Commission's public datasets:

https://www.nrc.gov/reading-rm/doc-collections/datasets/index.html

> The U.S. Nuclear Regulatory Commission is in the process of rescinding or revising guidance and policies posted on this webpage in accordance with Executive Order 14151 Ending Radical and Wasteful Government DEI Programs and Preferencing, and Executive Order 14168 Defending Women From Gender Ideology Extremism and Restoring Biological Truth to the Federal Government. In the interim, any previously issued diversity, equity, inclusion, or gender-related guidance on this webpage should be considered rescinded that is inconsistent with these Executive Orders.
