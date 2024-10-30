package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/prnndk/pbkk-go/helper"
)

var conferenceName = "Seminar PBKK Kelas D"

const conferenceTicket uint = 50

var remainingTicket uint = conferenceTicket

var bookings = make([]UserData, 0)

type UserData struct {
	firstName  string
	lastName   string
	email      string
	userTicket uint
}

var wg = sync.WaitGroup{}

func main() {
	greetUsers()

	firstName, lastName, email, userTicket := getUserInput()

	isValidName, isValidEmail, isValidTicket := helper.ValidateUserInput(firstName, lastName, email, userTicket, remainingTicket)

	if isValidEmail && isValidName && isValidTicket {

		bookTicket(userTicket, firstName, lastName, email)
		wg.Add(1)
		go sendTicket(userTicket, firstName, lastName, email)

		firstNames := getFirstName()
		fmt.Printf("These Are All Our Bookings: %v\n", firstNames)

		noTicketRemaining := remainingTicket == 0
		if noTicketRemaining {
			fmt.Printf("%v is Sold Out. Come back next century.\n", conferenceName)
		}

	} else {
		if !isValidName {
			fmt.Println("Your name is too short")
		}

		if !isValidEmail {
			fmt.Println("Invalid email address")
		}

		if !isValidTicket {
			fmt.Println("Invalid Number of Tickets")
		}
	}
	wg.Wait()

	// // --- CONTOH LAIN DARI LOOPING
	// jumlah_iterasi := 5
	// for i := 0; i < jumlah_iterasi; i++ {
	// 	fmt.Println("iterasi nomer ke: ", i)
	// }

	// books := []string{"Book 1", "Book 2", "Book 3"}
	// for _, book := range books {
	// 	fmt.Println(book)
	// }

	// // --- SWITCH CASE
	// fmt.Println("Tolong masukkan Plat Nomor anda")
	// var plat string
	// fmt.Scan(&plat)

	// plat = strings.ToUpper(plat)

	// switch plat {
	// case "B":
	// 	fmt.Println("Halo Jakarta, Depok dan Bekasi")
	// case "F":
	// 	fmt.Println("Halo Bogor")
	// case "D":
	// 	fmt.Println("Halo Bandung")
	// case "AG":
	// 	fmt.Println("Halo karesidenan Kediri")
	// case "E":
	// 	fmt.Println("Halo karesidenan Cirebon")
	// case "L":
	// 	fmt.Println("Halo Surabaya")
	// case "W":
	// 	fmt.Println("Halo Sidoarjo")
	// case "Z":
	// 	fmt.Println("Halo Sumedang")
	// default:
	// 	fmt.Println("Maaf plat kamu gaada di database")
	// }

}

func greetUsers() {
	fmt.Println("======== PBKK GOLANG PROGRAMMING PROGRESS ========")
	fmt.Printf("Selamat datang ke Aplikasi %v booking system\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTicket, remainingTicket)
	fmt.Printf("Get Your Ticket Now!\n")
}

func getFirstName() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTicket uint

	fmt.Println("Enter Your First Name")
	fmt.Scan(&firstName)

	fmt.Println("Enter Your Last Name")
	fmt.Scan(&lastName)

	fmt.Println("Enter Your Email Address")
	fmt.Scan(&email)

	fmt.Println("Enter Number of Tickets:")
	fmt.Scan(&userTicket)

	return firstName, lastName, email, userTicket
}

func bookTicket(userTicket uint, firstName string, lastName string, email string) {
	remainingTicket -= userTicket

	var userData = UserData{
		firstName:  firstName,
		lastName:   lastName,
		email:      email,
		userTicket: userTicket,
	}

	bookings = append(bookings, userData)

	fmt.Printf("List of booking is %v\n", bookings)

	fmt.Printf("Thank You %v %v for Booking %v Tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTicket, email)
	fmt.Printf("%v Tickets Remaining at %v Booking Application\n", remainingTicket, conferenceName)
}

func sendTicket(userTicket uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var waktu = time.Now().Format("02/01/2006 15:04:05")
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTicket, firstName, lastName)
	fmt.Println("########################################################################################")
	fmt.Printf("Sending ticket:\n    %v to email address %v at %v\n", ticket, email, waktu)
	fmt.Println("########################################################################################")
	wg.Done()
}
