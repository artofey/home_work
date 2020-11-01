package hw06_pipeline_execution //nolint:golint,stylecheck

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func doneStreamer(done In, in In) Out {
	out := make(Bi)
	go func() {
		defer close(out)
		select {
		case <-done:
			return
		default:
		}

		for v := range in {
			select {
			case <-done:
				return
			case out <- v:
			}
		}
	}()
	return out
}

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	inSt := doneStreamer(done, in)

	for _, stage := range stages {
		inSt = stage(inSt)
	}
	return inSt
}
