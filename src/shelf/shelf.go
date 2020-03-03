package shelf

// Shelf シェルフ
type Shelf interface {
	All() []interface{}
	Append([]interface{}) []interface{}
}
