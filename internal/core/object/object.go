package object

type Object interface {
	Type() string
	Size() int
	Serialize() []byte
	Deserialize([]byte) ([]byte, error)
}
