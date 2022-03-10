package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	for {
		fmt.Println("请按s开始连点,按e结束连点")
		var sw sync.WaitGroup
		ch := make(chan bool, 8)
		start := robotgo.AddEvent("s")

		go func() {
			exit := robotgo.AddEvent("e")
			if exit {
				for i := 0; i < 8; i++ {
					ch <- true
				}
			}

		}()

		if start {

			// go func() {
			// 	time.Sleep(1 * time.Second)
			// 	for i := 0; i < 8; i++ {
			// 		ch <- true
			// 	}
			// }()
			fmt.Println("开始连点")
			for i := 0; i < 8; i++ {
				sw.Add(1)
				go func() {
					for {
						select {
						case <-ch:
							sw.Done()
							return
						default:
							// x, y := robotgo.GetMousePos()
							robotgo.Click(`left`, false)
							time.Sleep(1 * time.Millisecond)
						}
					}
				}()
			}

		}
		sw.Wait()
		fmt.Println("停止连点")
	}
}
