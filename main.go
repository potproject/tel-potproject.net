package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	telnet "github.com/reiver/go-telnet"
)

var counter int = 0

func main() {
	counter = loadCounter()
	var handler telnet.Handler = MainHandler{}
	go func() {
		err := telnet.ListenAndServe(":23", handler)
		if nil != err {
			panic(err)
		}
	}()
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt, os.Kill)
	defer stop()
	<-ctx.Done()

	saveCounter(counter)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
}

func loadCounter() int {
	// load file counter.txt
	// if not exists, create it
	c, err := os.OpenFile("counter.txt", os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer c.Close()
	// read counter
	b, err := ioutil.ReadAll(c)
	if err != nil {
		panic(err)
	}
	str := string(b)
	if str == "" {
		str = "0"
	}
	// convert to int
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}

func saveCounter(counter int) {
	// save file counter.txt
	// if not exists, create it
	c, err := os.OpenFile("counter.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer c.Close()
	// write counter
	_, err = c.Write([]byte(strconv.Itoa(counter)))
	if err != nil {
		panic(err)
	}
	fmt.Println("saved counter:", counter)
}

type MainHandler struct{}

func (handler MainHandler) ServeTELNET(ctx telnet.Context, w telnet.Writer, r telnet.Reader) {
	counter++
	str := `
----------
You must not use Shift_JIS encoding to view this page!!!
Please use UTF-8 encoding.
----------
□□□□ あなたは ` + fmt.Sprintf("%05d", counter) + ` 人目の訪問者です □□□□
----------
 ∧＿＿∧ 
( ´・ω・）   _                   _           _                    _
|ﾉ| ﾉ)ﾉ)|ゝ | |                 (_)         | |                  | |
ﾒ＿ > ❙ ﾚﾉ_ | |_ _ __  _ __ ___  _  ___  ___| |_       _ __   ___| |_
| '_ \ / _ \| __| '_ \| '__/ _ \| |/ _ \/ __| __|     | '_ \ / _ \ __|
| |_) | (_) | |_| |_) | | | (_) | |  __/ (__| |_   _  | | | |  __/ |_
| .__/ \___/ \__| .__/|_|  \___/| |\___|\___|\__| (_) |_| |_|\___|\__|
| |             | |            _/ |
|_|             |_|           |__/

# Profile
Name: potpro (ぽとぷろ)
Blog: blog.potproject.net (https://blog.potproject.net)
Github: @potproject (https://github.com/potproject)
X(Twitter): @potpro (https://twitter.com/potpro)
Mastodon: https://mastodon.potproject.net/@potpro
Bluesky: https://bsky.app/profile/potp.ro
----------

`
	for _, c := range str {
		w.Write([]byte(string(c)))
		time.Sleep(50 * time.Millisecond)
	}

}
