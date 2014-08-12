package euler

// Une dois canais 
func FanIn( in1, in2 <-chan int64 ) <-chan int64 {
	out := make(chan int64)
	go func() {
		for {
			select {
			case s := <-in1: out <-s
			case s := <-in2: out <-s
			}
		}
	}()
	return out
}