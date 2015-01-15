// Package gotools implements a simple golang tools package.
package gotools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type language map[string]string

type I18n struct {
	path string
	data map[string]language
}

func NewI18n(path string) *I18n {
	i := I18n{}
	if filepath.IsAbs(path) {
		i.SetPath(path)
	} else {
		absPath, err := filepath.Abs(path)
		if err == nil {
			i.SetPath(absPath)
		}
	}

	i.Load()
	return &i
}

func (l language) Get(key string) string {
	value, ok := l[key]
	if ok == false {
		return ""
	}
	return value
}

// 设置多国语言文件根目录
func (i *I18n) SetPath(path string) {
	i.path = path
}

// 加载多国语言
func (i *I18n) Load() {
	i.data = make(map[string]language)
	filepath.Walk(i.path, i.walkI18nFile)
}

// 获取某个语言包
func (i *I18n) GetLanguage(langName string) language {
	lang, ok := i.data[langName]
	if ok == false {
		return make(map[string]string)
	}
	return lang
}

// 获取某个语言的某个值
func (i *I18n) Get(langName string, key string) string {
	return i.GetLanguage(langName).Get(key)
}

// 加载制定文件，文件名为语言名
func (i *I18n) walkI18nFile(path string, info os.FileInfo, err error) error {
	if err != nil {
		return nil
	}
	if info.Mode().IsDir() == true {
		return nil
	}
	fileName := info.Name()
	langName := fileName[:strings.LastIndex(fileName, ".")]
	fileBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}
	var lang language
	json.Unmarshal(fileBytes, &lang)
	i.data[langName] = lang
	fmt.Println("i18n add language :[", langName, "] OK!")
	return nil
}
