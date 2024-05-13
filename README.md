# Tech Challenge 1

The Tech Challenge 1 aims to do a solution for a Fast Food restaurant. With this software, the rastaurant can do a all the process of for a place
that makes all steps of a fast food dish order, as:

- Products creation/manipulation by the restaurant owner
- Customer identification
- Payment process
- Order tracking by the chef
- Order tracking by the waiter and the customer

All the Endpoints can be called by accessing `http://localhost:3210/api` API url.

To build and run this project. Follow the Docker section

## Docker build and run

This project was built using Docker and Docker Compose. So, to build and run the API, we need to run:

```
$ docker compose build
```

After the image build finish, run run:

```
$ docker compose up -d
```

After the containers shows these status:

```
 ✔ Container fastfood-database  Started
 ✔ Container fastfood-app       Started 
```

we can access `http://localhost:3210/api` endpoints

## How to use

To use all the endpoints in this API, we can follow these sequence to simulate a customer making an order in a restaurant.
We can separate in three moments.

- Restaurant products injestion. This is used by the restaurant owner to create all the product portfolio with its images and prices
- Customer self service. This is used by the customer to choose the products, pay for it and create an order 
- Order preparing and deliver. This is used by the chef and waiter to check the order status

We will divide in 2 sections: **Restaurant** owner and Customer **order**

## Section 1 - Restaurant owner

This will be added later

## Section 2 - Customer order

This section will use all the Endpoints to make a entire order flow.

### 1. User identification (Customer view)

- Call the GET `http://localhost:3210/api/customers/{cep}` to get this [Customer ID]
or
- Cal the POST `http://localhost:3210/api/customers` to create a Customer and retrieve the [Customer ID]

### 2. List all the categories (Customer view)

- Call the GET `http://localhost:3210/api/products/categories` to get a string array with all created categories

### 3. List products by the chosen category (Customer view)

- Call the GET `http://localhost:3210/api/products/categories/{category}` to get all products by a category
or
- Call the GET `http://localhost:3210/api/products/combo` if the customer wants to see the Combo products

With this endpoints we can simulate a screen producst selection by chosing all products IDs we want to deal and create a Order

### 4. Pay the products amount (Customer view)

- Call the GET `http://localhost:3210/api/payments/types` to show to customer which payment type to choose
- Call the POST `http://localhost:3210/api/payments` to pay for the amount and receive the [Payment ID]

### 5. Create an order (Customer view)

- Call the POST `http://localhost:3210/api/orders` with:
- - All the [Products IDs] chosen [*required]
- - The [Payment ID] [*required]
- - The [Customer ID] [optional]
- - Total price for the all products sum

### 6. List order to prepate (Chef view)

- Call the GET `http://localhost:3210/api/orders/to-prepare` to list the Orders with its [Order ID]

### 7. Update order to preparing (Chef view)

- Call the PUT `http://localhost:3210/api/orders/{id}/preparing` to set Preparing status

### 8. Update order to done (Chef view)

- Call the PUT `http://localhost:3210/api/orders/{id}/done` to set Done status

### 9. Update order to delivered (Waiter view)

- Call the PUT `http://localhost:3210/api/orders/{id}/delivered` to set Delivered status to indicate that customer receive the meal. 
This is used to 'finish' the order and can be used to track some convertion rate

### 9. Update order to not delivered (Waiter view)

- Call the PUT `http://localhost:3210/api/orders/{id}/not-delivered` to set Not Delivered status to indicate that customer doesn not receive the meal.
This is used to 'finish' the order and can be used to track some convertion rate

## Swagger

This project uses Swagger to show an site with all Endpoints used by this project to make an order in a Fast Food place. 
To create/update all Endpoints documentation just run `swag init -g cmd/api/main.go`. By doing this, we can see the documentation in
two different ways:

### Swagger

http://localhost:3210/swagger/index.html

### Redoc

http://localhost:3211/docs
