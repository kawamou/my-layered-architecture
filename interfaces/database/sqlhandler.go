package interfaces

type SqlHandler interface {
	Find(interface{}, ...interface{}) interface{}
	Create(interface{}) interface{}
}