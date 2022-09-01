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