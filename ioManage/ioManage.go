package ioManage

type IOManage interface {
	ReadLine() ([]string, error)
	WriteResult(data any) error
}
