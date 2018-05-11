package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Computer struct {
	w *sync.WaitGroup
}

func (c *Computer) ImportantCompute() {
	for i := 1; i <= 10; i++ {
		c.w.Add(1)
		go func(a int) {
			defer c.w.Done()
			for i := 1; i <= 10; i++ {
				fmt.Printf("compute%d step: %d\n", a, i)
				time.Sleep(1 * time.Second)
			}
		}(i)
	}

}

func (c *Computer) Shutdown() {
	c.w.Wait()
}

func main() {
	fmt.Println("start...")
	c := &Computer{
		w: &sync.WaitGroup{},
	}
	go c.ImportantCompute()

	defer c.Shutdown()

	shutdow_observer := make(chan os.Signal, 1)
	signal.Notify(shutdow_observer, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	select {
	case s := <-shutdow_observer:
		fmt.Println("shutdown by.... ", s)
		return
	}
}
