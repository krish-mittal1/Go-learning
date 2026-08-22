package main

import(
	"fmt"
)

type Vertex struct{ // struct is a collection of fields
	X int
	Y int
}

func main(){
	fmt.Println(Vertex{1,2})
	main2()
	main3()
	
}

func main2(){
	v := Vertex{1 , 2}
	v.X = 4
	fmt.Println(v.X)
}

func main3(){
	v1 := Vertex{1,2}
	v2 := Vertex{X: 1}
	v3 := Vertex{}
	p := &Vertex{1,2}

	fmt.Println(v1 , p , v2 , v3)
}