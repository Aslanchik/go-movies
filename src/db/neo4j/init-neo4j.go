package neo4j

import "github.com/neo4j/neo4j-go-driver/neo4j"

type Neo4jInstance struct {
	Driver  *neo4j.Driver
	Session *neo4j.Session
}

var Instance Neo4jInstance

const dbURI = "bolt://localhost:7687"

func ConfigureAndConnect() error {

	driver, err := neo4j.NewDriver(dbURI, neo4j.NoAuth(), func(conf *neo4j.Config) { conf.Encrypted = false })
	if err != nil {
		return nil
	}

	session, err := driver.Session(neo4j.AccessModeWrite)
	if err != nil {
		return nil
	}

	Instance = Neo4jInstance{
		Driver:  &driver,
		Session: &session,
	}

	return nil
}
