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
    ![Screenshot 2024-02-25 at 6.03.49 PM.png](..%2F..%2F..%2F..%2F..%2FDesktop%2Fcreate%2FScreenshot%202024-02-25%20at%206.03.49%20PM.png)
    ![Screenshot 2024-02-25 at 6.04.09 PM.png](..%2F..%2F..%2F..%2F..%2FDesktop%2Fcreate%2FScreenshot%202024-02-25%20at%206.04.09%20PM.png)
    ![Screenshot 2024-02-25 at 6.04.31 PM.png](..%2F..%2F..%2F..%2F..%2FDesktop%2Fcreate%2FScreenshot%202024-02-25%20at%206.04.31%20PM.png)
    ![Screenshot 2024-02-25 at 6.04.47 PM.png](..%2F..%2F..%2F..%2F..%2FDesktop%2Fcreate%2FScreenshot%202024-02-25%20at%206.04.47%20PM.png)
* Login:
    ![Screenshot 2024-02-25 at 6.06.13 PM.png](..%2F..%2F..%2F..%2F..%2FDesktop%2Flogin%2FScreenshot%202024-02-25%20at%206.06.13%20PM.png)
    ![Screenshot 2024-02-25 at 6.06.42 PM.png](..%2F..%2F..%2F..%2F..%2FDesktop%2Flogin%2FScreenshot%202024-02-25%20at%206.06.42%20PM.png)
* Logout:
    ![Screenshot 2024-02-25 at 6.07.13 PM.png](..%2F..%2F..%2F..%2F..%2FDesktop%2Flogout%2FScreenshot%202024-02-25%20at%206.07.13%20PM.png)
* GetUsers:
    ![Screenshot 2024-02-25 at 6.08.04 PM.png](..%2F..%2F..%2F..%2F..%2FDesktop%2Flistusers%2FScreenshot%202024-02-25%20at%206.08.04%20PM.png)
* SendMessage:
    ![Screenshot 2024-02-25 at 6.08.51 PM.png](..%2F..%2F..%2F..%2F..%2FDesktop%2FScreenshot%202024-02-25%20at%206.08.51%20PM.png)
    ![Screenshot 2024-02-25 at 6.11.34 PM.png](..%2F..%2F..%2F..%2F..%2FDesktop%2FScreenshot%202024-02-25%20at%206.11.34%20PM.png)
    ![Screenshot 2024-02-25 at 6.12.09 PM.png](..%2F..%2F..%2F..%2F..%2FDesktop%2FScreenshot%202024-02-25%20at%206.12.09%20PM.png)
* GetMessages:
    ![Screenshot 2024-02-25 at 6.13.28 PM.png](..%2F..%2F..%2F..%2F..%2FDesktop%2FScreenshot%202024-02-25%20at%206.13.28%20PM.png)

# Future Work
* In future if we fail scale issues we can use Redis for caching.
* We can also migrate to NoSQL databases like DynamoDB for better performance.