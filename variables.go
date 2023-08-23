package main
import(
	"fmt"
	"math/cmplx"
)

func main(){
	// var a, b, c int
	// var x = 2
	// var y = 2
	// var p, q = 7, 8
	// var r, s int = 10, 11
	// var j, k, l = 123, "Sujal", true
	// fmt.Println(a, b, c, x, y, p, q, r, s, j, k, l);	
	// r := foo()   // ok, declare a new variable r
	// r, m := bar()   // ok, declare a new variable m and assign r a new value
	// r, m := bar2()  //compile error: no new variables
	var(
		ToBe bool = false
		name string = "sujal gupta"
		z complex128 = cmplx.Sqrt(-5 + 2i)
	)
	fmt.Println(ToBe, name, z)
}