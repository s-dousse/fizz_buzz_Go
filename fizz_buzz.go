package main

func FizzBuzz(number int) (outcome string) {
	if number%15 == 0 {
		outcome = "FizzBuzz"
	} else if number%3 == 0 {
		outcome = "Fizz"
	}
	return
}