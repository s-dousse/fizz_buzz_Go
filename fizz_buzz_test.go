package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	assertErrorMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("it returns 'Fizz' when passed 3", func(t *testing.T) {
		got := FizzBuzz(3)
		want :=  "Fizz"

		assertErrorMessage(t, got, want)
	})

	t.Run("it returns 'FizzBuzz' when passed 15", func(t *testing.T) {
		got := FizzBuzz(15)
		want :=  "FizzBuzz"

		assertErrorMessage(t, got, want)
	})

	t.Run("it returns 'Buzz' when passed 5", func(t *testing.T) {
		got := FizzBuzz(5)
		want :=  "Buzz"

		assertErrorMessage(t, got, want)
	})
}