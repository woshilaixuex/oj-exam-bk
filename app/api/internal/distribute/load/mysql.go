package load

import (
	"github/lyr1cs/v0/oj-exam-backend/app/model/exam_account"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGormDB(dns string) *gorm.DB {
	db, _ := gorm.Open(mysql.Open(dns))
	return db
}
func FindExamUser(dns string, userId string) *exam_account.Account {
	db := NewGormDB(dns)
	var account exam_account.Account
	db.Model(exam_account.Account{}).Where("student_id != ?", 1).Find(&account)
	db.Model(exam_account.Account{}).Where("account = ?", account.Account).Update("student_id", userId)
	return &account
}

func FindStudenIdtByAccountAndPassword(dns string, account string) string {
	db := NewGormDB(dns)
	var student exam_account.Account
	db.Model(exam_account.Account{}).Where("account = ?", account).Find(&student)
	return student.StudentId
}

func SaveScore(dns string, studentId string, studentScore int) {
	db := NewGormDB(dns)
	db.Model(exam_account.Account{}).Where("student_id = ?", studentId).Update("score", studentScore)
	return
}
