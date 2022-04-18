package models

type Item struct {
	ID          int64  `bun:",pk,autoincrement"`
	Description string `bun:"size:255"`
	Completed   bool   `bun:"default:false"`
}
