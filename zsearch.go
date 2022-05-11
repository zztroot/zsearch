package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"zsearch/options"
	"zsearch/search"

	"github.com/jessevdk/go-flags"
)

func main() {
	opts := new(options.CmdOptions)
	parser := flags.NewParser(opts, flags.Default)
	parser.Name = search.Name
	parser.Usage = search.Usage
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
		fmt.Println(search.Version)
		return
	}
	if opts.About {
		fmt.Println(search.About)
		return
	}
	if opts.Ext {
		r := bytes.Buffer{}
		for ext := range search.ExtMap {
			_, _ = r.WriteString(ext + search.ExtSplit)
		}
		fmt.Println(r.String())
		return
	}
	search := search.NewSearch(opts)
	if err := search.Search(); err != nil {
		fmt.Println(err)
	}
}
