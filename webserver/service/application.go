package service

import (
	"errors"

	"github.com/google/uuid"
	"lxtend.com/friend_link/DB"
	db_model "lxtend.com/friend_link/webserver/model"
)

func NewApplication(name string, url string, description string, avatar string, email string) (string, string, error) {
	db := DB.GetDB()
	//test if existed
	query_sel := "Select * from FriendApplications where name = ? and url = ? and description = ? and avatar = ?"
	rows, err := db.Query(query_sel, name, url, description, avatar)
	if err != nil {
		return "", "", err
	}
	if rows.Next() {
		return "", "", nil
	}
	query := "INSERT INTO FriendApplications (name, url, description, avatar, changeToken,approveToken) VALUES (?, ?, ?, ?, ?,?)"
	changeToken := uuid.New().String()
	approveToken := uuid.New().String()
	_, err = db.Exec(query, name, url, description, avatar, changeToken, approveToken)
	return changeToken, approveToken, err
}

func EditApplication(changeToken string, name string, url string, description string, avatar string) error {
	db := DB.GetDB()
	query := "UPDATE FriendApplications SET name = ?, url = ?, description = ?, avatar = ? WHERE changeToken = ?"
	_, err := db.Exec(query, name, url, description, avatar, changeToken)
	return err
}

func DeleteApplication(approveToken string) error {
	db := DB.GetDB()
	query := "DELETE FROM FriendApplications WHERE approveToken = ?"
	_, err := db.Exec(query, approveToken)
	return err
}

func GetApplicationByApproveToken(approveToken string) (db_model.ApplicationData, error) {
	db := DB.GetDB()
	query := "SELECT name, email, url, description, avatar FROM FriendApplications WHERE approveToken = ?"
	rows, err := db.Query(query, approveToken)
	if err != nil {
		return db_model.ApplicationData{}, err
	}
	var name, email, url, description, avatar string
	for rows.Next() {
		err = rows.Scan(&name, &email, &url, &description, &avatar)
		if err != nil {
			return db_model.ApplicationData{}, err
		}
	}
	return db_model.ApplicationData{Name: name, Email: email, Url: url, Description: description, Avatar: avatar}, nil
}

func ApproveApplication(approveToken string) error {
	db := DB.GetDB()
	query := "SELECT name, url, description, avatar FROM FriendApplications WHERE approveToken = ?"
	rows, err := db.Query(query, approveToken)
	if rows == nil {
		return nil
	}
	if err != nil {
		return err
	}
	var name, url, description, avatar string

	for rows.Next() {
		err = rows.Scan(&name, &url, &description, &avatar)
		if err != nil {
			return err
		}
	}

	if name == "" || url == "" || description == "" || avatar == "" {
		return errors.New("empty result")
	}

	query = "INSERT INTO Friends (name, url, description, avatar) VALUES (?, ?, ?, ?)"
	_, err = db.Exec(query, name, url, description, avatar)
	if err != nil {
		return err
	}
	query = "DELETE FROM FriendApplications WHERE approveToken = ?"
	_, err = db.Exec(query, approveToken)
	return err
}
