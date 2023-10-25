package server

import (
	"fmt"
	"github.com/imneov/PostgreCRD/internal/config"
	"github.com/imneov/PostgreCRD/internal/gosql"
	"github.com/imneov/PostgreCRD/internal/tdtl"
	"strings"
)

type database struct {
	cfg    *config.Config
	mb     *gosql.MemoryBackend
	parser gosql.Parser
}

func NewDatabase(cfg *config.Config) *database {
	d := &database{
		parser: gosql.Parser{},
		mb:     gosql.NewMemoryBackend(),
		cfg:    cfg,
	}
	d.Init()
	return d
}

func (d *database) Parse(sql string) (*gosql.Ast, error) {
	return d.parser.Parse(sql)
}

func (d *database) Init() {
	d.CreatTables()
}

func (d *database) CreatTables() {
	for _, database := range d.cfg.Databases {
		for name, table := range database.TableMapping {
			ast, err := d.Parse(genCreateTableSQL(name, table))
			//"CREATE TABLE users (id INT, name TEXT); INSERT INTO users VALUES (1, 'Admin'); SELECT id, name FROM users")
			if err != nil {
				panic(err)
			}
			if len(ast.Statements) == 1 {
				d.mb.CreateTable(ast.Statements[0].CreateTableStatement)
			}
		}

	}
}

func (d *database) InsertToTables(tableName string, tm config.TableMapping, data []byte) {

	//INSERT INTO users VALUES (1, 'Admin')
}

func genCreateTableSQL(tableName string, tm config.TableMapping) string {
	fields := []string{}
	for k, e := range tm.Rules {
		fields = append(fields, fmt.Sprintf("%s %s", k, e.Type))
	}
	if len(fields) == 0 {
		return ""
	}
	return fmt.Sprintf("CREATE TABLE %s (%s)", tableName, strings.Join(fields, ", "))
}

func genInsertSQL(tableName string, tm config.TableMapping, data []byte) string {
	dt := tdtl.New(data)
	fields := []string{}
	values := []string{}
	for k, e := range tm.Rules {
		path := e.Source
		fields = append(fields, k)
		values = append(values, dt.Get(path).String())
	}
	if len(fields) == 0 {
		return ""
	}
	//INSERT INTO users VALUES (1, 'Admin')
	return fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, strings.Join(fields, ", "), strings.Join(values, ", "))
}
