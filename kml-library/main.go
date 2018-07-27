 package main

 import ("github.com/niketa/gorepo/kml-library/kmllib"
			"fmt")

 
 func main() {

        coordinates := kmllib.ReadCoordinates("./polylines.kml")

        fmt.Println("coordinates:",coordinates)
        
    }

