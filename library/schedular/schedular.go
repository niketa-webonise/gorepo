package schedular

import ("time")

func ExecuteAtIntervals(n time.Duration,execute func()) {

	for _= range time.Tick(n*time.Second){

		execute()

	}
}