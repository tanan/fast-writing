package database

import (
	"errors"
	"fast-writing-api/database/model"
	"fast-writing-api/domain"
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
