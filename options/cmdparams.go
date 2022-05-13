package options

type CmdOptions struct {
	Value   string `short:"d" long:"data" description:"Value to be searched [required] (待搜索词，必填项)"`
	Path    string `short:"p" long:"path" description:"Specifies the search path, which defaults to the current path (指定搜索目录，默认搜索当前目录)"`
	File    bool   `short:"f" long:"file" description:"Search file name (按文件名匹配搜索)"`
	Content bool   `short:"c" long:"content" description:"Search in file content (按文件内容匹配搜索)"`
	NoAa    bool   `short:"n" long:"noaa" description:"Case insensitive (不区分大小写匹配搜索，默认区分)"`
	Version bool   `short:"v" long:"version" description:"Print version info (版本)"`
	About   bool   `short:"a" long:"about" description:"View author information (关于)"`
	Ext     bool   `short:"e" long:"ext" description:"Supported file formats for reading"`
}
