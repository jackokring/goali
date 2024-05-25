// This package is a cell based game framework.
//
// Written in a normal form, and later expanded for speed memory trade-offs.
// The speed form may be cache serialized also.
package zone

//=================================
//******** Game Framework *********
//=================================

import (
	"context"
	"os/user"

	"github.com/jackc/pgx/v5"
	fe "github.com/jackokring/goali/filerr"
	za "github.com/jackokring/goali/zoneauto"
)

func username() string {
	me, err := user.Current()
	fe.Fatal(err)
	return me.Username
}

func OpenConnection() (queries *za.Queries, close func()) {
	me := username()
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "user="+me+" dbname="+me+" sslmode=verify-full")
	fe.Fatal(err)

	queries = za.New(conn)
	return queries, func() {
		conn.Close(ctx)
	}
}
