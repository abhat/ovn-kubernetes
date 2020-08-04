package libovsdb

import (
	"fmt"
	"io"
)

// DatabaseSchema is a database schema according to RFC7047
type DatabaseSchema struct {
	Name    string                 `json:"name"`
	Version string                 `json:"version"`
	Tables  map[string]TableSchema `json:"tables"`
}

// TableSchema is a table schema according to RFC7047
type TableSchema struct {
	Columns map[string]ColumnSchema `json:"columns"`
	Indexes [][]string              `json:"indexes,omitempty"`
}

// ColumnSchema is a column schema according to RFC7047
type ColumnSchema struct {
	Name      string      `json:"name"`
	Type      interface{} `json:"type"`
	Ephemeral bool        `json:"ephemeral,omitempty"`
	Mutable   bool        `json:"mutable,omitempty"`
}

// Print will print the contents of the DatabaseSchema
func (schema DatabaseSchema) Print(w io.Writer) {
	fmt.Fprintf(w, "%s, (%s)\n", schema.Name, schema.Version)
	for table, tableSchema := range schema.Tables {
		fmt.Fprintf(w, "\t %s\n", table)
		for column, columnSchema := range tableSchema.Columns {
			fmt.Fprintf(w, "\t\t %s => %v\n", column, columnSchema)
		}
	}
}

// Basic validation for operations against Database Schema
func (schema DatabaseSchema) validateOperations(operations ...Operation) bool {
	for _, op := range operations {
		table, ok := schema.Tables[op.Table]
		if ok {
			for column := range op.Row {
				if _, ok := table.Columns[column]; !ok {
					if column != "_uuid" && column != "_version" {
						return false
					}
				}
			}
			for _, row := range op.Rows {
				for column := range row {
					if _, ok := table.Columns[column]; !ok {
						if column != "_uuid" && column != "_version" {
							return false
						}
					}
				}
			}
			for _, column := range op.Columns {
				if _, ok := table.Columns[column]; !ok {
					if column != "_uuid" && column != "_version" {
						return false
					}
				}
			}
		} else {
			return false
		}
	}
	return true
}