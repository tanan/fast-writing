package database

import (
	"errors"
	"fast-writing-api/database/model"
	"fast-writing-api/domain"
	"github.com/google/uuid"
	"time"
)

func (h *SQLHandler) FindContentsList(limit int32, offset int32) ([]*domain.Contents, error) {
	var m []model.Contents
	db := h.Conn.Where("id > ?", offset).Order("id").Limit(limit).Find(&m)
	if db.Error != nil {
		return []*domain.Contents{}, errors.New("cannot find contents list: " + db.Error.Error())
	}

	var contentsList []*domain.Contents
	for _, v := range m {
		var user model.User
		userDb := h.Conn.Where("id = UUID_TO_BIN(?)", v.UserId.String()).First(&user)
		if userDb.Error != nil {
			return []*domain.Contents{}, errors.New("cannot find user by id: " + db.Error.Error())
		}
		contents := domain.Contents{
			ContentsId:  domain.ContentsId(v.Id),
			Creator:     user.Name,
			Title:       v.Title,
			QuizList:    nil,
			LastUpdated: v.LastUpdated,
		}
		contentsList = append(contentsList, &contents)
	}
	return contentsList, nil
}

func (h *SQLHandler) FindContentsById(id domain.ContentsId) (domain.Contents, error) {
	var contents model.Contents
	db := h.Conn.Where("id = ?", id).First(&contents)
	if db.Error != nil {
		return domain.Contents{}, errors.New("cannot find contents by id: " + db.Error.Error())
	}
	var quiz []model.Quiz
	quizDb := h.Conn.Where("contents_id = ?", id).Find(&quiz)
	if quizDb.Error != nil {
		return domain.Contents{}, errors.New("cannot find quiz by contents_id: " + db.Error.Error())
	}

	var user model.User
	println(contents.UserId.String())
	userDb := h.Conn.Where("id = UUID_TO_BIN(?)", contents.UserId.String()).First(&user)
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

func (h *SQLHandler) CreateContents(contents domain.Contents, userId domain.UserId) (int, error) {
	uid, err := uuid.Parse(string(userId))
	if err != nil {
		return 0, errors.New(string(userId) + " is not valid:" + err.Error())
	}
	db := h.Conn.Exec("INSERT INTO `writing_contents` (`user_id`,`title`,`last_updated`) VALUES (UUID_TO_BIN(?),?,?)", uid, contents.Title, time.Now())
	if db.Error != nil {
		return 0, errors.New("cannot create contents: " + db.Error.Error())
	}
	return 1, nil
}

func (h *SQLHandler) CreateQuiz(contentsId domain.ContentsId, quiz domain.Quiz) (int, error) {
	m := model.Quiz{
		ContentsId: int64(contentsId),
		Question:   quiz.Question,
		Answer:     quiz.Answer,
	}
	db := h.Conn.Create(&m)
	if db.Error != nil {
		return 0, errors.New("cannot create quiz: " + db.Error.Error())
	}
	return 1, nil
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
