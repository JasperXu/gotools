package oss

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"io"
	"net/http"
	"sort"
	"strings"
)

func addAuthorization(uf UploadFile, r *http.Request) {
	Method := r.Method
	ContentMD5 := r.Header.Get("Content-Md5")
	ContentType := r.Header.Get("Content-Type")
	Date := r.Header.Get("Date")
	CanonicalizedOSSHeaders := getCanonicalizedOSSHeaders(r)
	CanonicalizedResource := "/" + uf.BucketName + r.URL.Path
	Authorization := "OSS " + uf.AccessKeyId + ":" + computeSignature(uf.AccessKeyId, uf.AccessKeySecret, Method, ContentMD5, ContentType, Date, CanonicalizedOSSHeaders, CanonicalizedResource)
	r.Header.Add("Authorization", Authorization)
}

func getCanonicalizedOSSHeaders(r *http.Request) string {
	CanonicalizedOSSHeaders := ""

	var keyStrings []string
	Params := make(map[string]string)
	for k, v := range r.Header {
		if strings.HasPrefix(strings.ToLower(k), "x-oss-") {
			Params[strings.ToLower(k)] = join(v, ",")
			keyStrings = append(keyStrings, strings.ToLower(k))
		}
	}

	sort.Strings(keyStrings)
	for i := 0; i < len(keyStrings); i++ {
		CanonicalizedOSSHeaders += keyStrings[i] + ":" + Params[keyStrings[i]] + "\n"
	}

	return CanonicalizedOSSHeaders
}

func join(stringArray []string, separator string) string {
	result := ""
	for i, v := range stringArray {
		if i == 0 {
			result += v
		} else {
			result = separator + result
		}
	}
	return result
}

func computeSignature(AccessKeyId string, AccessKeySecret string, Method string, ContentMD5 string, ContentType string, Date string, CanonicalizedOSSHeaders string, CanonicalizedResource string) string {
	Signature := Method + "\n" + ContentMD5 + "\n" + ContentType + "\n" + Date + "\n" + CanonicalizedOSSHeaders + CanonicalizedResource
	mac := hmac.New(sha1.New, []byte(AccessKeySecret))
	io.WriteString(mac, Signature)
	Signature = base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return Signature
}

func computeMd5(data []byte) string {
	h := md5.New()
	h.Write(data)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
