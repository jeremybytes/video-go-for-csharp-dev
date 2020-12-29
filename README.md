Video: Go for the C# Developer
=======================
Code samples to go along with "A Tour of Go for the C# Developer" video (link below). The aim is to provide an overview of Go (golang) language features for developers familiar with the C# language.

Motivation
----------
Learning other programming languages enhances our work in our primary language. From the perspective of a C# developer, the Go language (golang) has many interesting ideas. Go is opinionated on some things (such as where curly braces go and what items are capitalized). Declaring an unused variable causes a compile failure; the use of "blank identifiers" (or "discards" in C#) are common. Concurrency is baked right in to the language through goroutines and channels. Programming by exception is discouraged; it's actually called a "panic" in Go. Instead, errors are treated as states to be handled like any other data state. We'll explore these features (and others) by building an application that uses concurrent operations to get data from a service. These ideas make us think about the way we program and how we can improve our day-to-day work (in C# or elsewhere).  

**Video Walkthrough**: [A Tour of Go (golang) for the C# Developer](https://www.youtube.com/watch?v=NW-8WpnGQtE)

Project Layout
--------------

**/async** contains the Go program  
**/net-core-people-service** contains a .NET Core 3.1 service (used by the Go program)  

The Go program is a console application that calls the .NET Core service and displays the output. In order to show concurrency, the application gets each record individually.

*Note: This project assumes that you have both "go" (this was created with version 1.14.5) and "dotnet" (this was created with 3.1.302) installed. Visit [https://golang.org/doc/install](https://golang.org/doc/install) and [https://dotnet.microsoft.com/download](https://dotnet.microsoft.com/download) to download the tools. In addition, I am using Visual Studio Code with the "Go" extension. Samples should work on all platforms supported by the runtimes (Windows, macOS, and Linux).*

Running the Service
-------------------  
The .NET Core service can be started from the command line by navigating to the ".../net-core-people-service" directory and typing `dotnet run`. This provides endpoints at the following locations:

* http://localhost:9874/people  
Provides a list of "Person" objects. This service will delay for 3 seconds before responding. Sample result:

```json
[{"id":1,"givenName":"John","familyName":"Koenig","startDate":"1975-10-17T00:00:00-07:00","rating":6,"formatString":null},  
{"id":2,"givenName":"Dylan","familyName":"Hunt","startDate":"2000-10-02T00:00:00-07:00","rating":8,"formatString":null}, 
{...}]
```

* http://localhost:9874/people/ids  
Provides a list of "id" values for the collection. Sample:  

```json
[1,2,3,4,5,6,7,8,9]
```

* http://localhost:9874/people/1  
Provides an individual "Person" record based on the "id" value. This service will delay for 1 second before responding. Sample record:

```json
{"id":1,"givenName":"John","familyName":"Koenig","startDate":"1975-10-17T00:00:00-07:00","rating":6,"formatString":null}
```

The Go Sample Program
---------------------
The **/async** folder contains the "main.go" file which is the completed project.  

In addition, subfolders provide the "main.go" file at intermediate steps along the way.

**/step01**  
Video Chapter: [Basics](https://youtu.be/NW-8WpnGQtE?t=160)  
Basics including package, import, functions, braces, package exports, variable assignment, named return values, and bare returns. Also, the Go extension for Visual Studio Code that handles formatting, imports, and linting.

**/step02**  
Video Chapter: [Calling a web service](https://youtu.be/NW-8WpnGQtE?t=740)  
Calls a web service to get a list of IDs. Concepts include multiple return values, handling errors, parsing JSON, and using "defer" to run functions.

**/step03**  
Video Chapter: [Parsing JSON](https://youtu.be/NW-8WpnGQtE?t=1415)  
Calls a web service to get an individual "Person" record. Concepts include creating a struct type, exposing properties, creating formatted strings, another way to parse JSON, and printing to the console.

**/step04**  
Video Chapter: ["for" loops](https://youtu.be/NW-8WpnGQtE?t=2186)  
Outputs a complete collection of "Person" records. Concepts include the for loop (indexers and range) as well as blank identifiers

**/step05**  
Video Chapter: [Interfaces and methods](https://youtu.be/NW-8WpnGQtE?t=2460)  
Creates a default formatter for the "person" type. Concepts include interfaces, methods, and method receivers.

**/step 06**  
Video Chapter: [Time and Args](https://youtu.be/NW-8WpnGQtE?t=3005)  
Adds an elapsed time to output and reads a command-line argument. Concepts include time, duration, duration output, and checking for application arguments.

**/step07**  
Video Chapter: [Concurrency](https://youtu.be/NW-8WpnGQtE?t=3310)  
Gets all of the "Person" records at the same time by making multiple concurrent service calls. Concepts include goroutines, channels, make(), writing to channels, and reading from channels.

**/step08**  
Video Chapter: [Errors](https://youtu.be/NW-8WpnGQtE?t=4030)  
Adds an "error" return value to the "getIDs" method. Concepts include creating error objects, appending error messages, multiple return values, exiting on error, and using "log.Fatalf()".

**/step09**  
Video Chapter: [Concurrency and errors](https://youtu.be/NW-8WpnGQtE?t=4480)  
Adds an "error" return value to the "getPerson" method. Concepts include checking HTTP status codes, short-circuiting loops, more on writing & reading channels.

**Other Topics**  
This is by no means an exhaustive look at Go. Additional topics and topics to look into further include packages, exports, project structure, types, interfaces, pointers, inline goroutines, closures, and "sync.WaitGroup".
