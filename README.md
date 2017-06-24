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

9. `\q` to exit mysql cli, then `exit` to exit the docker containers bash

10. `docker-compose up`

# Now navigate to http://127.0.0.1:3000, and start hacking on the go file!