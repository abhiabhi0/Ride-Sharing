package service

import (
	"reflect"
	"ride-sharing/pkg/contract"
	"testing"
)

func TestOfferRide(t *testing.T) {
	users := make(map[string]contract.User)
	offerRideDetail := make(map[string][]contract.Vehicle)
	userService := NewUserService(users)
	rideService := NewRideService(users, offerRideDetail)

	_ = userService.AddUser("Alice, F, 30")
	_ = userService.AddVehicle("Alice, Toyota, XYZ123")
	err := rideService.OfferRide("Alice, Origin=CityA, Available Seats=3, Vehicle=Toyota, XYZ123, Destination=CityB")
	if err != nil {
		t.Errorf("OfferRide failed: %v", err)
	}

	key := "CityA-CityB"
	if rides, exists := offerRideDetail[key]; !exists || len(rides) != 1 {
		t.Errorf("Ride was not offered correctly")
	}

	expectedRide := contract.Vehicle{
		Name:            "Toyota",
		RegistrationNum: "XYZ123",
		Ride: contract.Ride{
			Origin:      "CityA",
			Destination: "CityB",
		},
		AvailableSeats: 3,
		Offered:        true,
	}

	if !reflect.DeepEqual(offerRideDetail[key][0], expectedRide) {
		t.Errorf("Ride details do not match: got %+v, want %+v", offerRideDetail[key][0], expectedRide)
	}
}

func TestSelectRide(t *testing.T) {
	users := make(map[string]contract.User)
	offerRideDetail := make(map[string][]contract.Vehicle)
	userService := NewUserService(users)
	rideService := NewRideService(users, offerRideDetail)

	_ = userService.AddUser("Alice, F, 30")
	_ = userService.AddVehicle("Alice, Toyota, XYZ123")
	_ = rideService.OfferRide("Alice, Origin=CityA, Available Seats=3, Vehicle=Toyota, XYZ123, Destination=CityB")

	_ = userService.AddUser("Bob, Male, 25")
	ride, err := rideService.SelectRide("Bob, Origin=CityA, Destination=CityB, Seats=1, Most Vacant")
	if err != nil {
		t.Errorf("SelectRide failed: %v", err)
	}

	expectedRide := contract.Vehicle{
		Name:            "Toyota",
		RegistrationNum: "XYZ123",
		Ride: contract.Ride{
			Origin:      "CityA",
			Destination: "CityB",
		},
		AvailableSeats: 3,
		Offered:        true,
	}

	if !reflect.DeepEqual(ride, expectedRide) {
		t.Errorf("Selected ride does not match: got %+v, want %+v", ride, expectedRide)
	}

	user, exists := users["Bob"]
	if !exists {
		t.Errorf("User Bob was not found")
	}

	if user.RidesTaken != 1 {
		t.Errorf("User rides taken should be 1: got %v", user.RidesTaken)
	}
}
