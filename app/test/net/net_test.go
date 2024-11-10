package test

import (
	"context"
	"database/sql"
	"flag"
	"github/lyr1cs/v0/oj-exam-backend/app/model/enroll"
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
var configFile = flag.String("f", "../../api/etc/exam-api.yaml", "the config file")

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
	result := db.WithContext(context.Background()).Where("student_id = ?", "2200720230").First(&enroll)
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
