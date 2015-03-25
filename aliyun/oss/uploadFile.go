package oss

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type UploadFile struct {
	Host            string //BucketName.oss-cn-hangzhou.aliyuncs.com
	BucketName      string
	AccessKeyId     string
	AccessKeySecret string
}

func (uf UploadFile) Upload(fileName string, uploadName string) (string, error) {
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	MD5 := computeMd5(fileBytes)
	Date := time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT")
	ContentType := http.DetectContentType(fileBytes)
	buf := bytes.NewBuffer(fileBytes)

	client := &http.Client{}
	req, _ := http.NewRequest("PUT", "http://"+uf.Host+uploadName, buf)

	req.Header.Set("Host", uf.Host)
	req.Header.Set("Date", Date)
	req.Header.Set("Content-Length", strconv.Itoa(int(req.ContentLength)))
	req.Header.Set("Content-Md5", MD5)
	req.Header.Set("Content-Type", ContentType)
	//req.Header.Set("X-OSS-Meta-Author", "sorex@163.com")
	addAuthorization(uf, req)

	fmt.Println(req.Header)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode == 200 {
		return string(body), nil
	}
	return "", errors.New("[" + strconv.Itoa(resp.StatusCode) + "]" + string(body))
}
