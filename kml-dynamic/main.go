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

func hsin(theta float64) float64 {
        return math.Pow(math.Sin(theta/2), 2)
}

// Distance function returns the distance (in meters) between two points of
//     a given longitude and latitude relatively accurately (using a spherical
//     approximation of the Earth) through the Haversin Distance Formula for
//     great arc distance on a sphere with accuracy for small distances
//
// point coordinates are supplied in degrees and converted into rad. in the func
//
// distance returned is METERS!!!!!!
// http://en.wikipedia.org/wiki/Haversine_formula
func Distance(lat1, lon1, lat2, lon2 float64) float64 {
        // convert to radians
        // must cast radius as float to multiply later
        var la1, lo1, la2, lo2, r float64
        la1 = lat1 * math.Pi / 180
        lo1 = lon1 * math.Pi / 180
        la2 = lat2 * math.Pi / 180
        lo2 = lon2 * math.Pi / 180

        r = 6378100 // Earth radius in METERS

        // calculate
        h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

        return 2 * r * math.Asin(math.Sqrt(h))
}