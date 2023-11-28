package main

import "fmt"

type old struct {
	v int
}

type some struct {
	o  old
	op *old
}

func (s some) double() {
	s.o.v = s.o.v * 2
	s.op.v = s.op.v * 2
}

func (s *some) doubleP() {
	s.o.v = s.o.v * 2
	s.op.v = s.op.v * 2
}

type option interface {
	double()
}

type optionP interface {
	doubleP()
	double()
}

func main() {
	var aa some = some{o: old{1}, op: &old{3}}
	fmt.Printf("%v %d\n", aa, aa.op.v)
	aa.double()
	fmt.Printf("%v %d\n", aa, aa.op.v)
	aa.doubleP()
	fmt.Printf("%v %d\n", aa, aa.op.v)

	//var ab some
	//fmt.Printf("%v %d\n", ab, ab.op.v)

	var ac *some = &some{o: old{2}, op: &old{4}}
	fmt.Printf("%v %d\n", ac, ac.op.v)
	ac.double()
	fmt.Printf("%v %d\n", ac, ac.op.v)
	ac.doubleP()
	fmt.Printf("%v %d\n", ac, ac.op.v)

	//var ad *some
	//fmt.Printf("%v %d\n", ad, ad.op.v)

	var ae option = some{o: old{3}, op: &old{6}}
	fmt.Printf("%v\n", ae)
	ae.double()
	fmt.Printf("%v\n", ae)
	
	var af optionP = &some{o: old{3}, op: &old{6}}
	fmt.Printf("%v\n", af)
	af.double()
	fmt.Printf("%v\n", af)
	af.doubleP()
	fmt.Printf("%v\n", af)
}
