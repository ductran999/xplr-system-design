package main

import (
	"fmt"
	"sync"
	"time"
)

type TicketSpace struct {
	sync.RWMutex
	Available int
	Sold      int
}

type ProcessingUnit struct {
	space *TicketSpace
}

func (pu *ProcessingUnit) ProcessOrder(userID int) bool {
	pu.space.Lock()
	defer pu.space.Unlock()

	if pu.space.Available > 0 {
		pu.space.Available--
		pu.space.Sold++
		return true
	}
	return false
}

func DataPump(space *TicketSpace) {
	for {
		time.Sleep(2 * time.Second) // Pump data to db interval
		space.RLock()
		sold := space.Sold
		remain := space.Available
		space.RUnlock()

		fmt.Printf("[Data Pump] Synced:  Sole %d, Remain %d\n", sold, remain)
	}
}

func main() {
	// Init 1000 ticket on RAM
	space := &TicketSpace{Available: 1000}
	pu := &ProcessingUnit{space: space}

	// Background job sync data
	go DataPump(space)

	fmt.Println("SBA Simulation Started. 10.000 Users claim ticket concurrency...")

	var wg sync.WaitGroup
	start := time.Now()

	for i := 1; i <= 10000; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			pu.ProcessOrder(id)
		}(i)
	}

	wg.Wait()
	fmt.Println("============= FLASH SALE ==================")
	fmt.Printf("10.000 request on RAM need: %v\n", time.Since(start))

	time.Sleep(3 * time.Second)
}
