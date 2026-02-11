package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	parentCtx, parentCancel := context.WithCancel(context.Background())
	middleCtx, middleCancel := context.WithCancel(parentCtx)
	childCtx, childCancel := context.WithCancel(middleCtx)

	go ParentsGorutines(parentCtx)
	go ParentsGorutines(parentCtx)
	go ParentsGorutines(parentCtx)
	go ParentsGorutines(parentCtx)

	go MiddleGorutines(middleCtx)
	go MiddleGorutines(middleCtx)

	go ChildGorutines(childCtx)
	go ChildGorutines(childCtx)
	go ChildGorutines(childCtx)

	time.Sleep(1 * time.Second)
	middleCancel()

	time.Sleep(1 * time.Second)
	childCancel()

	time.Sleep(1 * time.Second)
	parentCancel()

	time.Sleep(5 * time.Second)

}

func ParentsGorutines(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Завершился родительский контекст!")
			return
		default:
			fmt.Println("Работает горутина первой группы")
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func MiddleGorutines(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Завершился средний контекст!")
			return
		default:
			fmt.Println("Работает горутина второй группы")
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func ChildGorutines(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Завершился дочерний контекст!")
			return
		default:
			fmt.Println("Работает горутина третей группы")
		}
		time.Sleep(100 * time.Millisecond)
	}
}
