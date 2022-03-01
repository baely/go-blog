package blog

import "fmt"

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (a Author) DisplayName() string {
	return fmt.Sprintf("id: %d, name: %s", a.Id, a.Name)
}
