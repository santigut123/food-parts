package parser

import (
)

type FileParser struct {
	filename string
	fileType string
}
func NewFileParser(fileName string, fileType string) *FileParser{
	nfp:=FileParser{
		filename: fileName,
		fileType: fileType,
	}
	return &nfp
}
func(fp *FileParser) Get
