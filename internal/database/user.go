package database

import "ecommerce/api/models"

func (d *DB) CheckUser(email, password string) (error) {
	_, err := d.conn.Exec("SELECT email, password FROM users where email=? AND password=?", email, password)
	if err != nil {
		return err
	}

	return nil
}

func (d *DB) AddUser(user models.User) (error) {
	_, err := d.conn.Exec("INSERT INTO users(username, password, email, address) VALUES(?,?,?,?)", user.Username, user.Password, user.Email, user.Address)
	if err != nil {
		return err
	}
	return nil
}
