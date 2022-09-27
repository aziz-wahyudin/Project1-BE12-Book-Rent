package controller

import (
	"Alterra/Project1-BE12-Book-Rent/model"
)

type UserControll struct {
	Model model.UserModel
}

func (uc UserControll) GetAll() ([]model.User, error) {
	res, err := uc.Model.GetAll()
	if err != nil {
		return nil, err
	}
	return res, nil
}