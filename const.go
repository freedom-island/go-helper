package go_helper

import "errors"

type Err error

var (
	ErrFileNotFound Err = errors.New("文件不存在")
)
