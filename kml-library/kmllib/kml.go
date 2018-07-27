package kmllib

import ( "fmt"
         "io/ioutil"
         "encoding/xml"
         "strings"
         "os")

type LineString struct{
    XMLName xml.Name `xml:"LineString"`
    Coordinates string `xml:"coordinates"`
}

type Placemark struct{
    XMLName xml.Name `xml:"Placemark"`
    Name string `xml:"name"`
    Description string `xml:"description"`
    LineString LineString `xml:"LineString"`
}

type Document struct {
    XMLName xml.Name `xml:"Document"`
    Placemark Placemark `xml:"Placemark"`
}

type Kml struct{
    XMLName xml.Name `xml:"kml"`
    Document Document `xml:"Document"`
}


func ReadCoordinates(filename string)string{

     kmlFile,err := os.Open(filename)
      if err != nil {
                fmt.Println("Error opening file:", err)
             
        }
        defer kmlFile.Close()

        KMLdata, _ := ioutil.ReadAll(kmlFile)

       var kml Kml
        xml.Unmarshal(KMLdata,&kml)

        polylineCoordinatesStr := kml.Document.Placemark.LineString.Coordinates

       polylineCoordinates := strings.TrimSpace(polylineCoordinatesStr)
 
       return polylineCoordinates
}