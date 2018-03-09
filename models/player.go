package models

//Player data struct
type Player struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Score int    `json:"score"`
}
