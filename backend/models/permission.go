package models

import "github.com/azurramas/ums_app/services"

type Permission struct {
	Code        string `db:"code" json:"code"`
	Description string `db:"description" json:"desc"`
	ID          int64  `db:"id" json:"id"`
}

type UserPermissions struct {
	UserID int `db:"userID" json:"userID"`
	PermID int `db:"permID" json:"permID"`
}

var permsSchema = `CREATE TABLE permissions (
	description varchar(255) NOT NULL,
	code varchar(50) NOT NULL,
	id int NOT NULL AUTO_INCREMENT,
	PRIMARY KEY (id)
	);`

var userPermsSchema = `CREATE TABLE user_permissions (
		userID int NOT NULL,
		permID int NOT NULL
		);`

func ListPermissions() ([]Permission, error) {
	var permissions []Permission

	db := services.AppInstance.GetDBConn()

	err := db.Select(&permissions, "SELECT * FROM permissions")

	return permissions, err
}

func GetUserPermissions(userID int) ([]UserPermissions, error) {
	var permissions []UserPermissions

	db := services.AppInstance.GetDBConn()

	err := db.Select(&permissions, "SELECT * FROM user_permissions WHERE userID = ?", userID)

	return permissions, err
}

func AddPermToUser(permID int, userID int) error {
	db := services.AppInstance.GetDBConn()

	_, err := db.Exec(`INSERT INTO user_permissions (permID, userID) 
					VALUES (?,?)`, permID, userID)
	if err != nil {
		return err
	}

	return err
}

func DeletePermissionByID(permID int, userID int) error {
	db := services.AppInstance.GetDBConn()

	_, err := db.Exec("DELETE FROM user_permissions WHERE permID = ? AND userID = ?", permID, userID)

	return err
}
