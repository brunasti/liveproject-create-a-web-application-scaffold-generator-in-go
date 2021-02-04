# liveproject-create-a-web-application-scaffold-generator-in-go
---


Manning - liveProject - Create a Web Application Scaffold Generator in Go


## Objective

Create a command-line application which will parse command-line arguments specified by the user. These command-line arguments will allow the user to configure the project name, location on local filesystem, a repository URL, and whether the project will be performing the role of an API backend or a website.

## Workflow

- Create a new directory and a file main.go inside it.

- Use the flag package to set up flags to accept the project name (-n), location on disk (-d), remote repository URL (-r), and whether the project will only be for an HTTP API backend or a website (-s). The -s flag is optional and should default to a Boolean value, false.

- Parse the input and create a structure to store that information.

- Write unit tests to ensure that passing in the right flags creates the configuration structure correctly and displays a message to the user saying that the project scaffold will be created.

- Write unit tests to ensure that passing in invalid flags returns an error to the user.

- Write unit tests to ensure that not specifying all the required data will return an error message to the user.

## Deliverable

At the end of this milestone, you will have a command-line application which should behave as follows. Assume that scaffold-gen is the name of the built application (go build -o scaffold-gen); passing the -h option should print a help text:

    $ ./application -h
    Usage of scaffold-gen:
        -d string
            Project location on disk
        -n string
            Project name
        -r string
            Project remote repository URL
        -s    Project will have static assets or not

If you specify all the arguments, it will print a message like the following:

    $ ./scaffold-gen -d ./project1 -n Project1 -r github.com/username/project1
    Generating scaffold for project Project1 in ./project1

If any of the required options are missing, it will print an error message:

    $ ./scaffold-gen -n Project1         
    Project path cannot be empty
    Project repository URL cannot be empty

## Importance to project

In this milestone, you lay the foundations of the command-line application. You set up the command-line interface that the user will interact with, and add appropriate validation to ensure that the user is informed when they have not specified any required data. You will write code which uses the Input/Output interface, and write tests which can test the program’s standard output.

In Milestone 3, you will extend this application to generate the web application scaffold that you will create in Milestone 2.

