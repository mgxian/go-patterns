package observer

import (
	"fmt"
	"io"
	"os"
)

var outputWriter io.Writer = os.Stdout

type stockObserver interface {
	update(string)
	getBoughtStocks() []stock
}

type stock struct {
	name  string
	price int
}

type stockBuyer struct {
	name   string
	stocks []stock
}

func newStockBuyer(name string) *stockBuyer {
	return &stockBuyer{
		name:   name,
		stocks: make([]stock, 0),
	}
}

func (s *stockBuyer) buy(aStock stock) {
	s.stocks = append(s.stocks, aStock)
}

func (s *stockBuyer) update(msg string) {
	fmt.Fprintln(outputWriter, "hello "+s.name+" "+msg)
}

func (s *stockBuyer) getBoughtStocks() []stock {
	return s.stocks
}

type stockKeeper interface {
	addStockObserver(obs stockObserver)
	removeStockObserver(obs stockObserver)
	notify(stockEvent)
}

type stockBird struct {
	stockBuyers []stockObserver
}

func newStockBird() *stockBird {
	return &stockBird{
		stockBuyers: make([]stockObserver, 0),
	}
}

func (sb *stockBird) addStockObserver(sobs stockObserver) {
	sb.stockBuyers = append(sb.stockBuyers, sobs)
}

func (sb *stockBird) removeStockObserver(sobs stockObserver) {}

func (sb *stockBird) generateNotifyMessage(se stockEvent) (msg string) {
	aStock := se.s
	if se.et == Up {
		return fmt.Sprintf("stock %s up and price is %d", aStock.name, aStock.price)
	}

	if se.et == Down {
		return fmt.Sprintf("stock %s down and price is %d", aStock.name, aStock.price)
	}

	return fmt.Sprintf("unkonw event type")
}

func (sb *stockBird) sendMessageTo(o stockObserver, se stockEvent) {
	for _, s := range o.getBoughtStocks() {
		if se.s.name == s.name {
			o.update(sb.generateNotifyMessage(se))
			return
		}
	}
}

func (sb *stockBird) notify(se stockEvent) {
	for _, o := range sb.stockBuyers {
		sb.sendMessageTo(o, se)
	}
}

type eventType int

// stock event type
const (
	Up eventType = iota
	Down
)

type stockEvent struct {
	s  stock
	et eventType
}

func newStockEvent(s stock, et eventType) stockEvent {
	return stockEvent{
		s:  s,
		et: et,
	}
}

type stockDog struct {
	sk     stockKeeper
	stocks map[string]stock
}

func newStockDog() *stockDog {
	return &stockDog{
		stocks: make(map[string]stock, 0),
	}
}

func (sd *stockDog) addStock(s stock) {
	sd.stocks[s.name] = s
}

func (sd *stockDog) check(aStock stock) {
	st, ok := sd.stocks[aStock.name]
	if !ok {
		return
	}

	amountOfIncrease := float64(aStock.price-st.price) / float64(st.price)
	if amountOfIncrease > 0 && amountOfIncrease > 0.05 {
		se := newStockEvent(aStock, Up)
		sd.sk.notify(se)
		return
	}

	if amountOfIncrease < 0 && amountOfIncrease < -0.05 {
		se := newStockEvent(aStock, Down)
		sd.sk.notify(se)
	}
}
