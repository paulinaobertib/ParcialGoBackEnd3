package tickets

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Ticket struct {
	Id          int64
	Name        string
	Email       string
	Destination string
	Time        string
	Price       float64
}

var infoTickets []Ticket

// requerimiento 1
func GetTotalTickets(destination string) (int, error) {
	if (destination == "") {
		return 0, errors.New("El destino no puede estar vacio")
	}

	cont := 0

	for _, ticket := range infoTickets {
		if ticket.Destination == destination {
			cont++
		}
	}

	if (cont==0) {
		return 0, errors.New("Destino invalido.")
	}

	return cont, nil
}

// requerimiento 2
func GetCountByPeriod(timeFrame string) (int, error) {
	if (timeFrame == "") {
		return 0, errors.New("Debe pasar un periodo de tiempo")
	}

	contEarlyMorning := 0
	contMorning := 0
	contAfternoon := 0
	contNight := 0

	for _, ticket := range infoTickets {
		timeSplit := strings.Split(ticket.Time, ":")
		hour, err := strconv.ParseInt(timeSplit[0], 10, 64);
		if (err != nil) {
			return 0, errors.New("Error al convertir la hora a entero")
		}
		
		if (hour >= 0 && hour <= 6) {
			contEarlyMorning++
		} else if (hour >= 7 && hour <= 12) {
			contMorning++
		} else if (hour <= 13 && hour >= 19) {
			contAfternoon++
		} else if (hour >= 20 && hour <= 23) {
			contNight++
		}
	}

	switch timeFrame {
	case "Early morning":
		return contEarlyMorning, nil
	case "Morning":
		return contMorning, nil
	case "Afternoon":
		return contAfternoon, nil
	case "Night":
		return contNight, nil
	default:
		return 0, errors.New("Debe pasar un periodo valido")
	}
}

// requerimiento 3
func AverageDestination(destination string) (float64, error) {
	if (destination == "") {
		return 0, errors.New("Debe pasar un destino")
	}
	
	total, err := GetTotalTickets(destination)
	if (err != nil) {
		return 0, err
	}

	percentage :=  float64(total) / float64(len(infoTickets)) * 100

	return percentage, nil
}

// funciones para obtener los datos del archivo .csv
func GetTickets() {
	data := strings.Split(readFile("./tickets.csv"), "\n")

	for i:=0; i<len(data)-1; i++ {
		var line = strings.Split(string(data[i]), ",")

		id, err := strconv.ParseInt(line[0], 10, 64)
		if err != nil {
			fmt.Println("Problemas para parsear el id")
		}

		price, err := strconv.ParseFloat(line[5], 64)
		if err != nil {
			fmt.Println("Problemas para parsear el precio")
		}

		ticket := Ticket {
			Id:          id,
			Name:        line[1],
			Email:       line[2],
			Destination: line[3],
			Time:        line[4],
			Price:       price,
		}

		infoTickets = append(infoTickets, ticket)
	}
}

func readFile(name string) string{
    defer func() {
        if err := recover(); err != nil {
            fmt.Println(err)
        }
    }()

    res, err := os.ReadFile(name)
    
    if err != nil {
        panic("No se ha podido leer el archivo")
    }

	return string(res);
}