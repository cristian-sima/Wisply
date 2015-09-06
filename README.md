# Wisply
Using open data to enhance education



# Set up database
Wisply needs a SQL database where to store its data. We will show which are the steps to set up a database and make Wisply working

## Ubuntu

### Install Golang

Wisply is created in Go (or Golang). You can find a lot of information about this language on the [official web page](http://golang.org/). Google even provide a virtual tour of the language [here](https://tour.golang.org/welcome/1)

How to install golang?

1. Open a terminal window
2. Type `sudo apt-get install golang`
3. Type `Y` and press enter

### Install 

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
  
### Install the application

Wisply comes with a special installer which helps you to quickly set up and run the application. In order to use the installer, make sure you have installed the MySQL server (see above)

These steps will help you to use the installer

1. Open a terminal window
2. Type `cd /the/path/to/Wisply/directory/` where *path/to/Wisply/directory* is the path on your server to the Wisply directory
3. We need to allow the installer to execute. Type `chmod u+x install/installer.sh`
4. Type `chmod u+x install` 
4. Type `cd install` to go in the installer directory
5. Type `bash ./installer.sh` in order to run the installer
6. Follow the steps of the installer
