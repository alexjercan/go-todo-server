package models

type Item struct {
	ID          int64 `bun:",pk,autoincrement"`
	Description string
	Completed   bool `bun:"default:false"`
}
