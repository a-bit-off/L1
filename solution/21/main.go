/*
Реализовать паттерн «адаптер» на любом примере.
*/
package main

// первый интерфейс и его имплементация структурой File
type Reader interface {
	Read() string
}

type File struct{}

func (f *File) Read() string {
	return "data from file"
}

// второй интерфейс и структура которая адаптиркует первую ко второму интерфейсу
type Writer interface {
	Writer(string)
}

type FileWriter struct {
	file *File
}

func (fw *FileWriter) Write(data string) {
	fw.file.Read()
}

func main() {
	writer := &FileWriter{file: &File{}}
	writer.Write("Some data")
}
