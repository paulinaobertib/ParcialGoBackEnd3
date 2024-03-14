package main

import (
	"ctd-b3/internal/tickets"
	"fmt"
)

func main() {
	tickets.GetTickets()

	destintation := make(chan int)
	period := make(chan int)
	average := make(chan float64)

	var dest string
	fmt.Println("Ingrese un destino: ")
	fmt.Scan(&dest)

	var timeFrame string
	var option int
	for {
		fmt.Println("Seleccione una franja horaria:")
		fmt.Println("1. Madrugada")
		fmt.Println("2. Mañana")
		fmt.Println("3. Tarde")
		fmt.Println("4. Noche")
		fmt.Print("Ingrese el número correspondiente a la franja horaria: ")
		fmt.Scan(&option)

		switch option {
		case 1:
			timeFrame = "Early morning"
		case 2:
			timeFrame = "Morning"
		case 3:
			timeFrame = "Afternoon"
		case 4:
			timeFrame = "Night"
		default:
			fmt.Println("Selección no válida. Por favor, ingrese un número del 1 al 4.")
			continue
		}

		break
	}

	var percentage string
	fmt.Println("Ingrese el destino para calcular su porcentaje: ")
	fmt.Scan(&percentage)

	go func(dest string) {
		totalDestination, err := tickets.GetTotalTickets(dest)
		if err != nil {
			fmt.Println("Ocurrió un error calculando el total de tickets. Error:", err)
			destintation <- 0
		} else {
			destintation <- totalDestination
		}
	}(dest)

	go func(timeFrame string) {
		totalPeriod, err := tickets.GetCountByPeriod(timeFrame)
		if err != nil {
			fmt.Println("Ocurrió un error calculando el total de tickets. Error:", err)
			period <- 0
		} else {
			period <- totalPeriod
		}
	}(timeFrame)

	go func(percentage string) {
		totalAverage, err := tickets.AverageDestination(percentage)
		if err != nil {
			fmt.Println("Ocurrió un error calculando el total de tickets. Error:", err)
			average <- 0
		} else {
			average <- totalAverage
		}
	}(percentage)

	fmt.Printf("El total de tickets vendidos a destino fueron %d\n", <-destintation)
	fmt.Printf("El total de tickets vendidos en ese periodo fueron %d\n", <-period)
	fmt.Printf("En promedio, viajaron %.2f personas\n", <-average)
}
