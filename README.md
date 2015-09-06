![Wisply](http://wisply.me/static/img/logo.jpg) # Wisply



Using open data to enhance education


## Use Wisply on Ubuntu

### Golang

Wisply is created in Go (or Golang). You can find a lot of information about this language on the [official web page](http://golang.org/). Google even provides a virtual tour for the language [here](https://tour.golang.org/welcome/1)

#### How to install Go?

1. Open a terminal window
2. Type `sudo apt-get install golang`
3. Type `Y` and press enter


### Install Wisply

#### Get Wisply

Wisply can be downloaded using the `git` command throught Go

1. Open a terminal window
2. Type `go get github.com/cristian-sima/Wisply`
  * In case you have problems with this step, it means you do not have Git. 
  * DigitalOcean has a very good tutorial about how to install Git on Ubuntu [here](https://www.digitalocean.com/community/tutorials/how-to-install-git-on-ubuntu-14-04).
3. You now have Wisply. Check next steps to see how to install it.

#### Install begoo tool

Wisply is using [beego](http://beego.me/) framework. This framework provides a tool which simplifies the process of maintaining Wisply. In order to install it:

1. Open a terminal window
2. Type `go get github.com/beego/bee`. If the script ends and there is nothing displayed, it means it is working
2. For more information go [beego tool website](http://beego.me/docs/install/bee.md)

#### Set up the SQL database

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
  
#### Set up Wisply

It's very quick to do it.

Wisply comes with a special installer which helps you to quickly set up and run the application. In order to use the installer, make sure you have installed the MySQL server (see above)

These steps will help you to use the installer

1. Open a terminal window
2. Type `cd /the/path/to/Wisply/directory/` where *path/to/Wisply/directory* is the path on your server to the Wisply directory
3. We need to allow the installer to execute. Type `chmod u+x install/installer.sh`
4. Type `chmod u+x install` 
4. Type `cd install` to go in the installer directory
5. Type `bash ./installer.sh` in order to run the installer
6. Follow the steps of the installer
