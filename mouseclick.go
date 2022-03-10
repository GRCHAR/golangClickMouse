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
		resultCh := make(chan int, 8)
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
			fmt.Println("开始连点")
			for i := 0; i < 8; i++ {
				sw.Add(1)

				go func() {
					count := 0
					for {
						select {
						case <-ch:
							resultCh <- count
							sw.Done()
							return
						default:
							// x, y := robotgo.GetMousePos()
							robotgo.Click(`left`, false)
							count++
							time.Sleep(1 * time.Millisecond)
						}
					}
				}()

			}

		}

		sw.Wait()
		totalCount := 0

		for i := 0; i < 8; i++ {
			totalCount += <-resultCh
		}

		close(resultCh)
		close(ch)
		fmt.Println("停止连点", totalCount)
	}
}
