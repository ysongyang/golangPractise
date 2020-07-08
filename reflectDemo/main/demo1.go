package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(a interface{}) {
	//先获取到reflect.Type

	rtyp := reflect.TypeOf(a)
	fmt.Println("Kind = ", rtyp.Kind()) //Kind返回该接口的具体分类
	fmt.Println("Name = ", rtyp.Name())

	rVal := reflect.ValueOf(a)

	n1 := 100 + rVal.Int()

	fmt.Printf("n1 = %v \n", n1)

	fmt.Printf("rVal = %v , rVal Type = %T \n", rVal, rVal)

	iV := rVal.Interface()

	num := iV.(int)

	fmt.Printf("num = %v \n", num)
}

func reflectTest02(a interface{}) {
	//先获取到reflect.Type

	rtyp := reflect.TypeOf(a)
	fmt.Println("Kind = ", rtyp.Kind()) //Kind返回该接口的具体分类
	fmt.Println("Name = ", rtyp.Name())

	rVal := reflect.ValueOf(a)

	iV := rVal.Interface()
	fmt.Printf("iV = %v , iV Type = %T \n", iV, iV)
	str := iV.(Student)

	fmt.Printf("str = %v name = %v age = %d \n", str, str.Name, str.Age)
}

type Student struct {
	Name string
	Age  int
}

func main() {
	/*var num int = 100
	reflectTest01(num)*/
	reflectTest02(Student{"张一博", 32})

	var str string = "tom"
	fs := reflect.ValueOf(&str) //地址传入
	//fs.SetString("jack") //error
	fs.Elem().SetString("jack")
	fmt.Println("str = ", str)

}
