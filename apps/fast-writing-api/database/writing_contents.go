package database

import (
	"errors"
	"fast-writing-api/database/model"
	"fast-writing-api/domain"
	"gorm.io/gorm/clause"
	"time"
)

func (h *SQLHandler) FindContentsList(limit int32, offset int32) ([]*domain.Contents, error) {
	var m []model.Contents
	db := h.Conn.Where("id > ? and scope = ?", offset, "public").Order("id").Limit(int(limit)).Find(&m)
	if db.Error != nil {
		return []*domain.Contents{}, errors.New("cannot find contents list: " + db.Error.Error())
	}
	return h.toContentsList(m)
}

func (h *SQLHandler) FindContentsListByUserId(userId domain.UserId, limit int32, offset int32) ([]*domain.Contents, error) {
	var m []model.Contents
	db := h.Conn.Where("id > ? and user_id = ?", offset, userId).Order("id").Limit(int(limit)).Find(&m)
	if db.Error != nil {
		return []*domain.Contents{}, errors.New("cannot find contents list by user: " + db.Error.Error())
	}
	return h.toContentsList(m)
}

func (h *SQLHandler) toContentsList(m []model.Contents) ([]*domain.Contents, error) {
	var contentsList []*domain.Contents
	for _, v := range m {
		var user model.User
		db := h.Conn.Where("id = ?", v.UserId).First(&user)
		if db.Error != nil {
			return []*domain.Contents{}, errors.New("cannot find user by id: " + db.Error.Error())
		}
		contents := domain.Contents{
			ContentsId:  domain.ContentsId(v.Id),
			Creator:     user.Name,
			Title:       v.Title,
			Description: v.Description,
			Scope:       v.Scope,
			QuizList:    nil,
			LastUpdated: v.LastUpdated,
		}
		contentsList = append(contentsList, &contents)
	}
	return contentsList, nil
}

func (h *SQLHandler) FindContentsById(id domain.ContentsId) (domain.Contents, error) {
	var m model.Contents
	db := h.Conn.Where("id = ?", id).First(&m)
	if db.Error != nil {
		return domain.Contents{}, errors.New("cannot find contents by id: " + db.Error.Error())
	}
	var quiz []model.Quiz
	quizDb := h.Conn.Where("contents_id = ?", id).Find(&quiz)
	if quizDb.Error != nil {
		return domain.Contents{}, errors.New("cannot find quiz by contents_id: " + db.Error.Error())
	}

	var user model.User
	userDb := h.Conn.Where("id = ?", m.UserId).First(&user)
	if userDb.Error != nil {
		return domain.Contents{}, errors.New("cannot find user by id: " + db.Error.Error())
	}
	contents := domain.Contents{
		ContentsId:  domain.ContentsId(m.Id),
		Creator:     user.Name,
		Title:       m.Title,
		Description: m.Description,
		Scope:       m.Scope,
		QuizList:    h.toQuizList(quiz),
		LastUpdated: m.LastUpdated,
	}
	return contents, nil
}

func (h *SQLHandler) CreateContents(contents domain.Contents, userId domain.UserId) (domain.Contents, error) {
	m := model.Contents{
		Id:          int64(contents.ContentsId),
		UserId:      string(userId),
		Title:       contents.Title,
		Description: contents.Description,
		Scope:       contents.Scope,
		LastUpdated: time.Now(),
	}
	if db := h.Conn.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"user_id": m.UserId, "title": m.Title, "description": m.Description, "scope": m.Scope, "last_updated": time.Now()}),
	}).Create(&m); db.Error != nil {
		return domain.Contents{}, errors.New("cannot create contents: " + db.Error.Error())
	}

	for i, v := range contents.QuizList {
		if quizId, err := h.CreateQuiz(userId, v); err != nil {
			return domain.Contents{}, err
		} else {
			contents.QuizList[i] = domain.Quiz{
				Id:         domain.QuizId(quizId),
				Question:   v.Question,
				Answer:     v.Answer,
				ContentsId: v.ContentsId,
			}
		}
	}
	contents.ContentsId = domain.ContentsId(m.Id)
	contents.LastUpdated = m.LastUpdated
	return contents, nil
}

func (h *SQLHandler) DeleteContents(userId domain.UserId, contentsId domain.ContentsId) (int64, error) {
	db := h.Conn.Where("user_id = ?", userId).Delete(&model.Contents{}, contentsId)
	if db.Error != nil {
		return 0, errors.New("cannot delete contents: " + db.Error.Error())
	}
	return 1, nil
}

func (h *SQLHandler) CreateQuiz(userId domain.UserId, quiz domain.Quiz) (int64, error) {
	m := model.Quiz{
		Id:          int64(quiz.Id),
		ContentsId:  int64(quiz.ContentsId),
		UserId:      string(userId),
		Question:    quiz.Question,
		Answer:      quiz.Answer,
		LastUpdated: time.Now(),
	}
	db := h.Conn.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"contents_id": m.ContentsId, "user_id": m.UserId, "question": m.Question, "answer": m.Answer, "last_updated": time.Now()}),
	}).Create(&m)
	if db.Error != nil {
		return 0, errors.New("cannot create quiz: " + db.Error.Error())
	}
	return m.Id, nil
}

func (h *SQLHandler) toQuizList(l []model.Quiz) []domain.Quiz {
	var quizList []domain.Quiz
	for _, v := range l {
		quizList = append(quizList, domain.Quiz{
			Id:         domain.QuizId(v.Id),
			Question:   v.Question,
			Answer:     v.Answer,
			ContentsId: domain.ContentsId(v.ContentsId),
		})
	}
	return quizList
}

func (h *SQLHandler) DeleteQuiz(userId domain.UserId, contentsId domain.ContentsId, quizId domain.QuizId) (int64, error) {
	db := h.Conn.Where("user_id = ? and contents_id = ?", userId, contentsId).Delete(&model.Quiz{}, quizId)
	if db.Error != nil {
		return 0, errors.New("cannot delete quiz: " + db.Error.Error())
	}
	return 1, nil
}
