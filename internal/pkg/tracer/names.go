package tracer

import "fmt"

const (
	// ServiceName название сервиса
	ServiceName = "Task manager"

	Name = "Task manager tracer"
)

// MakeSpanName подставляет в название спана название сервиса
func MakeSpanName(name string) string {
	return fmt.Sprintf("%s/%s", ServiceName, name)
}
