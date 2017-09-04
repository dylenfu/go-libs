package base

import (
	"reflect"
	"fmt"
	"log"
)

type Foo struct {
	X string
	Y int
}

func (f Foo) Test1() {
	fmt.Printf("X is %s,Y is %d", f.X, f.Y)
}

// show use of typeof.name,string...
func ReflectDemo1() {
	var i int = 123
	var s string = "hi"
	var a []string = []string{"ab", "abc", "abcd"}
	var f Foo

	fmt.Printf("i is type of %s\n", reflect.TypeOf(i))
	fmt.Printf("s is type of %s\n", reflect.TypeOf(s).Name())
	fmt.Printf("a is type of %s\n", reflect.TypeOf(a).String())
	fmt.Printf("f is type of %s\n", reflect.TypeOf(f).Name())
}

// show use of num fields
func ReflectDemo2() {
	var f Foo
	typ := reflect.TypeOf(f)
	for i:=0;i<typ.NumField();i++ {
		field := typ.Field(i)
		fmt.Printf("%s type is :%s\n", field.Name, field.Type)
	}

	field1,_:= typ.FieldByName("X")
	fmt.Printf("Foo.X field name is %s and type is %s and index is %d\n", field1.Name, field1.Type, field1.Index)
}

// show use of reflect method,method should not be internal
func ReflectDemo3() {
	var f Foo
	typ := reflect.TypeOf(f)

	m := typ.Method(0)
	fmt.Printf("Foo method num is %d,method name is %s,method type is %s\n", typ.NumMethod(), m.Name, m.Type)

	fmt.Println("Foo method is ", m.Func.String())
}

// show use of reflect value
func ReflectDemo4() {
	var i int= 12
	var s string = "hi"
	var f Foo = Foo{X:"ha", Y:12}

	fmt.Println("value of i:", reflect.ValueOf(i))
	fmt.Println("value of s:", reflect.ValueOf(s))
	fmt.Println("value of f:", reflect.ValueOf(f))
}

// show use of reflect interface
func ReflectDemo5()  {
	var i int = 123
	fmt.Println(reflect.ValueOf(i).Interface()) //123

	var f = Foo{"abc", 123}
	fmt.Println(f) //{abc 123}
	fmt.Println(reflect.ValueOf(f).Interface() == f)  //true
	fmt.Println(reflect.ValueOf(f).Interface())  //{abc 123}
}

// show use of reflect method and call
func ReflectDemo6() {
	f := &Foo{"abc", 1}

	var params []reflect.Value
	params = append(params, reflect.ValueOf(3))
	params = append(params, reflect.ValueOf(4))
	sum := reflect.ValueOf(f).MethodByName("Do1").Call(params)
	println(sum[0].Int())

	// we use []reflect.Value{} replace of nil
	// reflect.ValueOf(f).MethodByName("Do2").Call()
	reflect.ValueOf(f).MethodByName("Do2").Call([]reflect.Value{})
}

func (f Foo) Do1(num1,num2 int) int {
	return num1 + num2 + f.Y
}

func (f *Foo) Do2() {
	println("hello Foo.X:" + f.X)
}


type Human struct {
	Name string
	Age  int
}

type Student struct {
	Human
	Class string
}

func JudgeType() {
	h := Human{Name:"dylenfu", Age: 12}
	s := &Student{h,"big"}

	if reflect.TypeOf(h) == reflect.TypeOf(Human{}) {
		log.Println("h is human")
	}

	if reflect.TypeOf(s) == reflect.TypeOf(&Student{}) {
		log.Println("s is student")
	}
}