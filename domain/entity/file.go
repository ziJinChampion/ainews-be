package entity

type File struct {
	ID        int64
	Name      string
	Extension string
	Content   []byte
}
