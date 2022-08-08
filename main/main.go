//http.go
package main

import (
	practiceInterview "golangAlgorithmPractice/practiceInterview"
	"log"
	"sync"
)

func main() {
	log.SetFlags(log.Lshortfile)

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(1)

	// 练习
	go practiceInterview.Test(&waitGroup)

	waitGroup.Wait()

}
