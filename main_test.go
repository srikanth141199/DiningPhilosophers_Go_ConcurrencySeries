package main

import (
	"testing"
	"time"
)

func Test_dine(t *testing.T) {
	// Run the dine function in a separate goroutine

	eatTime = 100 * time.Millisecond
	sleepTime = 100 * time.Millisecond
	thinkTime = 100 * time.Millisecond

	for i := 0; i < 10; i++ {
		orderFinished = []string{}
		dine()
		// Check that all philosophers have finished eating
		if len(orderFinished) != len(philosophers) {
			t.Errorf("Expected %d philosophers to finish eating, but got %d", len(philosophers), len(orderFinished))
		}
	}
}
