package exercises

type Regions []string

func (es Exercises) GetAllPossibleRegions() Regions {
	regions := Regions{}
	for _, e := range es {
		hasRegion := false
		for _, r := range regions {
			if e.Region == r {
				hasRegion = true
			}
		}
		if !hasRegion {
			regions = append(regions, e.Region)
		}
	}
	return regions
}
