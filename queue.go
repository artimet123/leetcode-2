package main 

import "fmt"

type Queue struct {
	List []interface{}
}

func NewQueue()*Queue{
	list := []interface{}{}
	return &Queue{list}
}

func (q *Queue)Enqueue(val interface{}){
	q.List = append(q.List, val)
}

func (q *Queue)Dequeue()(interface{}, bool){
	if q.Len() == 0 {
		return "", false
	}

	val := q.List[0]
	q.List = q.List[1:]
	return val, true
}

func (q *Queue) Len() int {
    return len(q.List)
}

func main(){
	q := NewQueue()
	q.Enqueue(123)
	q.Enqueue("mao")

	val, ok := q.Dequeue()

	fmt.Println(val, ok)
	val2, ok2 := q.Dequeue()
	fmt.Println(val2, ok2)
}