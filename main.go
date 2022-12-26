package main

import (
	"fmt"
	"time"
)

func on_tick(t *Timer) {
	t.show()
}

func on_finish(t *Timer) {
	fmt.Printf("Finished timer!")
}

func main() {
	timer := create_timer(on_tick, on_finish)
	timer.set(TEST_TIMER)
	timer.start()

	time.Sleep(20 * time.Second)
}
