package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	// "github.com/jackc/pgx/v5/pgconn" // for on_notice??
)

type VolleyDB struct {
	connString string
}

func CreateDBConnection(connString string) VolleyDB {
	vdb := VolleyDB{connString}
	return vdb
}

func (v *VolleyDB) PingDB() (ok bool) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, v.connString)
	if err != nil {
		log.Fatalln("Could not connect to DB: ", err)
	}
	defer conn.Close(ctx)

	err = conn.Ping(ctx)
	if err != nil {
		log.Fatalln("Ping failed:", err)
	}

	return true
}
