# Wisply
Using open data to enhance education



# Set up database
Wisply needs a SQL database where to store its data. We will show which are the steps to set up a database and make Wisply working

## Ubuntu


### Install MySQL server

1. Connect to hosting server and open a terminal window
2. Type `sudo apt-get install mysql-server`. This command will get the SQL package
3. Type `Y` to confirm the installation
4. You will be asked to enter a password for root user (administrator). Write down or remember the password. I STRONGLY recommend you to enter a strong one. You can use a password generator like [this](https://strongpasswordgenerator.com/).
5. Repeat the password and wait for the system to install MySQL. When it stops, you can see
  start/running, process 28412
6. Once the installation is complete, the MySQL server should be started automatically. You can run the following command from a terminal prompt to check whether the MySQL server is running:
  `sudo netstat -tap | grep mysql`
  You will see a list of details regarding the service such as port, process id
  
### Set up database

1. Now you have to create the databse.
