package main

/*
* using a hashmap to store the matched numbers and for easier lookup of intersections
*/
func intersection(x,y []int) (z []int) {
	lookup := make(map[int]bool)
	
	//convert one of the array in to map for easier lookup
	for _,i := range x {
		lookup[i] = false
	}

	// mark the element as true if it matches with other array
	for _,i := range y {
		if _, ok := lookup[i]; ok {
			lookup[i] = true
		}
	}
	//filter out all the matched elements
	for key,i := range lookup {
		if i {
			z = append(z,key)
		}
	}
	return 
	
}