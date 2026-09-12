package validater

type TodoValidater struct {
}

func NewTodoValidater() *TodoValidater {
	return &TodoValidater{}
}
func (t *TodoValidater) Struct(s interface{}) error {
	return nil
}
