package jumpgameiii

type Queue struct {
	Elements []int
}

func (q *Queue) Enqueue(e int) {
	q.Elements = append(q.Elements, e)
}

func (q *Queue) Len() int {
	return len(q.Elements)
}

func (q *Queue) Dequeue() int {
	e := q.Elements[0]
	q.Elements[0] = 0
	q.Elements = q.Elements[1:]
	return e
}

func CanReach(arr []int, start int) bool {
	minIndex, maxIndex := 0, len(arr)-1

	q := Queue{}
	isMarked := make(map[int]bool)

	q.Enqueue(start)
	isMarked[start] = true

	for q.Len() > 0 {
		e := q.Dequeue()
		if arr[e] == 0 {
			return true
		}

		index := e + arr[e]
		if index >= minIndex && index <= maxIndex && !isMarked[index] {
			q.Enqueue(index)
			isMarked[index] = true
		}

		index = e - arr[e]
		if index >= minIndex && index <= maxIndex && !isMarked[index] {
			q.Enqueue(index)
			isMarked[index] = true
		}
	}

	return false
}
