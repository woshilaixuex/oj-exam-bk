package load

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"github/lyr1cs/v0/oj-exam-backend/common/constm"
	"io"
	"math/rand"
	"os"
	"strings"
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
	FileIo           *os.File
	FileMutex        sync.RWMutex
	InitRedisService LoadService
)

// 初始化服务
func InitLoadServer(ctx context.Context, path string) {
	data := openFile(path)
	// 关闭文件
	go func() {
		<-ctx.Done()
		FileIo.Close()
	}()
	// 解析数据
	btlist, err := ParseCSVData(data)
	if err != nil {
		logx.Errorf("data parse slince err: %v", err)
	}
	err = InitRedisService.UploadExamUsersToRedis("csd_test_l", btlist)
	if err != nil {
		logx.Errorf("redis push list err: %v", err)
	}
}

// 生产随机数
func generateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// 生产随机密码
func generateRandomPassword(length int) string {
	return generateRandomString(length)
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

	initialContent := []byte("")
	initialContent = append(initialContent, ProductData(2)...)
	_, err = FileIo.Write(initialContent)
	if err != nil {
		logx.Errorf("file write error: %v", err)
		panic(err)
	}
	return initialContent
}

// 生产数据除了初始化一遍不需要写入csv文件
func ProductData(num int) []byte {
	var buffer bytes.Buffer

	for i := 0; i < num; i++ {
		var randomSuffix string
		for {
			randomSuffix = generateRandomString(5)
			added, err := InitRedisService.AddToUniqueSet(constm.UNIQUESET_ID, randomSuffix)
			if err != nil {
				logx.Error("Error interacting with Redis:", err)
				return nil
			}
			if added == 1 {
				break
			}
		}

		account := "csd" + randomSuffix
		password := generateRandomPassword(8)
		email := fmt.Sprintf("%s@exam.com", account)
		name := account
		buffer.WriteString(fmt.Sprintf("%s,%s,%s,%s\n", account, password, email, name))
	}

	return buffer.Bytes()
}

func ParseCSVData(data []byte) ([]constm.ExamUser, error) {
	reader := csv.NewReader(strings.NewReader(string(data)))

	var users []constm.ExamUser
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			logx.Errorf("failed to read CSV record: %v", err)
			return nil, err
		}
		// CSV 的列顺序为：account, password, email, name
		if len(record) >= 4 {
			user := constm.ExamUser{
				Account:  record[0],
				Password: record[1],
				Email:    record[2],
				Name:     record[3],
			}
			users = append(users, user)
		} else {
			logx.Statf("skipping invalid CSV record: %v", record)
			continue
		}
	}
	return users, nil
}
