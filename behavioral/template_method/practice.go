package template_method

import (
	"fmt"
	"io"
	"os"
)

var outputWriter io.Writer = os.Stdout

type db interface {
	connect()
	open()
	query(sql string)
	update(sql string)
	close()
}

type baseDB struct {
	db
}

func (bd *baseDB) execute(sql string, isUpdated bool) {
	bd.connect()
	bd.open()

	if isUpdated {
		bd.update(sql)
	} else {
		bd.query(sql)
	}

	bd.close()
}

func newDB() *baseDB {
	return &baseDB{}
}

func (bd *baseDB) use(aDB db) {
	bd.db = aDB
}

type mysql struct {
}

func (m *mysql) connect() {
	fmt.Fprintln(outputWriter, "mysql connect")
}

func (m *mysql) open() {
	fmt.Fprintln(outputWriter, "mysql open")
}

func (m *mysql) query(sql string) {
	fmt.Fprintln(outputWriter, "mysql query "+sql)
}

func (m *mysql) update(sql string) {
	fmt.Fprintln(outputWriter, "mysql update "+sql)
}

func (m *mysql) close() {
	fmt.Fprintln(outputWriter, "mysql close")
}

type tidb struct {
}

func (t *tidb) connect() {
	fmt.Fprintln(outputWriter, "tidb connect")
}

func (t *tidb) open() {
	fmt.Fprintln(outputWriter, "tidb open")
}

func (t *tidb) query(sql string) {
	fmt.Fprintln(outputWriter, "tidb query "+sql)
}

func (t *tidb) update(sql string) {
	fmt.Fprintln(outputWriter, "tidb update "+sql)
}

func (t *tidb) close() {
	fmt.Fprintln(outputWriter, "tidb close")
}
