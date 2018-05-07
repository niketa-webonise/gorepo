 package main

 import (
         "fmt"
         "io/ioutil"
         "os"
         "encoding/xml"
         "strings"
         "math"
         "strconv"
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

 type PolyStyle struct {
        XMLName xml.Name `xml:"PolyStyle"`
        Color string `xml:"color"`
 }

 type LineStyle struct {
        XMLName xml.Name `xml:"LineStyle"`
        Width string `xml:"width"`
 }

 type Style struct {
        XMLName xml.Name `xml:"Style"`
        Id int `xml:"id,attr"`
        PolyStyle PolyStyle `xml:"PolyStyle"`
        LineStyle LineStyle `xml:"LineStyle"`
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
        xmlFile, err := os.Open("./polylines.kml")
        if err != nil {
                fmt.Println("Error opening file:", err)
                return
        }
        defer xmlFile.Close()

        XMLdata, _ := ioutil.ReadAll(xmlFile)

        var kml Kml
        xml.Unmarshal(XMLdata, &kml)


        polylineCoordinatesStr := strings.TrimSpace(kml.Document.Placemark.LineString.Coordinates)
        //fmt.Println(polylineCoordinatesStr)
        coordinates := strings.Split(polylineCoordinatesStr, " ")
        //fmt.Println(coordinates)

        
        //convert to float64
        lat1, err := strconv.ParseFloat(x1, 64) 
        lon1, err := strconv.ParseFloat(y1, 64)
        lat2, err := strconv.ParseFloat(x2, 64) 
        lon2, err := strconv.ParseFloat(y2, 64)  

        dist := Distance(lat1,lon1,lat2,lon2)

        distInKML := dist/1000

        fmt.Printf("Distance between two polylines coordinates: %.3f"+"km ",distInKML)

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
