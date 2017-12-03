package DAO

import (
	"errors"
	"fmt"
	"strconv"
	"store/model"
	"store/util"
)

func CreateUser(u *model.User) (int, error) {
	if u.Name == "" {
		return 0, errors.New("name is obligatory")
	}
	if u.Login == "" {
		return 0, errors.New("login is obligatory")
	}
	if u.Pass == "" {
		return 0, errors.New("pass is obligatory")
	}

	res, err := db.Exec("INSERT INTO user (name, login, pass) VALUES (?, ?, ?)",
		u.Name, u.Login, u.Pass)
	util.CheckErr(err)

	id, err := res.LastInsertId()

	return int(id), err
}

func ReadUser(u *model.User, start int, quantity int) (users []model.User, total int, err error) {
	conditions := ""

	if u.Iduser != 0 {
		conditions += "AND iduser='" + fmt.Sprint(u.Iduser) + "' "
	}
	if u.Name != "" {
		conditions += "AND name LIKE '%" + u.Name + "%' "
	}
	if u.Login != "" {
		conditions += "AND login = '" + u.Login + "' "
	}
	if u.Pass != "" {
		conditions += "AND pass = '" + u.Pass + "' "
	}

	query := "SELECT COUNT(*) FROM user WHERE TRUE " + conditions

	row := db.QueryRow(query)
	err = row.Scan(&total)
	util.CheckErr(err)

	query = "SELECT * FROM user WHERE TRUE " + conditions +
		"LIMIT " + strconv.Itoa(start) + "," + strconv.Itoa(quantity)

	rows, err := db.Query(query)
	defer rows.Close()
	util.CheckErr(err)

	for rows.Next() {
		u = &model.User{}

		err := rows.Scan(&u.Iduser, &u.Name, &u.Login, &u.Pass)
		util.CheckErr(err)

		users = append(users, *u)
	}
	err = rows.Err()
	util.CheckErr(err)
	return
}

func UpdateUser(u *model.User) (err error) {
	initialQuery := "UPDATE user SET "
	query := initialQuery

	if u.Iduser == 0 {
		return errors.New("user id is obligatory")
	}
	if u.Name != "" {
		query += "name = '" + u.Name + "' "
	}
	if u.Login != "" {
		if query != initialQuery {
			query += ","
		}
		query += "login = '" + u.Login + "' "
	}
	if u.Pass != "" {
		if query != initialQuery {
			query += ","
		}
		query += "pass = '" + u.Pass + "' "
	}

	if query == initialQuery {
		return errors.New("nothing to update")
	}

	query += "WHERE iduser = '" + strconv.Itoa(u.Iduser) + "'"

	_, err = db.Exec(query)
	util.CheckErr(err)

	return
}

func DeleteUser(u *model.User) error {
	if u.Iduser == 0 {
		return errors.New("user id is obligatory")
	}

	_, err := db.Exec("DELETE FROM user WHERE iduser = ?", u.Iduser)
	util.CheckErr(err)

	return err
}
