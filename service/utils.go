package service

func parseCitiesToString(cities []*City) (result string) {
	for position, city := range cities {
		if position != len(cities)-1 {
			result += city.Name + " - "
		} else {
			result += city.Name
		}
	}
	return result
}
