package codeCover

import (
	"testing"
)

func TestFibo1(t *testing.T) {
	if fibo1(1) != 1 {
		t.Errorf("Error fibo1(1): %d\n", fibo1(1))
	}
}

func TestFibo2(t *testing.T) {
	if fibo2(0) != 0 {
		t.Errorf("Error fibo2(0): %d\n", fibo1(0))
	}
}

func TestFibo1_10(t *testing.T) {
	if fibo1(10) == 1 {
		t.Errorf("Error fibo1(1): %d\n", fibo1(1))
	}
}

func TestFibo2_10(t *testing.T) {
	if fibo2(10) != 0 {
		t.Errorf("Error fibo2(0): %d\n", fibo1(0))
	}
}
