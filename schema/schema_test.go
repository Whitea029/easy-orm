package schema

import (
	"testing"

	"github.com/Whitea029/easy-orm/dialect"
)

type User struct {
	ID   int `easyorm:"PRIMARY KEY"`
	Name string
	Age  int
}

var TestDial, _ = dialect.GetDialect("sqlite3")

func TestParse(t *testing.T) {
	schema := Pares(&User{}, TestDial)
	if schema.Name != "User" || len(schema.FieldNames) != 3 {
		t.Fatal("failed to parse User struct")
	}
	if schema.GetField("ID").Tag != "PRIMARY KEY" {
		t.Fatal("failed to parse primary key")
	}
}
