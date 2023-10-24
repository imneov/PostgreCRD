package main

import (
	"context"
	"flag"
	"fmt"
	wire "github.com/imneov/PostgreCRD"
	"github.com/jackc/pgtype"
	shopspring "github.com/jackc/pgtype/ext/shopspring-numeric"
	"github.com/lib/pq/oid"
	"github.com/shopspring/decimal"
	"log"
	"strings"
	"time"
)

var (
	balance decimal.Decimal
	err     error
	data    [][]any
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()

	types := wire.ExtendTypes(func(info *pgtype.ConnInfo) {
		info.RegisterDataType(pgtype.DataType{
			Value: &shopspring.Numeric{},
			Name:  "numeric",
			OID:   pgtype.NumericOID,
		})
	})

	srv, err := wire.NewServer(handler, types, wire.Version("10.4"))
	if err != nil {
		panic(err)
	}

	balance, err = decimal.NewFromString("256.23")
	if err != nil {
		panic(err)
	}
	data = append(data, []any{balance, "John", true, 29})
	data = append(data, []any{balance, "Marry", false, 21})

	log.Println("PostgreSQL server is up and running at [127.0.0.1:5432]")
	srv.ListenAndServe("127.0.0.1:5432")
}

var table = wire.Columns{
	{
		Table:  0,
		Name:   "account_balance",
		Oid:    oid.T_numeric,
		Width:  1,
		Format: wire.TextFormat,
	},
	{
		Table:  0,
		Name:   "name",
		Oid:    oid.T_text,
		Width:  256,
		Format: wire.TextFormat,
	},
	{
		Table:  0,
		Name:   "member",
		Oid:    oid.T_bool,
		Width:  1,
		Format: wire.TextFormat,
	},
	{
		Table:  0,
		Name:   "age",
		Oid:    oid.T_int4,
		Width:  1,
		Format: wire.TextFormat,
	},
}

func handler(ctx context.Context, query string) (wire.PreparedStatementFn, []oid.Oid, wire.Columns, error) {
	log.Println("incoming SQL query:", query)

	if query == "BEGIN" {
		statement := func(ctx context.Context, writer wire.DataWriter, parameters []string) error {
			return writer.Complete(query)
		}

		return statement, wire.ParseParameters(query), nil, nil
	}

	if strings.Index(query, "INSERT") != -1 {
		statement := func(ctx context.Context, writer wire.DataWriter, parameters []string) error {
			return writer.Complete(query)
		}
		data = append(data, []any{balance, fmt.Sprintf("Marry-%d", time.Now().Second()), false, 21})

		return statement, wire.ParseParameters(query), nil, nil
	}

	if strings.Index(query, "DELETE") != -1 {
		statement := func(ctx context.Context, writer wire.DataWriter, parameters []string) error {
			return writer.Complete(query)
		}
		data = data[1:]

		return statement, wire.ParseParameters(query), nil, nil
	}

	if strings.Index(query, "UPDATE") != -1 {
		statement := func(ctx context.Context, writer wire.DataWriter, parameters []string) error {
			return writer.Complete(query)
		}

		return statement, wire.ParseParameters(query), nil, nil
	}

	if strings.Index(query, "COMMIT") != -1 {
		statement := func(ctx context.Context, writer wire.DataWriter, parameters []string) error {
			return writer.Complete(query)
		}

		return statement, wire.ParseParameters(query), nil, nil
	}

	statement := func(ctx context.Context, writer wire.DataWriter, parameters []string) error {
		for _, d := range data {
			writer.Row(d)
		}
		return writer.Complete(fmt.Sprintf("SELECT %d", len(data)))
	}

	return statement, wire.ParseParameters(query), table, nil
}
