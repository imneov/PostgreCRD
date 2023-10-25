package config

type Config struct {
	Databases []Database `json:"databases,omitempty"`
}

type Database struct {
	PGServer     PGServer                `json:"PGServer"`
	TableMapping map[string]TableMapping `json:"tableMapping,omitempty"`
}

// PGServer
// DATABASE := "postgres://username:password@localhost:5432/database_name"
type PGServer struct {
	DATABASE string `yaml:"database"`
}

// TableMapping
// Rules -> [TableName]Rule
type TableMapping struct {
	// Type is "DataTable/RelationTable"
	Type  string          `json:"type,omitempty"`
	Rules map[string]Rule `json:"rules,omitempty" json:"rules,omitempty"`
}

// Rule
// SourcePath: spec.instance
type Rule struct {
	Source   string `json:"source,omitempty"`
	Type     string `json:"type,omitempty"`
	Nullable string `json:"nullable,omitempty"`
}
