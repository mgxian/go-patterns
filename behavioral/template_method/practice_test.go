package template_method

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDBUseMySQL(t *testing.T) {
	oriOutputWriter := outputWriter
	outputWriter = new(strings.Builder)
	defer func() { outputWriter = oriOutputWriter }()

	db := newDB()
	db.use(&mysql{})
	sql := "select * from mgxian limit 1"
	isUpdated := false
	db.execute(sql, isUpdated)

	expected := "mysql connect\n" +
		"mysql open\n" +
		"mysql query select * from mgxian limit 1\n" +
		"mysql close\n"
	assert.Equal(t, expected, outputWriter.(*strings.Builder).String())
	outputWriter.(*strings.Builder).Reset()

	sql = "update mgxian set name=will where name=mgxian"
	isUpdated = true
	db.execute(sql, isUpdated)

	expected = "mysql connect\n" +
		"mysql open\n" +
		"mysql update update mgxian set name=will where name=mgxian\n" +
		"mysql close\n"
	assert.Equal(t, expected, outputWriter.(*strings.Builder).String())
}

func TestDBUseTidb(t *testing.T) {
	oriOutputWriter := outputWriter
	outputWriter = new(strings.Builder)
	defer func() { outputWriter = oriOutputWriter }()

	db := newDB()
	db.use(&tidb{})
	sql := "select * from mgxian limit 1"
	isUpdated := false
	db.execute(sql, isUpdated)

	expected := "tidb connect\n" +
		"tidb open\n" +
		"tidb query select * from mgxian limit 1\n" +
		"tidb close\n"
	assert.Equal(t, expected, outputWriter.(*strings.Builder).String())
	outputWriter.(*strings.Builder).Reset()

	sql = "update mgxian set name=will where name=mgxian"
	isUpdated = true
	db.execute(sql, isUpdated)

	expected = "tidb connect\n" +
		"tidb open\n" +
		"tidb update update mgxian set name=will where name=mgxian\n" +
		"tidb close\n"
	assert.Equal(t, expected, outputWriter.(*strings.Builder).String())
}
