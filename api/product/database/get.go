package database

import (
	"fmt"

	md "github.com/SG143s/DGRPL/api/product/model"
)

func SAll(am int) []md.Sm {
	var data []md.Sm
	var temp md.Sm

	row, err := db.Query("CALL selectall(?)", am)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer row.Close()

	for row.Next() {
		err := row.Scan(&temp.ID, &temp.Own, &temp.Name, &temp.Price, &temp.Image)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		data = append(data, temp)
	}
	return data
}

func Search(val string) []md.Sm {
	var data []md.Sm
	var temp md.Sm

	row, err := db.Query("CALL search(?)", val)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer row.Close()

	for row.Next() {
		err := row.Scan(&temp.ID, &temp.Own, &temp.Name, &temp.Price, &temp.Image)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		data = append(data, temp)
	}
	return data
}

func Pr(id string) md.Pr {
	var temp md.Pr

	row, err := db.Query("CALL getdat(?)", id)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer row.Close()

	for row.Next() {
		err := row.Scan(&temp.ID, &temp.Own, &temp.Name, &temp.Desc, &temp.TimeUsed, &temp.Stock, &temp.Price)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
	}
	row, err = db.Query("CALL getdimgs(?)", id)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	var im string
	for row.Next() {
		err := row.Scan(&im)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		temp.Image = append(temp.Image, im)
	}
	return temp
}

func SearchOwn(val string) []md.Sm {
	var data []md.Sm
	var temp md.Sm

	row, err := db.Query("CALL getown(?)", val)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer row.Close()

	for row.Next() {
		err := row.Scan(&temp.ID, &temp.Own, &temp.Name, &temp.Price, &temp.Image)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		data = append(data, temp)
	}
	return data
}
