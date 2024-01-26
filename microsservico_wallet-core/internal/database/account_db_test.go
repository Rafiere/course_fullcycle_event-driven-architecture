package database

import (
	"database/sql"
	"github.com.br/Rafiere/course_fullcycle_arquitetura-baseada-em-microsservicos/microsservico_wallet-core/internal/entity"
	"github.com/stretchr/testify/suite"
	"testing"
)

type AccountDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	accountDB *AccountDB
	client    *entity.Client
}

func (s *AccountDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	s.db = db

	db.Exec("CREATE TABLE clients (id VARCHAR(255), name VARCHAR(255), email VARCHAR(255), created_at TIMESTAMP)")
	db.Exec("CREATE TABLE accounts (id VARCHAR(255), client_id VARCHAR(255), balance DECIMAL(10,2), created_at TIMESTAMP)")
	s.accountDB = NewAccountDB(db)
	s.client, _ = entity.NewClient("John", "john@email.com")
}

func (s *AccountDBTestSuite) TearDownSuite() {
	defer s.db.Close()

	_, err := s.db.Exec("DROP TABLE clients")
	if err != nil {
		return
	}

	_, err = s.db.Exec("DROP TABLE accounts")
	if err != nil {
		return
	}
}

func TestAccountDBTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDBTestSuite))
}

func (s *AccountDBTestSuite) TestAccountDB_FindByID() {

	s.db.Exec("INSERT INTO clients (id, name, email, created_at) VALUES (?, ?, ?, ?)", s.client.ID, s.client.Name, s.client.Email, s.client.CreatedAt)

	account := entity.NewAccount(s.client)

	err := s.accountDB.Save(account)

	s.Nil(err)

	accountDB, err := s.accountDB.Get(account.ID)

	s.Nil(err)
	s.Equal(account.ID, accountDB.ID)
	s.Equal(account.Client.ID, accountDB.Client.ID)
	s.Equal(account.Balance, accountDB.Balance)
	s.Equal(account.Client.ID, accountDB.Client.ID)
	s.Equal(account.Client.Name, accountDB.Client.Name)
	s.Equal(account.Client.Email, accountDB.Client.Email)
}

func (s *AccountDBTestSuite) TestAccountDB_Save() {

	account := entity.NewAccount(s.client)

	err := s.accountDB.Save(account)
	s.Nil(err)
}
