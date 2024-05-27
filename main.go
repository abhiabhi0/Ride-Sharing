package main

import (
	"fmt"
	"ride-sharing/pkg/contract"
	"ride-sharing/pkg/service"
)

func main() {

	users := make(map[string]contract.User)
	offerRideDetail := make(map[string][]contract.Vehicle)

	userService := service.NewUserService(users)
	rideService := service.NewRideService(users, offerRideDetail)

	err := userService.AddUser("Rohan, M, 36")
	if err != nil {
		fmt.Printf("error adding user, err: %v\n", err)
	}
	err = userService.AddVehicle("Rohan, Swift, KA-01-12345")
	if err != nil {
		fmt.Printf("error adding vehicle, err: %v\n", err)
	}
	err = userService.AddUser("Shashank, M, 29")
	if err != nil {
		fmt.Printf("error adding user, err: %v\n", err)
	}
	err = userService.AddVehicle("Shashank, Baleno, TS-05-62395")
	if err != nil {
		fmt.Printf("error adding vehicle, err: %v\n", err)
	}
	err = userService.AddUser("Nandini, F, 29")
	if err != nil {
		fmt.Printf("error adding user, err: %v\n", err)
	}
	err = userService.AddUser("Shipra, F, 27")
	if err != nil {
		fmt.Printf("error adding user, err: %v\n", err)
	}
	err = userService.AddVehicle("Shipra, Polo, KA-05-41491")
	if err != nil {
		fmt.Printf("error adding vehicle, err: %v\n", err)
	}
	err = userService.AddVehicle("Shipra, Activa, KA-12-12332")
	if err != nil {
		fmt.Printf("error adding vehicle, err: %v\n", err)
	}
	err = userService.AddUser("Gaurav, M, 29")
	if err != nil {
		fmt.Printf("error adding user, err: %v\n", err)
	}
	err = userService.AddUser("Rahul, M, 35")
	if err != nil {
		fmt.Printf("error adding user, err: %v\n", err)
	}
	err = userService.AddVehicle("Rahul, XUV, KA-05-1234")
	if err != nil {
		fmt.Printf("error adding vehicle, err: %v\n", err)
	}

	err = rideService.OfferRide("Rohan, Origin=Hyderabad, Available Seats=1, Vehicle=Swift, KA-01-12345, Destination=Bangalore")
	if err != nil {
		fmt.Printf("error offering ride, err: %v\n", err)
	}
	err = rideService.OfferRide("Shipra, Origin=Bangalore, Available Seats=1, Vehicle=Activa, KA-12-12332, Destination=Mysore")
	if err != nil {
		fmt.Printf("error offering ride, err: %v\n", err)
	}
	err = rideService.OfferRide("Shipra, Origin=Bangalore, Available Seats=2, Vehicle=Polo, KA-05-41491, Destination=Mysore")
	if err != nil {
		fmt.Printf("error offering ride, err: %v\n", err)
	}
	err = rideService.OfferRide("Shashank, Origin=Hyderabad, Available Seats=2, Vehicle=Baleno, TS-05-62395, Destination=Bangalore")
	if err != nil {
		fmt.Printf("error offering ride, err: %v\n", err)
	}
	err = rideService.OfferRide("Rahul, Origin=Hyderabad, Available Seats=5, Vehicle=XUV, KA-05-1234, Destination=Bangalore")
	if err != nil {
		fmt.Printf("error offering ride, err: %v\n", err)
	}
	err = rideService.OfferRide("Rohan, Origin=Bangalore, Available Seats=1, Vehicle=Swift, KA-01-12345, Destination=Pune")
	if err != nil {
		fmt.Printf("error offering ride, err: %v\n", err)
	}

	ride, err := rideService.SelectRide("Nandini, Origin=Bangalore, Destination=Mysore, Seats=1, Most Vacant")
	if err != nil {
		fmt.Printf("error selecting ride, err: %v\n", err)
	}

	fmt.Printf("ride selected: %v\n", ride)

	ride, err = rideService.SelectRide("Gaurav, Origin=Bangalore, Destination=Mysore, Seats=1, Preferred Vehicle=Activa")
	if err != nil {
		fmt.Printf("error selecting ride, err: %v\n", err)
	}

	fmt.Printf("ride selected: %v\n", ride)

	ride, err = rideService.SelectRide("Shashank, Origin=Mumbai, Destination=Bangalore, Seats=1, Most Vacant")
	if err != nil {
		fmt.Printf("error selecting ride, err: %v\n", err)
	}

	fmt.Printf("ride selected: %v\n", ride)

	ride, err = rideService.SelectRide("Rohan, Origin=Hyderabad, Destination=Bangalore, Seats=1, Preferred Vehicle=Baleno")
	if err != nil {
		fmt.Printf("error selecting ride, err: %v\n", err)
	}

	fmt.Printf("ride selected: %v\n", ride)

	ride, err = rideService.SelectRide("Shashank, Origin=Hyderabad, Destination=Bangalore, Seats=1, Preferred Vehicle=Polo")
	if err != nil {
		fmt.Printf("error selecting ride, err: %v\n", err)
	}

	fmt.Printf("ride selected: %v\n", ride)

	userService.PrintRideStats()
}
