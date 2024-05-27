package service

import (
	"fmt"
	"ride-sharing/pkg/contract"
	"strconv"
	"strings"
)

type UserService interface {
	AddUser(userDetail string) error
	AddVehicle(vehicleDetail string) error
	PrintRideStats()
}

type userServiceImpl struct {
	users map[string]contract.User
}

func NewUserService(users map[string]contract.User) UserService {
	return &userServiceImpl{
		users: users,
	}
}

func (u *userServiceImpl) AddUser(userDetail string) error {
	parts := strings.Split(userDetail, ", ")
	if len(parts) != 3 {
		return fmt.Errorf("invalid user detail format. Expected format: 'Name, Gender, Age'")
	}

	name := parts[0]
	gender := parts[1]
	age, err := strconv.Atoi(parts[2])
	if err != nil {
		return fmt.Errorf("invalid age format: %v", err)
	}

	user := contract.User{
		Name:         name,
		Gender:       gender,
		Age:          age,
		Role:         contract.Passenger, // Default role as Passenger
		Vehicles:     make([]contract.Vehicle, 0),
		RidesTaken:   0,
		RidesOffered: 0,
	}
	u.users[name] = user

	//fmt.Printf("User added: %+v\n", user)
	return nil
}

func (u *userServiceImpl) AddVehicle(vehicleDetail string) error {
	parts := strings.Split(vehicleDetail, ", ")
	if len(parts) != 3 {
		return fmt.Errorf("invalid vehicle detail format. Expected format: 'Name, VehicleName, RegistrationNum'")
	}

	name := parts[0]
	vehicleName := parts[1]
	registrationNum := parts[2]

	if user, exists := u.users[name]; exists {
		vehicle := contract.Vehicle{
			Name:            vehicleName,
			RegistrationNum: registrationNum,
			Offered:         false,
		}
		user.Vehicles = append(user.Vehicles, vehicle)
		user.Role = contract.Driver
		u.users[name] = user
		//fmt.Printf("Vehicle added for user %s: %+v\n", name, vehicle)
	}
	return nil
}

func (u *userServiceImpl) PrintRideStats() {
	for _, details := range u.users {
		fmt.Printf("%s: %d Taken, %d Offered\n", details.Name, details.RidesTaken, details.RidesOffered)
	}
}
