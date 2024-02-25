## Chat Service Backend

This is the backend for the chat service. 

# Design
 Database used is MySQL. It has two tables, one for users and one for messages. 
 * User table has the following columns:
    * username - PK
    * password 
    * logged_in
 * Message table has the following columns:
    * id - PK
    * sender
    * receiver
    * message
    * time_stamp

The flow for the APIs is as follows:
* CreateUser: 
    * The user sends a POST request to the server with the username and password. 
    * The server checks if the username is already taken. If not, it creates a new user in the user table.
* Login:
    * The user sends a POST request to the server with the username and password. 
    * The server checks if the username and password match. If yes, it updates the logged_in column in the user table.
* Logout:
    * The user sends a POST request to the server with the username. 
    * The server updates the logged_in column in the user table.
* GetUsers:
    * The user sends a GET request to the server. 
    * The server returns all the users in the user table.
* SendMessage:
    * The user sends a POST request to the server with the sender, receiver and message. 
    * The server adds the message to the message table.
* GetMessages:
    * The user sends a POST request to the server with the sender and receiver. 
    * The server returns all the messages between the sender and receiver.

# How to run
* Clone the repository
* Run the following commands:
    * `docker-compose up --build`
    * `/bin/bash init.sh`

# Screenshot for the curls and output
* CreateUser:
  <img width="1397" alt="Screenshot 2024-02-25 at 6 04 47 PM" src="https://github.com/musafir-V/chatbackend/assets/60506099/443ac705-5db4-4bce-b962-7c4c9e86a292">
 <img width="1392" alt="Screenshot 2024-02-25 at 6 04 31 PM" src="https://github.com/musafir-V/chatbackend/assets/60506099/b2d1434f-aaa8-41ca-9607-4d6374158f2d">
 <img width="1432" alt="Screenshot 2024-02-25 at 6 04 09 PM" src="https://github.com/musafir-V/chatbackend/assets/60506099/dd5a9736-52fc-4c21-8511-b4c672689bfb">
 <img width="1390" alt="Screenshot 2024-02-25 at 6 03 49 PM" src="https://github.com/musafir-V/chatbackend/assets/60506099/0f2065f7-bd41-4988-aae6-ec072581d2e5">

* Login:
   <img width="1401" alt="Screenshot 2024-02-25 at 6 06 42 PM" src="https://github.com/musafir-V/chatbackend/assets/60506099/fc66ea24-03b4-4a56-925a-29be5034d433">
   <img width="1395" alt="Screenshot 2024-02-25 at 6 06 13 PM" src="https://github.com/musafir-V/chatbackend/assets/60506099/082fae4e-14ce-4bb2-aed4-9afaee2226bb">
* Logout:
    <img width="1124" alt="Screenshot 2024-02-25 at 6 07 13 PM" src="https://github.com/musafir-V/chatbackend/assets/60506099/5d6196f3-eafd-4440-a123-c57c1433d399">

* GetUsers:
    <img width="575" alt="Screenshot 2024-02-25 at 6 08 04 PM" src="https://github.com/musafir-V/chatbackend/assets/60506099/94286024-962d-4315-ba1c-e06b863497e8">
* SendMessage:
   <img width="1446" alt="Screenshot 2024-02-25 at 6 12 09 PM" src="https://github.com/musafir-V/chatbackend/assets/60506099/c1f09c5b-616a-47c3-b749-9a1ed6c15375">
<img width="1196" alt="Screenshot 2024-02-25 at 6 11 34 PM" src="https://github.com/musafir-V/chatbackend/assets/60506099/b21e5f6b-83f3-4eba-8430-3f504dc3561d">
<img width="1128" alt="Screenshot 2024-02-25 at 6 08 51 PM" src="https://github.com/musafir-V/chatbackend/assets/60506099/51920870-3365-4cf4-89da-13b589723eed">
* GetMessages:
  <img width="1276" alt="Screenshot 2024-02-25 at 6 13 28 PM" src="https://github.com/musafir-V/chatbackend/assets/60506099/21bf7d70-b3b2-4b28-a77e-f3ddb848c0bc">
# Future Work
* In future if we fail scale issues we can use Redis for caching.
* We can also migrate to NoSQL databases like DynamoDB for better performance.
