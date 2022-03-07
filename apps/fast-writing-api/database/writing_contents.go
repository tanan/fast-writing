package database

import (
	"errors"
	"fast-writing-api/database/model"
	"fast-writing-api/domain"
)

func (h *SQLHandler) FindContentsById(id domain.ContentsId) (domain.Contents, error) {
	var contents model.Contents
	db := h.Conn.Where("id = ?", id).First(&contents)
	if db.Error != nil {
		return domain.Contents{}, errors.New("cannot find contents by id: " + db.Error.Error())
	}
	var quiz []model.Quiz
	quizDb := h.Conn.Where("contents_id = ?", id).First(&quiz)
	if quizDb.Error != nil {
		return domain.Contents{}, errors.New("cannot find quiz by contents_id: " + db.Error.Error())
	}

	var user model.User
	userDb := h.Conn.Where("id = UUID_TO_BIN(?)", id).First(&user)
	if userDb.Error != nil {
		return domain.Contents{}, errors.New("cannot find user by id: " + db.Error.Error())
	}
	return domain.Contents{
		ContentsId:  domain.ContentsId(contents.Id),
		Creator:     user.Name,
		Title:       contents.Title,
		QuizList:    h.toQuizList(quiz),
		LastUpdated: contents.LastUpdated,
	}, nil
}

func (h *SQLHandler) toQuizList(l []model.Quiz) []domain.Quiz {
	var quizList []domain.Quiz
	for _, v := range l {
		quizList = append(quizList, domain.Quiz{
			Id:         v.Id,
			Question:   v.Question,
			Answer:     v.Answer,
			ContentsId: domain.ContentsId(v.ContentsId),
		})
	}
	return quizList
}
