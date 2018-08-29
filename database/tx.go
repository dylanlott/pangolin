package db

import "sync"

type Transaction struct {
	db   DB
	lock sync.Mutex
}

func NewTransaction() (*Transaction, error) {
	return &Transaction{}, nil
}

func (t *Transaction) Action() error {
	return nil
}

func (t *Transaction) Completed() error {
	return nil
}
