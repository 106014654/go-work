package api

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	syscalls "syscall"
	"time"
)

func main(){
	g, ctx := errgroup.WithContext(context.Background()) //创建空白上下文

	mux := http.NewServeMux() //创建多个http访问


	serverOut := make(chan struct{}) //服务断联信号
	mux.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		serverOut <- struct{}{}
	})
	server := http.Server{
		Handler: mux,
		Addr:    ":xxxx",
	}

	g.Go(func() error {
		select {
		case <-ctx.Done(): //直接终止程序
			log.Println("errgroup exit...")
		case <-serverOut: //手动停止
			log.Println("server will out...")
		}

		timeoutCtx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
		defer cancel()

		log.Println("shutting down server...")
		return server.Shutdown(timeoutCtx) //关闭请求连接
	})

	g.Go(func() error {
		quit := make(chan os.Signal)
		signal.Notify(quit, syscalls.SIGINT, syscalls.SIGTERM) //监听 中断，结束信号

		select {
		case <-ctx.Done():
			return ctx.Err()
		case sig := <-quit:
			return errors.Errorf("get os signal: %v", sig)
		}
	})

	if err := g.Wait(); err != nil {
		fmt.Println("group error: ", err)
	}
}
