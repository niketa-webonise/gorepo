package sum

import "testing"

func TestSum(t *testing.T) {

	SumTests := []struct{
		x int
		y int
		result int
	}{
		{1,1,2},
		{10,10,20},
	}

	for _,test1 := range SumTests {

		total := Sum(test1.x,test1.y)

		if total != test1.result {

			
			t.Errorf("Sum of (%d+%d) was incorrect, got: %d, want: %d.", test1.x, test1.y, total,test1.result)
		}
	}
}