package enroll

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ EnrollTableModel = (*customEnrollTableModel)(nil)

type (
	// EnrollTableModel is an interface to be customized, add more methods here,
	// and implement the added methods in customEnrollTableModel.
	EnrollTableModel interface {
		enrollTableModel
		withSession(session sqlx.Session) EnrollTableModel
	}

	customEnrollTableModel struct {
		*defaultEnrollTableModel
	}
)

// NewEnrollTableModel returns a model for the database table.
func NewEnrollTableModel(conn sqlx.SqlConn) EnrollTableModel {
	return &customEnrollTableModel{
		defaultEnrollTableModel: newEnrollTableModel(conn),
	}
}

func (m *customEnrollTableModel) withSession(session sqlx.Session) EnrollTableModel {
	return NewEnrollTableModel(sqlx.NewSqlConnFromSession(session))
}
