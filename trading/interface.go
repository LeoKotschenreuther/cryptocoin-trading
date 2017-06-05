package trading

import (
	"time"
	"github.com/preichenberger/go-coinbase-exchange"
)

type ticker interface {
	Subscribe(algorithm algorithm, interval time.Duration)
}

type algorithm interface {
	PriceUpdate(rate coinbase.HistoricRate, duration time.Duration)
}