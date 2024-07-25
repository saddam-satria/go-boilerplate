package configs

import (
	"bufio"
	"os"
	"path/filepath"
)

type Fs struct {
	FileName   string
	FolderPath string
}

func (ctx *Fs) Write(message string) error {

	filePath := filepath.Join(ctx.FolderPath, ctx.FileName)

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		os.Create(filePath)
	}

	defer file.Close()

	file.Write([]byte(message))

	return nil
}

func (ctx *Fs) Read() (string, error) {
	filePath := filepath.Join(ctx.FolderPath, ctx.FileName)
	var result string

	file, err := os.OpenFile(filePath, os.O_RDONLY, 0666)

	if err != nil {
		return result, err
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		result += scanner.Text()
	}

	return result, nil
}

func (ctx *Fs) GetFile(flag int) (*os.File, error) {
	filePath := filepath.Join(ctx.FolderPath, ctx.FileName)
	file, err := os.OpenFile(filePath, flag, 0666)

	if err != nil {
		return nil, err
	}

	return file, nil
}
