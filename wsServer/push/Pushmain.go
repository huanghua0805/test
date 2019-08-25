package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/alfred-zhong/wserver"
)

func main() {
	pushURL := "http://127.0.0.1:9090/push"
	contentType := "application/json"



	for {
		pm := wserver.PushMessage{
			UserID:  "jack",
			Event:   "topic1",
			Message: fmt.Sprintf("Hello in "),
		}
		b, _ := json.Marshal(pm)



		fmt.Println(b)
		fmt.Println(string(b))
		http.DefaultClient.Post(pushURL, contentType, bytes.NewReader(b))
		fmt.Println(bytes.NewReader(b))

		time.Sleep(time.Second)
	}
}
