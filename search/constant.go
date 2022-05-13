package search

const (
	Name    = "ZSEARCH"
	Version = "v0.0.5"
	Usage   = "[OPTIONS VALUE]||[OPTIONS=VALUE]"
	About   = "Name:zhongtian zhao | Email:1033141032@qq.com"
)

// 默认值
const (
	DefaultPath = "."
)

// 格式化输出
const (
	Tail = `StartTime:%s | EndTime:%s | FileNum:%d | DirNum:%d | Second:%s | FoundNum:%d`
	// 文件路径
	ResultsFileName = `[f] %s`
	// 文件路径 -> 文件内容匹配到多少数量
	ResultsFileContent = `[c] %s -> %d`
)

// 错误信息
const (
	ErrorDirectory = "Please check whether the directory is correct"
)

// 格式化
const (
	FormatTime = "2006-01-02 15:04:05"
	ExtSplit   = " | "
)
