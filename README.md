# Parking Lot System - VocaGame Test

A simple command-line interface (CLI) application written in Go to manage a parking lot system. This tool allows you to create a parking lot, park cars, handle departures with fee calculation, and view current parking status.

## Prerequisites

- [Go (Golang)](https://go.dev/dl/) installed on your machine.

## How to Run

1. Clone the repository or navigate to the project directory.
2. Run the application using the following command:

```bash
go run main.go
```

## Usage Guide

Once the application is running, you will see a menu with the following commands. Type the command name to proceed.

### Available Commands

1.  **`create_parking_lot`**
    -   Initializes the parking lot with a specific number of slots.
    -   **Input**: Capacity (integer).
    -   **Example**: `6`

2.  **`park`**
    -   Allocates the nearest available slot to a car.
    -   **Input**: Car Number (string).
    -   **Charge**: Default parking charge is $10 for the first 2 hours.

3.  **`leave`**
    -   Removes a car from the parking lot and calculates the total fee.
    -   **Input**:
        -   Car Number (string).
        -   Hours Parked (integer).
    -   **Calculation**: $10 for the first 2 hours. Additional hours are charged at $10 per hour.

4.  **`status`**
    -   Displays the current status of all parking slots.
    -   Shows Slot Number, Status (park/leave), Car Number, and Hours parked.

5.  **`exit`**
    -   Exits the application.

## Screenshots

### 1. Initialize Parking Lot
<img width="270" height="172" alt="1" src="https://github.com/user-attachments/assets/5e33e093-7fb9-4e52-8de0-38ac6bb8f90b" />


### 2. Park a Car
<img width="221" height="344" alt="2" src="https://github.com/user-attachments/assets/e75e8907-996d-4555-b39d-262512b102fa" />


### 3. Parking Lot Full
<img width="257" height="169" alt="3" src="https://github.com/user-attachments/assets/ddafe768-0f2e-47e7-a7a4-dcbb272815de" />


### 4. Status Check
<img width="199" height="180" alt="4" src="https://github.com/user-attachments/assets/19a02a95-06ef-43cf-82ca-ccbc0a5aa135" />


### 5. Leave Parking Lot
<img width="532" height="235" alt="5" src="https://github.com/user-attachments/assets/86c66049-e6a8-4685-9fe4-245ef26994c3" />


### 6. Status After Leaving
<img width="168" height="171" alt="6" src="https://github.com/user-attachments/assets/6a9c9796-c1d4-4240-a509-59bd8ade2150" />


### 7. Re-allocate Slot
<img width="207" height="244" alt="7" src="https://github.com/user-attachments/assets/0caf339b-7de8-449e-b557-ddbab3747009" />

