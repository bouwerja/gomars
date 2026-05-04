package roulette

import (
	"runtime"
	"sync"
)

func Run(trials int) TrialsResult {
	var r TrialsResult
	r.Trial = make([]Trials, 0, trials)

	jobs := make(chan struct{}, trials)
	resultsChan := make(chan Trials, trials)

	numWorkers := runtime.NumCPU()
	var wg sync.WaitGroup

	for range numWorkers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range jobs {
				resultsChan <- RunGame()
			}
		}()
	}

	go func() {
		for range trials {
			jobs <- struct{}{}
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	for t := range resultsChan {
		r.Trial = append(r.Trial, t)
	}

	return r
}
