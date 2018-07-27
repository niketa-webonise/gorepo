package main

import ("fmt"
		"os"
        "io/ioutil"
        "encoding/xml"
       /* "strings"
        "math"
        "strconv"*/)


type kml struct{

    XMLName xml.Name `xml:"kml"`
    Document Document `xml:"Document"`
} 

type Document struct {
    
    XMLName xml.Name `xml:"Document"`
    Folder Folder `xml:"Folder"`
    //Placemarks []Placemark `xml:"Placemark"`
}

type Folder struct{

    XMLName xml.Name `xml:"Folder"`
    name string `xml:"name"`
    description string `xml:"description"`
    Folder []Placemark `xml:"Placemark"`
}

type Placemark struct {

    XMLName xml.Name `xml:"Placemark"`
    Name string `xml:"name"`
    Description string `xml:"description"`
    StyleUrl string `xml:"styleUrl"`
    LineString LineString `xml:"LineString"`
}

type LineString struct {

    XMLName xml.Name `xml:"LineString"`
    Coordinates string `xml:"coordinates"`
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

       var  placemarks Folder
        
       xml.Unmarshal(kmlFile,&placemarks)

       for i := 0;i<len(placemarks.Folder);i++ {
    
        fmt.Println("placemark name:"+ placemarks.Folder[i].Name)
        fmt.Println("placemark description:"+placemarks.Folder[i].StyleUrl)

       }
   }

    /*

       if len(os.Args) < 2 {
        fmt.Println("Missing parameter, provide file name!")
        return
    }

    kmlFile, err := ioutil.ReadFile(os.Args[1])
    if err != nil {
       fmt.Println("Can't read file:", os.Args[1])
        panic(err)
    }

       var kml kml
        
       xml.Unmarshal(kmlFile,&kml)


     placemark1CoordinatesStr := strings.TrimSpace(kml.Document.Placemarks[1].LineString.Coordinates)
     placemark2CoordinatesStr := strings.TrimSpace(kml.Document.Placemarks[2].LineString.Coordinates)
     placemark3CoordinatesStr := strings.TrimSpace(kml.Document.Placemarks[3].LineString.Coordinates)
    // placemark4CoordinatesStr := strings.TrimSpace(kml.Document.Placemarks[1].LineString.Coordinates)


     points1 := strings.Split(placemark1CoordinatesStr," ")
     points2 := strings.Split(placemark2CoordinatesStr," ")
     points3 := strings.Split(placemark3CoordinatesStr," ")
     //points4 := strings.Split(placemark4CoordinatesStr," ")
     
     var placemark1xcors,placemark1ycors []string

     for i:=0;i<len(points1);i++ {

       pointCors := strings.Split(points1[i],",")
             
             xcor := pointCors[0]
             ycor := pointCors[1]

             if(xcor >= "85.28520449768512") {
	              placemark1xcors = append(placemark1xcors,xcor)
	              placemark1ycors = append(placemark1ycors,ycor)
	      		}
     }


       placemark1TotalDist := 0.0
         for j:=0;j<len(placemark1xcors)-1;j++ {
             

             lon1, _ := strconv.ParseFloat(placemark1xcors[j],64)
             lat1, _ := strconv.ParseFloat(placemark1ycors[j],64)
             lon2, _ := strconv.ParseFloat(placemark1xcors[j+1],64) 
             lat2, _ := strconv.ParseFloat(placemark1ycors[j+1],64)
             
             placemark1TotalDist = placemark1TotalDist + Distance(lat1,lon1,lat2,lon2)
          
         
      }
       fmt.Printf("\nplacemark1TotalDist: %.3f"+"km",placemark1TotalDist/1000)

     var placemark2xcors,placemark2ycors []string

     for i:=0;i<len(points2);i++ {

       pointCors := strings.Split(points2[i],",")
             
             xcor := pointCors[0]
             ycor := pointCors[1]

             placemark2xcors = append(placemark2xcors,xcor)
             placemark2ycors = append(placemark2ycors,ycor)
     
     }

     placemark2TotalDist := 0.0 
         for j:=0;j<len(placemark2xcors)-1;j++ {
             

             lon1, _ := strconv.ParseFloat(placemark2xcors[j],64)
             lat1, _ := strconv.ParseFloat(placemark2ycors[j],64)
             lon2, _ := strconv.ParseFloat(placemark2xcors[j+1],64) 
             lat2, _ := strconv.ParseFloat(placemark2ycors[j+1],64)
             
             placemark2TotalDist = placemark2TotalDist + Distance(lat1,lon1,lat2,lon2)
             
         }

         fmt.Printf("\nplacemark2TotalDist: %.3f"+"km",placemark2TotalDist/1000)


         var placemark3xcors,placemark3ycors []string

         for i:=0;i<len(points3);i++ {

         	pointCors := strings.Split(points3[i],",")

         	xcor := pointCors[0]
         	ycor := pointCors[1]

         	if xcor <= "85.29511243841535" {

         		placemark3xcors = append(placemark3xcors,xcor)
         		placemark3ycors = append(placemark3ycors,ycor)
         	}

         
         }

         placemark3TotalDist := 0.0
         for j:=0;j<len(placemark3xcors)-1;j++ {

         	lon1, _ := strconv.ParseFloat(placemark3xcors[j],64)
            lat1, _ := strconv.ParseFloat(placemark3ycors[j],64)
            lon2, _ := strconv.ParseFloat(placemark3xcors[j+1],64) 
            lat2, _ := strconv.ParseFloat(placemark3ycors[j+1],64)

             
            placemark3TotalDist = placemark3TotalDist + Distance(lat1,lon1,lat2,lon2)

            
         }

         fmt.Printf("\nplacemark3TotalDist: %.3f"+"km",placemark3TotalDist/1000)

      
        totalD := placemark1TotalDist + placemark2TotalDist + placemark3TotalDist
         fmt.Printf("\nTotal distance from A4 to A8: %.3f"+"km",totalD/1000)

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
}*/