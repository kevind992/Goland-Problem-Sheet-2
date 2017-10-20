//Author: Kevin Delassus
//Problem Sheet 2
//This problem set is for you to learn the fundamentals of creating a web application in Go. 
//Create a single Git repository as your submission, complete with README and gitignore files. 
//NB: after completing each exercise, commit your code - there should be at least one commit per exercise. 
//dYou be will required to submit a URL to the repository and the use of GitHub is recommended for this purpose. 
//All code should be fully commented, and the README should explain how to clone your repository and run the code.

package main

import (

  //"fmt"
  "html/template"
  "log"
  "net/http"
  //"bytes"
)
/*func requestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "r.Method:           ",  r.Method           )
	fmt.Fprintln(w, "r.URL:              ",  r.URL              )
	fmt.Fprintln(w, "r.Proto:            ",  r.Proto            )
	fmt.Fprintln(w, "r.ContentLength:    ",  r.ContentLength    )
	fmt.Fprintln(w, "r.TransferEncoding: ",  r.TransferEncoding )
	fmt.Fprintln(w, "r.Close:            ",  r.Close            )
	fmt.Fprintln(w, "r.Host:             ",  r.Host             )
	fmt.Fprintln(w, "r.Form:             ",  r.Form             )
	fmt.Fprintln(w, "r.PostForm:         ",  r.PostForm         )
	fmt.Fprintln(w, "r.RemoteAddr:       ",  r.RemoteAddr       )
	fmt.Fprintln(w, "r.RequestURI:       ",  r.RequestURI       )

	fmt.Fprintln(w, "r.URL.Opaque:       ", r.URL.Opaque        )
	fmt.Fprintln(w, "r.URL.Scheme:       ", r.URL.Scheme        )
	fmt.Fprintln(w, "r.URL.Host:         ", r.URL.Host          )
	fmt.Fprintln(w, "r.URL.Path:         ", r.URL.Path          )
	fmt.Fprintln(w, "r.URL.RawPath:      ", r.URL.RawPath       )
	fmt.Fprintln(w, "r.URL.RawQuery:     ", r.URL.RawQuery      )
	fmt.Fprintln(w, "r.URL.Fragment:     ", r.URL.Fragment      )

	for key, value := range r.Header {
		fmt.Fprintln(w, "\t" + key + ":", value)
	}

	body := new(bytes.Buffer)
	body.ReadFrom(r.Body)
	fmt.Fprintln(w, "r.Body:             ",  body.String())
}*/

type PageVariables struct {
	Title string
}

func main() {

  //http.HandleFunc("/", requestHandler)
	http.HandleFunc("/", HomePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HomePage(w http.ResponseWriter, r *http.Request){

    title := "Guessing Game"
    HomePageVars := PageVariables{ //store the title in a struct
      Title : title,
    }

    t, err := template.ParseFiles("Q3.html") //parse the html file homepage.html
    if err != nil { // if there is an error
  	  log.Print("template parsing error: ", err) // log it
  	}
    err = t.Execute(w, HomePageVars) //execute the template and pass it the HomePageVars struct to fill in the gaps
    if err != nil { // if there is an error
  	  log.Print("template executing error: ", err) //log it
  	}
}