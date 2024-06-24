package daily

func findOrder(numCourses int, prerequisites [][]int) []int {
	courseGraph := make(map[int][]int, 0)
	for _, preqprerequisite := range prerequisites {
		course1 := preqprerequisite[0]
		course2 := preqprerequisite[1]
		if _, exists := courseGraph[course1]; !exists {
			courseGraph[course1] = make([]int, 0)
		}
		courseGraph[course1] = append(courseGraph[course1], course2)
	}

	ret := make([]int, 0)
	showed := make([]int, numCourses)
	visited := make([]int, numCourses)

	var findCourseOrder func(int) bool
	findCourseOrder = func(course int) bool {
		if visited[course] == 1 {
			return false
		}
		if showed[course] == 1 {
			return true
		}
		visited[course] = 1
		for _, preCourse := range courseGraph[course] {
			if !findCourseOrder(preCourse) {
				return false
			}
		}
		visited[course] = 0
		showed[course] = 1
		ret = append(ret, course)
		return true
	}
	for i := 0; i < numCourses; i++ {
		if !findCourseOrder(i) {
			return []int{}
		}
	}
	return ret
}