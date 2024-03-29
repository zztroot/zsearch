package search

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
	"unicode"

	"zsearch/options"
)

type search struct {
	Value    string // 搜索值
	FileName bool   // 搜索文件名
	Content  bool   // 搜索文件内容
	Path     string // 搜索路径
	ALL      bool   // 搜索全部，包括文件名和文件内容
	NoAa     bool   // 搜索是否区分大小写
	Wait     sync.WaitGroup
	TD       TailData
}

type TailData struct {
	StartTime string
	EndTime   string
	FileNum   int32
	DirNum    int32
	Second    int64
	FoundNum  int32
}

func NewSearch(opts *options.CmdOptions) *search {
	search := new(search)
	search.Path = DefaultPath
	if opts.Path != "" {
		search.Path = opts.Path
	}
	if !opts.Content && !opts.File {
		search.ALL = true
	} else {
		search.Content = opts.Content
		search.FileName = opts.File
	}
	search.Value = opts.Value
	if opts.NoAa {
		search.NoAa = opts.NoAa
		search.Value = strings.Map(unicode.ToLower, opts.Value)
	}
	return search
}

func (s *search) Search() error {
	defer func() {
		fmt.Println("\n" + fmt.Sprintf(Tail, s.TD.StartTime, s.TD.EndTime, s.TD.FileNum, s.TD.DirNum, fmt.Sprintf("%dms", s.TD.Second), s.TD.FoundNum))
	}()
	cpuNum := runtime.NumCPU()
	c := make(chan struct{}, cpuNum)
	start := time.Now()
	s.TD.StartTime = start.Format(FormatTime)
	defer func() {
		s.TD.EndTime = time.Now().Format(FormatTime)
		s.TD.Second = time.Since(start).Milliseconds()
	}()
	// 获取目录下所有文件
	if s.Path == "" {
		s.Path, _ = os.Getwd()
	}
	_, err := os.Open(filepath.FromSlash(s.Path))
	if err != nil {
		return err
	}
	if err := filepath.Walk(filepath.FromSlash(s.Path), func(path string, info fs.FileInfo, err error) error {
		if info == nil {
			return nil
		}
		if !info.IsDir() {
			s.TD.FileNum++
			if s.ALL {
				// 查询全部
				s.Wait.Add(1)
				c <- struct{}{}
				go s.fileName(c, path, info.Name())
				s.Wait.Add(1)
				c <- struct{}{}
				go s.fileContent(c, path)
			} else if s.FileName {
				// 查询文件名
				s.Wait.Add(1)
				c <- struct{}{}
				go s.fileName(c, path, info.Name())
			} else {
				// 查询文件内容
				s.Wait.Add(1)
				c <- struct{}{}
				go s.fileContent(c, path)
			}
		} else {
			s.TD.DirNum++
		}
		return nil
	}); err != nil {
		return Fail(ErrorDirectory)
	}
	s.Wait.Wait()
	return nil
}

// 搜索文件名
func (s *search) fileName(c chan struct{}, path, name string) {
	defer func() {
		// 不用处理recover错误
		recover()
	}()
	defer func() {
		<-c
		s.Wait.Done()
	}()
	if s.NoAa {
		name = strings.Map(unicode.ToLower, name)
	}
	if strings.Contains(name, s.Value) {
		s.TD.FoundNum++
		fmt.Println(fmt.Sprintf(ResultsFileName, path))
	}
}

// 搜索文件内容
func (s *search) fileContent(c chan struct{}, p string) {
	defer func() {
		// 不用处理recover错误
		recover()
	}()
	defer func() {
		<-c
		s.Wait.Done()
	}()
	ext := path.Ext(p)
	if _, b := ExtMap[ext]; !b {
		return
	}
	read, err := ioutil.ReadFile(p)
	if err != nil {
		s.TD.FoundNum++
		fmt.Println(fmt.Sprintf(ResultsFileContent+" %v", p, 0, err))
		return
	}
	text := string(read)
	if s.NoAa {
		text = strings.ToLower(text)
	}
	num := strings.Count(text, s.Value)
	if num > 0 {
		s.TD.FoundNum++
		fmt.Println(fmt.Sprintf(ResultsFileContent, p, num))
	}
}
