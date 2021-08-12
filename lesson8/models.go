package lesson8

// для хранения флагов
type programFlag struct {
	DeleteDuplicate bool
	ConfirmDelete   bool
}

// Название и размер файла
type FileInfo struct {
	Name string
	Size int64
}
