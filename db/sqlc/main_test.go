package db

import {
	"log"
	"database/sql"
	"testing"
	"os"

	_"github.com/lib/pq"
}

const {
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
}

var testQueries = Queries

func TestMain(m *testing.M) {
	conn, err := sql.open(dbDriver, dbSource)

	if err != nil {
		log.Fatal ("cannot connect to db: ", err)
	}

	testQueries = New(conn)

	os.Exist(m.Run())
}