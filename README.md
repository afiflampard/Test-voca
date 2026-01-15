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
![Init Parking Lot](https://github.com/user-attachments/assets/28928135-6beb-41c1-9051-e3b7af1f668a)

### 2. Park a Car
![Park Car](https://github.com/user-attachments/assets/3f1e2408-e0c2-4e61-8891-5c7d4792de21)

### 3. Parking Lot Full
![Parking Full](https://github.com/user-attachments/assets/c0eb2b30-dae5-4ff2-a653-892df54ff5a0)

### 4. Status Check
![Status](https://github.com/user-attachments/assets/9cc1794a-9b89-473b-a8e4-60305e65e412)

### 5. Leave Parking Lot
![Leave](https://github.com/user-attachments/assets/4d94fc9d-ba85-475d-9c6f-5af684c560d6)

### 6. Status After Leaving
![Status After Leave](https://github.com/user-attachments/assets/8315110d-8443-4514-a866-1f0b5fa4fa87)

### 7. Re-allocate Slot
![Re-park](https://github.com/user-attachments/assets/15ff5fce-3160-4ead-9e3c-3692455bb18b)
