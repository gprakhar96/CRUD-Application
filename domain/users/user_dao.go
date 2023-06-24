package users

import (
	"fmt"
	"mypackage/datasources/mysql/userdb"
	"mypackage/utils/errors"
)

const (
	queryInsertUser = "INSERT INTO users.user_table(first_name, last_name, email, date_created) VALUES ($1, $2, $3, $4) RETURNING id;"
	queryGetUser    = "SELECT id, first_name, last_name, email, date_created FROM users.user_table WHERE id = $1;"
	queryUpdateUser = "UPDATE users.user_table SET first_name = $1, last_name = $2, email = $3 WHERE id = $4;"
	queryDeleteUser = "DELETE FROM users.user_table WHERE id = $1;"
)

func (user *User) Save() *errors.RestErrors {
	stmt, queryErr := userdb.DBClient.Prepare(queryInsertUser)
	if queryErr != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Error while trying to prepare statement %s", queryErr.Error()))
	}

	defer stmt.Close()

	var insertedId int64
	insertErr := stmt.QueryRow(user.FirstName, user.LastName, user.Email, user.DateCreated).Scan(&insertedId)
	if insertErr != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Error while trying to insert user [%s]", insertErr))
	}

	user.ID = insertedId
	return nil
}

func (user *User) Get() *errors.RestErrors {
	stmt, queryErr := userdb.DBClient.Prepare(queryGetUser)
	if queryErr != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Error while trying to prepare statement %s", queryErr.Error()))
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	scanErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated)
	if scanErr != nil {
		return errors.New404Error(fmt.Sprintf("User not found  [%s]", scanErr.Error()))
	}
	return nil
}

func (user *User) Update() *errors.RestErrors {
	stmt, queryErr := userdb.DBClient.Prepare(queryUpdateUser)
	if queryErr != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Error while trying to prepare statement %s", queryErr.Error()))
	}
	defer stmt.Close()

	_, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Unable to update user %s", err.Error()))
	}

	return nil
}

func (user *User) Delete() *errors.RestErrors {
	stmt, queryErr := userdb.DBClient.Prepare(queryDeleteUser)
	if queryErr != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Error while trying to prepare statement %s", queryErr.Error()))
	}
	defer stmt.Close()

	_, err := stmt.Exec(user.ID)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Unable to delete user %s", err.Error()))
	}
	return nil
}
