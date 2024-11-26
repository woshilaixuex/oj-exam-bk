package exam_account

type Account struct {
	ID        int64
	Account   string
	Password  string
	StudentId string `gorm:"column:student_id"`
	Score     int64  `gorm:"column:score"`
}

func (Account) TableName() string {
	return "exam_account"
}
