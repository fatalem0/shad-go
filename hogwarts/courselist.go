//go:build !solution

package hogwarts

func exploreCoursePrereqs(
	course string,
	courses map[string][]string,
	learned map[string]bool,
	entered map[string]bool,
	res []string,
) []string {
	if entered[course] {
		panic("The course has been entered")
	}

	if learned[course] {
		return res
	}

	entered[course] = true

	for _, prereq := range courses[course] {
		res = exploreCoursePrereqs(prereq, courses, learned, entered, res)
	}

	learned[course] = true
	entered[course] = false
	res = append(res, course)

	return res
}

func GetCourseList(prereqs map[string][]string) []string {
	learned := make(map[string]bool)
	entered := make(map[string]bool)
	res := make([]string, 0)

	for course := range prereqs {
		res = exploreCoursePrereqs(course, prereqs, learned, entered, res)
	}

	return res
}
