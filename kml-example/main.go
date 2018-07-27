package main

import  (
         "fmt"
         "io/ioutil"
         "os"
         "encoding/xml"
          )

type Point struct {
	xmlName xml.Name `xml:"Point"`
	Coordinates string `xml:"coordinates"`
}

type Placemark struct {
	XMLName xml.Name `xml:"Placemark"`
	name string `xml:"name"`
	Point Point `xml:"Point"`
}
type Document struct {

	XMLName xml.Name `xml:"Document"`
	Placemark Placemark `xml:"Placemark"`
}
type kml struct {

	XMLName xml.Name `xml:"kml"`
	Document Document `xml:"Document"`
}

func main(){

	 xmlFile, err := os.Open("./tennis-lines.kml")
        if err != nil {
                fmt.Println("Error opening file:", err)
                return
        }
        defer xmlFile.Close()

        XMLdata, _ := ioutil.ReadAll(xmlFile)

        var k kml
        xml.Unmarshal(XMLdata, &k)

        fmt.Println(k.Document.Placemark.Point.Coordinates)
}