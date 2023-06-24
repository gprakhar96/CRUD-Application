package services

import (
	"mypackage/domain/users"
	"mypackage/utils/errors"
)

func CreateNewUser(newUser users.User) (*users.User, *errors.RestErrors) {
	// TODO : Some sort of validation requirements may exist on the bound data object.
	err := newUser.Validate()
	if err != nil {
		return nil, err
	}

	saveErr := newUser.Save()
	if saveErr != nil {
		return nil, saveErr
	}
	return &newUser, nil
}

func GetUserByID(currentUser users.User) (*users.User, *errors.RestErrors) {
	if currentUser.ID <= 0 {
		return nil, errors.NewBadRequestError("User id negative")
	}

	getErr := currentUser.Get()
	if getErr != nil {
		return nil, getErr
	}
	return &currentUser, nil
}

func UpdateUserByID(userToUpdate users.User) (*users.User, *errors.RestErrors) {
	existingUser, err := GetUserByID(userToUpdate)
	if err != nil {
		return nil, err
	}
	if userToUpdate.FirstName != "" {
		existingUser.FirstName = userToUpdate.FirstName
	}
	if userToUpdate.LastName != "" {
		existingUser.LastName = userToUpdate.LastName
	}
	if userToUpdate.Email != "" {
		existingUser.Email = userToUpdate.Email
	}

	updateErr := existingUser.Update()
	if updateErr != nil {
		return nil, updateErr
	}
	return existingUser, nil
}

func DeleteUserByID(userToDelete users.User) (*users.User, *errors.RestErrors) {
	existingUser, err := GetUserByID(userToDelete)
	if err != nil {
		return nil, err
	}
	deleteErr := existingUser.Delete()
	if deleteErr != nil {
		return nil, deleteErr
	}
	return existingUser, nil
}
