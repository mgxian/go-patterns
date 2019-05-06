package bridge

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransformFromMySQLToMultipleFileType(t *testing.T) {
	oriOutputWriter := outputWriter
	outputWriter = new(strings.Builder)
	defer func() { outputWriter = oriOutputWriter }()

	tf := new(mysqlDBTransformer)

	tf.setFile(new(txtFile))
	tf.transform()
	assert.Equal(t, "query data from mysql\nwrite to txt file\n", outputWriter.(*strings.Builder).String())
	outputWriter.(*strings.Builder).Reset()

	tf.setFile(new(xmlFile))
	tf.transform()
	assert.Equal(t, "query data from mysql\nwrite to xml file\n", outputWriter.(*strings.Builder).String())
	outputWriter.(*strings.Builder).Reset()

	tf.setFile(new(pdfFile))
	tf.transform()
	assert.Equal(t, "query data from mysql\nwrite to pdf file\n", outputWriter.(*strings.Builder).String())
}

func TestTransformFromPostgresqlToMultipleFileType(t *testing.T) {
	oriOutputWriter := outputWriter
	outputWriter = new(strings.Builder)
	defer func() { outputWriter = oriOutputWriter }()

	tf := new(postgresqlDBTransformer)

	tf.setFile(new(txtFile))
	tf.transform()
	assert.Equal(t, "query data from postgresql\nwrite to txt file\n", outputWriter.(*strings.Builder).String())
	outputWriter.(*strings.Builder).Reset()

	tf.setFile(new(xmlFile))
	tf.transform()
	assert.Equal(t, "query data from postgresql\nwrite to xml file\n", outputWriter.(*strings.Builder).String())
	outputWriter.(*strings.Builder).Reset()

	tf.setFile(new(pdfFile))
	tf.transform()
	assert.Equal(t, "query data from postgresql\nwrite to pdf file\n", outputWriter.(*strings.Builder).String())
}

func TestTransformFromHbaseToMultipleFileType(t *testing.T) {
	oriOutputWriter := outputWriter
	outputWriter = new(strings.Builder)
	defer func() { outputWriter = oriOutputWriter }()

	tf := new(hbaseDBTransformer)

	tf.setFile(new(txtFile))
	tf.transform()
	assert.Equal(t, "query data from hbase\nwrite to txt file\n", outputWriter.(*strings.Builder).String())
	outputWriter.(*strings.Builder).Reset()

	tf.setFile(new(xmlFile))
	tf.transform()
	assert.Equal(t, "query data from hbase\nwrite to xml file\n", outputWriter.(*strings.Builder).String())
	outputWriter.(*strings.Builder).Reset()

	tf.setFile(new(pdfFile))
	tf.transform()
	assert.Equal(t, "query data from hbase\nwrite to pdf file\n", outputWriter.(*strings.Builder).String())
}
