package main

import (
	"fmt"
	"github.com/psoho/fastposter-client-go/fastposter"
)

func main() {

	// 创建海报客户端
	client := fastposter.ClientWithEndpoint("ApfrIzxCoK1DwNZOEJCwlrnv6QZ0PCdv", "http://127.0.0.1:5000")

	// 设置参数
	params := map[string]interface{}{
		"name": "测试文本撒旦法",
	}

	// 生成海报
	poster, err := client.BuildPoster("de9a1007d3dbffbe", params, "png")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	//fmt.Println("B64:", poster.B64String())
	poster.Save()
}
