package main

import "fmt"

func main() {
	var appName = "BOOKING APP"
	const totalTickets = 100
	var remainingTickets uint = 100
	fmt.Println("GOLANG ", appName)
	fmt.Println("A simple Go application")
	fmt.Printf("We have a total of %v  Tickets and %v remaining. \n", totalTickets, remainingTickets)

	var (
		fName string
		sName string
		noTickets uint
		email string
	)

	fmt.Println("Hi there, Enter you First Name :")
	fmt.Scan(&fName)

	fmt.Printf("Nice %v, Enter you Second Name :\n",fName)
	fmt.Scan(&sName)

        fmt.Println("How many tickets do you want ? :")
	fmt.Scan(&noTickets)

        fmt.Println("Lastly, Enter an Email Address :")
	fmt.Scan(&email)

	fmt.Printf("Thank You %v %v, you have succesfully booked %v tickets \n", fName, sName, noTickets)

	//Calculate the remaining tickets
	newTickets := remainingTickets - noTickets

	fmt.Printf("There are now %v tickets left in the sytem", newTickets)


	

}
