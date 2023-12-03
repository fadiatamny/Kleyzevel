# Tuneless Treasures

##### The typical REST API Service
The goal of this project is to explore in the selected language the ways to implement the following checklist

1. Full CRUD Operations
2. HTML Landing page serving ( maybe interactive )
3. DB Connectivity 
4. Authorization and limitation on specific routes
5. Docs

### The Idea
We want to create a instrument/music shop center. This site will serve its users by allowing them to order various instruments from different categories that will be serviced and prepped by professional workers in each field. The user will be able to see the status of his order and the list of instruments he ordered through a well designed web interface that will be reactive, the interface will be connected in real time to all changes and notifications happening to the orders.

#### Requirements
- The shop has a list of instruments, each instrument has a name, a price, a description and a category
- Each instrument belongs to a manufacturer, each manufacturer has a name, a description and a list of instruments
- The shop has workers, each worker has a name, a surname, a category that he can work on
- The shop saves a list of orders, each order has a list of instruments, a worker, a date, a status (pending, completed, canceled) and a total price
- Each order is attached to a customer, each customer has a name, a surname, a list of orders and a list of instruments and an email
- Authentication should be maid either by a JWT token or by a session cookie or OTP
- The landing page will contain a login form
- Once logged in the user will be redirected to the main page where he can see the list of instruments and the list of manufacturers
- The user can order an instrument, the order will be maid by a worker, the worker will be selected by the system based on the category of the instrument and spreading of workloads
- The user can see the list of orders and the list of instruments he ordered
- If a worker logs in he will be redirected to the main page where he can see the list of orders and the list of instruments he ordered
- If a user is logged in he can get a push notification when the status of his order changes
- Orders will be accepted into a queue and processed by a cron to ingest them into the system


## Tech Stack

- Language of your choice
- DB: Postgres
- RedPanda ( Kafka )
- Redis
- OTP ( email code ) or use Keycloak
- Docker & Docker Compose
- HTMX ( for the landing page )
- Redocly for the docs