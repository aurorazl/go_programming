package main

/**
txt转markdown需求：
1。顶格内容 --》 一级标题
2。\t的数量表示标题的级数，数量越多，标题级数越大。
	允许标题\t的数量有限，超过后视为内容
	标题判定条件：
		以数字开头且不以句号结尾（否则视为要展示列表形式的内容）
		or 包含冒号，且以冒号结尾
3。为标题和内容去掉制表符\t
	内容对齐第一行，后续行更多的\t不再去掉
4。内容前后添加多行引号

*/

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"unicode"
)

func listDirs(dirPath string) []fs.FileInfo {
	fileInfos, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	return fileInfos
}

func listDirAndGetAllFilePath(dirPath string) (ret []string) {
	fileInfos := listDirs(dirPath)
	for _, file := range fileInfos {
		if !file.IsDir() {
			ret = append(ret, path.Join(dirPath, file.Name()))
		}
	}
	return
}

func isEmptyLine(line []byte) bool {
	if len(strings.TrimSpace(string(line))) == 0 {
		return true
	}
	return false
}

var targetList = [][]byte{{'#', '#', 32}, {'#', '#', '#', 32}, {'#', '#', '#', '#', 32}, {'#', '#', '#', '#', '#', 32}, {'#', '#', '#', '#', '#', '#', 32}}
var lastTitleIndentation = 0
var isPrevLineTitle = false

// 满足title标题的条件：
// 数字开头、包含冒号且以冒号结尾
// 新规则：上一行是title，那么这行缩进可以更多。上一行不是，那么这行要求缩进不大于。
func isTitle(line []byte, indentation int) bool {
	str := strings.TrimSpace(string(line))
	ret := (isPrevLineTitle || lastTitleIndentation >= indentation) &&
		((unicode.IsNumber(rune(str[0])) && (!strings.HasSuffix(str, ".") &&
			!strings.HasSuffix(str, "。"))) ||
			(strings.HasSuffix(str, "?") || strings.HasSuffix(str, "？")) ||
			strings.HasSuffix(str, "：") || strings.HasSuffix(str, ":") ||
			strings.HasSuffix(str, "？") || strings.HasSuffix(str, "?"))
	if ret {
		lastTitleIndentation = indentation
	}
	isPrevLineTitle = ret
	return ret
}

func isImageLine(line []byte) bool {
	str := strings.TrimSpace(string(line))
	return strings.HasPrefix(str, "![image-")
}

var addQuotationHead = false
var addQuotationTail = false
var trimIndex = 0

func convertTabToChar(line []byte) []byte {
	i := 0
	for i < len(line) {
		if line[i] == '\t' {
			i++
		} else {
			break
		}
	}
	if !isTitle(line, i) {
		//非标题
		if addQuotationTail {
			// 前面已经添加了闭引号，现在是新环。
			// 前面一行是标题，当前行是内容
			trimIndex = i
			line = line[trimIndex:]
			if !isImageLine(line) {
				line = append([]byte{'`', '`', '`', '\n'}, line...)
				addQuotationHead = true
				addQuotationTail = false
			}
		} else {
			//去掉第一行内容的\t数量的\t
			line = line[min(trimIndex, i):]
		}

	} else {
		// 每行去掉全部的\t
		line = line[i:]
		//支持有限个#
		if i < len(targetList) {
			tmp := make([]byte, len(line)+len(targetList[i]))
			copy(tmp, targetList[i])
			copy(tmp[len(targetList[i]):], line)
			line = tmp
		}
		if addQuotationHead {
			// 已经添加了开引号
			// 前面一行是内容，当前行是标题，需要闭环
			line = append([]byte{'`', '`', '`', '\n'}, line...)
			addQuotationHead = false
		}
		trimIndex = 0
		addQuotationTail = true
	}
	return line
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func copyToBuffer(dst, src []byte, start int) []byte {
	if start+len(src) > len(dst)-1 {
		copy(dst[start:], src[:len(dst)-start])
		dst = append(dst, src[len(dst)-start:]...)
	} else {
		copy(dst[start:], src)
	}
	return dst
}

func initVariable() {
	lastTitleIndentation = 0
	isPrevLineTitle = false
	addQuotationHead = false
	addQuotationTail = false
	trimIndex = 0
}
func convertTxtToMarkdown(txtPath, mdPath string) {
	initVariable()
	txtFile, err := os.Open(txtPath)
	if err != nil {
		log.Fatalf("open txt file error : %v", err)
	}
	mdFile, err := os.OpenFile(mdPath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer txtFile.Close()
	defer mdFile.Close()
	reader := bufio.NewReader(txtFile)
	writer := bufio.NewWriter(mdFile)
	tmp := make([]byte, 1024)
	index := 0
	for {
		line, prefix, err := reader.ReadLine()
		if err == io.EOF {
			if addQuotationHead {
				tmp = append(tmp[:index], []byte{'`', '`', '`'}...)
				index += 3
			}
			writer.Write(tmp[:index])
			break
		}
		if isEmptyLine(line) {
			line = []byte{'\n'}
			isPrevLineTitle = false
		} else {
			line = convertTabToChar(line)
			if !prefix {
				line = append(line, '\n')
			}
		}
		tmp = copyToBuffer(tmp, line, index)
		index += len(line)
		// cache已满，写入文件
		if index > 1023 {
			writer.Write(tmp[:index])
			tmp = make([]byte, 1024)
			index = 0
		}
	}
	writer.Flush()
}

func convertDir(dirPath string) {
	allFilePath := listDirAndGetAllFilePath(dirPath)
	for _, filePath := range allFilePath {
		if strings.HasSuffix(filePath, ".txt") {
			splits := strings.Split(filePath, "/")
			fileName := path.Join(dirPath, "markdown",
				strings.Split(splits[len(splits)-1], ".")[0]+".md")
			fmt.Printf("start convert file: %s to %s\n", filePath, fileName)
			convertTxtToMarkdown(filePath, fileName)
		}
	}
}

func main() {
	//禁用，因为部分md文件已经修改。
	//dirPath := "/Users/kelvin/Documents/workstate/Python/notebook/"
	//convertDir(dirPath)

	//只能用该方法
	name := "架构设计"
	convertTxtToMarkdown(fmt.Sprintf("/Users/kelvin/Documents/workstate/Python/notebook/%s.txt", name),
		fmt.Sprintf("/Users/kelvin/Documents/workstate/Python/notebook/markdown/%s.md", name))
}
