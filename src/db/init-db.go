package db

import (
	"fmt"
	"go-movies/src/db/mongodb"
	"go-movies/src/db/neo4j"
)

func InitDatabases() {
	var err error
	err = mongodb.ConfigureAndConnect()
	if err == nil {
		fmt.Println("MongoDB Connection Established!")
	}

	err = neo4j.ConfigureAndConnect()
	if err == nil {
		fmt.Println("Neo4j Connection Established!")
	}
}
