package load

import (
	"context"
	"fmt"
	"io"
	"math/rand"
	"os"
	"sync"

	"github.com/zeromicro/go-zero/core/logx"
)

/*
 * @Author: lyr1cs
 * @Email: linyugang7295@gmail.com
 * @Description: 数据生成:加载到csv文件，redis，数据库？ 数据消费:从redis获取用户，mysql中绑定
 * @Date: 2024-11-09 21:49
 */
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var (
	FileIo    *os.File
	FileMutex sync.RWMutex
	ProNum    int = 50
	StartNum  int = 1
)

func generateRandomPassword(length int) string {
	password := make([]byte, length)
	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}
	return string(password)
}

func InitLoadServer(ctx context.Context, path string) ([]ExamUser, error) {
	data := openFile(path)
	// 关闭文件
	go func() {
		<-ctx.Done()
		FileIo.Close()
	}()
	return parseCSVData(data)
}

// 初始化打开文件
func openFile(path string) []byte {
	FileIo, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return creatFile(path)
		}
		logx.Errorf("file start err: %v", err)
		panic(err)
	}

	logx.Stat("start load csv file")
	content, err := io.ReadAll(FileIo)
	if err != nil {
		logx.Errorf("file read error: %v", err)
		panic(err)
	}
	return content
}

// 初始化创建文件
func creatFile(path string) []byte {
	FileIo, err := os.Create(path)
	if err != nil {
		logx.Errorf("file creation error: %v", err)
		panic(err)
	}

	// Optional: Write some initial content to the file
	initialContent := []byte("account,password,email,name\n")
	initialContent = append(initialContent, ProductData(StartNum)...)
	_, err = FileIo.Write(initialContent)
	if err != nil {
		logx.Errorf("file write error: %v", err)
		panic(err)
	}
	return initialContent
}

func ProductData(start int) []byte {
	var result []byte
	for i := start; i < start+ProNum; i++ {
		account := "csd" + fmt.Sprintf("%05d", i)
		password := generateRandomPassword(8)
		email := fmt.Sprintf("%s@exam.com", account)
		name := account
		line := fmt.Sprintf("%s,%s,%s,%s\n", account, password, email, name)
		result = append(result, []byte(line)...)
	}
	return result
}
