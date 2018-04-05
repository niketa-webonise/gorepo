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

        polylineCoordinatesStr := kml.Document.Placemark.LineString.Coordinates

        fmt.Println(polylineCoordinatesStr)

        poly_split := strings.Split(polylineCoordinatesStr,",")

        x1 := poly_split[0]
        y1 := poly_split[1]

        x2 := poly_split[3]
        y2 := poly_split[4]

        x3 := poly_split[6]
        y3 := poly_split[7]

     
    //convert to float64
        lat1, err := strconv.ParseFloat(x1, 64) 
        lon1, err := strconv.ParseFloat(y1, 64)
        lat2, err := strconv.ParseFloat(x2, 64) 
        lon2, err := strconv.ParseFloat(y2, 64)  
        lat3, err := strconv.ParseFloat(x3,64)
        lon3, err := strconv.ParseFloat(y3,64)

        dist1 := Distance(lat1,lon1,lat2,lon2)
        dist2 := Distance(lat2,lon2,lat3,lon3)
        dist3 := Distance(lat1,lon1,lat3,lon3)

        dist_one := dist1/1000
        dist_two := dist2/1000
        dist_three := dist3/1000

        fmt.Printf("\nDistance between x1,y1 and x2,y2 two polylines coordinates: %.3f"+"km ",dist_one)
        fmt.Printf("\nDistance between x2,y2 and x3,y3 two polylines coordinates: %.3f"+"km ",dist_two)
        fmt.Printf("\nDistance between x1,y1 and x3,y3 two polylines coordinates: %.3f"+"km ",dist_three)

    }


    // haversin(θ) function
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
