package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avg_prep_time int) int{
	if avg_prep_time == 0 {
        avg_prep_time = 2
    }    

    return len(layers) * avg_prep_time
}


// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	for _, layer := range layers {
		if layer == "noodles" {
			noodles += 50
		} else if layer == "sauce" {
			sauce += 0.2
		}
	}
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(ingreFriend []string, ingreMe []string) {
    lastItemIndex := len(ingreFriend) - 1
    ingreMe[len(ingreMe) - 1] = ingreFriend[lastItemIndex]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(portions []float64, portionCount int) [](float64){
    var s []float64
    for _, portion := range(portions) {
		x := portion / 2.0
        s = append(s, x * float64(portionCount))
    }
    return s
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
