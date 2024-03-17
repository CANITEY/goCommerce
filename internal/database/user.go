package database

import (
	"ecommerce/api/models"

	"github.com/google/uuid"
)

func (d *DB) CheckUser(email, password string) (bool) {
	data := d.conn.QueryRow("SELECT uuid FROM users where email=? AND password=?", email, password)
	uuid := ""
	if err := data.Scan(&uuid); err != nil {
		return false
	}
	if uuid == "" {
		return false
	}
	return true
}

func (d *DB) AddUser(user models.User) error {
	// getting the last rowid
	id := new(int)
	row := d.conn.QueryRow("SELECT MAX(rowid) FROM users")
	if err := row.Scan(&id); err != nil {
		return err
	}
	*id = *id + 1

	// generating UUID
	uid, err := uuid.NewDCESecurity(uuid.Person, uint32(*id))
	if err != nil {
		return err
	}

	_, err = d.conn.Exec("INSERT INTO users(email, password, username, address, phone, uuid) VALUES(?,?,?,?,?,?)", user.Email, user.Password, user.Username, user.Address, user.Phone, uid.String())
	if err != nil {
		return err
	}
	return nil
}

func (d *DB) GetUser(email string) (*models.User, error) {
	user := new(models.User)
	row := d.conn.QueryRow("SELECT ROWID, * FROM users where email=?", email)
	n := ""
	if err := row.Scan(&user.ID, &user.Email, &n, &user.Username, &user.Address, &user.Phone, &user.UUID); err != nil {
		return nil, err
	}
	return user, nil
}
