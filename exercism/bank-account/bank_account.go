package account

import "sync"

type Account struct {
	balance int
	active  bool
	mu      sync.Mutex
}

func Open(amount int) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{balance: amount, active: true}
}

func (a *Account) Close() (int, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if !a.active {
		return 0, false
	}
	balance := a.balance
	a.balance = 0
	a.active = false
	return balance, true
}

func (a *Account) Balance() (int, bool) {
	return a.balance, a.active
}

func (a *Account) Deposit(amount int) (int, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if !a.active || a.balance+amount < 0 {
		return 0, false
	}
	a.balance = a.balance + amount
	return a.balance, true
}
