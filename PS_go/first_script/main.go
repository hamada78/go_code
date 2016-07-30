/* Generate some reports that will help keep track of the company's power plants and how much demand the power grid is exerting on those plants.
as a cool little console application that demonstrates a lot of the key concepts that are involved when developing with Go.
*/

package main

import (
	"fmt"
)

func main() {
	plantCapacities := []float64{30, 30, 30, 60, 60, 100}
	capacity := plantCapacities[0] + plantCapacities[1] + plantCapacities[2] + plantCapacities[3] + plantCapacities[4] + plantCapacities[5]
	gridload := 75.0
	utilization := gridload / capacity

	fmt.Printf("%-15s%.0f\n", "Capacity: ", capacity)
	fmt.Printf("%-15s%.0f\n", "Load: ", gridload)
	fmt.Printf("%-15s%.1f%%\n", "Utilization:", utilization*100)
}
