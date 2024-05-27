package service

import (
	"fmt"
	"ride-sharing/pkg/contract"
	"strconv"
	"strings"
)

type RideService interface {
	OfferRide(rideDetail string) error
	SelectRide(rideDetail string) (contract.Vehicle, error)
}

type rideServiceImpl struct {
	users           map[string]contract.User
	offerRideDetail map[string][]contract.Vehicle
}

func NewRideService(users map[string]contract.User, offerRideDetail map[string][]contract.Vehicle) RideService {
	return &rideServiceImpl{
		users:           users,
		offerRideDetail: offerRideDetail,
	}
}

func (r *rideServiceImpl) OfferRide(rideDetail string) error {
	parts := strings.Split(rideDetail, ", ")
	if len(parts) != 6 {
		return fmt.Errorf("invalid ride detail format. Expected format: 'Name, Origin, Available Seats, Vehicle Name, Registration Num, Destination'")
	}

	name := parts[0]
	origin := strings.Split(parts[1], "=")[1]
	availableSeatsStr := strings.Split(parts[2], "=")[1]
	vehicleName := strings.Split(parts[3], "=")[1]
	registrationNum := strings.Split(parts[4], ", ")[0]
	destination := strings.Split(parts[5], "=")[1]

	availableSeats, err := strconv.Atoi(availableSeatsStr)
	if err != nil {
		return fmt.Errorf("invalid available seats format: %v", err)
	}

	if user, exists := r.users[name]; exists {
		for i := range user.Vehicles {
			if user.Vehicles[i].Name == vehicleName && user.Vehicles[i].RegistrationNum == registrationNum {
				if !user.Vehicles[i].Offered {
					ride := contract.Ride{
						Origin:      origin,
						Destination: destination,
					}
					user.Vehicles[i].Offered = true
					user.Vehicles[i].AvailableSeats = availableSeats
					user.Vehicles[i].Ride = ride

					key := origin + "-" + destination
					if rides, exist := r.offerRideDetail[key]; exist {
						rides = append(rides, user.Vehicles[i])
						r.offerRideDetail[key] = rides
					} else {
						rides = make([]contract.Vehicle, 0)
						rides = append(rides, user.Vehicles[i])
						r.offerRideDetail[key] = rides
					}
					user.RidesOffered += 1

				} else {
					return fmt.Errorf("already offered: %v", rideDetail)
				}
			}
		}
		r.users[name] = user
	}
	return nil
}

func (r *rideServiceImpl) SelectRide(rideDetail string) (contract.Vehicle, error) {
	var selectedRide contract.Vehicle

	parts := strings.Split(rideDetail, ", ")
	if len(parts) != 5 {
		return selectedRide, fmt.Errorf("invalid ride detail format. Expected format: 'Name, Origin, Destination, Seats, Selection Strategy'")
	}

	name := parts[0]
	origin := strings.Split(parts[1], "=")[1]
	destination := strings.Split(parts[2], "=")[1]
	seatsStr := strings.Split(parts[3], "=")[1]
	selectionStrategy := parts[4]

	seats, err := strconv.Atoi(seatsStr)
	if err != nil {
		return selectedRide, fmt.Errorf("invalid seats format: %v", err)
	}
	foundRide := false

	if selectionStrategy == "Most Vacant" {
		maxSeats := -1
		key := origin + "-" + destination

		rides := r.offerRideDetail[key]
		for _, r := range rides {
			if r.AvailableSeats >= seats && r.AvailableSeats > maxSeats {
				maxSeats = r.AvailableSeats
				selectedRide = r
				foundRide = true
			}
		}
	} else if strings.Contains(selectionStrategy, "Preferred Vehicle") {
		preferredVehicle := strings.Split(selectionStrategy, "=")[1]
		key := origin + "-" + destination

		rides := r.offerRideDetail[key]
		for _, r := range rides {
			if r.AvailableSeats >= seats && r.Name == preferredVehicle {
				selectedRide = r
				foundRide = true
				break
			}
		}
	}

	if user, exists := r.users[name]; exists {
		user.RidesTaken += 1
		r.users[name] = user
	} else {
		return selectedRide, fmt.Errorf("invalid user")
	}

	if !foundRide {
		return selectedRide, fmt.Errorf("no rides found")
	}

	fmt.Printf("selected ride is %v\n", selectedRide)
	return selectedRide, nil
}
