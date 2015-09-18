# Wisply

![Wisply](http://wisply.me/static/img/wisply/logo/jpg.jpg)


Building the hive of education

## Get Wisply

Wisply can be downloaded using the `git` command through Go

1. Open a terminal window
2. Type `go get github.com/cristian-sima/Wisply`
  * In case you have problems with this step, it means you do not have Git.
  * DigitalOcean has a very good tutorial about how to install Git on Ubuntu [here](https://www.digitalocean.com/community/tutorials/how-to-install-git-on-ubuntu-14-04).
3. Now, you have Wisply. Check the next steps to see how to install it.


## Getting started

We tested Wisply and we ensure you that it works on two main operating systems: Microsoft Windows (10) and Ubuntu. We provide bellow the details of how to set up and install Wisply on your system.  Until you can start Wisply you need to install the language, framework and the server for database. Please follow the next steps.

## Ubuntu

Wisply comes with special tools which help you to quickly set up and run the application.

#### Golang

Wisply is created in Go (or Golang). You can find a lot of information about this language on the [official web page](http://golang.org/). Google even provides a virtual tour for the language [here](https://tour.golang.org/welcome/1)

##### How to install Go?

1. Open a terminal window
2. Type `sudo apt-get install golang`
3. Type `Y` and press enter

##### Configure Golang

1. `GOROOT` is the path where Go is installed. You do not have to change it, but in case you want type `export GOROOT=/usr/local/go`
2. `GOPATH` is the path to the Go workspace. Here Go will download Wisply. So, type `export GOPATH=$HOME/path` where `path` is the path where you want to store Wisply
3. You have to save the changes. Type `export PATH=$PATH:$GOROOT/bin:$GOPATH/bin`

##### Install begoo tool

Wisply is using [beego](http://beego.me/) framework. This framework provides a tool which simplifies the process of maintaining Wisply. In order to install it:

1. Open a terminal window
2. Type `go get github.com/beego/bee`. If the script ends and there is nothing displayed, it means it is working
2. For more information about this tool, please go to [official beego tool website](http://beego.me/docs/install/bee.md)

#### SQL Database

Wisply needs a SQL database in order to store its data. We will show you the steps to set up a database.

##### Install MySQL server

If you already have a MySQL server, skip this step.

1. Connect to your hosting server
2. Open a terminal window
3. In order to get the package, type `sudo apt-get install mysql-server`
4. Type `Y` to confirm
5. You will be asked to enter a password for root user (administrator). Write down or remember the password. I **strongly** recommend you to enter a strong password (at least 8 characters and to contain at least one number or a special character). If you want to generate a strong password, you can use a password generator like [this](https://strongpasswordgenerator.com/).
6. Repeat the password and wait for the system to install MySQL. When it stops, you can see
  `start/running, process 28412`
7. Once the installation is completed, the MySQL server should be started automatically. You can run the following command from a terminal prompt to check whether the MySQL server is running:
  `sudo netstat -tap | grep mysql`
  You will see a list of details regarding the service such as port and process id.


#### Install Wisply

It's very quick to do it.

Wisply comes with a special installer which helps you to quickly set up and run the application. In order to use the installer, make sure you have a MySQL server (see above)

How to use the installer?

1. Open a terminal window
2. Type `cd /the/path/to/Wisply/directory/` where *path/to/Wisply/directory* is the path on your server to the Wisply directory
3. We need to allow the installer to execute. Type `chmod u+x util/ubuntu/install/installer.sh` and type `chmod u+x util/ubuntu/install`
4. In order to run the installer, type `bash /util/ubuntu/installer.sh`
5. Follow the steps shown by installer:

![Installer Example](http://wisply.me/static/img/wisply/example/installer.jpg)

#### Run Wisply

In order to start Wisply, go to Wisply directory and type `bash util/ubuntu/start.sh`. You will see the wizard for running. This script detects if there was any previous version of Wisply started and it tells you.

#### Stop Wisply

If you want to stop Wisply, go to Wisply directory and type `bash util/ubuntu/stop.sh`.

## Windows

Wisply provides several utilities which you can use during the development process. You can find these utilities in the directory `util/windows/Utilities`.

* **Format Wisply** - It can be used to format the code of all the packages. It closes itself after formatting.
* **Start bee go tool** - this script runs the Wisply server and displays a command prompt. Also, it re-builds Wisply after each file has been modified.
* **Test Wisply** - It shows a command promt with the result of testing Wisply. In case the tests were not good, the user can type `y` to re-test Wisply, or type any other character to exit.
*  **Start goconvey** - It is a shortcut to start the [goconvey](https://github.com/smartystreets/goconvey) server. It can be used for a more detailed testing (you need to install this first).
*  **CommitSQL** - When it is executed it exports the entire SQL database scheme (without data) to the file `util/ubuntu/install/src/sql/Wisply.sql`. Thus, it can be used to quickly update the database schema.
* **GenerateJsDoc** - A shortcut to the JSDoc generated which generates the HTML documentation for JavaScript code. (Please edit the `conf.json` and change the paths to yours)


#### Set up utilities

1. Go to directory `util/windows/Utilities`
2. Edit every file and change the path where the Wisply directory is
3. In order to start a script, double click on the file


#### MySQL server

I recommend XAMPP. This software contains the MySQL server and it has a good user interface. You can download XAMPP from [here](https://www.apachefriends.org/download.html).

### Install Wisply

1. Open XAMPP, start Apache and MySQL modules
3. For the MySQL module, click ADMIN.
4. Go to `Databases`
5. Create a new database with the name `wisply`
6. Go to Users and create a user
7. Assign the user all the privileges for the database `wisply`
8. Click on database `wisply`
9. Choose `import`
10. Upload the file from `/util/ubuntu/install/src/sql/Wisply.sql`
11. Update the file `conf/database/default.json` with the details of this database (username, password)
12. Double click `/util/windows/Start bee go tool.bat`
13. You are now running Wisply!


## Update Wisply

Wisply is evolving. You may want to get the last version on your server, but installing each version is time consuming. Thus, in order to quickly update Wisply, please follow the next steps:

1. Open a terminal window
2. Type `go get -u github.com/cristian-sima/Wisply`
3. If there was no message shown, it means the update was successful

The update is not affecting the current configuration (information about database, about server) and does not delete the existing data from database.

**Note**: In case the new version has a different SQL schema, there is a need to install the application again (follow the installer steps, highlighted above)

# Documentation

Wisply considers that documentation is an essential element to maintain an application. The software provides a lot of information regarding documentation.

## Go

Every method or property that is exported is well documented. You can see all the documentation by runnning the `go fmt`
The documentation is according to the official google style. You can find it  [here](	http://blog.golang.org/godoc-documenting-go-code).

## JavaScript

JavaScript has a lot of documentation. The documentation passes the guidelines of [JSDoc](http://usejsdoc.org/). The documentation can be accessed in browser (see bellow). The HTML files have been generated using [JSDoc 3](https://github.com/jsdoc3/jsdoc). The html file are in the directory `util/doc/js`.

**Note!**
* Due to security reasons, the JavaScript documentation *can't* be accessed using the Wisply server (for instance http://wisply.me/doc/js). You can access it by using the absolute path in your browser (C:\path\util\doc\js\index.html for Windows)


## Utilities
The Shell code is documentated according to the guidelines presented by [Inquisitor](http://www.inquisitor.ru/doc/shelldoc.html). Also, the utilities for Windows have comments in order to adapt them to your system


# Tips (optional)

This section contains information regarding the development progress which you may find it useful.

## Editor

If you want to use Golang, but you do not what editor to choose, I recommend (Atom)[]. It is simple, easy to use and fast. Also, I suggest these plugins (the description is the offial one):

* (**atom-pretiffy**)[https://atom.io/packages/atom-prettify] Pretiffy your HTML
* (**auto-indent**)[https://atom.io/packages/auto-indent] This package will allow you to auto-indent your current file. Use the auto-indent:apply command.
* (**docblockr**)[https://atom.io/packages/docblockr] A helper package for writing documentation.
* (**go-plus**)[https://atom.io/packages/go-plus] *It saves you a lot of time* Adds `gocode`, `gofmt`, `goimports`, `go vet`, `golint`, `go build` and `go test` functionality for the go language.
* (**linter**)[https://atom.io/packages/linter] A Base Linter with Cow Powers
* (**linter-jshint**)[https://atom.io/packages/linter-jshint] Linter plugin for JavaScript, using jshint


Do you have a prefered editor or a plugin which I did not mention, please add it.
