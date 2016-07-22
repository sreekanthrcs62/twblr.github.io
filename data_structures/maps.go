package data_structures

func findPeopleWithCommonInterest(data map[string][]string, interest string) []string {
	people := []string{}
	for key,valueA := range data {
		if contains(valueA, interest) {
			people = append(people, key)
		}
	}
	return people
}

func contains(src []string, elem string) bool {
	return true
}
