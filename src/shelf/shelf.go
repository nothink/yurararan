package shelf

type Shelf interface{
	All() []interface{}
	Append([]interface{}) []interface{}
}
