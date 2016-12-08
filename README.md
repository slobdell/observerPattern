```
package main

import (
	"./observerPattern"
	"fmt"
	"time"
)

func registerAndListen(observable observerPattern.ObservableSubject) {
	listenChannel := make(chan []interface{})
	observable.Register(listenChannel)

	for {
		var publishedValues []interface{} = <-listenChannel

		for _, val := range publishedValues {
			asString := val.(string)
			fmt.Printf("Value was just published: %s\n", asString)
		}
	}
}

func blockIndefinitely() {
	for {
		time.Sleep(1 * time.Second)
	}
	select {}
}

func doStuff() {
	observable := observerPattern.NewConcreteObservable()
	for i := 0; i < 10; i++ {
		go registerAndListen(observable)
	}
	time.Sleep(1 * time.Second)
	observable.NotifyObservers("Hello")
}

func main() {
	doStuff()
	blockIndefinitely()
}
```
