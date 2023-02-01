package add

import "fmt"

// Write a Go program that implements a simple calculator that can add,
// subtract, multiply and divide two numbers. The program should take
// the operation to be performed as a command-line argument, along with the numbers.
// Note : make another function to read from the cli then pass the input as a parameter
// 		to the calculator function

func Calculator(operator string, number ...float64) float64 {

	var result float64=number[0]
	switch operator{
	case"add":
		for _,i:=range number[1:]{
			result=result+i

		}
	
	case"subtract":
		for _,i:=range number[1:]{
			result=result-i

		}
		return result
	case"multiply":
		for _,i:=range number[1:]{
			result=result*i

		}
	case"divide":
		for _,i:=range number[1:]{
			result=result/i

		}
		
	}
	//TODO: your code here
	return result
}

func readFromCommandLine() {
	var x,y float64
	var opr string
	fmt.Scanln(&opr)
	fmt.Scanln(&x)
	fmt.Scanln(&y)
	result:=Calculator(opr,x,y)
	fmt.Printf("the sum of and is: %f",result)


	// should accept minimum 3 arguments
	// call the calculator with the arguments

	//TODO: your code here
}
