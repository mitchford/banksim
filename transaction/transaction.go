package transaction

import (
	"time"
)

type Transaction struct {
	TxnId int
	ExternalSource bool
	FromAccountId int
	ToAccountId int
	Amount int
	Date time.Time
}
