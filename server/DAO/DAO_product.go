package DAO

import (
	"errors"
	"fmt"
	"strconv"
	"store/model"
	"store/util"
)

func CreateProduct(p model.Product) (int, error) {
	if p.Name == "" {
		return 0, errors.New("name is obligatory")
	}
	if p.Price == 0 {
		return 0, errors.New("price is obligatory")
	}

	res, err := db.Exec("INSERT INTO product (name, price) VALUES (?, ?)", p.Name, p.Price)
	util.CheckErr(err)

	id, err := res.LastInsertId()

	return int(id), err
}

func ReadProduct(p model.Product, start int, quantity int) (products []model.Product, err error) {
	query := "SELECT * FROM product WHERE TRUE "

	if p.Idproduct != 0 {
		query += "AND idproduct='" + fmt.Sprint(p.Idproduct) + "' "
	}
	if p.Name != "" {
		query += "AND name LIKE '%" + p.Name + "%' "
	}
	if p.Price != 0 {
		query += "AND price='" + fmt.Sprint(p.Price) + "' "
	}

	query += "LIMIT " + strconv.Itoa(start) + "," + strconv.Itoa(quantity)

	rows, err := db.Query(query)
	defer rows.Close()
	util.CheckErr(err)

	for rows.Next() {
		p = model.Product{}

		err := rows.Scan(&p.Idproduct, &p.Name, &p.Price)
		util.CheckErr(err)

		products = append(products, p)
	}
	err = rows.Err()
	util.CheckErr(err)
	return
}

func UpdateProduct(p model.Product) (err error) {
	initialQuery := "UPDATE product SET "
	query := initialQuery

	if p.Idproduct == 0 {
		return errors.New("product id is obligatory")
	}
	if p.Name != "" {
		query += "name = '" + p.Name + "' "
	}
	if p.Price != 0 {
		if query != initialQuery {
			query += ","
		}
		query += "price = '" + fmt.Sprint(p.Price) + "' "
	}

	if query == initialQuery {
		return errors.New("nothing to update")
	}

	query += "WHERE idproduct = '" + strconv.Itoa(p.Idproduct) + "'"

	_, err = db.Exec(query)
	util.CheckErr(err)

	return
}

func DeleteProduct(p model.Product) error {
	if p.Idproduct == 0 {
		return errors.New("product id is obligatory")
	}

	_, err := db.Exec("DELETE FROM product WHERE idproduct = ?", p.Idproduct)
	util.CheckErr(err)

	return err
}