package models

import (
	"github.com/azurramas/ums_app/services"
)

type User struct {
	FirstName string `db:"firstName" json:"firstName"`
	LastName  string `db:"lastName" json:"lastName"`
	Username  string `db:"username" json:"username"`
	Password  string `db:"password" json:"password"`
	Email     string `db:"email" json:"email"`
	Status    string `db:"status" json:"status"`
	ID        int64  `db:"id" json:"id"`
}

var userSchema = `CREATE TABLE users (
firstName varchar(50) NOT NULL,
lastName varchar(50) NOT NULL,
username varchar(50) NOT NULL,
password varchar(255) NOT NULL,
email varchar(50) NOT NULL,
status varchar(50) NOT NULL,
id int NOT NULL AUTO_INCREMENT,
PRIMARY KEY (id)
);`

func ListUsers() ([]User, error) {
	var users []User

	db := services.AppInstance.GetDBConn()

	err := db.Select(&users, "SELECT * FROM users")

	return users, err
}

func GetUser(id int) (User, error) {
	var user User

	db := services.AppInstance.GetDBConn()

	err := db.Get(&user, "SELECT * FROM users WHERE id = ?", id)

	return user, err
}

func (u *User) Create() error {
	db := services.AppInstance.GetDBConn()

	user, err := db.NamedExec(`INSERT INTO users (firstName, lastName, username, password, email, status) 
					VALUES (:firstName, :lastName, :username, :password, :email, :status)`, u)
	if err != nil {
		return err
	}

	u.ID, err = user.LastInsertId()
	if err != nil {
		return err
	}

	return err
}

func (u *User) Update() error {
	db := services.AppInstance.GetDBConn()

	_, err := db.NamedExec(`UPDATE users SET firstName=:firstName, lastName=:lastName, username=:username, 
							email=:email, status=:status WHERE id=:id`, u)
	if err != nil {
		return err
	}

	return err
}

func (u User) Delete() error {
	db := services.AppInstance.GetDBConn()

	_, err := db.Exec("DELETE FROM users WHERE id = ?", u.ID)

	return err
}

func (u User) ExistsByUsernameOrEmail() (bool, error) {
	var count int

	db := services.AppInstance.GetDBConn()

	err := db.Get(&count, "SELECT COUNT(*) FROM users WHERE username = ?", u.Username)

	if count > 0 {
		return count > 0, err
	}

	err = db.Get(&count, "SELECT COUNT(*)  FROM users WHERE email = ?", u.Email)
	return count > 0, err

}

func (u User) ExistsByUsernameOrEmailWhereID() (bool, error) {
	var count int

	db := services.AppInstance.GetDBConn()

	err := db.Get(&count, "SELECT COUNT(*) FROM users WHERE username = ? AND id != ?", u.Username, u.ID)

	if count > 0 {
		return count > 0, err
	}

	err = db.Get(&count, "SELECT COUNT(*)  FROM users WHERE email = ? AND id != ?", u.Email, u.ID)
	return count > 0, err

}

func UserExistsByID(id int64) (bool, error) {
	var count int

	db := services.AppInstance.GetDBConn()

	err := db.Get(&count, "SELECT COUNT(*) FROM users WHERE id = ?", id)
	return count > 0, err
}
