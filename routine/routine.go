package concurrence

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func Add(a, b int) int {
	fmt.Println("Add func")
	return a + b
}

func SimpleGoroutine() {
	fmt.Printf("logic processer nums: %d\n", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU() / 2) // set the number of core that current process can use

	fmt.Printf("routine nums: %d\n", runtime.NumGoroutine())

	go Add(3, 5)
	go Add(3, 5)

	time.Sleep(1 * time.Second) // sleep for 1 second
}

func WaitGroup() {
	const N = 10
	wg := sync.WaitGroup{}
	wg.Add(N)
	var i int = 0
	for ; i < N; i++ {
		go func(a, b int) {
			defer wg.Done()
			time.Sleep(10 * time.Millisecond) // simulate some work
			fmt.Printf("Hello from goroutine %d\n", a)
			time.Sleep(10 * time.Millisecond) // simulate some work
			_ = a + b
		}(i, i+1)
	}

	fmt.Printf("current routine nums: %d\n", runtime.NumGoroutine()) //N+1
	wg.Wait()                                                        // wait for all goroutines to finish
	fmt.Printf("current routine nums: %d\n", runtime.NumGoroutine()) //1
}

var lock syn.Mutex
var count int

func inc3() {
	lock.Lock()
	count++
	lock.Unlock()
}

var (
	mu sync.RWMutex
)

func ReentranceRlock(n int) {
	mu.RLock()
	defer mu.RUnlock()
	fmt.Printf("ReentranceRlock %d\n", n)
	if n > 0 {
		ReentranceRlock(n - 1)
	}
	time.Sleep(1 * time.Second)
}
