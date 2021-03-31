package models

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type DataBaseList struct {
	ID           int64  `json:"id"`
	DatabaseName string `json:"db_name"`
	DBUser       string `json:"db_user"`
	DBPass       string `json:"db_password"`
	DBType       string `json:"db_type"`
	Port         int    `json:"port"`
	Host         string `json:"host"`
	Status       bool   `json:"status"`
}

type DBDetail struct {
	Database []string
	DBVersion string

}
