
This library can be used for reading kml file and return the coordinates string

Example
Polylines.kml

<?xml version="1.0" encoding="UTF-8"?>
<kml xmlns="http://www.opengis.net/kml/2.2">
<Document>
<name>Shapes</name>
<Style id="thickLine">
<LineStyle>
<width>2.5</width>
</LineStyle>
</Style>
<Style id="transparent50Poly">
<PolyStyle>
<color>7fffffff</color>
</PolyStyle>
</Style>
<Placemark>
<name>Shape</name>
<description>Shape</description>
<LineString><coordinates>74.0643310546875,18.58377568837094,0 74.1412353515625,18.620218991632978,0 </coordinates></LineString><styleUrl>#thickLine</styleUrl></Placemark>
</Document>
</kml>

This library can be used by the function name i.e ReadKML with the given user parameter.

The code to read the above file can be written as

package main

import ("github.com/niketa/gorepo/kml-library/kmllib"
"fmt")


func main() {

        kmllib.ReadKML("./polylines.kml") //this will read kml file and return coordinates     
   }