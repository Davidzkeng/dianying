package db

import (
	"database/sql"
	"dianying/src/user-srv/entity"
	"time"
)

func FindUserByEmail(email string) (*entity.User,error) {
	user := entity.User{}
	err := db.Get(&user,"select * from user where email = ?",email)
	if err == sql.ErrNoRows {
		return nil,nil
	}
	return &user,err
}

func InsertUser(userName string,password string,email string) error {
	today := time.Now().Format("2006-01-02")
	_,err := db.Exec("insert into user (user_name,password,created_at,email) values(?,?,?,?)",userName,password,today,email)
	return err
}

func FindUserByEmailPassword(email,password string) (*entity.User,error) {
	user := entity.User{}
	err := db.Get("select * from user where email = ? and password = ?",email,password)
	if err == sql.ErrNoRows{
		return nil,nil
	}
	return &user,err
}

func UpdateUserEmailProfile(userId int64,email string) error {
	_,err := db.Exec("update user set email = ? where user_id = ?",email,userId)
	if err == sql.ErrNoRows{
		return nil
	}
	return err
}

func UpdateUserNameProfile(userId int64,userName string) error {
	_,err := db.Exec("update user set user_name = ? where user_id = ?",userName,userId)
	if err == sql.ErrNoRows{
		return nil
	}
	return err
}

func UpdateUserPhoneProfile(userId int64,userPhone string) error {
	_,err := db.Exec("update user set phone = ? where user_id = ?",userPhone,userId)
	if err == sql.ErrNoRows{
		return nil
	}
	return err
}
