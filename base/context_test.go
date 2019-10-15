package base

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// cancel必须自己主动实现，cancel()方法只是发送信号，通知goroutine return
func TestSimpleContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, "sf1", "hahah")

	go sf1(ctx)
	time.Sleep(3 * time.Second)
	cancel()
	fmt.Println("----stop sf1")

	time.Sleep(5 * time.Second)
}

func sf1(ctx context.Context) {
	for {
		fmt.Println(ctx.Value("sf1"))
		time.Sleep(1 * time.Second)
	}
}

// context.TODO&BACKGROUD底层实现时emptyContext，他是一个int值，代表同一个context,使用context.todo一般是要生成
// child
func TestTodoSingleton(t *testing.T) {
	ctx1 := context.TODO()
	ctx2 := context.TODO()
	if ctx1 == ctx2 {
		t.Log("context todo is singleton")
	} else {
		t.Log("context todo not singleton")
	}
}

func TestContextChildren(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	t.Logf("%v", ctx)
	ctx1 := context.WithValue(ctx, "f1", "haha1")
	t.Logf("%v", ctx1)
	ctx2 := context.WithValue(ctx1, "f2", "heihei2")
	t.Logf("%v", ctx2)
	if ctx == ctx1 {
		t.Log("ctx == ctx1")
	} else {
		t.Log("ctx != ctx1")
	}
	if ctx2 == ctx1 {
		t.Log("ctx2 == ctx1")
	} else {
		t.Log("ctx2 != ctx1")
	}

	go f1(ctx2)
	go f2(ctx2)

	time.Sleep(10 * time.Second)
	cancel()

	time.Sleep(2 * time.Second)
}

func f1(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("f1", ctx.Value("f1"))
			time.Sleep(time.Second)
		}
	}
}

func f2(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("f2", ctx.Value("f2"))
			time.Sleep(2 * time.Second)
		}
	}
}
