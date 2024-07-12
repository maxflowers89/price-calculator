package iomanager

type IOManager interface {
	ReadLines() ([]string, error)
	WriteData(data interface{}) error
}
