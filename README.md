
# DigiCart

An ecommerce application for homebakers.
This application is a backend server for a bakery ecom application made for homebakers who out of hobby sell homemade bakery items. This helps them monetize and give them presence on a platform.

## Architecture
The backend Architecture is made up of four microservices.

**Auth** - Authenticates the user and returns the auth_token necessary to interact with other services.

**Orders** - Deals with placing and showing the existing orders to the user.

**Payments** - Deals with making the payments for the orders and showing the payment history.

**Notify** - All the above microservices interact with this service using **RabbitMQ** to place their request to notify the user through email.

## Database
**Postgres** is used as the database.
DB_USER is picked from .env.
Create a DB named digicart for this application.
Note : you might need to change the 
dsn in connect.go if your postgres user requires password for login. For quick setup, use the user with no password.

## API Reference

### Auth microservice (http://localhost:8080)
#### User signup

```http
  POST /signup
```
Signup the user and returns the auth_token (expires after 24 hours)

Authorization : None
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `email` | `string` | **Required** either email or phone |
| `password` | `string` | **Required**|
| `firstname` | `string` | Your first name |
| `lastname` | `string` | lastname |
| `phone` | `string` | **Required** either email or phone |
| `address` | `address` | **Required** { street : "", city : "", country: ""}|

#### User Login

```http
  POST /login
```
Login the user and returns the auth_token (expires after 24 hours)

Authorization : None
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `email or phone`      | `string` | **Required**|
| `password`      | `string` | **Required**|

### Orders microservice (http://localhost:8090)

#### Get All Orders
```http
  GET /orders
```
Returns all the orders that belong to the user.

Authorization : Bearer <JWT_TOKEN>

#### Get Order
```http
  GET /orders/{id}
```
Returns the order with {id} that belong to the user.

Authorization : Bearer <JWT_TOKEN>

#### Place Order
```http
  POST /orders/new
```
Places a order for the user and creates a pending payment for user by calling payment microservice

Authorization : Bearer <JWT_TOKEN>

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `productID`      | `string` | **Required** the product id for  which the order needs to be placed|

### Payments microservice (http://localhost:8070)

#### Pay Order
```http
  GET /pay/{id}
```
Complete the payment for the payment id.

Authorization : Bearer <JWT_TOKEN>

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `productID`      | `string` | **Required** the product id for  which the order needs to be placed|

#### Get All Payments
```http
  GET /payments
```
Returns all the payments for the user.

Authorization : Bearer <JWT_TOKEN>

#### Create Payment
```http
  POST /payments
```
Used by order microservice to create payment for the order.

Authorization : Bearer <JWT_TOKEN>


## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`JWT_SECRET_KEY Ex - "ddfd"`

`PAYMENT_SERVICE_URL`

`AMQP_SERVER_URL default: amqp://guest:guest@localhost:5672/`

`SENDER_EMAIL`

`SENDER_PASSWORD`

`SMTP_HOST default: smtp.gmail.com`

`SMTP_PORT default:587`
## FAQ

#### Is testing done for this project ?

The notify microservice contain some unit tests.
Other microservice contain db related functions and suitable unit tests could not be written.



