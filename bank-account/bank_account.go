package account

import (
	"sync"
)

const testVersion = 1

type Account struct {
	mutex   sync.Mutex
	balance int64
	closed  bool
}

func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{balance: initialDeposit, closed: false}
}

func (a *Account) Close() (payout int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if a.closed {
		return
	}

	payout = a.balance
	a.balance = 0
	ok, a.closed = true, true
	return
}

func (a *Account) Balance() (balance int64, ok bool) {
	if a.closed {
		return
	}

	balance, ok = a.balance, true
	return
}

func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if a.closed {
		return
	}

	if a.balance+amount < 0 {
		return
	}

	a.balance += amount
	newBalance, ok = a.balance, true
	return
}
