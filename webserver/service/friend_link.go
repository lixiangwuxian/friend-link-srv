package service

import (
	"math/rand"

	"lxtend.com/friend_link/DB"
	db_model "lxtend.com/friend_link/webserver/model"
)

func GetFriends() ([]db_model.FriendItem, error) {

	db := DB.GetDB()
	shuffleSlice := func(slice []db_model.FriendItem) {
		for i := len(slice) - 1; i > 0; i-- {
			j := rand.Intn(i + 1)
			slice[i], slice[j] = slice[j], slice[i]
		}
	}
	query := "SELECT name, url, description, avatar FROM Friends"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	var frienditems []db_model.FriendItem
	var name, url, description, avatar string
	for rows.Next() {
		err = rows.Scan(&name, &url, &description, &avatar)
		if err != nil {
			return nil, err
		}
		frienditems = append(frienditems, db_model.FriendItem{Name: name, Url: url, Description: description, Avatar: avatar})
	}
	shuffleSlice(frienditems)
	return frienditems, nil
}

func DeleteFriend(name string) error {
	db := DB.GetDB()
	query := "DELETE FROM Friends WHERE name = ?"
	if _, err := db.Exec(query, name); err != nil {
		return err
	}
	return nil
}
