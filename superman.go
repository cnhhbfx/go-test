package main

import (
	"fmt"
	"time"
)

// 01
type SuperMan interface {
	Fly()
}

type IntMan struct {
	Speed int
}

func (i *IntMan) Fly() {
	fmt.Println("IntMan can fly, speed: ", i.Speed)
}

type FuncMan struct {
	SpeedFunc func(int, int) int
	X         int
	Y         int
}

func (f *FuncMan) Fly() {
	fmt.Println("FuncMan can fly, speed: ", f.SpeedFunc(f.X, f.Y))
}

func JiaFa(x int, y int) int {
	return x + y
}

func ChengFa(x int, y int) int {
	return x * y
}

// 02
func GoSafe(fn func()) {
	go RunSafe(fn)
}

func RunSafe(fn func()) {
	defer Recover()

	fn()
}

func Recover(cleanups ...func()) {
	for _, cleanup := range cleanups {
		cleanup()
	}

	if p := recover(); p != nil {
		fmt.Println(p)
	}
}

func Plus() {
	fmt.Println("result ", 1/0)
}

func main() {
	//01
	var k SuperMan = &IntMan{Speed: 7}
	k.Fly()

	var jiaFa SuperMan = &FuncMan{SpeedFunc: JiaFa, X: 7, Y: 7}
	jiaFa.Fly()

	var chengFa SuperMan = &FuncMan{SpeedFunc: ChengFa, X: 7, Y: 7}
	chengFa.Fly()

	// 02
	GoSafe(Plus)

	time.Sleep(5)
}
