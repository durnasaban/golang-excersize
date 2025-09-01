package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, preparationTime int) int {
    if preparationTime == 0 {
        return len(layers) * 2
    }

    return len(layers) * preparationTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64){
    noodlesForEach := 50
    litersOfSauce := 0.2

    for _, layer := range layers {
        switch layer {
            case "noodles":
            	noodles += noodlesForEach
            case "sauce":
            	sauce += litersOfSauce
        }
    }

    return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(recipeOfFriend, ownRecipe []string) {
    ownRecipe[len(ownRecipe)-1] = recipeOfFriend[len(recipeOfFriend)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) (sliceOfPortions []float64) {
    scaledQuantities := make([]float64, len(quantities))

    for index, value := range quantities {
        scaledQuantities[index] = value * float64(portions) / float64(2)
    }

    return scaledQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
