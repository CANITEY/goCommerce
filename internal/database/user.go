package database

import (
	"ecommerce/api/models"

	"github.com/google/uuid"
)

func (d *DB) CheckUser(email, password string) error {
	_, err := d.conn.Exec("SELECT email, password FROM users where email=? AND password=?", email, password)
	if err != nil {
		return err
	}

	return nil
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

	_, err = d.conn.Exec("INSERT INTO users(email, password, username, address, phone, uuid) VALUES(?,?,?,?,?,?)", user.Username, user.Password, user.Email, user.Address, uid.String())
	if err != nil {
		return err
	}
	return nil
}

func (d *DB) GetUser(email string) (*models.User, error) {
	user := new(models.User)
	row := d.conn.QueryRow("SELECT ROWID, * FROM users where email=?", email)
	if err := row.Scan(&user.ID, &user.Email, nil, &user.Username, &user.Address, &user.UUID); err != nil {
		return nil, err
	}
	return user, nil
}
