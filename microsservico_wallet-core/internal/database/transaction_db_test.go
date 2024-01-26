package database

import (
	"database/sql"
	"github.com.br/Rafiere/course_fullcycle_arquitetura-baseada-em-microsservicos/microsservico_wallet-core/internal/entity"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TransactionDBTestSuite struct {
	suite.Suite
	db            *sql.DB
	client        *entity.Client
	client2       *entity.Client
	accountFrom   *entity.Account
	accountTo     *entity.Account
	transactionDB *TransactionDB
}

func (s *TransactionDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)

	s.db = db

	db.Exec("CREATE TABLE clients (id VARCHAR(255), name VARCHAR(255), email VARCHAR(255), created_at TIMESTAMP)")
	db.Exec("CREATE TABLE accounts (id VARCHAR(255), client_id VARCHAR(255), balance DECIMAL(10,2), created_at TIMESTAMP)")
	db.Exec("CREATE TABLE transactions (id VARCHAR(255), account_id_from VARCHAR(255), account_id_to VARCHAR(255), amount DECIMAL(10,2), created_at TIMESTAMP)")

	client, err := entity.NewClient("John", "john@email.com")
	s.client = client
	s.Nil(err)

	client2, err := entity.NewClient("John2", "john2@email.com")
	s.client2 = client2
	s.Nil(err)

	accountFrom := entity.NewAccount(s.client)
	accountFrom.Balance = 1000
	s.accountFrom = accountFrom

	accountTo := entity.NewAccount(s.client2)
	accountTo.Balance = 1000
	s.accountTo = accountTo

	s.transactionDB = NewTransactionDB(db)
}

func (s *TransactionDBTestSuite) TearDownSuite() {
	defer s.db.Close()

	_, err := s.db.Exec("DROP TABLE clients")
	if err != nil {
		return
	}

	_, err = s.db.Exec("DROP TABLE accounts")
	if err != nil {
		return
	}

	_, err = s.db.Exec("DROP TABLE transactions")
	if err != nil {
		return
	}
}

func TestTransactionDBTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionDBTestSuite))
}

/* Abaixo, estamos criando e persistindo uma transação no banco de dados. */
func (s *TransactionDBTestSuite) TestTransactionDB_Create() {
	transaction, err := entity.NewTransaction(s.accountFrom, s.accountTo, 100)
	s.Nil(err)

	err = s.transactionDB.Create(transaction)
	s.Nil(err)
}
