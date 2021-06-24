package main

import (
	"runtime"
	"sync")


func main() {
	var l sync.Mutex

    l.Lock()
}