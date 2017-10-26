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
  //"log"
  "net/http"
  //"strings"
  "strconv"
  "time"
  //"math/rand"
)

type Templatedata struct {
	Time string
	Author string
}

func guessHandler(w http.ResponseWriter, r *http.Request){
    
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

    //Parsing in the guess Template
    t, _ := template.ParseFiles("template/guess.html")
    t.Execute(w, Templatedata{Time: "HH,MM,SS", Author: "Kevin"})
    
}

func main(){

  http.Handle("/", http.FileServer(http.Dir("./static")))

  http.HandleFunc("/guess",guessHandler)

  http.ListenAndServe(":8080", nil)
}