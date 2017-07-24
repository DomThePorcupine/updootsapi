## Welcome to the updoots™ API

# To get up and running follow these steps:

1. `git clone https://github.com/DomThePorcupine/updootsapi.git`

2. `docker-compose build`

3. `docker-compose run updoots_web bash` (you are now inside a docker container)

4. `mysql -h updoots_db -pnpassword` (you are now inside the SQL command line)

5. `create database testdb;`

6. `grant all privileges on testdb.* to 'nuser';`
  
7. `\q` (you are now back in docker bash again)

8. `mysql -h updoots_db -unuser -pnpassword testdb` (you are now inside the SQL command line again)

9. Paste the following into the SQL shell:
    ```
    create table `posts` (
        `id` INT NOT NULL AUTO_INCREMENT,
        `message` VARCHAR(250),
        `userid` VARCHAR(20),
        `created` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        PRIMARY KEY (`id`)
    );
    then:
    create table `votes` ( 
        `message` INT NOT NULL, 
        `userid` VARCHAR(20), 
        `updoot` TINYINT(1) NOT NULL default 0, 
        `downdoot` TINYINT(1) NOT NULL default 0
    );
    then:
    create table `users` (
        `userid` VARCHAR(20) NOT NULL UNIQUE,
        `admin` TINYINT(1) NOT NULL default 0
    );
    ```

10. `\q` to exit mysql cli, then `exit` to exit the docker containers bash

11. `docker-compose up`

## Now navigate to http://127.0.0.1:3000, and start hacking on the go file!