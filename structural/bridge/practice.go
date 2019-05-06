package bridge

import (
	"fmt"
	"io"
	"os"
)

var outputWriter io.Writer = os.Stdout

type outPutFile interface {
	write()
}

type transformer struct {
	file outPutFile
}

func (t *transformer) setFile(file outPutFile) {
	t.file = file
}

func (t *transformer) transform() {
}

type mysqlDBTransformer struct {
	transformer
}

func (m *mysqlDBTransformer) query() {
	fmt.Fprintln(outputWriter, "query data from mysql")
}

func (m *mysqlDBTransformer) transform() {
	m.query()
	m.file.write()
}

type postgresqlDBTransformer struct {
	transformer
}

func (p *postgresqlDBTransformer) query() {
	fmt.Fprintln(outputWriter, "query data from postgresql")
}

func (p *postgresqlDBTransformer) transform() {
	p.query()
	p.file.write()
}

type hbaseDBTransformer struct {
	transformer
}

func (h *hbaseDBTransformer) query() {
	fmt.Fprintln(outputWriter, "query data from hbase")
}

func (h *hbaseDBTransformer) transform() {
	h.query()
	h.file.write()
}

type txtFile struct {
}

func (t *txtFile) write() {
	fmt.Fprintln(outputWriter, "write to txt file")
}

type xmlFile struct {
}

func (t *xmlFile) write() {
	fmt.Fprintln(outputWriter, "write to xml file")
}

type pdfFile struct {
}

func (t *pdfFile) write() {
	fmt.Fprintln(outputWriter, "write to pdf file")
}
