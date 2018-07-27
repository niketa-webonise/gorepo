package controllers

import ("net/http"
        "os"
        "io/ioutil"
        "fmt"
        "bytes"
        )




 func FileURL(w http.ResponseWriter, r *http.Request) {

          // grab the generated receipt.pdf file and stream it to browser
          streamPDFbytes, err := ioutil.ReadFile("./downloads/UserStories.pdf")

          if err != nil {
                  fmt.Println(err)
                  os.Exit(1)
          }

          b := bytes.NewBuffer(streamPDFbytes)


          // stream straight to client(browser)
          
          w.Header().Set("Content-Disposition","attachment;filename="+"UserStories.pdf")
          
          if _, err := b.WriteTo(w); err != nil { 
                  fmt.Fprintf(w, "%s", err)
          }

          w.Write([]byte("PDF Generated"))
  }

