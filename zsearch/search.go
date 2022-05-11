package zsearch

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

	"zsearch/options"
)

type search struct {
	Value    string // 搜索值
	FileName bool   // 搜索文件名
	Content  bool   // 搜索文件内容
	Path     string // 搜索路径
	ALL      bool   // 搜索全部，包括文件名和文件内容
	Wait     sync.WaitGroup
	TD       TailData
}

type fileMap struct {
	Name  string
	Path  string
	Size  int64
	IsDir bool
}

var fm []fileMap

type TailData struct {
	StartTime string
	EndTime   string
	FileNum   int32
	DirNum    int32
	Second    int64
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
	return search
}

func (s *search) Search() error {
	defer func() {
		fmt.Println("\n" + fmt.Sprintf(Tail, s.TD.StartTime, s.TD.EndTime, s.TD.FileNum, s.TD.DirNum, fmt.Sprintf("%dms", s.TD.Second)))
	}()
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
		if !info.IsDir() {
			s.TD.FileNum++
			fm = append(fm, fileMap{
				Path:  path,
				Name:  info.Name(),
				Size:  info.Size(),
				IsDir: info.IsDir(),
			})
		} else {
			s.TD.DirNum++
		}
		return nil
	}); err != nil {
		return Fail(ErrorDirectory)
	}
	// 搜索
	s.run()
	return nil
}

func (s *search) run() {
	cupNum := runtime.NumCPU()
	c := make(chan struct{}, cupNum)
	if s.ALL {
		// 查询全部
		for i := range fm {
			s.Wait.Add(2)
			c <- struct{}{}
			go s.fileName(c, fm[i].Path, &fm[i])
			c <- struct{}{}
			go s.fileContent(c, fm[i].Path)
		}
	} else if s.FileName {
		// 查询文件名
		for i := range fm {
			s.Wait.Add(1)
			c <- struct{}{}
			go s.fileName(c, fm[i].Path, &fm[i])
		}
	} else {
		// 查询文件内容
		for i := range fm {
			s.Wait.Add(1)
			c <- struct{}{}
			go s.fileContent(c, fm[i].Path)
		}
	}
	s.Wait.Wait()
}

// 搜索文件名
func (s *search) fileName(c chan struct{}, path string, info *fileMap) {
	defer func() {
		<-c
		s.Wait.Done()
	}()
	if strings.Contains(info.Name, s.Value) {
		fmt.Println(fmt.Sprintf(ResultsFileName, path))
	}
}

// 搜索文件内容
func (s *search) fileContent(c chan struct{}, p string) {
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
		fmt.Println(fmt.Sprintf(ResultsFileContent+" %v", p, 0, err))
		return
	}
	num := strings.Count(string(read), s.Value)
	if num > 0 {
		fmt.Println(fmt.Sprintf(ResultsFileContent, p, num))
	}
}
