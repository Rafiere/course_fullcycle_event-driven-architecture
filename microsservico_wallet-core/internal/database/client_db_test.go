package database

import (
	"database/sql"
	"github.com.br/Rafiere/course_fullcycle_arquitetura-baseada-em-microsservicos/microsservico_wallet-core/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ClientDBTestSuite struct {
	suite.Suite
	db       *sql.DB
	clientDB *ClientDB
}

/* Para cada teste que seja executado o método abaixo também será executado. */
func (s *ClientDBTestSuite) SetupSuite() {

	db, err := sql.Open("sqlite3", ":memory:")

	s.Nil(err)

	s.db = db

	db.Exec("CREATE TABLE clients (id VARCHAR(255), name VARCHAR(255), email VARCHAR(255), created_at TIMESTAMP)")

	s.clientDB = NewClientDB(db)
}

func (s *ClientDBTestSuite) TearDownSuite() {
	defer s.db.Close()

	_, err := s.db.Exec("DROP TABLE clients")
	if err != nil {
		return
	}
}

/* Esse teste será o responsável de executar todos os testes. */
func TestClientDBTestSuite(t *testing.T) {
	suite.Run(t, new(ClientDBTestSuite))
}

func (s *ClientDBTestSuite) TestClientDB_Get() {

	client, _ := entity.NewClient("John", "john@email.com")

	s.clientDB.Save(client)

	clientDB, err := s.clientDB.Get(client.ID)

	s.Nil(err)
	s.Equal(client.ID, clientDB.ID)
	s.Equal(client.Name, clientDB.Name)
	s.Equal(client.Email, clientDB.Email)
}

func (s *ClientDBTestSuite) TestClientDB_Save() {

	client, _ := entity.NewClient("John", "john@email.com")

	err := s.clientDB.Save(client)
	s.Nil(err)
}
