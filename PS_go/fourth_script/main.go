/* Generate some reports that will help keep track of the company's power plants and how much demand the power grid is exerting on those plants.
as a cool little console application that demonstrates a lot of the key concepts that are involved when developing with Go.
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	plants := []PowerPlant{
		PowerPlant{hydro, 300, active},
		PowerPlant{wind, 30, active},
		PowerPlant{wind, 25, inactive},
		PowerPlant{wind, 35, active},
		PowerPlant{solar, 45, unavailable},
		PowerPlant{solar, 40, inactive},
	}
	grid := PowerGrid{300, plants}

	fmt.Println("1) Generate Power Plant Report")
	fmt.Println("2) Generate Power GridReport")
	fmt.Print("Please choose an option: ")

	var option string
	fmt.Scanln(&option)

	switch option {
	case "1":
		grid.generatePlantReport()
	case "2":
		grid.generatePowerGridReport()
	default:
		fmt.Println("Unkown option, please choose 1 or 2 only")
	}
}

type PlantType string

const (
	hydro PlantType = "Hydro"
	wind  PlantType = "Wind"
	solar PlantType = "Solar"
)

type Plantstatus string

const (
	active      Plantstatus = "Active"
	inactive    Plantstatus = "inactive"
	unavailable Plantstatus = "unavailable"
)

type PowerPlant struct {
	PlantType PlantType
	capacity  float64
	status    Plantstatus
}

type PowerGrid struct {
	load   float64
	plants []PowerPlant
}

func (pg *PowerGrid) generatePlantReport() {
	for idx, p := range pg.plants {
		label := fmt.Sprintf("%s%d", "Plant #", idx)
		fmt.Println(label)
		fmt.Println(strings.Repeat("-", len(label)))
		fmt.Printf("%-15s%s\n", "Type:", p.PlantType)
		fmt.Printf("%-15s%.0f\n", "Capacity:", p.capacity)
		fmt.Printf("%-15s%s\n", "Status", p.status)
		fmt.Println("")
	}
}

func (pg *PowerGrid) generatePowerGridReport() {
	capacity := 0.
	for _, p := range pg.plants {
		if p.status == active {
			capacity += p.capacity
		}
	}
	label := "Power Grid Report"
	fmt.Println(label)
	fmt.Println(strings.Repeat("-", len(label)))
	fmt.Printf("%-15s%.0f\n", "Capacity: ", capacity)
	fmt.Printf("%-15s%.0f\n", "Load: ", pg.load)
	fmt.Printf("%-15s%.1f%%\n", "Utilization:", pg.load/capacity*100)
}
