package main

import ("os"
		"io/ioutil"
		"fmt"
		"encoding/xml"
		"strings"
		"strconv"
		"math"
		)


type LineString struct{
	XMLName xml.Name `xml:"LineString"`
	Coordinates string `xml:"coordinates"`
}

type Placemark struct{
	XMLName xml.Name `xml:"Placemark"`
	name string `xml:"name"`
	description string `xml:"description"`
	LineString LineString `xml:"LineString"`
}
type Document struct{
	XMLName xml.Name `xml:"Document"`
	name string `xml:"name"`
	Placemark Placemark `xml:"Placemark"`

}

type Kml struct{
	XMLName xml.Name `xml:"kml"`
	Document Document `xml:"Document"`
} 

const (
	earthRadiusMi = 3958 // radius of the earth in miles.
	earthRaidusKm = 6371 // radius of the earth in kilometers.
)


func main(){
 
 if len(os.Args) < 2 {
 	fmt.Println("missing parameter!,please provide file name")
 	return
 }

  kmlFile, err := ioutil.ReadFile(os.Args[1])
    if err != nil {
       fmt.Println("Can't read file:", os.Args[1])
        panic(err)
    }

       var kml Kml
        
       xml.Unmarshal(kmlFile,&kml)

       polyCoordinatesStr := strings.TrimSpace(kml.Document.Placemark.LineString.Coordinates)
       
       if polyCoordinatesStr != "" {

       	polyCoordinates := strings.Split(polyCoordinatesStr," ")

       var xcors,ycors []string
	   for i:=0;i<len(polyCoordinates);i++{

	   		points := strings.Split(polyCoordinates[i],",")
	   		xcor := points[0]
	   		ycor := points[1]

	   		xcors = append(xcors,xcor)
	   		ycors = append(ycors,ycor)

	   		
	   }

	  
	   for j:=0;j<len(polyCoordinates)-1;j++{

	   		lat1,_ := strconv.ParseFloat(xcors[j],64)
	   		lon1,_ := strconv.ParseFloat(ycors[j],64)
	   		lat2,_ := strconv.ParseFloat(xcors[j+1],64)
	   		lon2,_ := strconv.ParseFloat(ycors[j+1],64)

	   		TotalDistInMiles,TotalDistInKM := Distance(lat1,lon1,lat2,lon2)

	   		fmt.Println("Distance in miles:",TotalDistInMiles)
	   		fmt.Println("Distance in km:",TotalDistInKM)
	   } 	
   
       }else{

       	fmt.Println("no points")
       	return 
       }

}


// degreesToRadians converts from degrees to radians.
func degreesToRadians(d float64) float64 {
	return d * math.Pi / 180
}

// Distance calculates the shortest path between two coordinates on the surface
// of the Earth. This function returns two units of measure, the first is the
// distance in miles, the second is the distance in kilometers.
func Distance(Lat1,Lon1,Lat2,Lon2 float64) (mi, km float64) {

	lat1 := degreesToRadians(Lat1)
	lon1 := degreesToRadians(Lon1)
	lat2 := degreesToRadians(Lat2)
	lon2 := degreesToRadians(Lon2)

	diffLat := lat2 - lat1
	diffLon := lon2 - lon1

	a := math.Pow(math.Sin(diffLat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(diffLon/2), 2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	mi = c * earthRadiusMi
	km = c * earthRaidusKm

	return mi, km
}






