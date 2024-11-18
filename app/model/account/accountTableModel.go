package account

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ AccountTableModel = (*customAccountTableModel)(nil)

type (
	// AccountTableModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAccountTableModel.
	AccountTableModel interface {
		accountTableModel
		withSession(session sqlx.Session) AccountTableModel
	}

	customAccountTableModel struct {
		*defaultAccountTableModel
	}
)

// NewAccountTableModel returns a model for the database table.
func NewAccountTableModel(conn sqlx.SqlConn) AccountTableModel {
	return &customAccountTableModel{
		defaultAccountTableModel: newAccountTableModel(conn),
	}
}

func (m *customAccountTableModel) withSession(session sqlx.Session) AccountTableModel {
	return NewAccountTableModel(sqlx.NewSqlConnFromSession(session))
}
