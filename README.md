# Parking Lot Management System

## Overview
This project implements a Parking Lot Management System in Golang that allows customers to use an automated parking system. It supports basic functionalities like allocating parking slots, releasing them, and calculating parking charges.

## Features
- **Create Parking Lot:** Initialize the parking lot with a specified number of slots.
- **Park:** Allocate the nearest available parking slot to a car.
- **Leave:** Free up a parking slot using the car's registration number and calculate the parking charge based on the hours parked.
- **Status:** Display the current status of the parking lot.
- **Case-insensitive Commands:** Accept commands in any case (e.g., `PARK`, `park`, `Park`).
- **Command File Input:** Read and process commands from a file.

## Commands
The system supports the following commands:

### 1. `create_parking_lot <n>`
Creates a parking lot with `n` slots.
- Example: `create_parking_lot 6`

### 2. `park <registration_number>`
Parks a car with the given registration number and color in the nearest available slot.
- Example: `park KA-01-HH-1234`

### 3. `leave <registration_number> <hours_parked>`
Frees up the slot occupied by the car with the specified registration number and calculates the parking charge.
- Parking charges:
  - $10 for the first 2 hours.
  - $10 for each additional hour.
- Example: `leave KA-01-HH-1234 3`

### 4. `status`
Displays the current parking lot status, including slot number, registration number, and color of parked cars.
- Example: `status`

## How to Use

### Prerequisites
- Install Golang (https://golang.org/dl/).

### Setup
1. Clone the repository or copy the code to a file named `parking_lot.go`.
2. Create a commands file (e.g., `commands.txt`) with the desired commands.

### Running the Program
1. Compile and run the program with the commands file as an argument:
   ```bash
   go run parking_lot.go commands.txt
   ```

### Example Input (`commands.txt`):
```text
create_parking_lot 6
park KA-01-HH-1234
park KA-01-HH-9999
leave KA-01-HH-1234 3
status
```

### Expected Output:
```text
Created a parking lot with 6 slots
Allocated slot number: 1
Allocated slot number: 2
Slot number 1 is free. Charge: $20
Slot No.    Registration No
2           KA-01-HH-9999
```

## Project Structure
- **`parking_lot.go`**: Main program file containing the implementation.
- **`commands.txt`**: Example file with input commands.

## Limitations
- The parking charges are calculated only for whole hours.
- The system does not persist data; all data is reset when the program exits.

## Future Enhancements
- Add support for fractional hour charges.
- Implement persistence for parking lot data.
- Add a web or mobile interface for user interaction.

## License
This project is licensed under the MIT License.

