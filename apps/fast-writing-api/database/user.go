package database

import (
	"errors"
	"fast-writing-api/database/model"
	"fast-writing-api/domain"
	"github.com/google/uuid"
)

func (h *SQLHandler) FindUserById(id domain.UserId) (domain.User, error) {
	var user model.User
	db := h.Conn.Where("id = UUID_TO_BIN(?)", id).First(&user)
	if db.Error != nil {
		return domain.User{}, errors.New("cannot find user by id: " + db.Error.Error())
	}

	return domain.User{
		Id:    domain.UserId(user.Id.String()),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (h *SQLHandler) CreateUser(user domain.User) (domain.UserId, error) {
	id, err := uuid.Parse(string(user.Id))
	if err != nil {
		return domain.UserId(""), errors.New("cannot parse uuid:" + string(user.Id))
	}
	model := model.User{
		Id:    model.MysqlUUID(id),
		Name:  user.Name,
		Email: user.Email,
	}
	db := h.Conn.Create(&model)
	if db.Error != nil {
		return "", errors.New("cannot create user:" + db.Error.Error())
	}
	return user.Id, nil
}
