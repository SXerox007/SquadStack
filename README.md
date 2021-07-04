# SquadStack
Parking Lot Problem

## Install Docker to run through Docker
```
//For Mac
https://docs.docker.com/docker-for-mac/install/

//For Linux
https://docs.docker.com/engine/install/ubuntu/

//For Windows
https://docs.docker.com/docker-for-windows/install/
```

## Install golang if run manual (optional). Select OS which you are using.
``` https://golang.org/doc/install ```

## Ways to run code (as need)
### 1. Direct run ```sh run.sh``` file 
This contains docker commands. Nothing needs to run anything. If have Docker installed.

### 2. Run through Docker commands
```
//build docker file
docker build -t sellerapp  -f Dockerfile .
// then
docker-compose build && docker-compose up
```

### 3. Run from `Make` file. If you installed golang 
```
make app
```
It run the `go run` command

### 4. Run Direct Binary file 
```
./squadstack
```

## Go Server:
```
server run on localhost:50051
```

## Added `sample.txt` file in codebase that contains your sample file 

### Run API CURL for Default file run i.e sample.txt
Request:
```
curl -X POST -F 'type=default' http://:50051/squadstack/v1/file/parking-lot
```
Response:
```
{
    "success": true,
    "message": "Success",
    "code": 200,
    "data": [
        "Created parking of 6 slots",
        "Car with vehicle registration number KA-01-HH-1234 has been parked at slot number 1",
        "Car with vehicle registration number PB-01-HH-1234 has been parked at slot number 2",
        "1,2",
        "No parked car matches the query",
        "Car with vehicle registration number PB-01-TG-2341 has been parked at slot number 3",
        "Car with vehicle registration number PB-01-HH-1234 has slot number 2",
        "Slot number 2 vacated, the car with vehicle registration number PB-01-HH-1234 left the space, the driver of the car was of age 21",
        "Car with vehicle registration number HR-29-TG-3098 has been parked at slot number 2",
        "No parked car matches the query",
        "No parked car matches the query"
    ]
}
```

### Run API CURL for Input file
Request:
``` 
curl -X POST http://:50051/squadstack/v1/file/parking-lot --form 'file=@"<file-path>"' 
```
Request Example:
``` 
curl -X POST http://:50051/squadstack/v1/file/parking-lot --form 'file=@"/Users/sumitthakur/go/src/SquadStack/sample.txt"' 
```

Result: 
```
You will get the result in response and as well as terminal also.
```


## Need Help:
Email: sumitthakur769@gmail.com 






