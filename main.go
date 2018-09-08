package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"github.com/xiaozefeng/qiniu-client/clipboard"
	"github.com/xiaozefeng/qiniu-client/model"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"
)

var pathVar string

func init() {
	flag.StringVar(&pathVar, "path", "", "upload file path")
}

var config = ".qiniu_client/conf.json"

func main() {
	homePath, err := unixHome()
	if err != nil {
		panic("读取用户目录失败")
	}
	// 读取配置文件
	confPath := fmt.Sprintf("%s/%s", homePath, config)
	data, err := ioutil.ReadFile(confPath)
	if err != nil {
		panic("读取配置文件失败")
	}
	var account model.Account
	err = json.Unmarshal(data, &account)
	if err != nil {
		panic("解析配置失败")
	}

	flag.Parse()
	fmt.Printf("path: %s\n", pathVar)
	if pathVar == "" {
		fileName, err := clipboard.Pop()
		if err != nil {
			panic("execute pbpaste commond error")
		}
		if fileName == "" {
			panic("executed pbpaste get nothing!")
		}
		pathVar = fileName
	}
	// 获取文件拓展名
	fileExt := path.Ext(pathVar)
	localFile := fmt.Sprintf("%s%s", account.RootPath, pathVar)
	key := fmt.Sprintf("qiniu_client%s%s", time.Now().Format("20060102150405"), fileExt)
	log.Printf("key: %s\n", key)

	// 初始化七牛云client
	policy := storage.PutPolicy{
		Scope: account.Bucket,
	}
	mac := qbox.NewMac(account.AccessKey, account.SecretKey)
	token := policy.UploadToken(mac)

	cfg := storage.Config{}
	cfg.Zone = &storage.ZoneHuadong
	cfg.UseHTTPS = false
	cfg.UseCdnDomains = false

	// 构建表单对象
	uploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	b, err := ioutil.ReadFile(localFile)
	if err != nil {
		panic(err)
	}

	// 上传文件
	err = uploader.Put(context.Background(), &ret, token, key, bytes.NewReader(b), int64(len(b)), nil)
	if err != nil {
		panic(err)
	}
	mdUrl := fmt.Sprintf("![](%s/%s)", account.Prefix, ret.Key)
	fmt.Printf("%s\n", mdUrl)

	// 将文件写入剪切板
	tempPath := homePath + "/.qiniu_client/temp.txt"
	err= ioutil.WriteFile(tempPath, []byte(mdUrl), 0777)
	if err != nil {
		panic(err)
	}
	err = clipboard.Push(tempPath)
	if err != nil {
		panic(err)
	}
}

// 获取unix系统的用户目录
func unixHome() (homePath string, err error) {
	// 获取环境变量$HOME
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}

	// If that filed. try to shell
	output, err := exec.Command("sh", "-c", "eval echo -$USER").Output()
	if err != nil {
		return "", err
	}

	result := strings.TrimSpace(string(output))
	if result == "" {
		return "", errors.New("blank output when reading home directory")
	}

	return result, nil
}
