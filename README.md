# GO - Angular 2 QuickStart Source

See [Angular 2 Quickstart](angular/README.md) for the original copy of the README.md
See [Github Angular 2 Quickstart](https://github.com/angular/quickstart) for the live original repository.

This repository just copies the contents of the original in a angular/ directory
and adds a goserver/ directory with a simple program that returns the angular 2 content

# Go server integraion

This main program will do the trick.
I chose github.com/gorilla/mux for the mux instead of the standard library because it is great and I use it for all projects.
Compile and run from the goservrer directory, 
notice the ".." in `r.PathPrefix("/").Handler(http.FileServer(http.Dir("../angular")))`

    git clone ...
    cd angular
    npm install
    npm install -g tsc
    npm run tsc

At this point you could run as described in the original turorial.  
But lets start things up using go

    cd ../goserver
    go get -v
    go build
    ./goserver

