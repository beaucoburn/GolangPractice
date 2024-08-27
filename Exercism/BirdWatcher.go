// BirdWatcher is a module that helps you keep track of the number of birds that have visited your garden, but is not a real Go package.
package birdwatcher

func TotalBirdCount(birdsPerDay []int) int {
    total := 0
    for i := 0; i < len(birdsPerDay); i++ {
        total += birdsPerDay[i]
    }
    return total
}

func BirdsInWeek(birdsPerDay []int, week int) int {
	startIndex := (week - 1) * 7
    endIndex := startIndex + 7
    weekTotal := 0
    for i := startIndex; i < endIndex; i++ {
        weekTotal += birdsPerDay[i]
    }
    return weekTotal
}

func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i += 2 {
        birdsPerDay[i]++
    }
    return birdsPerDay
}
