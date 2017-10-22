//Author: Kevin Delassus
//Problem Sheet 2
//This problem set is for you to learn the fundamentals of creating a web application in Go. 
//Create a single Git repository as your submission, complete with README and gitignore files. 
//NB: after completing each exercise, commit your code - there should be at least one commit per exercise. 
//dYou be will required to submit a URL to the repository and the use of GitHub is recommended for this purpose. 
//All code should be fully commented, and the README should explain how to clone your repository and run the code.

package main

import (

  "fmt"
  "html/template"
  "log"
  "net/http"
  //"strings"
  "strconv"
  "time"
  "math/rand"
)

type PageVariables struct {
  Title string
}

type MessageVariables struct{
  Message string
}

func HomePage(w http.ResponseWriter, r *http.Request){
    
    count := 0
  
    // Try to read the cookie.
    var cookie, err = r.Cookie("count")
    if err == nil {
      // If we could read it, try to convert its value to an int.
      count, _ = strconv.Atoi(cookie.Value)
    }
  
    // Increase count by 1 either way.
    count += 1
  
    // Create a cookie instance and set the cookie.
    // You can delete the Expires line (and the time import) to make a session cookie.
    cookie = &http.Cookie{
      Name:    "count",
      Value:   strconv.Itoa(count),
      Expires: time.Now().Add(72 * time.Hour),
    }
    http.SetCookie(w, cookie)  

    title := "Guessing Game"
    HomePageVars := PageVariables{ //store the title in a struct
      Title : title,
    }
  
    t, err := template.ParseFiles("index.html") //parse the html file homepage.html

    if err != nil { // if there is an error
  	  log.Print("template parsing error: ", err) // log it
  	}
    err = t.Execute(w, HomePageVars)
    if err != nil { // if there is an error
  	  log.Print("template executing error: ", err) //log it
  	}
}
func GuessPage(w http.ResponseWriter, r *http.Request){

    target := 0
  
    // Try to read the cookie.
    var cookie, err = r.Cookie("target")
    if err == nil {
      // If we could read it, try to convert its value to an int.
      target, _ = strconv.Atoi(cookie.Value)
    }
  
    // Increase count by 1 either way.
    if (target == 0){

        target = rand.Intn(10)
    }
    fmt.Println(target)
    // Create a cookie instance and set the cookie.
    // You can delete the Expires line (and the time import) to make a session cookie.
    cookie = &http.Cookie{
      Name:    "target",
      Value:   strconv.Itoa(target),
      Expires: time.Now().Add(72 * time.Hour),
    }
    http.SetCookie(w, cookie)


      message := "Guess a number Between 1 and 10"
      Message := MessageVariables{ //store the title in a struct
        Message : message,
      }
  
      t, error := template.ParseFiles("guess.html") //parse the html file homepage.html
  
      if error != nil { // if there is an error
        log.Print("template parsing error: ", error) // log it
      }
      error = t.Execute(w,  Message) //execute the template and pass it the HomePageVars struct to fill in the gaps
      if error != nil { // if there is an error
        log.Print("template executing error: ", error) //log it
      }
}

func main() {
 
	// Send a 404 for favicon.ico
  http.Handle("/favicon.ico", http.NotFoundHandler())
  
  http.HandleFunc("/", HomePage)
  http.HandleFunc("/guess", GuessPage)
  
  
	log.Fatal(http.ListenAndServe(":8080", nil))
}