package main

import "fmt"

type Employee int

func (e Employee) Demo(){
	fmt.Println("Welcome in codelangs",e)
}

func main() {
	//using struct
	type Employee struct{
		name string
		age int
	}

	var employee Employee = Employee{name: "Raju",age: 22}
	fmt.Println(employee)


	//using map
	// type x map[string]int
	// var y x = make(x)
	// y["ram"] = 10
	// y["Rani"] = 20
	// fmt.Println(y)

	//type defination using slice
	// type x []int
	// var y x =[]int{10,20,30}
	// fmt.Println(y)

	//type defination using array
	// type x [3]int
	// var y x = [3]int{10,20,30}
	// fmt.Println(y)

	//type defination using boolean
	// type x bool
	// var y x = false
	// fmt.Println(y)


//type defination in complex64

// type x complex64
// var y x = 1 + 3i
// fmt.Println(y)





// type defination in integer
	// type x int
	// var y x = 10
	// fmt.Println(y)






	//golang only allow explict conversion

	// type Demo int
	// var y Demo = 20
	// //var x int = y //emplict conversation //error 
	// var z int = int(y)//explict conversation
	// fmt.Println(z)



	//type safety

	// type x int 
	// type y int

	// var z x = 10
	// var m y = 20
	// var d x = 30

	// z = d
	// fmt.Println(z,m,d)






	// var x Employee = 10
	// //fmt.Println(x)
	// x.Demo()

	// type name string //type defination declaration

	// var nam name = "Ram"
	// fmt.Println(nam)

}