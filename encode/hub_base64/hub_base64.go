// Package hub_base64 base64 encode / decode
package hub_base64

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
)

func Encode(src []byte) []byte {
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(src)))
	base64.StdEncoding.Encode(dst, src)
	return dst
}

func EncodeString(src string) string {
	return string(Encode([]byte(src)))
}

func Decode(data []byte) ([]byte, error) {
	src := make([]byte, base64.StdEncoding.DecodedLen(len(data)))
	n, err := base64.StdEncoding.Decode(src, data)
	if err != nil {
		err = errors.New("base64.StdEncoding.Decode failed")
	}
	return src[:n], err
}

func DecodeString(data string) (string, error) {
	ret, err := Decode([]byte(data))
	return string(ret), err
}

func EncodeFile(path string) ([]byte, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		err = fmt.Errorf(`failed read filename "%s"`, path)
		return nil, err
	}
	return Encode(content), nil
}

func EncodeFileString(path string) (string, error) {
	content, err := EncodeFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
