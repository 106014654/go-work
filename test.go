package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"io"
	"net/http"
	"os"
	"os/signal"
	"time"
)

//启动 HTTP server
func StartHttpServer(server *http.Server) error {

	http.HandleFunc("/hello", HelloServer2)
	fmt.Println("http server start")
	err := server.ListenAndServe()
	return err
}

func HelloServer2(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	group, errCtx := errgroup.WithContext(ctx)

	server := &http.Server{Addr: ":8080"}

	group.Go(func() error {
		return StartHttpServer(server)
	})

	group.Go(func() error {
		time.Sleep(5 * time.Second)
		<-errCtx.Done() //cancel、timeout、deadline 导致 Done 被 close
		fmt.Println("http server stop")
		return server.Shutdown(errCtx)
	})

	chanel := make(chan os.Signal, 1) //这里要用 buffer 为1的 chan
	signal.Notify(chanel)

	group.Go(func() error {
		for {
			select {
			case <-errCtx.Done(): // 因为 cancel、timeout、deadline 都可能导致 Done 被 close
				return errCtx.Err()
			case <-chanel: // 因为 kill -9 或其他而终止
				cancel()
			}
		}
		return nil
	})

	if err := group.Wait(); err != nil {
		fmt.Println("group error: ", err)
	}
	fmt.Println("all group done!")

}
