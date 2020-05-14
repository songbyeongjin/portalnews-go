package db

import(
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"

)

func TestGetDbConnector(t *testing.T){
	assertion := assert.New(t)
	dbConnector, err := getDbConnector()

	assertion.NotNil(dbConnector)
	assertion.Nil(err)
}

func TestGetConnectString(t *testing.T){
	assertion := assert.New(t)

	connector := &Connector{
		Dialect:"testDialect",
		Host:"testHost",
		Dbname:"testDbName",
		Port:"testPort",
		User:"testUser",
		Password:"testPassword",
	}

	compareStr := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s",
		connector.Host,
		connector.Port,
		connector.User,
		connector.Dbname,
		connector.Password)

	assertion.Equal(connector.GetConnectString(), compareStr)
}

func TestNewDbHandler(t *testing.T){
	assertion := assert.New(t)

	handler := NewDbHandler()
	assertion.NotNil(handler)
}