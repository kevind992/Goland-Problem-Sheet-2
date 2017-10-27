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
  "net/http"
  "strconv"
  "time"
  "math/rand"
)
//Setting Struct
type GuessData struct {
  Message string
  ContratsMes string
  Guess int
  CountGos int
}

//Setting Variables
var target int = 0
var countGos = 0

func guessHandler(w http.ResponseWriter, r *http.Request){
   
    //This will be displayed in the H1 on the guess.tmpl
    headerMessage := "Guess a number between 1 and 20"
   
    guessReply := " "

    // Try to read the cookie.
    var cookie, err = r.Cookie("target")
    if err == nil {
      // If we could read it, try to convert its value to an int.
      target, _ = strconv.Atoi(cookie.Value)
    }
    //if random number is equeled to 0 rand.Intn is generating a random number between 1 & 20
    if (target == 0){
      rand.Seed(time.Now().UTC().UnixNano())
      target = rand.Intn(20)
    }
    // Create a cookie instance and set the cookie.
    // You can delete the Expires line (and the time import) to make a session cookie.
    cookie = &http.Cookie{
      Name:    "target",
      Value:   strconv.Itoa(target),
      Expires: time.Now().Add(72 * time.Hour),
    }
    http.SetCookie(w, cookie)
    
    r.ParseForm()

    //taking the user input from guess.tmpl and casting it to an int into guess
    guess,_:=strconv.Atoi(r.Form.Get("guessinput"))
  
    //Filtering the different responces by what the user guess is
    if guess > 0 {
      if guess == target{
          guessReply = "Correct Answer!"
          target = 0;
          countGos++
      }else if guess > target{
          guessReply = "Your to High, Try Again"
          countGos++
      }else if guess < target{
          guessReply = "Your to Low, Try Again"
      }
    }else{
      guessReply = " "
    }
    //Parsing the guess.tmpl file
    t, _ := template.ParseFiles("template/guess.tmpl")

    t.Execute(w, &GuessData{Message: headerMessage, Guess:guess, ContratsMes: guessReply, CountGos: countGos })
}
func main(){
  //handling everything in the static folder
  http.Handle("/", http.FileServer(http.Dir("./static")))

  http.HandleFunc("/guess",guessHandler)

  http.ListenAndServe(":8080", nil)
}