package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

// Stage ...
type Stage func(in In) (out Out)

func doneStreamer(done In, in In) Out {
	out := make(Bi)
	go func() {
		defer close(out)

		for {
			select {
			case <-done:
				return
			case i, ok := <-in:
				if !ok {
					return
				}
				select {
				case <-done:
					return
				case out <- i:
				}
			}
		}
	}()
	return out
}

// ExecutePipeline ...
func ExecutePipeline(in In, done In, stages ...Stage) Out {
	for _, stage := range stages {
		in = stage(doneStreamer(done, in))
	}
	return in
}
