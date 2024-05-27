package service

import (
	"reflect"
	"ride-sharing/pkg/contract"
	"testing"
)

func TestAddUser(t *testing.T) {
	users := make(map[string]contract.User)
	userService := NewUserService(users)

	err := userService.AddUser("Alice, F, 30")
	if err != nil {
		t.Errorf("AddUser failed: %v", err)
	}

	user, exists := users["Alice"]
	if !exists {
		t.Errorf("User Alice was not added")
	}

	expectedUser := contract.User{
		Name:         "Alice",
		Gender:       "F",
		Age:          30,
		Role:         contract.Passenger,
		Vehicles:     make([]contract.Vehicle, 0),
		RidesTaken:   0,
		RidesOffered: 0,
	}

	if !reflect.DeepEqual(user, expectedUser) {
		t.Errorf("User details do not match: got %+v, want %+v", user, expectedUser)
	}
}

func TestAddVehicle(t *testing.T) {
	users := make(map[string]contract.User)
	userService := NewUserService(users)

	_ = userService.AddUser("Alice, F, 30")
	err := userService.AddVehicle("Alice, Toyota, XYZ123")
	if err != nil {
		t.Errorf("AddVehicle failed: %v", err)
	}

	user, exists := users["Alice"]
	if !exists {
		t.Errorf("User Alice was not found")
	}

	expectedVehicle := contract.Vehicle{
		Name:            "Toyota",
		RegistrationNum: "XYZ123",
		Offered:         false,
	}

	if len(user.Vehicles) != 1 || !reflect.DeepEqual(user.Vehicles[0], expectedVehicle) {
		t.Errorf("Vehicle details do not match: got %+v, want %+v", user.Vehicles, []contract.Vehicle{expectedVehicle})
	}

	if user.Role != contract.Driver {
		t.Errorf("User role should be Driver: got %v", user.Role)
	}
}
