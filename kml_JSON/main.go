package main

import ("fmt"
        "os"
        "io/ioutil"
        "encoding/xml"
        "encoding/json")

type LineString struct {

    Tessellate string `xml:"tessellate"`
    Coordinates string `xml:"coordinates"`
}

type Point struct {

    Coordinates string `xml:"coordinates"`
    GdrawOrder string `xml:"http://www.google.com/kml/ext/2.2 drawOrder"`
}

type LookAt struct {

    Longitude string `xml:"longitude"`
    Latitude string `xml:"latitude"`
    Altitude string `xml:"altitude"`
    Heading string `xml:"heading"`
    Tilt string `xml:"tilt"`
    Range string `xml:"range"`
    GaltitudeMode string `xml:"http://www.google.com/kml/ext/2.2 altitudeMode"`
}

type Placemark struct {

    Name string `xml:"name"`
    Description string `xml:"description"`
    StyleUrl string `xml:"styleUrl"`
    Point Point `xml:"Point"`
    LineString LineString `xml:"LineString"`
    LookAt LookAt `xml:"LookAt"`

}


type Folder struct{

    Name string `xml:"name"`
    Description string `xml:"description"`
    Placemarks []Placemark `xml:"Placemark"`
}


type Pair struct{

   Key string `xml:"key"`
   StyleUrl string `xml:"styleUrl"`
}

type StyleMap struct{

   ID string `xml:"id,attr"`
   Pairs []Pair `xml:"Pair"`
}

type LineStyle struct{

   Color string `xml:"color"`

}

type HotSpot struct{

    X string `xml:"x,attr"`
    Y string `xml:"y,attr"`
    Xunits string `xml:"xunits,attr"`
    Yunits string `xml:"yunits,attr"`
}

type Icon struct {

   Href string `xml:"href"`
}

type IconStyle struct{

   Scale string `xml:"scale"`
   Icon Icon `xml:"Icon"`
   HotSpot HotSpot `xml:"hotSpot"`

}

type Style struct {
   
   ID string `xml:"id,attr"`
   IconStyle IconStyle `xml:"IconStyle"`
   LineStyle LineStyle `xml:"LineStyle"`
}

type Document struct {
   
   Name string `xml:"name"`
   Open string `xml:"open"`
   Styles []Style `xml:"Style"`
   StyleMaps []StyleMap `xml:"StyleMap"` 
   Folders []Folder `xml:"Folder"`
   Placemarks []Placemark `xml:"Placemark"`
}


type Kml struct{

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
       jsonData, err := json.Marshal(kml)

       if err != nil {

        fmt.Println(err)
        os.Exit(1)
       }

     
        jsonFile, err := os.Create("./kmlJSON.json")

         if err != nil {
                 fmt.Println(err)
         }
         defer jsonFile.Close()

         jsonFile.Write(jsonData)
         jsonFile.Close()
}