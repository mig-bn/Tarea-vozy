package user_service

import (
	m "Tarea-vozy/models"
	userRepository "Tarea-vozy/repositories/repository"
)

func Create(user m.User) error {

	err := userRepository.Create(user)

	if err != nil {
		return err
	}

	return nil

}

func Read() (m.Users, error) {

	users, err := userRepository.Read()

	if err != nil {
		return nil, err
	}

	return users, nil

}

func Update(user m.User, userID string) error {

	err := userRepository.Update(user, userID)

	if err != nil {
		return err
	}

	return nil

}

func Delete(userID string) error {

	err := userRepository.Delete(userID)

	if err != nil {
		return err
	}

	return nil

}
