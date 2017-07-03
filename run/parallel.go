package run

import floc "github.com/workanator/go-floc"

// Parallel runs jobs in parallel and waits until all of them done.
func Parallel(jobs ...floc.Job) floc.Job {
	return func(flow floc.Flow, state floc.State, update floc.Update) {
		// Do not start parallel jobs if the execution is finished
		if flow.IsFinished() {
			return
		}

		// Create channel which is used for back counting of finished jobs
		done := make(chan int, len(jobs))
		defer close(done)

		// Run jobs in parallel
		running := 0
		for index, job := range jobs {
			running++

			go func(index int, job floc.Job) {
				// Write the index of the finished job
				defer func() { done <- index }()
				// Do the job
				job(flow, state, update)
			}(index, job)
		}

		// Wait until all jobs done
		for running > 0 {
			select {
			case <-flow.Done():
				// The execution is finished
				return

			case <-done:
				// One of the jobs finished
				running--
			}
		}
	}
}
