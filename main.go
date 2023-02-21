package main

import (
	"fmt"
	"net/http"
	"os"

	gogpt "github.com/sashabaranov/go-gpt3"
)

var client = gogpt.NewClient("自己账号的apikey")

func main() {
	// 第一个接口，返回html文件内容
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		f, _ := os.ReadFile("index.html")
		w.Write([]byte(f))
	})

	// 第二个接口，请求chatgpt的接口
	http.HandleFunc("/api", func(w http.ResponseWriter, req *http.Request) {
		text := req.FormValue("text")
		fmt.Println(text)

		if text == "" {
			w.WriteHeader(400)
			w.Write([]byte("请输入文本"))
			return
		}

		// 发起请求
		ret, err := client.CreateCompletion(req.Context(), gogpt.CompletionRequest{
			Model:            "text-davinci-003",
			MaxTokens:        512,
			TopP:             1,
			FrequencyPenalty: 0,
			PresencePenalty:  0.6,
			Prompt:           text,
		})

		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}

		w.Write([]byte(ret.Choices[0].Text))
	})

	// 启动http服务
	fmt.Println("Listen on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
