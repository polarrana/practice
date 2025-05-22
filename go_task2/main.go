package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	//指针1
	i := 1
	index1(&i)
	fmt.Println(i)
	//指针2
	slice := []int{1, 2, 3, 4, 5}
	index2(slice)
	fmt.Println(slice)
	//面向对象1
	r := Rectangle{Width: 5, Height: 3}
	fmt.Println(r.Area())
	fmt.Println(r.Perimeter())
	c := Circle{Radius: 2}
	fmt.Println(c.Area())
	fmt.Println(c.Perimeter())
	//面向对象2
	e := Employee{
		Person: Person{
			Name: "John",
			Age:  30,
		},
		EmployeeID: 123,
	}
	e.PrintInfo()
	//channel1
	channel1 := make(chan int)
	flag1 := true
	go func() {
		for i := 1; i < 11; i++ {
			channel1 <- i
		}
		flag1 = false
	}()
	for flag1 {
		fmt.Println(<-channel1)
	}
	//channel2
	channel2 := make(chan int, 100)
	channel3 := make(chan int)
	//生产者协程
	go func() {
		for i := 1; i <= 100; i++ {
			channel2 <- i
		}
	}()
	//消费者协程
	go func() {
		for {
			num := <-channel2
			fmt.Println(num)
			if num == 100 {
				channel3 <- 0
			}
		}
	}()
	<-channel3
	//锁机制1编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
	var (
		mutex   sync.Mutex
		counter int
		wg      sync.WaitGroup
	)
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mutex.Lock()
				counter++
				mutex.Unlock()
			}
		}()
	}
	wg.Wait() // 等待所有协程完成
	fmt.Println(counter)
	//锁机制2使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
	var counter2 int64
	wg2 := sync.WaitGroup{}
	wg2.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg2.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter2, 1)
			}
		}()
	}
	wg2.Wait()
	fmt.Println(counter2)

}

// 1.指针1 编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
func index1(p *int) {
	*p += 10
}

// 2.指针2 实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
func index2(slice []int) {
	for i := 0; i < len(slice); i++ {
		slice[i] *= 2
	}
}
