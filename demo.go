package main

import (
	"fmt"
	"github.com/psoho/fastposter-client-go/fastposter"
)

func main() {
	client := fastposter.Client("26d64614342acb1a40b3")

	params := map[string]interface{}{
		"name": "测试文本",
	}

	poster, err := client.BuildPoster("5c11c65cca7840ed", params, "png")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	//fmt.Println("Error:", poster.B64String())
	err = poster.SaveTo("demo.png")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
