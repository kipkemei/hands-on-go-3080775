// concurrency/sync/begin/main.go
package main

import (
	"fmt"
	"sync"
)

func main() {
	// given a list of names
	names := []string{"Sam", "Maxine", "Habiba"}

	// initialize a map to store the length of each name
	namesMap := make(map[string]int)

	// initialize a wait group for the goroutines that will be launched
	var wg sync.WaitGroup

	// set the counter.
	wg.Add(len(names))

	// initialize a mutex
	var mu sync.Mutex

	// launch a goroutine for each name we want to process.
	for _, name := range names {
		// launch a goroutine to fetch the url.
		go func(name string) {
			// decrement the counter when the goroutine completes.
			defer wg.Done()

			// call lock inside the goroutine
			mu.Lock()

			// to ensure that the lock gets released regardless of what happens inside the goroutine.
			defer mu.Unlock()

			// update the map using the name as a key and the length of name as the value
			namesMap[name] = len(name)

			/* pass in a copy of the name to avoid the risk of the goroutines closing 
			over under the wrong name value as the for-loop iteration continues.
			It is a common practice to avoid shadowing the wrong variable.
			*/ 
		}(name)  
	}
	// wait for all goroutines to finish
	wg.Wait()

	// print the map
	fmt.Println(namesMap)
}
