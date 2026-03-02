package services

type services struct {
	Transaction
}

var Service *services

func init() {
	Service = &services{
		Transaction: Transaction{},
	}
}
