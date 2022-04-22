package database

import (
	"errors"
	"fast-writing-api/database/model"
	"fast-writing-api/domain"
	"time"
)

func (h *SQLHandler) FindUserById(id domain.UserId) (domain.User, error) {
	var user model.User
	db := h.Conn.Where("id = ?", id).First(&user)
	if db.Error != nil {
		return domain.User{}, errors.New("cannot find user by id: " + db.Error.Error())
	}

	return domain.User{
		Id:    domain.UserId(user.Id),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (h *SQLHandler) UpdateUser(user domain.User) (domain.UserId, error) {
	model := model.User{
		Id:          string(user.Id),
		Name:        user.Name,
		Email:       user.Email,
		LastUpdated: time.Now(),
	}
	db := h.Conn.Save(&model)
	if db.Error != nil {
		return "", errors.New("failed to update user:" + db.Error.Error())
	}
	return user.Id, nil
}

func (h *SQLHandler) CreateUser(user domain.User) (domain.UserId, error) {
	model := model.User{
		Id:          string(user.Id),
		Name:        user.Name,
		Email:       user.Email,
		LastUpdated: time.Now(),
	}
	db := h.Conn.Create(&model)
	if db.Error != nil {
		return "", errors.New("cannot create user:" + db.Error.Error())
	}
	return user.Id, nil
}
