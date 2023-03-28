package _19

import "container/list"

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
	res := 0

	var idEmpMap = make(map[int]*Employee)
	for _, e := range employees {
		idEmpMap[e.Id] = e
	}

	var emp = idEmpMap[id]
	if emp == nil {
		return 0
	}

	queue := list.New()
	queue.PushBack(emp)

	for queue.Len() > 0 {
		e := queue.Remove(queue.Front()).(*Employee)
		res += e.Importance
		for _, subordinate := range e.Subordinates {
			queue.PushBack(idEmpMap[subordinate])
		}
	}

	return res
}
