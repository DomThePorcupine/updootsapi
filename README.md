## Welcome to the updoots™ API

# To get up and running follow these steps:

1. `git clone https://github.com/DomThePorcupine/updootsapi.git`

2. `docker-compose build`

3. `docker-compose run updoots_web bash` (you are now inside a docker container)

4. `mysql -h updoots_db -pnpassword` (you are now inside the SQL command line)

5. `create database testdb;`

6. `grant all privileges on testdb.* to 'nuser';`

7.  Paste the following into the SQL shell:
    ```

    create table `messages` (
        `id` INT NOT NULL AUTO_INCREMENT,
        `message` VARCHAR(250),
        `userid` VARCHAR(20),
        `updoots` INT,
        PRIMARY KEY (`id`)
    );

    ```

8. `\q` (you are now back in docker bash again)

9. `mysql -h updoots_db -unuser -pnpassword testdb` (you are now inside the SQL command line again)

1. `\q` to exit mysql cli, then `exit` to exit the docker containers bash

11. `docker-compose up`

## Now navigate to http://127.0.0.1:3000, and start hacking on the go file!