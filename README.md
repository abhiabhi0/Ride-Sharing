# Ride-Sharing

Problem Statement
Develop a ride-sharing applica3on with the following features.
Features:
  • User Roles: Users can either offer a shared ride (Driver) or consume a shared ride (Passenger).
  • Ride Selec3on: Users can search and select from mul3ple available rides on a route with the same source and des3na3on.

Requirements:
  • User Management:
    Onboarding: Implement func3onality to add user details `add_user(user_detail)`
  
  • Vehicle Management:
    Implement func3onality to add vehicle details for users. `add_vehicle(vehicle_detail)`

    ```
    add_user("Rohan, M, 36"); add_vehicle("Rohan, Swift, KA-01-12345")
    add_user("Shashank, M, 29"); add_vehicle("Shashank, Baleno, TS-05-62395")
    add_user("Nandini, F, 29")
    add_user("Shipra, F, 27"); add_vehicle("Shipra, Polo, KA-05-41491"); add_vehicle("Shipra, Activa, KA-12-12332")
    add_user("Gaurav, M, 29")
    add_user("Rahul, M, 35"); add_vehicle("Rahul, XUV, KA-05-1234")
    ```

  • Ride Offering:
    Offer Ride: Allow users to offer a shared ride on a route with specific details. `offer_ride(ride_detail)`

    ```
    offer_ride("Rohan, Origin=Hyderabad, Available Seats=1, Vehicle-Swift, KA-01-12345, Destination=Bangalore")
    offer_ride("Shipra, Origin-Bangalore, Available Seats-1, Vehicle-Activa, KA-12-12332, Destination-Mysore")
    offer_ride("Shipra, Origin=Bangalore, Available Seats=2, Vehicle-Polo, KA-05-41491, Destination-Mysore")
    offer_ride("Shashank, Origin-Hyderabad, Available Seats=2, Vehicle-Baleno, TS-05-62395, Destination-Bangalore")
    offer_ride("Rahul, Origin Hyderabad, Available Seats=5, Vehicle-XUV, KA-05-1234, Destination-Bangalore")
    offer_ride("Rohan, Origin=Bangalore, Available Seats=1, Vehicle-Swift, KA-01-12345, Destination=Pune") (This call should fail since a ride has already been offered by this user for this vehicle)
    ```

  • Ride Selection:
    Select Ride: Users can select a ride from mul3ple offered rides using a selection strategy based on preferred vehicle or most vacant seats `select_ride(source, destination, seats, selection_strategy)`
    ```
    select_ride("Nandini, Origin-Bangalore, Destination-Mysore, Seats=1, Most Vacant") (2(c) is the desired output)
    select_ride("Gaurav, Origin=Bangalore, Destination Mysore, Seats 1, Preferred Vehicle Activa") (2(b) is the desired output)
    select_ride("Shashank, Origin=Mumbai, Destination-Bangalore, Seats-1, Most Vacant") (No rides found)
    select_ride("Rohan, Origin=Hyderabad, Destination-Bangalore, Seats-1, Preferred Vehicle-Baleno") (2(d) is the desired output)
    select_ride("Shashank, Origin=Hyderabad, Destination-Bangalore, Seats=1, Preferred Vehicle Polo") (No rides found)
  ```

  • Statistics:
    Print Ride Stats: Retrieve and display total rides offered/taken by all users. `print_ride_stats()`
