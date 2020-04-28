todo
====
Simple demonstration of a Todo List using Golang, MySQL and Angular

# How to Install
These instructions assume that you've already created a useable MySQL database for your application, along with having the required credentials. If you do not need a database, you can ignore the database credentials or set them as placeholders for later. If you need help on creating a database, you can [learn how, here](https://dev.mysql.com/doc/mysql-installation-excerpt/5.7/en/).

* Download the git repository
    * `git clone https://github.com/jesse-greathouse/todo`
* Change to the todo directory
    * `cd todo`

## Install Golang
To get the detailed instructions, on how to install golang, [check here](https://golang.org/dl/).

## Installing the App
`$ bin/install.sh`

## Configuring the app
`$ bin/configure.sh`
    -- or --
`$ bin/configure-osx.sh` -- if you're on macOS (sed works differently on BSD based systems :-/)

This is an interactive script that will prompt you for all of the required configuration strings to run the app.

## Run the app
`$ bin/run.sh`

The configuration script creates the run script which will run the application with the required configuration.

## Check it in your web browser
At this point you should be able to check the application online at the port you specified.
