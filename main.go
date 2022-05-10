package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"zsearch/options"
	"zsearch/zsearch"

	"github.com/jessevdk/go-flags"
)

func main() {
	opts := new(options.CmdOptions)
	parser := flags.NewParser(opts, flags.Default)
	parser.Name = zsearch.Name
	parser.Usage = zsearch.Usage
	parser.LongDescription = "author:zzt"
	_, err := parser.Parse()
	if err != nil {
		return
	}
	// 判断是否存在参数
	if len(os.Args[1:]) <= 0 {
		// 返回帮助信息
		parser.WriteHelp(os.Stdout)
		return
	}
	if !strings.Contains(os.Args[1], "-") {
		// 返回帮助信息
		parser.WriteHelp(os.Stdout)
		return
	}
	// 是否显示版本
	if opts.Version {
		fmt.Println(zsearch.Version)
		return
	}
	if opts.About {
		fmt.Println(zsearch.About)
		return
	}
	if opts.Ext {
		r := bytes.Buffer{}
		for ext := range zsearch.ExtMap {
			_, _ = r.WriteString(ext + zsearch.ExtSplit)
		}
		fmt.Println(r.String())
		return
	}
	search := zsearch.NewSearch(opts)
	if err := search.Search(); err != nil {
		fmt.Println(err)
	}
}
