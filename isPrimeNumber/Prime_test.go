package isPrimeNumber

import "testing"

func TestPrimeNumber(t * testing.T) {
	TestCases := []struct{
		input int
		output bool
	}{
		{0, false},
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{7, true},
		{8, false},
	}
	for index, testcase := range TestCases {
		if res := isPrimeNum(testcase.input); res != testcase.output {
			t.Errorf("The case -> %v, Expected %v, Got %v\n", testcase, testcase.output, res)
			t.Fatalf("ERROR IN %v test\n", index)
		} else {
			t.Log("PASSED")
		}
	}
}