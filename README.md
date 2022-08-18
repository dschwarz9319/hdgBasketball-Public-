# Hdg Basketball

## Overview
Hdg Basketball is a site intended to be a source of communication and reference for parents and children
enrolled in the youth Harford County Basketball league.

## Setup
In order to properly run this app, First you need Golang, NPM, and MongoDB installed on your local machine.
- [Golang](https://golang.org/doc/install)
- [NPM](https://www.npmjs.com/get-npm)
- [MongoDB](https://docs.mongodb.com/manual/installation/)

You will also need to have ports 27017, 8000, and 3000 unallocated and available for use on your machine.

# Starting the API
In a terminal on your local computer, from the root of the project, navigate to the `/api` folder, then run the command:
```
go run server.go
```

Afterwards you should see a list of endpoints and HTTP requests in the terminal waiting for requests. If you wish to test the functionality of these requests, in [postman](https://www.postman.com/downloads/), import the collection titled `Hdg-Basketball.postman_collection.json` located in the `/dev` folder. The collection you imported contains a list of postman requests that should test the various endpoints of the api.

NOTE: The API functionality is dependent on an instance of MongoDB running on your machine and listening on port `27017`.

# Starting the Front-end/UI
In a terminal on your local computer (preferably a different one than the terminal used for running the API), from the root of the project, navigate to the `/ui` folder, then run the commands:
```
npm install node-sass
npm i
npm start
```

From there a browser window should pop up with the website homepage loaded in it. Be aware that some functionality of the UI is dependant on having the API currently running on port 3000 of your local machine.

## Admin Features (CRUD Functionality)

If you wish to test the `admin` side of the website, navigate to the website's URL, (likely localhost:3000 if you're running the app locally) and append `/admin/home` to that URL. You should see a menu with options to add or remove news, as well as add, remove, or update a team's standings. Each change you make in the admin portion of the website should immediately show up in the corresponding page of the UI.

<br/>
--------------------------------
<br/>

### Desktop View
<p align="center">
  <img src="https://user-images.githubusercontent.com/51220736/185450151-59020aa0-2859-4095-82db-1f5b02fbd63e.png" />
</p>

<br/><br/>

### Mobile View
<p align="center">
  <img src="https://user-images.githubusercontent.com/51220736/185450437-da9f693f-86ba-4281-8f47-1b58b086d9cf.png" />
</p>
