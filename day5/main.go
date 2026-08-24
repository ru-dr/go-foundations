package main

import "fmt"

type Logger interface {
	Log(message string)
	Level() string
}

type ConsoleLogger struct{}

func (c ConsoleLogger) Log(message string) {
	fmt.Println(message)
}

func (c ConsoleLogger) Level() string {
	return "INFO"
}

type FileLogger struct{}

func (f FileLogger) Log(message string) {
	fmt.Println("[FILE] " + message)
}

func (f FileLogger) Level() string {
	return "ERROR"
}

func PrintLog(l Logger, msg string) {
	fmt.Printf("[%s] %s\n", l.Level(), msg)
	l.Log(msg)
}

type Storer interface {
	Save(data string)
	Load() string
	Clear()
}

type MemoryStore struct {
	data string
}

func (m *MemoryStore) Save(data string) { m.data = data }
func (m *MemoryStore) Load() string     { return m.data }
func (m *MemoryStore) Clear()           { m.data = "" }

type FileStore struct {
	contents string
}

func (f *FileStore) Save(data string) { f.contents = "[file] " + data }
func (f *FileStore) Load() string     { return f.contents }
func (f *FileStore) Clear()           { f.contents = "" }

func UseStore(s Storer) {
	s.Save("hello")
	fmt.Println(s.Load())
	s.Clear()
	fmt.Printf("after clear: %q\n", s.Load())
}

func main() {
	PrintLog(ConsoleLogger{}, "Hello from console")
	PrintLog(FileLogger{}, "Error from file")

	UseStore(&MemoryStore{})
	UseStore(&FileStore{})
}
