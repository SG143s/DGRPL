package database

//"fmt"
import (
	md "github.com/SG143s/DGRPL/api/authentication/model"
)

func Getidup(username string, password string) string {
	var info string
	row, _ := db.Query("CALL getidup(?, ?)", username, password)
	for row.Next() {
		row.Scan(&info)
	}
	defer row.Close()
	return info
}

func Getdat(id string) md.User {
	var data md.User
	row, err := db.Query("CALL getdata(?)", id)
	if err != nil {
		panic(err)
	}
	for row.Next() {
		err = row.Scan(&data.ID, &data.Uname, &data.Email, &data.Password, &data.Name, &data.Pic)
		if err != nil {
			panic(err)
		}
	}
	return data
}
