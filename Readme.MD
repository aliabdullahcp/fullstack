# Golang_REST_API
## Using Golang, Docker, Postgres

### How to RUN
1. Install Docker Desktop and make it up and running
2. Open Terminal in Project's Root Directory (i.e where docker-compose.yml file is located)
3. Run the following Command to Start the project. Wait for it to get up and running
    ```sh
    docker-compose up --build
    ```
4. To Stop the Application you can run the following command.
    ```sh
    docker-compose down --remove-orphans --volumes
    ```

### How to Use Database Using PGAdmin
1. While the project is running, Go to browser and open
    ```sh
    http://localhost:5050
    ```
2. Use the following credentials
    ```sh
    email: live@admin.com
    passowrd: password
    ```
3. Right Click on the "Servers" to create a new server. Choose "Create" then "Server"
4. Fill in any name that you want.
5. Click on connection tab
    Here you will need the IPAddress of the PGAdmin container.
Use the following Command to get
    ```sh
    docker container ls
    ```
    Copy the ID of the full_db_postgres container
    ```sh
    docker inspect <container_id> | grep IPAddress
    ```
6. Once got the IP Address, Fill in the connection tab with Hostname/Address and Username as "postgres" and Password as "password"
7. Click Save.