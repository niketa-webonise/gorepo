package wordy

import "testing"

var wordyTests = []struct {
	n string
	expected int
}{
	{"what is 5 plus 13?",18},
	{"what is 53 plus 2?",55},
	//{"What is -1 plus -10?",-11},
	{"What is 123 plus 5?",128},
	{"What is 10 minus 2?",8},
	{"What is 33 multiplied by 2?",66},
	{"what is 10 divided by 2?",5},
	{"what is 5 plus 10 plus 2?",17},
	{"What is 5 plus 2 minus 1?",6},
	{"What is 5 power to 2?",25},
	{"What is 20 minus 4 minus 13?",3},


}

func TestWordy(t *testing.T) {

	for _,result := range wordyTests {
		actual := Wordy(result.n)
		if actual != result.expected {
			t.Errorf("Wordy(%s): expected %d, actual %d", result.n, result.expected, actual)
		}
	}
}
