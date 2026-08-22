package main

import(
	"fmt"
	"math"
	"runtime"
)	

func main(){
	sum := 0;
	for i:= 0 ; i <10 ; i++{ // go has only one looping construct which is for loop and it can be used as while and do while loop as well
		sum += i
	}
	fmt.Println(sum)

	main2() // other function call from main functiion

	fmt.Println(if_else(2) , if_else(-4)) // calling if_else function

	fmt.Println(
		pow(3, 2 , 10),
		pow(3 , 3 , 20),
	)

	fmt.Println(
		pow2(3, 2 , 10),
		pow2(3 , 3 , 20),
	)

	switch_case() // calling switch_case function

	defer_func() // calling defer_func function
	defer2() // calling defer2 function
}

func main2(){
	sum := 1
	for sum < 1000{ // for loop can be used as while loop as well
		sum += sum
	}
	fmt.Println(sum)
}

func if_else(x float64) string{
	if x < 0{
		return if_else(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n ,lim float64) float64{
	if v := math.Pow(x, 2); v< lim{ // we can also define variables inside if statement and it will be available only inside the if block
		return v
	}
	return lim
}

func pow2(x , n , lim float64) float64{ 
	if v := math.Pow(x , n); v < lim{
		return v
	} else{
		fmt.Printf("%g >= %g\n", v, lim)
	}
		return lim
}

func switch_case(){ // switch statement is used to select one of the many code bloacks to be executed. It is similar to if-else statement but it is more readable and efficient than if-else statement
	fmt.Print("Go runs on ")
	switch os:= runtime.GOOS; os{ // runtime.GOOS -> returns the operating system name , ed windows , linuc , macOS
	case "darwin":
		fmt.Println("macOs.")
		case "linux":
			fmt.Println("Linux.")
		default:
			fmt.Printf("%s.\n", os)
	}
}

func defer_func(){
	defer fmt.Println("world") // defer statement is used to delay the execution of a function until the surrounding function returns. It is used to ensure that a function call is performed later in a program’s execution, usually for purposes of cleanup. defer is often used where e.g. ensure and finally would be used in other languages.
	fmt.Println("hello")
}

func defer2(){
	fmt.Println("counting")
	for i := 0 ; i < 10 ; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}