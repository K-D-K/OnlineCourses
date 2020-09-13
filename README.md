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
