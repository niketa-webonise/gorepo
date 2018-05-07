 package main

 import (
         "fmt"
         "io/ioutil"
         "encoding/xml"
         "strings"
         "math"
         "strconv"
         "os"
         
 )

 type LineString struct {
        XMLName xml.Name `xml:"LineString"`
        Coordinates string `xml:"coordinates"`
 }

 type Placemark struct {
        XMLName xml.Name `xml:"Placemark"`
        Name string `xml:"name"`
        Description string `xml:"description"`
        LineString LineString `xml:"LineString"`
        StyleUrl string `xml:"styleUrl"`
 }

 type Document struct {
        XMLName xml.Name `xml:"Document"`
        Name string `xml:"name"`
        Placemark Placemark `xml:"Placemark"`
 }

 type Kml struct {
        XMLName xml.Name `xml:"kml"`
        Document Document `xml:"Document"`
 }


 func main() {

     if len(os.Args) < 2 {
        fmt.Println("Missing parameter, provide file name!")
        return
    }

    kmlFile, err := ioutil.ReadFile(os.Args[1])
    if err != nil {
       fmt.Println("Can't read file:", os.Args[1])
        panic(err)
    }

       var kml Kml
        xml.Unmarshal(kmlFile,&kml)

        polylineCoordinatesStr := kml.Document.Placemark.LineString.Coordinates

       polylineCoordinates := strings.TrimSpace(polylineCoordinatesStr)
     
        if polylineCoordinates != "" {

            points := strings.Split(polylineCoordinates," ")

           	var xcors,ycors []string
        
            for i:=0;i<len(points);i++ {


                pointCors := strings.Split(points[i],",")
                xcor := pointCors[0]
                ycor := pointCors[1]

                xcors = append(xcors,xcor)
                ycors = append(ycors,ycor)
            }
      
            totalDist := 0.0
            for j:=0;j<len(points)-1;j++ {
                
                lat1, _ := strconv.ParseFloat(xcors[j],64)
                lon1, _ := strconv.ParseFloat(ycors[j],64)
                lat2, _ := strconv.ParseFloat(xcors[j+1],64) 
                lon2, _ := strconv.ParseFloat(ycors[j+1],64)
                
                totalDist = Distance(lon1, lat1, lon2, lat2)
            	fmt.Println(totalDist/1000)
            }

        } else {
            fmt.Println("No points!")
            return
        }

}


const (
    earthRadiusMi = 3958 // radius of the earth in miles.
    earthRaidusKm = 6371 // radius of the earth in kilometers.
)

// Coord represents a geographic coordinate.
type Coord struct {
    Lat float64
    Lon float64
}

// degreesToRadians converts from degrees to radians.
func degreesToRadians(d float64) float64 {
    return d * math.Pi / 180
}

// Distance calculates the shortest path between two coordinates on the surface
// of the Earth. This function returns two units of measure, the first is the
// distance in miles, the second is the distance in kilometers.
func Distance(p, q Coord) (mi, km float64) {
    lat1 := degreesToRadians(p.Lat)
    lon1 := degreesToRadians(p.Lon)
    lat2 := degreesToRadians(q.Lat)
    lon2 := degreesToRadians(q.Lon)

    diffLat := lat2 - lat1
    diffLon := lon2 - lon1

    a := math.Pow(math.Sin(diffLat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*
        math.Pow(math.Sin(diffLon/2), 2)

    c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

    mi = c * earthRadiusMi
    km = c * earthRaidusKm

    return mi, km
}