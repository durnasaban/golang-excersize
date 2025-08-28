package purchase

import "fmt"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
    textInResult := "%s is clearly the better choice."
    
    if option1 < option2 {
        return fmt.Sprintf(textInResult, option1)
    }
    
	return fmt.Sprintf(textInResult, option2)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	ageInInt := int(age)
    
    switch {
        case ageInInt < 3:
        	return originalPrice * float64(80) / float64(100)
        case ageInInt >= 3 && ageInInt < 10:
        	return originalPrice * float64(70) / float64(100)
        default:
        	return originalPrice * float64(50) / float64(100)
    }
}
