## go-movies-crud
A simple CRUD API built using **Go** Lang.

### Features:
* A simple server *movies* has been created.
* Here, there is no database. Instead, **structs** and **slices** have been used to perform operation on data inside the *movies* server. 
* It is served on ```http://localhost:8000/```.
* The URL router and dispatcher package ```gorilla/mux``` is employed.
* There are five routes with corresponding functions. The routes are:
  * Get All Entries
  * Get Entries By ID
  * Create
  * Update
  * Delete
* Postman was used for testing the API.
