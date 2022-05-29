//http.go
package main

import (
	practiceInterview "golangAlgorithmPractice/practiceInterview"
)

func main() {
	// 练习
	practiceInterview.Test()

	chanStop := make(chan bool)
	chanStop <- true
}
