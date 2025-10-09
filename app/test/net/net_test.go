package test

import (
	"context"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"github/lyr1cs/v0/oj-exam-backend/app/model/enroll"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"testing"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
 * @Author: lyr1cs
 * @Email: linyugang7295@gmail.com
 * @Description: 测试sql io用时
 * @Date: 2024-11-09 11:13
 */
var configFile = flag.String("f", "../../../etc/exam-api.yaml", "the config file")

type Config struct {
	rest.RestConf
	JwtAuth struct {
		AccessSecret string
	}
	Redis redis.RedisConf
	DB    struct {
		DataSource string
	}
}

func TestSqlNet(t *testing.T) {
	flag.Parse()
	var c Config
	conf.MustLoad(*configFile, &c)
	conn := sqlx.NewMysql(c.DB.DataSource)
	model := enroll.NewEnrollTableModel(conn)
	info, err := model.FindOneByStudentId(context.Background(), "2200720230")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(info)
}

// 超时了？
func TestGromSqlNet(t *testing.T) {
	flag.Parse()
	var c Config
	conf.MustLoad(*configFile, &c)
	dsn := c.DB.DataSource
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	var enroll enroll.EnrollTable
	result := db.WithContext(context.Background()).Where("student_number = ?", "2200720230").First(&enroll)
	if result.Error != nil {
		t.Error(result.Error)
		return
	}

	t.Logf("Found student: %+v", enroll)
}

func TestSQLNet(t *testing.T) {
	flag.Parse()
	var c Config
	conf.MustLoad(*configFile, &c)

	db, err := sql.Open("mysql", c.DB.DataSource)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	var studentID string
	var createdAt sql.NullTime
	query := "SELECT student_number,created_at FROM enroll_table WHERE student_number = ?"
	err = db.QueryRowContext(context.Background(), query, "2200720230").Scan(&studentID, &createdAt)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("Found student ID: %s", studentID)
}

type Account struct {
	ID       int64
	Account  string
	Password string
	Name     string `gorm:"column:stu_name"`
}

func (Account) TableName() string {
	return "exam_account"
}

func TestMysql(t *testing.T) {
	flag.Parse()
	var c Config
	conf.MustLoad(*configFile, &c)
	db, err := gorm.Open(mysql.Open(c.DB.DataSource), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	prefix := "csd20250"
	for i := 1; i <= 100; i++ {
		account := prefix + strconv.Itoa(i)
		password := make([]byte, 8)
		rand.Read(password)
		passwordStr := base64.StdEncoding.EncodeToString(password)
		db.Create(&Account{
			Account:  account,
			Password: passwordStr,
			Name:     "1",
		})
	}
}

// 定义响应结构体
type ContestRankResponse struct {
	Error interface{} `json:"error"`
	Data  struct {
		Results []struct {
			User struct {
				Username string `json:"username"`
			} `json:"user"`
			TotalScore int `json:"total_score"`
		} `json:"results"`
	} `json:"data"`
}

func TestGetScore(t *testing.T) {
	url := "https://oj.dsstudio.tech/api/contest_rank?offset=0&limit=" + strconv.Itoa(200) + "&contest_id=23&force_refresh=0"
	// 创建请求
	req, _ := http.NewRequest("GET", url, nil)

	// 添加 Cookie 到请求头
	req.Header.Add("Cookie", "_ga=GA1.2.1938032973.1728895971; csrftoken=ztoxP736vfYDgSyH8N85Y2LoQziCIoOrwxPeedv44gScsbiDOsEHPTEHA6yv9hX7; sessionid=ugqylizesg3z6xr98hwsw4rr1d8pavev; _ga_59QEB25NR7=GS2.2.s1754807026$o22$g1$t1754811069$j60$l0$h0")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("请求失败: %v", err)
		return
	}
	defer resp.Body.Close() // 确保关闭响应体

	// 读取响应
	body, _ := ioutil.ReadAll(resp.Body)

	var response ContestRankResponse
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Printf("解析JSON失败: %v\n", err)
		return
	}

	// 打印每个学生的分数
	for _, result := range response.Data.Results {
		fmt.Printf("学生: %v, 分数: %d\n\n", result.User.Username, result.TotalScore)
	}

}
