# OnlineCourses
# Tech
 - golang
 - gorm
 - go-chi
 
# Development
 - Database - Postgres
 To Run Project
    ```sh
    $ DBPORT="{port_no}" DBUSER="${user_name}" DBNAME="${database_name}" go run .
    ```
    Example:
     - psql
     - create database onlinecourse ;
     - grant ALL PRIVILEGES on database onlinecourse to kdk ;
     - Kill Terminal
     - DBPORT="5432" DBUSER="kdk" DBNAME="onlinecourse" go run .

# Features
 - Create User
 - Create Course
 - Get Course
 - Enroll Course
 - Grant Permission
 
# API
 **Create User**
----
  Create List of users
  * **URL** <br>
   /users
 * **Method** <br>
  `POST`
* **Data Params**
```
[
    {
        "name": "Dharani Kumar Admin-3",
        "role": 1, // Admin = 1 , Student = 2
        "address": "Kangeyam",
        "phone": "7708407974"
    }
]
```
 * **Success Response:**
    * **Code:** 200 <br>
    **Content:**
```
[
    {
        "id": "5",
        "created_at": "2020-09-13T20:39:11.687910166+05:30",
        "updated_at": "2020-09-13T20:39:11.687910166+05:30",
        "deleted_at": null,
        "name": "Dharani Kumar Admin-3",
        "phone": "7708407974",
        "address": "Kangeyam",
        "role": 1
    }
]
```
 
