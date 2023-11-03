package common

import (
	"time"
)

type Executable interface {
	Run(chan bool)
}

func StartTicker(interval time.Duration, service Executable, tickerStopSignal chan bool) {
	serviceStopSignal := make(chan bool)
	ticker := time.NewTicker(interval)
	go func() {
		for {
			select {
			case <-serviceStopSignal:
				ticker.Stop()
				tickerStopSignal <- true
			case <-ticker.C:
				go service.Run(serviceStopSignal)
			}
		}
	}()
}
