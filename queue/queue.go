package queue

type Queue []interface{}

func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)
}

func (q *Queue) Poll() interface{} {
	i := q.Peek()
	*q = (*q)[1:]
	return i
}

func (q *Queue) Peek() interface{} {
	return (*q)[0]
}

func (q *Queue) IsEmpty() bool {
	return 0 == len(*q) || *q == nil
}
