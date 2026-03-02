package controllers

type controllers struct {
	Transaction
}

var Controller *controllers

func init() {
	Controller = &controllers{
		Transaction: Transaction{},
	}
}
