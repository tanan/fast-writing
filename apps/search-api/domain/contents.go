package domain

type ContentsId string

type Contents struct {
	Id       ContentsId
	Title    string
	Category string
	Username string
	Quiz     string
	Answer   string
}
