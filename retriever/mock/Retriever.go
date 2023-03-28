package mock

import "fmt"

type Retriever struct {
	Contents string
}

func (r *Retriever) String() string {
	return fmt.Sprintf("Retriever :{Contente:%s}", r.Contents)
}

func (r *Retriever) Post(url string, form map[string]string) string {
	r.Contents = form["Contents"]
	return "ok"
}

func (r *Retriever) Get(url string) string {
	return r.Contents
}
