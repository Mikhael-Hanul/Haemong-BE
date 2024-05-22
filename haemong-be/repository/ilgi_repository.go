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
	IlgiId  string
	Title   string
	Content string
	Date    string
	Weather string
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

func (r *IlgiRepositroy) DeleteIlgi(ilgiId string) error {
	_, err := r.db.Exec("DELETE FROM tbl_ilgi WHERE ilgiId = ?", ilgiId)
	if err != nil {
		return fmt.Errorf("일기 삭제에 실패함 : " + err.Error())
	}
	return nil
}

func (r *IlgiRepositroy) SearchIlgi(keyword string) ([]IlgiEntity, error) {
	rows, err := r.db.Query("SELECT ilgiId, title, content FROM tbl_ilgi WHERE content LIKE ?", "%"+keyword+"%")
	if err != nil {
		return nil, fmt.Errorf("일기 검색에 실패함 : " + err.Error())
	}
	defer rows.Close()

	var ilgis []IlgiEntity
	for rows.Next() {
		var ilgi IlgiEntity
		err := rows.Scan(&ilgi.IlgiId, &ilgi.Title, &ilgi.Content, &ilgi.Weather, &ilgi.Date)
		if err != nil {
			return nil, fmt.Errorf("일기 검색 결과 읽기에 실패함 : " + err.Error())
		}
		ilgis = append(ilgis, ilgi)
	}
	return ilgis, nil
}
