package cleanup

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	cleanupOnce   sync.Once
	fns           []func()
	m             sync.Mutex
	once          sync.Once
	signalChannel chan os.Signal
)

func Wait() {
	once.Do(func() {
		signalChannel = make(chan os.Signal, 1)
		signal.Notify(signalChannel, syscall.SIGTERM, syscall.SIGINT)
	})
	<-signalChannel
	log.Println("cleanup: receive a interrupt signal, exit")
	Run()
	os.Exit(0)
}

func Add(f func()) {
	m.Lock()
	fns = append(fns, f)
	m.Unlock()
}

func Run() {
	cleanupOnce.Do(func() {
		log.Printf("cleanup: performing %d cleanups", len(fns))
		for _, f := range fns {
			f()
		}
		log.Println("cleanup: all cleanup done.")
	})
}
