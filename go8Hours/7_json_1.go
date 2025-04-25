package main

import (
	"encoding/json"
	"fmt"
)

// 结构体标签的应用：json格式的转化，一般很少手动解析结构体标签
type Movie struct {
	Title  string   `json:title`
	Year   int      `json:year`
	Price  int      `json:price`
	Actors []string `json:actors`
}

func main() {
	movie := Movie{"喜剧之王", 2000, 10, []string{"ZhouXingchi", "ZhangBaizhi"}}

	// 编码的过程，将结构体struct转换成json格式
	jsonStr, err := json.Marshal(movie)

	if err != nil {
		fmt.Println("json marshal error", err)
		return
	}

	fmt.Printf("jsonStr = %s\n", jsonStr)

	// 解码的过程，将json转化成结构体struct
	myMovie := Movie{}
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println("json unmarshal is error", err)
		return
	}

	fmt.Printf("%v\n", myMovie)
}
