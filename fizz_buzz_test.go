package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	t.Run("it returns 'Fizz' when passed 3", func(t *testing.T) {
		got := FizzBuzz(3)
		want :=  "Fizz"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}