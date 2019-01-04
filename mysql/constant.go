package mysql

import "errors"

const (
	SelectUser      string = "select * from user "
	DeleteUser      string = "delete from user "
	InsertUser      string = "insert into user(name, pass_md5)values(?, ?)"
	UpdateUser      string = "update user set "
	CreateUserTable string = `
    create table if not exists Users(
	uid      int unsigned  auto_increment,
	name     varchar(20)   not null,
    pass_md5 varchar(50)   not null,
    primary key(uid))
    engine=InnoDB default charset=utf8`
)

var (
	NoSuchUser error = errors.New("No Such User !")
)
