package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	d *sql.DB
}

type User struct {
	Uid  int    `json:"uid"`
	Name string `json:"name"`
	Pass string `json:"pass"`
}

func New(dc DBConfig) (DB, error) {
	conn, err := sql.Open("mysql", dc.String())
	if err != nil {
		return DB{}, err
	}
	return DB{conn}, nil
}

func (d DB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return d.d.Exec(query, args...)
}

func rowsToUsers(rs *sql.Rows) ([]*User, error) {
	list := make([]*User, 0, 1)
	for rs.Next() {
		var x = &User{}
		err := rs.Scan(&x.Uid, &x.Name, &x.Pass)
		if err != nil {
			return nil, err
		}
		list = append(list, x)
	}
	return list, nil
}

func (d DB) AddUser(name, pass string) error {
	_, err := d.Exec(InsertUser, name, pass)
	return err
}

func (d DB) DeleteUser(uid int) error {
	_, err := d.Exec(DeleteUser + fmt.Sprintf("where uid=%d", uid))
	return err
}

func (d DB) CheckExists(name string) (bool, error) {
	rs, err := d.d.Query(SelectUser + "where name='" + name + "'")
	if err != nil {
		return false, err
	}
	return rs.Next(), nil
}

func (d DB) GetUser(name string) (*User, error) {
	rs, err := d.d.Query(SelectUser + "where name='" + name + "'")
	if err != nil {
		return nil, err
	}
	ls, err := rowsToUsers(rs)
	if err != nil {
		return nil, err
	}
	return ls[0], nil
}

func (d DB) Close() error {
	return d.d.Close()
}
