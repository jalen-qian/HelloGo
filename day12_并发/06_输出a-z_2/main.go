package main

import (
	"sync"
)

func printAToZ(wg *sync.WaitGroup, ch chan rune) {
	defer wg.Done()

}

func main() {
	ch := make(chan rune)
	wg := &sync.WaitGroup{}
	go printAToZ(wg, ch)
	go printAToZ(wg, ch)
	wg.Wait()
}
