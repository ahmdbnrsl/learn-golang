package database

var connection string

func init() {
    connection = "MySQL"
}

func GetDb() string {
    return connection
}