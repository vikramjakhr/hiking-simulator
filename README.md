# hiking-simulator

Hiking simulator for finding shortest time for crossing a barrier

## How to use it:

* See usage with:

```
hiking-simulator --help
```

* Run test cases and code coverage (go > 1.12)
```
go test --cover ./...
```

* Build the binary
```
go build
```

## How to run it:

* [Download the binary from here](https://github.com/vikramjakhr/hiking-simulator/releases/download/v1.0.0/hiking-simulator)
* [Download the input yaml from here](https://github.com/vikramjakhr/hiking-simulator/releases/download/v1.0.0/input.yaml)

Place the above downloaded files in any directory and run below command from that directory
```
./hiking-simulator 
```

* Run as a container
```
docker pull vikramjakhr/mine:hiking-simulator
docker run -it vikramjakhr/mine:hiking-simulator
```
