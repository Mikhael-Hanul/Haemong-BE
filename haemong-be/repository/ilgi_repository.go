package repository

import (
	"database/sql"
	"fmt"
)

type IlgiRepositroy struct {
	db *sql.DB
}

func NewIlgiRepository(db *sql.DB) *IlgiRepositroy {
	return &IlgiRepositroy{
		db: db,
	}
}

type IlgiEntity struct {
	ilgiId  string
	title   string
	content string
	date    string
	weather string
}

func (r *IlgiRepositroy) SaveIlgi(ilgiId, title, content, date, weather string) error {
	_, err := r.db.Exec("insert into tbl_ilgi value (?,?,?,?,?)", ilgiId, title, content, date, weather)
	if err != nil {
		return fmt.Errorf("일기 등록에 실패함 : " + err.Error())
	}
	return nil
}

func (r *IlgiRepositroy) ModifyIlgi(ilgiId, title, content, date, weather string) error {
	_, err := r.db.Exec("UPDATE tbl_ilgi SET title = ?, content = ?, date = ?, weather = ? WHERE ilgiId = ?", title, content, date, weather, ilgiId)
	if err != nil {
		return fmt.Errorf("일기 수정에 실패함 : " + err.Error())
	}
	return nil
}
