package grace

import (
	"os"
	"os/signal"
)

// Shutdown accept a callback function which is called
// after signal has been received.
func Shutdown(callback func(), signals ...os.Signal) {
	defer callback()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, signals...)
	<-ch
}
