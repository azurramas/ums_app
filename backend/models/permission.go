package models

import "github.com/azurramas/ums_app/services"

type Permission struct {
	Code        string `db:"code" json:"code"`
	Description string `db:"desc" json:"desc"`
	UserID      int64  `db:"userID" json:"userID"`
	ID          int64  `db:"id" json:"id"`
}

var permsSchema = `CREATE TABLE permissions (
	desc varchar(255) NOT NULL,
	code int NOT NULL,
	userID int NOT NULL,
	id int NOT NULL AUTO_INCREMENT,
	PRIMARY KEY (id)
	);`

func (p *Permission) Add() error {
	db := services.AppInstance.GetDBConn()

	perm, err := db.NamedExec(`INSERT INTO permissions (code, desc, userID) 
					VALUES (:code, :desc, :userID)`, p)
	if err != nil {
		return err
	}

	p.ID, err = perm.LastInsertId()
	if err != nil {
		return err
	}

	return err
}

func DeletePermissionByID(id int) error {
	db := services.AppInstance.GetDBConn()

	_, err := db.NamedExec("DELETE FROM permissions WHERE id = ?", id)

	return err
}

func PermissionExists(id int) (bool, error) {
	var count int

	db := services.AppInstance.GetDBConn()

	err := db.Get(&count, "SELECT COUNT(*) FROM permissions WHERE id = ?", id)

	return count > 0, err
}
