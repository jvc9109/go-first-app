package filewriter

type Writer interface {
	WriteToFile(filename string) error
}
