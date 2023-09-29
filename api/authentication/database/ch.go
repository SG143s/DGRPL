package database

import "fmt"

func ChUn(username string) bool {
	var res int
	row, err := db.Query("CALL chuname (?)", username)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer row.Close()
	for row.Next() {
		err = row.Scan(&res)
		if err != nil {
			fmt.Println(err)
			return false
		}
	}
	return res != 0
}

func ChLog(username string, password string) bool {
	var res int
	row, err := db.Query("CALL chlogin (?, ?)", username, password)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer row.Close()
	for row.Next() {
		err = row.Scan(&res)
		if err != nil {
			fmt.Println(err)
			return false
		}
	}
	fmt.Println(password)
	fmt.Println(res)
	return res != 0
}

func ChId(id string) bool {
	var res int
	row, err := db.Query("CALL chid (?)", id)
	if err != nil {
		fmt.Println(err)
		return true
	}
	defer row.Close()
	for row.Next() {
		err = row.Scan(&res)
		if err != nil {
			fmt.Println(err)
			return true
		}
	}
	return res != 1
}

func ChEmail(mail string) bool {
	var res int
	row, err := db.Query("CALL chemail (?)", mail)
	if err != nil {
		fmt.Println(err)
		return true
	}
	defer row.Close()
	for row.Next() {
		err = row.Scan(&res)
		if err != nil {
			fmt.Println(err)
			return true
		}
	}
	return res != 1
}
