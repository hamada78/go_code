/* Generate some reports that will help keep track of the company's power plants and how much demand the power grid is exerting on those plants.
as a cool little console application that demonstrates a lot of the key concepts that are involved when developing with Go.
*/

package main

import (
	"fmt"
)

func main() {

	plantCapacities := []float64{30, 30, 30, 60, 60, 100}
	activePlants := []int{0, 1}
	gridload := 75.0

	fmt.Println("1) Generate Power Plant Report")
	fmt.Println("2) Generate Power GridReport")
	fmt.Print("Please choose an option: ")

	var option string
	fmt.Scanln(&option)

	switch option {
	case "1":
		generatePlantCapacityReport(plantCapacities...)
	case "2":
		generatePowerGridReport(activePlants, plantCapacities, gridload)

	default:
		fmt.Println("Unkown option, please choose 1 or 2 only")
	}
}
func generatePlantCapacityReport(plantCapacities ...float64) {
	for idx, cap := range plantCapacities {
		fmt.Printf("Plant %d capacity: %.0f\n", idx, cap)
	}
}
func generatePowerGridReport(activePlants []int, plantCapacities []float64, gridload float64) {
	capacity := 0.
	for _, palntId := range activePlants {
		capacity += plantCapacities[palntId]
	}
	fmt.Printf("%-15s%.0f\n", "Capacity: ", capacity)
	fmt.Printf("%-15s%.0f\n", "Load: ", gridload)
	fmt.Printf("%-15s%.1f%%\n", "Utilization:", gridload/capacity*100)
}
