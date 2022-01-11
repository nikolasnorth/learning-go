package main

import (
	"log"
	"testing"
)

func TestParseInputToFloat(t *testing.T) {
	t.Run("with two floats", func(t *testing.T) {
		expectedA := 2.0
		expectedB := 1.15
		actualA, actualB, err := ParseInputToFloat("2.0", "1.15")
		if err != nil {
			log.Fatalf("got an error but didn't expect one %q", err.Error())
		}
		if actualA != expectedA {
			log.Fatalf("got %f want %f", actualA, expectedA)
		}
		if actualB != expectedB {
			log.Fatalf("got %f want %f", actualB, expectedB)
		}
	})
	t.Run("with one float and one integer", func(t *testing.T) {
		expectedA := 2.3
		expectedB := 2.0
		actualA, actualB, err := ParseInputToFloat("2.3", "2")
		if err != nil {
			log.Fatalf("got an error but didn't expect one %q", err.Error())
		}
		if actualA != expectedA {
			log.Fatalf("got %f want %f", actualA, expectedA)
		}
		if actualB != expectedB {
			log.Fatalf("got %f want %f", actualB, expectedB)
		}
	})
	t.Run("with two integers", func(t *testing.T) {
		expectedA := 3.0
		expectedB := 2.0
		actualA, actualB, err := ParseInputToFloat("3", "2")
		if err != nil {
			log.Fatalf("got an error but didn't expect one %q", err.Error())
		}
		if actualA != expectedA {
			log.Fatalf("got %f want %f", actualA, expectedA)
		}
		if actualB != expectedB {
			log.Fatalf("got %f want %f", actualB, expectedB)
		}
	})
}

func TestCalc(t *testing.T) {
	t.Run("with invalid operator", func(t *testing.T) {
		_, err := Calc("invalidop", 1, 1)
		if err == nil {
			log.Fatal("expected an error but didn't get one")
		}
	})

	t.Run("adding", func(t *testing.T) {
		t.Run("two integers", func(t *testing.T) {
			expected := 120.0
			actual, err := Calc(Add, 100, 20)
			if err != nil {
				log.Fatalf("got an error but didn't expect one %q", err.Error())
			}
			if actual != expected {
				log.Fatalf("got %f want %f", actual, expected)
			}
		})

		t.Run("one integer and one float", func(t *testing.T) {
			expected := 120.0
			actual, err := Calc(Add, 100, 20.0)
			if err != nil {
				log.Fatalf("got an error but didn't expect one %q", err.Error())
			}
			if actual != expected {
				log.Fatalf("got %f want %f", actual, expected)
			}
		})

		t.Run("two floats", func(t *testing.T) {
			expected := 120.0
			actual, err := Calc(Add, 100.0, 20.0)
			if err != nil {
				log.Fatalf("got an error but didn't expect one %q", err.Error())
			}
			if actual != expected {
				log.Fatalf("got %f want %f", actual, expected)
			}
		})
	})

	t.Run("subtracting", func(t *testing.T) {
		t.Run("two integers", func(t *testing.T) {
			expected := 100.0
			actual, err := Calc(Subtract, 150, 50)
			if err != nil {
				log.Fatalf("got an error but didn't expect one %q", err.Error())
			}
			if actual != expected {
				log.Fatalf("got %f want %f", actual, expected)
			}
		})

		t.Run("one integer and one float", func(t *testing.T) {
			expected := 100.0
			actual, err := Calc(Subtract, 150, 50.0)
			if err != nil {
				log.Fatalf("got an error but didn't expect one %q", err.Error())
			}
			if actual != expected {
				log.Fatalf("got %f want %f", actual, expected)
			}
		})

		t.Run("two floats", func(t *testing.T) {
			expected := 100.0
			actual, err := Calc(Subtract, 150.0, 50.0)
			if err != nil {
				log.Fatalf("got an error but didn't expect one %q", err.Error())
			}
			if actual != expected {
				log.Fatalf("got %f want %f", actual, expected)
			}
		})
	})

	t.Run("multiplying", func(t *testing.T) {
		t.Run("two integers", func(t *testing.T) {
			expected := 100.0
			actual, err := Calc(Multiply, 10, 10)
			if err != nil {
				log.Fatalf("got an error but didn't expect one %q", err.Error())
			}
			if actual != expected {
				log.Fatalf("got %f want %f", actual, expected)
			}
		})

		t.Run("one integer and one float", func(t *testing.T) {
			expected := 100.0
			actual, err := Calc(Multiply, 10, 10.0)
			if err != nil {
				log.Fatalf("got an error but didn't expect one %q", err.Error())
			}
			if actual != expected {
				log.Fatalf("got %f want %f", actual, expected)
			}
		})

		t.Run("two floats", func(t *testing.T) {
			expected := 100.0
			actual, err := Calc(Multiply, 10.0, 10.0)
			if err != nil {
				log.Fatalf("got an error but didn't expect one %q", err.Error())
			}
			if actual != expected {
				log.Fatalf("got %f want %f", actual, expected)
			}
		})
	})

	t.Run("dividing", func(t *testing.T) {
		t.Run("with a zero divisor", func(t *testing.T) {
			_, err := Calc(Divide, 100, 0)
			if err == nil {
				log.Fatal("expected an error but didn't get one")
			}
		})

		t.Run("two integers", func(t *testing.T) {
			expected := 10.0
			actual, err := Calc(Divide, 100, 10)
			if err != nil {
				log.Fatalf("got an error but didn't expect one %q", err.Error())
			}
			if actual != expected {
				log.Fatalf("got %f want %f", actual, expected)
			}
		})

		t.Run("one integer by one float", func(t *testing.T) {
			expected := 10.0
			actual, err := Calc(Divide, 100, 10.0)
			if err != nil {
				log.Fatalf("got an error but didn't expect one %q", err.Error())
			}
			if actual != expected {
				log.Fatalf("got %f want %f", actual, expected)
			}
		})

		t.Run("two floats", func(t *testing.T) {
			expected := 10.0
			actual, err := Calc(Divide, 100.0, 10.0)
			if err != nil {
				log.Fatalf("got an error but didn't expect one %q", err.Error())
			}
			if actual != expected {
				log.Fatalf("got %f want %f", actual, expected)
			}
		})
	})
}
