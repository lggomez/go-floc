package run

import floc "github.com/workanator/go-floc"

/*
If runs the job if the condition is met.

Summary:
	- Run jobs in goroutines : NO
	- Wait all jobs finish   : YES
	- Run order              : SEQUENCE
*/
func If(predicate floc.Predicate, job floc.Job) floc.Job {
	return func(flow floc.Flow, state floc.State, update floc.Update) {
		if predicate(flow, state) {
			job(flow, state, update)
		}
	}
}
