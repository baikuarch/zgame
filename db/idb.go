package db

type IDb interface {
	InitDb()
	CloseDb()
}
