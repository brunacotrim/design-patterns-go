package main

import "log"

func main() {

	test1 := &method_1{}
	t1 := template{
		it: test1,
	}
	t1.templateMethod()

	test2 := &method_2{}
	t2 := template{
		it: test2,
	}
	t2.templateMethod()

}

type iTemplate interface {
	operation_1()
	operation_3()
}

type template struct {
	it iTemplate
}

func (tm *template) operation_2() {
	log.Println("Template - Operation 2")
}

func (tm *template) operation_4() {
	log.Println("Template - Operation 4")
}

func (tm *template) templateMethod() {
	tm.it.operation_1()
	tm.operation_2()
	tm.it.operation_3()
	tm.operation_4()
}

type method_1 struct {
	template
}

type method_2 struct {
	template
}

func (m1 *method_1) operation_1() {
	log.Println("Method 1 - Operation 1")
}

func (m1 *method_1) operation_3() {
	log.Println("Method 1 - Operation 3")
}

func (m1 *method_1) hook() {
	log.Println("Method 1 - Hook")
}

func (m1 *method_2) operation_1() {
	log.Println("Method 2 - Operation 1")
}

func (m1 *method_2) operation_2() {
	log.Println("Method 2 - Operation 2")
}

func (m1 *method_2) operation_3() {
	log.Println("Method 2 - Operation 3")
}
