package main

import (
	"fmt"
	"github.com/psoho/fastposter-client-go/fastposter"
)

func main() {
	client := fastposter.Client("07657854eb3858269c76")

	params := map[string]interface{}{
		"name": "测试文本",
	}

	poster, err := client.BuildPoster("4b9423a28e594ac5", params, "png")
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
