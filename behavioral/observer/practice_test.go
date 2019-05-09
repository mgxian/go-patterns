package observer

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStockDog(t *testing.T) {
	oriOutputWriter := outputWriter
	outputWriter = new(strings.Builder)
	defer func() { outputWriter = oriOutputWriter }()

	apple := stock{"apple", 10}
	google := stock{"google", 20}
	facebook := stock{"facebook", 15}

	sb1 := newStockBuyer("mgxian")
	sb1.buy(google)
	sb1.buy(facebook)

	sb2 := newStockBuyer("will")
	sb2.buy(apple)
	sb2.buy(google)

	sb := newStockBird()
	sb.addStockObserver(sb1)
	sb.addStockObserver(sb2)

	sd := newStockDog()
	sd.sk = sb
	sd.addStock(apple)
	sd.addStock(google)
	sd.addStock(facebook)

	changedStock := stock{"apple", 15}
	sd.check(changedStock)
	assert.Equal(t, "hello will stock apple up and price is 15\n", outputWriter.(*strings.Builder).String())
	outputWriter.(*strings.Builder).Reset()

	changedStock = stock{"apple", 9}
	sd.check(changedStock)
	assert.Equal(t, "hello will stock apple down and price is 9\n", outputWriter.(*strings.Builder).String())
	outputWriter.(*strings.Builder).Reset()

	changedStock = stock{"google", 21}
	sd.check(changedStock)
	assert.Equal(t, "", outputWriter.(*strings.Builder).String())
	outputWriter.(*strings.Builder).Reset()

	changedStock = stock{"google", 22}
	sd.check(changedStock)
	expected := "hello mgxian stock google up and price is 22\nhello will stock google up and price is 22\n"
	assert.Equal(t, expected, outputWriter.(*strings.Builder).String())
	outputWriter.(*strings.Builder).Reset()
}
