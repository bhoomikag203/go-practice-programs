package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const countdownStart = 3
const finalWord = "Go!"

type Sleeper interface {
	Sleep()
}

// type SpySleeper struct {
// 	Calls int
// }

// func (s *SpySleeper) Sleep() {
// 	s.Calls++
// }

// type DefaultSleeper struct{}

// func (d *DefaultSleeper) Sleep() {
// 	time.Sleep(1 * time.Second)
// }

type CountdownOperationsSpy struct {
	Calls []string
}

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

type spyTime struct {
	durationSlept time.Duration
}

func (s *spyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
