package database

//"fmt"
//md "github.com/SG143s/DGRPL/api/authentication/model"

func Getidup(username string, password string) string {
	var info string
	row, _ := db.Query("CALL getidup(?, ?)", username, password)
	for row.Next() {
		row.Scan(&info)
	}
	defer row.Close()
	return info
}
