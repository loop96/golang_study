package model

type Movie struct {
	Name        string
	Tag         []string
	Country     string
	TimeMinutes float64
	Date        string
	Stars       float64
	Introduce   string
	Casts       []Cast
}
type Cast struct {
	ActorName string
	PlayName  string
	PicUrl    string
}
