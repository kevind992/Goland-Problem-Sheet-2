//Author: Kevin Delassus
//Problem Sheet 2
//This problem set is for you to learn the fundamentals of creating a web application in Go. 
//Create a single Git repository as your submission, complete with README and gitignore files. 
//NB: after completing each exercise, commit your code - there should be at least one commit per exercise. 
//dYou be will required to submit a URL to the repository and the use of GitHub is recommended for this purpose. 
//All code should be fully commented, and the README should explain how to clone your repository and run the code.

package main

import (
  "html/template"
  "log"
  "net/http"
)

type PageVariables struct {
	Title string
}

func main() {

  http.HandleFunc("/guess/", GuessPage)
  //http.HandleFunc("/index/", HomePage)
  
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/*func HomePage(w http.ResponseWriter, r *http.Request){

    title := "Guessing Tom"
    HomePageVars := PageVariables{ //store the title in a struct
      Title : title,
    }

    t, err := template.ParseFiles("index.html") //parse the html file homepage.html

    if err != nil { // if there is an error
  	  log.Print("template parsing error: ", err) // log it
  	}
    err = t.Execute(w, HomePageVars) //execute the template and pass it the HomePageVars struct to fill in the gaps
    if err != nil { // if there is an error
  	  log.Print("template executing error: ", err) //log it
  	}
}*/
func GuessPage(w http.ResponseWriter, r *http.Request){
  
      title := "Guessing Tom"
      GuessPageVars := PageVariables{ //store the title in a struct
        Title : title,
      }
  
      t, err := template.ParseFiles("guess.html") //parse the html file homepage.html
  
      if err != nil { // if there is an error
        log.Print("template parsing error: ", err) // log it
      }
      err = t.Execute(w, GuessPageVars) //execute the template and pass it the HomePageVars struct to fill in the gaps
      if err != nil { // if there is an error
        log.Print("template executing error: ", err) //log it
      }
  }