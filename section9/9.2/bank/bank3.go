package bank

import "sync"

// 互斥锁
var (
	mu      sync.Mutex // 保护balance
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	// balance = balance + amount
	// mu.Unlock()
	defer mu.Unlock()
	deposit(amount)
}

// func Balance() int {
// 	mu.Lock()
// 	b := balance
// 	mu.Unlock()
// 	return b
// }

func Balance() int {
	mu.Lock()
	defer mu.Unlock()
	return balance
}

func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if Balance() < 0 {
		deposit(amount)
		return false //余额不足
	}
	return true
}

// 这个函数要求已获取互斥锁
func deposit(amount int) { balance += amount }
