package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	wire "github.com/imneov/PostgreCRD"
	"github.com/lib/pq/oid"
)

func main() {
	srv, err := wire.NewServer(handler, wire.Session(session), wire.Version("10.4"))
	if err != nil {
		panic(err)
	}

	log.Println("PostgreSQL server is up and running at [127.0.0.1:5432]")
	srv.ListenAndServe("127.0.0.1:15432")
}

type key int

var mu sync.Mutex
var id = key(1)
var counter = 0

func session(ctx context.Context) (context.Context, error) {
	mu.Lock()
	counter++
	defer mu.Unlock()
	return context.WithValue(ctx, id, counter), nil
}

var table = wire.Columns{
	{
		Table:  0,
		Name:   "account_balance",
		Oid:    oid.T_numeric,
		Width:  1,
		Format: wire.TextFormat,
	},
}

func handler(ctx context.Context, query string) (wire.PreparedStatementFn, []oid.Oid, wire.Columns, error) {
	log.Println("incoming SQL query:", query)

	statement := func(ctx context.Context, writer wire.DataWriter, parameters []string) error {
		session := ctx.Value(id).(int)
		return writer.Complete(fmt.Sprintf("OK, session: %d", session))
	}

	return statement, wire.ParseParameters(query), nil, nil
}
