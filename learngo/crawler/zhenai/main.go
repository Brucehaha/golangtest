package main

//noinspection GoUnresolvedReference
import (
	"brucego.com/learngo/crawler/engine"
	"brucego.com/learngo/crawler/zhenai/parser"
)

func main() {
	request := engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	}
	engine.Run(request)
}
