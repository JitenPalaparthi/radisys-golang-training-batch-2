- Vendor Microservice

- Create Vendor
- Get Vendor By ID
- Update Vendor By ID
- Delete Vendor By ID
- Get All Vendors

- The database layer (DAL) must be plug and play.
- That means DAL layer should be loosely coupled with other layers.

- Database is Postgres
- Docker instance of postgres
- Message Broker Kafka

- Containerise it and run it in Kubernetes

## Database

- postgres

    ```docker run -d --name pgvendorsdb -e POSTGRES_PASSWORD=admin123 -e POSTGRES_USER=admin -e POSTGRES_DB=vendorsdb -p 5432:5432 postgres```

- kafka

    ```docker-compose -f docker/kafka-zookeeper.yaml up -d```
    
- once kafka is up and run exec the kafka container and run the below command

    ```docker exec -it dockers-kafka-1 bash```  - Note: use proper container name or id

    ```kafka-topics --create --if-not-exists --bootstrap-server kafka:29092 --partitions 1 --replication-factor 1 --topic vendor.created```

- What is a message broker?

SAGA Pattern
- Orchastration : orchastrator
- choreography  : each and every component that is involved in it

 - Whan a new vendor is created, I want to send an email, SMS and also notify
  other system that a new Vendor has been created.

  - Flight Ticket
    - booking.com
    - search for flights based on destination and also dates
    - select a flight
    - select a seat
    - select few other options
    - make payment
    - upon successful payment
    - a booking is made
    - a seat is allocated to you
    - deatils are send to your email, app notification and also message

- cluster: multiple nodes/brokers/state machines
- broker: server or a state machine
- topic : a subject that publisher and subscribers communicate on
- partition: messages spread across multiple brokers as multiple partitions
- replication: duplicating a message on multiple brokers
- publisher: who send messages on a topic
- subscriber: who subscribes to a topic and fetch messages those are published
- zookeeper: communication between , discovery of kafka brokers

- Go client
