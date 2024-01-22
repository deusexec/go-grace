package grace

import (
	"os"
	"os/signal"
)

func Shutdown(callback func(), signals ...os.Signal) {
	defer callback()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, signals...)
	<-ch
}
