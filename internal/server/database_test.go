package server

import (
	"github.com/imneov/PostgreCRD/internal/config"
	"testing"
)

func Test_genCreateTableSQL(t *testing.T) {
	tests := []struct {
		name string
		args config.TableMapping
		want string
	}{
		{
			name: "users", args: config.TableMapping{
				Rules: map[string]config.Rule{
					"id":          config.Rule{Type: "INT"},
					"instance_id": config.Rule{Type: "varchar"},
					"name":        config.Rule{Type: "text"},
				},
			}, want: "CREATE TABLE users (id INT, instance_id varchar, name text)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := genCreateTableSQL(tt.name, tt.args); got != tt.want {
				t.Errorf("genCreateTableSQL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_genInsertSQL(t *testing.T) {
	tests := []struct {
		tableName string
		args      config.TableMapping
		data      string
		want      string
	}{
		{
			tableName: "users", args: config.TableMapping{
				Rules: map[string]config.Rule{
					"id":          config.Rule{Type: "INT"},
					"instance_id": config.Rule{Type: "varchar"},
					"name":        config.Rule{Type: "text"},
				},
			},
			data: "{instance_id:'abc', name: 'def', id:123}",
			want: "CREATE TABLE users (id INT, instance_id varchar, name text)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.tableName, func(t *testing.T) {
			if got := genInsertSQL(tt.tableName, tt.args, []byte(tt.data)); got != tt.want {
				t.Errorf("genInsertSQL() = %v, want %v", got, tt.want)
			}
		})
	}
}
