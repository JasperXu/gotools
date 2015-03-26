package main

import (
	"fmt"

	"github.com/jasperxu/gotools/aliyun/oss"
)

func main() {
	o := oss.OSS{Host: "BucketName.oss-cn-hangzhou.aliyuncs.com", BucketName: "BucketName", AccessKeyId: "AccessKeyId", AccessKeySecret: "AccessKeySecret"}
	b, err := o.Upload("tmp/uuid/aaa.jpg", "/uuid/aaa.jpg")
	if err != nil {
		fmt.Println("[error]", err)
	}
	fmt.Println("[ok]", b)

	b, err = o.Delete("/bc961ebef1b74f828cd56d17ed430a7d/test2.jpg")
	if err != nil {
		fmt.Println("[error]", err)
	}
	fmt.Println("[ok]", b)

}
