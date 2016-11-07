package linenotcat

import (
	"fmt"
	"os"
	"time"
)

const (
	interval = 3 * time.Second
)

type stream struct {
	queue        *queue
	shutdown     chan os.Signal
	lineNotifier *lineNotifier
}

func newStream(token string) *stream {
	return &stream{
		queue:        newQueue(),
		shutdown:     make(chan os.Signal, 1),
		lineNotifier: &lineNotifier{token: token},
	}
}

func (s *stream) trap() {
	sigcount := 0
	for sig := range s.shutdown {
		if sigcount > 0 {
			fmt.Println("Aborted")
			os.Exit(1)
		}
		fmt.Printf("Got signal: %s\n", sig.String())
		fmt.Println("Press Ctrl+C again to exit immediately")
		sigcount++
		go s.exit()
	}
}

func (s *stream) exit() {
	for {
		if s.queue.isEmpty() {
			os.Exit(0)
		} else {
			fmt.Println("Flushing remaining messages...")
			time.Sleep(interval)
		}
	}
}

func (s *stream) processStreamQueue(tee bool) {
	if !s.queue.isEmpty() {
		lines := s.queue.flush()
		s.lineNotifier.notifyMessages(lines, tee)
	}
	time.Sleep(interval)
	s.processStreamQueue(tee)
}

func (s *stream) watchStdin() {
	for _ = range time.Tick(interval) {
		lines := make(chan string)
		go readFromStdin(lines)
		for line := range lines {
			s.queue.add(line)
		}
	}
}
