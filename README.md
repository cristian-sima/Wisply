# Wisply

![Wisply](http://wisply.me/static/img/logo.jpg) 


Building the hive of education

## Get Wisply

Wisply can be downloaded using the `git` command throught Go

1. Open a terminal window
2. Type `go get github.com/cristian-sima/Wisply`
  * In case you have problems with this step, it means you do not have Git. 
  * DigitalOcean has a very good tutorial about how to install Git on Ubuntu [here](https://www.digitalocean.com/community/tutorials/how-to-install-git-on-ubuntu-14-04).
3. You now have Wisply. Check next steps to see how to install it.


## Gettings started

Wisply has been tested and it can work on two main operating systems: Microsoft Windows (10) and Ubuntu. We provide the details for how to set up and install Wisply.

## Ubuntu

Wisply comes with a special installer which helps you to quickly set up and run the application. Until we can run the installer there must be installed the language, framework and the server for database. Please follow the next steps.

#### Golang

Wisply is created in Go (or Golang). You can find a lot of information about this language on the [official web page](http://golang.org/). Google even provides a virtual tour for the language [here](https://tour.golang.org/welcome/1)

##### How to install Go?

1. Open a terminal window
2. Type `sudo apt-get install golang`
3. Type `Y` and press enter

##### Configure Golang

1. `GOROOT` is the path where Go is installed. You do not have to change it, but in case you want type `export GOROOT=/usr/local/go`
2. `GOPATH` is the path to the Go workspace. Here Go will download Wisply. So, type `export GOPATH=$HOME/path` where `path` is the path where you want to store Wisply
3. You now have to save the changes. Type `export PATH=$PATH:$GOROOT/bin:$GOPATH/bin`

##### Install begoo tool

Wisply is using [beego](http://beego.me/) framework. This framework provides a tool which simplifies the process of maintaining Wisply. In order to install it:

1. Open a terminal window
2. Type `go get github.com/beego/bee`. If the script ends and there is nothing displayed, it means it is working
2. For more information go [beego tool website](http://beego.me/docs/install/bee.md)

#### SQL Database

Wisply needs a SQL database where to store its data. We will show which are the steps to set up a database and make Wisply working

##### Install MySQL server

If you already have MySQL, skip this step.

1. Connect to hosting server and open a terminal window
2. Type `sudo apt-get install mysql-server`. This command will get the SQL package
3. Type `Y` to confirm the installation
4. You will be asked to enter a password for root user (administrator). Write down or remember the password. I STRONGLY recommend you to enter a strong one. You can use a password generator like [this](https://strongpasswordgenerator.com/).
5. Repeat the password and wait for the system to install MySQL. When it stops, you can see
  `start/running, process 28412`
6. Once the installation is complete, the MySQL server should be started automatically. You can run the following command from a terminal prompt to check whether the MySQL server is running:
  `sudo netstat -tap | grep mysql`
  You will see a list of details regarding the service such as port, process id
  

#### Install Wisply

It's very quick to do it.

Wisply comes with a special installer which helps you to quickly set up and run the application. In order to use the installer, make sure you have installed the MySQL server (see above)

These steps will help you to use the installer

1. Open a terminal window
2. Type `cd /the/path/to/Wisply/directory/` where *path/to/Wisply/directory* is the path on your server to the Wisply directory
3. We need to allow the installer to execute. Type `chmod u+x util/ubuntu/install/installer.sh`
4. Type `chmod u+x util/ubuntu/install` 
5. In order to run the installer, typ `bash /util/ubuntu/installer.sh`
6. Follow the steps of the installer

#### Run Wisply

In order to start wisply, go to Wisply directory and type `bash util/ubuntu/start.sh`. You will see the wizard for running. This script detects if there was any previous version of Wisply started.


#### Stop Wisply

If you want to stop wisply, go to Wisply directory and type `bash util/ubuntu/stop.sh`. 

## Windows

Wisply provides several utilities which you can use during the development process. You can find these utilities in the directory `util/windows/Utilities`.

* **Format Wisply** - it can be used to format all the packages of Wisply. It closes itself after formating
* **Start bee go tool** - this script runs the Wisply server and displays a command promt with live log. Also, it re-builds Wisply after each modification
* **Test Wisply** - It shows a command promt with the result of testing Wisply. In case the test were not good, the user can type `y` to re-test Wisply, or type any other character to exit
*  **Start goconvey** - It is a shorcut to start the [goconvey](https://github.com/smartystreets/goconvey) server. It can be used for a more detailed testing (you need to install this first)
*  **CommitSQL** - When it is executed it exports the entire SQL database scheme (without data) to the file `util/ubuntu/install/src/sql/Wisply.sql`. Thus, it can be used to quickly update database

#### Set up utilities

1. Go to directory `util/windows/Utilities` 
2. Edit every file and change the path to the one where Wisply is located
3. In order to start them, just double click on the file


#### MySQL server

I recommend XAMPP. This software contain the MySQL server and it has a good user interface. You can download XAMPP from [here](https://www.apachefriends.org/download.html). 

### Set up Wisply

1. Open XAMPP, start Apache and MySQL modules
3. For the MySQL module, click ADMIN. 
4. Go to `Databases`
5. Create a new database with the name `wisply`
6. Go to Users and create a user
7. Assign the user all the priviledges for the database `wisply`
8. Click on database `wisply`
9. Choose `import`
10. Upload the file from `/util/ubuntu/install/src/sql/Wisply.sql`
11. Update the file `conf/database/default.json` with the details of this databse (username, password)
12. Double click `/util/windows/Start bee go tool.bat`
13. You are now running Wisply!
