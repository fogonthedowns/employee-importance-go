package main

/**
 * Definition for Employee.
 * type Employee struct {
 *     Id int
 *     Importance int
 *     Subordinates []int
 * }
 */

func getImportance(employees []*Employee, id int) int {
	emap := make(map[int]*Employee)
	for _, value := range employees {
		emap[value.Id] = value
	}

	return dfs(id, emap)
}

func dfs(eid int, m map[int]*Employee) int {
	e := m[eid]

	ans := e.Importance
	for _, value := range e.Subordinates {

		ans += dfs(value, m)
	}
	return ans
}
