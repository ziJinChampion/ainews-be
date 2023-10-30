package entity

type File struct {
	Id        int64
	Name      string
	Extension string
	Content   []byte
	Size      int64
	Url       string
}
