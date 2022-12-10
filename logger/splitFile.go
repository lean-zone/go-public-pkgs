package logger

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

var gBootTime = time.Now()

// SplitFile 文件群
type SplitFile struct {
	folderName   string
	filePrefix   string
	fileExt      string
	splitIndex   int
	splitSize    int
	file         *os.File
	fileMu       sync.RWMutex
	blockSize    int
	onceInitFile sync.Once
}

// NextFileName 获取下一个文件名
func (p *SplitFile) NextFileName() string {
	path := fmt.Sprintf("%s/%s", ".", p.folderName)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, 0755)
		if err != nil {
			path = path + "-"
		} else {
			path = path + "/"
		}
	} else {
		path = path + "/"
	}
	timeStr := gBootTime.Format("2006-01-02 15-04-05")
	indexStr := strconv.Itoa(p.splitIndex)
	p.splitIndex++
	return path + p.filePrefix + "[" + timeStr + "]-" + indexStr + "." + p.fileExt
}

// OpenNextFile 获取下一个文件
func (p *SplitFile) OpenNextFile() (*os.File, error) {
	return os.OpenFile(p.NextFileName(), os.O_WRONLY|os.O_CREATE, 0755)
}

// Write interface 接口
func (p *SplitFile) Write(buf []byte) (int, error) {
	var err error
	p.onceInitFile.Do(func() {
		if p.file == nil {
			p.file, err = p.OpenNextFile()
		}
	})
	if err != nil {
		return 0, err
	}
	//超过分片大小，创建新文件
	if p.blockSize >= p.splitSize*1024 {
		p.fileMu.Lock()
		nextFile, err := p.OpenNextFile()
		if err == nil {
			err := p.file.Close()
			if err != nil {
				return 0, err
			}
			p.file = nextFile
		}
		//无论新文件是否创建成功，都把blockSize设置为0，避免频繁打开新文件
		p.blockSize = 0
		p.fileMu.Unlock()
	}
	p.fileMu.RLock()
	lenOfWritten, err := p.file.Write(buf)
	p.blockSize += lenOfWritten
	p.fileMu.RUnlock()
	return lenOfWritten, err
}

// NewFileCtx 创建新的文件群
func NewFileCtx(folder, filePrefix, fileExt string) *SplitFile {
	return &SplitFile{
		folderName: folder,
		filePrefix: filePrefix,
		fileExt:    fileExt,
		splitIndex: 0,
		splitSize:  100 * 1024,
		file:       nil,
		blockSize:  0,
	}
}
