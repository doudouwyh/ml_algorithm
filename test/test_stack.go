package test

import (
	"wh_ml/lib"
	"fmt"
)

func TestStack() {
	stack := lib.NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	len := stack.Len()
	if len != 4 {
		fmt.Println("stack.Len() failed. Got:",len," expected 4.")
	}

	value := stack.Peak().(int)
	if value != 4 {
		fmt.Println("stack.Peak() failed. Got:",value," expected 4.")
	}

	value = stack.Pop().(int)
	if value != 4 {
		fmt.Println("stack.Pop() failed. Got:",value," expected 4.")
	}

	len = stack.Len()
	if len != 3 {
		fmt.Println("stack.Len() failed. Got:",len," expected 3.")
	}

	value = stack.Peak().(int)
	if value != 3 {
		fmt.Println("stack.Peak() failed. Got:",value," expected 3.")
	}

	value = stack.Pop().(int)
	if value != 3 {
		fmt.Println("stack.Pop() failed. Got:",value," expected 3.")
	}

	value = stack.Pop().(int)
	if value != 2 {
		fmt.Println("stack.Pop() failed. Got:",value," expected 2.")
	}

	empty := stack.Empty()
	if empty {
		fmt.Println("stack.Empty() failed. Got:",empty," expected false.")
	}

	value = stack.Pop().(int)
	if value != 1 {
		fmt.Println("stack.Pop() failed. Got:",value," expected 1.")
	}

	empty = stack.Empty()
	if !empty {
		fmt.Println("stack.Empty() failed. Got:",empty," expected true.")
	}

	nilValue := stack.Peak()
	if nilValue != nil {
		fmt.Println("stack.Peak() failed. Got:",nilValue," expected nil.")
	}

	nilValue = stack.Pop()
	if nilValue != nil {
		fmt.Println("stack.Pop() failed. Got:",nilValue," expected nil.")
	}

	len = stack.Len()
	if len != 0 {
		fmt.Println("stack.Len() failed. Got:",len," expected 0.")
	}
}
