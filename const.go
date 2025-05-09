package main

import "errors"

type Err error

var (
	ErrFileNotFound Err = errors.New("文件不存在")
)
