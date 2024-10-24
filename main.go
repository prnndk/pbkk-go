package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("======== PBKK GOLANG PROGRESS ========")

	conferenceName := "PBKK-seminar"
	const conferenceTickets = 10
	var remainingTickets uint = conferenceTickets
	bookings := []string{}

	fmt.Printf("Welcome to %v bookin system\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are tickets available\n", conferenceTickets, remainingTickets)

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// --- LOOPING
	for {
		fmt.Println("Please enter your first name")
		fmt.Scanln(&firstName)

		fmt.Println("Please enter your last name")
		fmt.Scanln(&lastName)

		fmt.Println("Please enter your email")
		fmt.Scanln(&email)

		fmt.Println("Please enter the number of tickets you want to book")
		fmt.Scanln(&userTickets)

		// --- VALIDATION INPUT
		validName := len(firstName) >= 2 && len(lastName) >= 2
		validEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
		validTickets := userTickets > 0 && userTickets <= remainingTickets

		if validEmail && validName && validTickets {
			remainingTickets -= userTickets

			fmt.Printf("Thank you %v %v for booking %v tickets for %v. You will receive a convirmation email at %v\n", firstName, lastName, userTickets, conferenceName, email)

			fmt.Printf("There are %v tickets remaining\n", remainingTickets)

			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Data type : %T\n", bookings)
			fmt.Printf("Data : %v\n", bookings)
			fmt.Printf("Data length : %v\n", len(bookings))
		} else {
			fmt.Println("There are many invalid inputs")
		}

		// --- CONDITIONAL
		if remainingTickets <= 0 {
			fmt.Println("All tickets are sold out")
			break
		}

	}

	// --- CONTOH LAIN DARI LOOPING
	jumlah_iterasi := 5
	for i := 0; i < jumlah_iterasi; i++ {
		fmt.Println("iterasi nomer ke: ", i)
	}

	books := []string{"Book 1", "Book 2", "Book 3"}
	for _, book := range books {
		fmt.Println(book)
	}

	// --- SWITCH CASE
	fmt.Println("Tolong masukkan Plat Nomor anda")
	var plat string
	fmt.Scan(&plat)

	plat = strings.ToUpper(plat)

	switch plat {
	case "B":
		fmt.Println("Halo Jakarta, Depok dan Bekasi")
	case "F":
		fmt.Println("Halo Bogor")
	case "D":
		fmt.Println("Halo Bandung")
	case "AG":
		fmt.Println("Halo karesidenan Kediri")
	case "E":
		fmt.Println("Halo karesidenan Cirebon")
	case "L":
		fmt.Println("Halo Surabaya")
	case "W":
		fmt.Println("Halo Sidoarjo")
	case "Z":
		fmt.Println("Halo Sumedang")
	default:
		fmt.Println("Maaf plat kamu gaada di database")
	}

}
