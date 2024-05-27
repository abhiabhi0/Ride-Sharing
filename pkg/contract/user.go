package contract

type UserRole string

const (
	Driver    UserRole = "driver"
	Passenger UserRole = "passenger"
)

type User struct {
	Name         string    `json:"name"`
	Gender       string    `json:"gender"`
	Age          int       `json:"age"`
	Role         UserRole  `json:"role"`
	Vehicles     []Vehicle `json:"vehicle"`
	RidesTaken   int       `json:"ridesTaken"`
	RidesOffered int       `json:"ridesOffered"`
}

type Vehicle struct {
	Name            string `json:"name"`
	RegistrationNum string `json:"registrationNum"`
	Ride            Ride   `json:"ride"`
	AvailableSeats  int    `json:"availableSeats"`
	Offered         bool   `json:"offered"`
}

type Ride struct {
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
}
