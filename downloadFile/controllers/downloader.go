package controllers

import ("net/http"
		"os"
		"io"
		"fmt"
		"strings"
		"github.com/dustin/go-humanize"
        "net/url"
		)



type WriteCounter struct {
	Total uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

func (wc WriteCounter) PrintProgress() {

	fmt.Printf("\r%s", strings.Repeat(" ", 35))

	
	fmt.Printf("\rDownloading... %s complete", humanize.Bytes(wc.Total))
}



func FileURL(w http.ResponseWriter,r *http.Request) {

	if r.Method == "GET"{	  

    fileUrl := "https://d1ohg4ss876yi2.cloudfront.net/golang-resize-image/big.jpg"

   /* err := DownloadFile("bigger.png", fileUrl,w)
    if err != nil {
        panic(err)
     }*/

     fileURL,err := url.Parse(fileUrl)

       if err != nil {
                 panic(err)
         } 

          path := fileURL.Path

         segments := strings.Split(path, "/")
        // fmt.Println("file segment:",segments)

         filePath := segments[2] // change the number to accommodate changes to the url.Path position
        // fmt.Println("file name",fileName) 

   err = DownloadFile(filePath,fileUrl,w)
    if err != nil {
        panic(err)
     }

   }
}


func DownloadFile(filepath string, url string,w http.ResponseWriter) error {

	
    out, err := os.Create(filepath+".tmp")
    if err != nil {
        return err
    }
    defer out.Close()

    // Get the data
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

  
	counter := &WriteCounter{}

   _, err = io.Copy(w, io.TeeReader(resp.Body, counter))
   
    if err != nil {
        return err
    }


    fmt.Print("\n")

	err = os.Rename(filepath+".tmp", filepath)
	if err != nil {
		return err
	}


   return nil
} 


