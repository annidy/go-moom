package main

import (
	"fmt"
	"go-moom/internal"
	"log"
	"time"
)

func main() {
	fmt.Println("moom supervisor start...")
	ch := make(chan int)
	go func() {
		for  {
			cf := internal.Config{}
			if err := cf.ReadConfig(); err != nil {
				log.Fatal(err)
			}
			log.Printf("count = %v\n", cf.Count)
			if cf.Count > 15 {
				cf.Count = 1
				cf.WriteConfig()
				internal.Restart()
			}
			time.Sleep(time.Minute * 10)
		}
	}()
	// 最简单的写法
	<-ch
}
