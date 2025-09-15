package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Deposit struct {
	Amount   int
	Customer string
}

func main() {
	workers := 5
	customers := 10

	// Channels
	depositCh := make(chan Deposit, customers)
	resultCh := make(chan Deposit, customers)

	// Independent random generator
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// WaitGroup for workers
	var wg sync.WaitGroup
	wg.Add(workers)

	// Start worker pool
	go startWorkers(workers, depositCh, resultCh, &wg)

	// Send deposits
	generateDeposits(customers, depositCh, r)
	close(depositCh)

	// Close result channel once workers are done
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	// Collect and print passbook
	passbook := collectPassbook(resultCh)
	printPassbook(passbook)
}

// startWorkers starts N worker goroutines
func startWorkers(n int, depositCh <-chan Deposit, resultCh chan<- Deposit, wg *sync.WaitGroup) {
	for i := 0; i < n; i++ {
		go processDeposit(depositCh, resultCh, wg)
	}
}

// processDeposit simulates deposit processing by a worker
func processDeposit(depositCh <-chan Deposit, resultCh chan<- Deposit, wg *sync.WaitGroup) {
	defer wg.Done()
	for d := range depositCh {
		// processing: fee
		d.Amount -= 10 // flat fee
		resultCh <- d
	}
}

// generateDeposits creates random deposits for customers
func generateDeposits(customers int, depositCh chan<- Deposit, r *rand.Rand) {
	for i := 0; i < customers; i++ {
		deposit := Deposit{
			Amount:   r.Intn(1000),                // random amount
			Customer: fmt.Sprintf("Cust-%d", i+1), // label
		}
		depositCh <- deposit
	}
}

// collectPassbook gathers all processed deposits into a slice
func collectPassbook(resultCh <-chan Deposit) []Deposit {
	var passbook []Deposit
	for d := range resultCh {
		passbook = append(passbook, d)
	}
	return passbook
}

// printPassbook prints all records neatly
func printPassbook(passbook []Deposit) {
	fmt.Println("📒 Passbook Records:")
	var total int
	for _, entry := range passbook {
		total += entry.Amount
		fmt.Printf("Customer: %s | Amount: %d\n", entry.Customer, entry.Amount)
	}
	fmt.Printf("Total Balance: %d\n", total)
}
