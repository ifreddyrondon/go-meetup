package capture

import (
	"encoding/json"
	"sort"
	"sync"
)

const WorkersNumber = 4

type indexCapture struct {
	index int
	*Capture
}

type job struct {
	index int
	data  json.RawMessage
}

func worker(wg *sync.WaitGroup, jobs <-chan job, results chan<- indexCapture) {
	for job := range jobs {
		var c Capture
		if err := json.Unmarshal(job.data, &c); err == nil {
			results <- indexCapture{index: job.index, Capture: &c}
		}
		wg.Done()
	}
}

// START OMIT
func (p *Captures) UnmarshalJSON(data []byte) error {
	var raw []json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		panic(err)
	}
	var wg sync.WaitGroup
	wg.Add(len(raw))
	jobs := make(chan job, len(raw))
	results := make(chan indexCapture, len(raw))
	for w := 0; w < WorkersNumber; w++ {
		go worker(&wg, jobs, results)
	}
	for i, v := range raw {
		jobs <- job{index: i, data: v}
	}
	close(jobs)
	wg.Wait()
	close(results)
	// ...
	// END OMIT
	processed := make([]indexCapture, 0, len(results))
	for data := range results {
		processed = append(processed, data)
	}
	sort.Slice(processed, func(i, j int) bool {
		return processed[i].index < processed[j].index
	})
	for _, v := range processed {
		*p = append(*p, *v.Capture)
	}
	return nil
}
