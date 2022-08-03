package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers() //вызов функции
	//fmt.Printf("conferensTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceName, conferenceTickets, remainingTickets) смотрим типы переменных

	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)                                              //увлеичивает счетчик потоков, завершения которых должен дождаться основной поток
		go sendTicket(userTickets, firstName, lastName, email) //параллелизм

		firstNames := FirstName()
		fmt.Printf("The first names if bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Our conferense is booked out")
			// break
		}
	} else {
		if !isValidName {
			fmt.Printf("First name or last name you entered is too short\n")
		}
		if !isValidEmail {
			fmt.Printf("Your email address does not contain @ sign\n")
		}
		if !isValidTicketNumber {
			fmt.Printf("Number of tickets you entered is invalid\n")
		}

		//continue //переходим к следующей итерации цикла
	}
	wg.Wait() // Тут ждем завершения вторичных потоков
}

func greetUsers() {
	fmt.Printf("Welcome to %v conference\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func FirstName() []string {
	firstNames := []string{}
	for _, booking := range bookings { // Если переменная не нужна, то можно использовать пустой идентификатор Blank identifier _
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var userTickets uint
	var lastName string
	var email string
	/*
		Указатели & - это специальные переменные, которые хранят адрес другой переменной в памяти. Они представленны в C++ тоже
		fmt.Println(conferenceName)
		fmt.Println(&conferenceName)
	*/
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("How many tickets you want?")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	//create a map for a user

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. We'll send confirmation at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("################################")
	fmt.Printf("Sending ticket:\n%v \nto email addrress %v\n", ticket, email)
	fmt.Println("################################")
	wg.Done() // выдаем сообщение основного потока о завершении работы вторичного, счтечик уменьшится
}
