package webfetcer

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

func Webfetcer(url string, resultChan chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	sleepTime := time.Duration(rand.Intn(401)+100) * time.Millisecond
	time.Sleep(sleepTime)

	resultChan<- fmt.Sprintf("Fetched: %s (took %v)", url, sleepTime)
}

func Receiver(resultChan <-chan string, done chan<- bool){
	for result := range resultChan{
			log.Println(result)
	}
	done <-true
}