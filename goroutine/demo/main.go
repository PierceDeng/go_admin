package main

import (
	"fmt"
)

var intChan chan int

type Dog struct {
	Name string
}

//func main() {
//
//	intChan = make(chan int, 10)
//	intChan <- 1
//	fmt.Println("intchan %v,address :%v", <-intChan, &intChan)
//	fmt.Println("intchan size: %v,cap :%v", len(intChan), cap(intChan))
//
//	allChan := make(chan interface{}, 10)
//
//	allChan <- &Dog{
//		Name: "wawa",
//	}
//
//	allChan <- 1
//	allChan <- "hhh"
//
//	//dog := <-allChan
//	//dogObj := dog.(*Dog)
//	//fmt.Println(dogObj.Name)
//
//	close(allChan)
//
//	for {
//		val, ok := <-allChan
//		if !ok {
//			break
//		}
//		fmt.Println("val:", val)
//	}
//
//}

//func main() {
//
//	intChan = make(chan int, 10)
//	exitChan := make(chan bool, 1)
//
//	go writeData(intChan)
//	go readDate(intChan, exitChan)
//
//	if <-exitChan {
//		return
//	}
//	fmt.Println("done.")
//
//	var waitGroup sync.WaitGroup
//	waitGroup.Add(2)
//
//	for {
//		select {
//		case <-intChan:
//			waitGroup.Done()
//		default:
//			waitGroup.Wait()
//		}
//	}
//
//}

func writeData(intChan chan int) {

	for i := 0; i < 50; i++ {
		var tempInt int
		tempInt = i
		intChan <- tempInt
		fmt.Println("writeData:", tempInt)
	}
	close(intChan)
}

func readDate(iniChan chan int, exitChan chan bool) {

	var count int
	for {
		val, ok := <-iniChan
		if !ok {
			count++
			break
		}
		fmt.Println("readDate:", val)
	}
	exitChan <- true
	close(exitChan)

}

var times = 100

func main() {

	dogChan := make(chan struct{}, 1)
	catChan := make(chan struct{}, 1)
	fishChan := make(chan struct{}, 1)
	exitChan := make(chan bool, 1)

	for i := 1; i <= times; i++ {
		go dog(fishChan, dogChan, i == 1)
		go cat(dogChan, catChan)
		go fish(catChan, fishChan, exitChan, i)
	}

	for {
		if <-exitChan {
			close(dogChan)
			close(catChan)
			close(fishChan)
			close(exitChan)
			break
		}
	}

	fmt.Println("done")

}

//func main() {
//	NewTimes := 3
//	for i := 1; i <= NewTimes; i++ {
//		var wg sync.WaitGroup
//		dogCh := make(chan struct{})
//		catCh := make(chan struct{})
//		fishCh := make(chan struct{})
//
//		wg.Add(3)
//		go func() {
//			defer wg.Done()
//			<-dogCh // 等待触发
//			fmt.Println("dog")
//			catCh <- struct{}{}
//		}()
//		go func() {
//			defer wg.Done()
//			<-catCh
//			fmt.Println("cat")
//			fishCh <- struct{}{}
//		}()
//		go func() {
//			defer wg.Done()
//			<-fishCh
//			fmt.Println("fish")
//		}()
//		// 启动第一个
//		dogCh <- struct{}{}
//		wg.Wait()
//	}
//	fmt.Println("done")
//}

func dog(fishChan chan struct{}, dogChan chan struct{}, first bool) {
	if !first {
		<-fishChan
	}

	fmt.Println("dog")
	dogChan <- struct{}{}
}

func cat(dogChan chan struct{}, catChan chan struct{}) {
	<-dogChan
	fmt.Println("cat")
	catChan <- struct{}{}
}

func fish(catChan chan struct{}, fishChan chan struct{}, exitChan chan bool, i int) {
	<-catChan
	fmt.Println("fish")
	fishChan <- struct{}{}
	if i == times {
		exitChan <- true
	}
	fmt.Println("打印循环第?次", i)

}
