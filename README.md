# Problem_Sheet_2_Web_Applications


This problem set was set up to learn the fundamentals of creating a web application in Go.

To run this program you first need to make sure you have the Go compiler on your Pc

You then need to clone the repository https://github.com/kevind992/Problem_Sheet_2_Web_Applications.git

Navigate to the folder location on the cmd

To compile the program  

> go build GuessingGame.go

To compile the program

> GuessingGame.exe

You then need to open a browser and navigate to 

http://localhost:8080/

You should now see the program working.

For the Answer to Part 1

if you open your cmd and run the command

curl http://localhost:8080/ -verbose

This will then give you the information

λ curl http://localhost:8080/ -verbose
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 8080 (#0)
> GET / HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.55.1
> Accept: */*
> Referer: rbose
>
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Content-Length: 703
< Content-Type: text/html; charset=utf-8
< Last-Modified: Fri, 27 Oct 2017 09:24:26 GMT
< Date: Fri, 27 Oct 2017 09:51:23 GMT
<
<!doctype html>


