// Purchase is a package that helps you purchase a vehicle, but it's not a real Go package.
package purchase

func NeedsLicense(kind string) bool {
	if kind == "car" {
		return true
	} else if kind == "truck" {
		return true
	} else {
		return false
	}
}

func ChooseVehicle(option1, option2 string) string {
	carChoice := option1 < option2
	if carChoice {
		return (option1 + " is clearly the better choice.")
	} else {
		return (option2 + " is clearly the better choice.")
	}
}

func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
		return originalPrice * .8
	} else if age < 10 {
		return originalPrice * .7
	} else {
		return originalPrice * .5
	}
}
