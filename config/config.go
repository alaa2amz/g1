package config

const(
	 JWTKey = "123abc"
	MySqlDBName =""

	/*AD
	dsn := `alaazak:0100ZakAD@tcp(mysql-alaazak.alwaysdata.net:3306)/alaazak_g1?`
		+`charset=utf8mb4&parseTime=True`
	*/

	//sqlite
	//DB, dberr = gorm.Open(sqlite.Open("db.sqlite?_foreign_keys=on"))

	//local mysql
	DSN = "alaazak:0100ZakAD@/alaazak_g1?charset=utf8mb4&parseTime=True"
)
