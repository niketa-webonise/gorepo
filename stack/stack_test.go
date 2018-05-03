package stack

import "testing"

func TestEmpty(t *testing.T){

	s := NewStack()

	if s.Pop() != nil {

		t.Errorf("expected empty but not empty")
	}
}



func TestStack(t *testing.T){

	s := NewStack()
	values := [5]int{10,20,30,40,50}

	for i:=0;i<len(values);i++{

		s.Push(&Node{values[i]})

	}



	for i:=len(values)-1;i>=0;i--{

		n := s.Pop()
		//fmt.Println("The elements after popping from stack",n)

		if n.Value != values[i] {

			t.Errorf("expected:%d","actual:%d",n.Value,values[i])
		} 
	}
}


func TestSize(t *testing.T){

	s := NewStack()
	values := [6]int{10,20,30,40,50,60}
	
	for i:=0;i<len(values);i++{

		s.Push(&Node{values[i]})

	}

	if len(values) != s.count {

		t.Errorf("Expected value:%d","actaul value:%d",len(values),s.count)
	}
}

func TestPeek(t *testing.T){

	s := NewStack()
	values := [3]int{10,20,30}

		for i:=0;i<len(values);i++{

		s.Push(&Node{values[i]})

	}

	n := s.Peek()

	if values[2] != n.Value {

		t.Errorf("Expected:%d","Actual:%d",values[2],n.Value)
	}
}

