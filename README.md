# GO - Angular 2 QuickStart Source

See [Angular 2 Quickstart](Angular2Quickstart.md) for the original copy of the README.md
See [Github Angular 2 Quickstart](https://github.com/angular/quickstart) for the live original repository.

This repository is focused on the following:

* Go server integration
* In memory data repository
* Mongodb data repository

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

    
