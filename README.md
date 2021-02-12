liveproject-create-a-web-application-scaffold-generator-in-go
---


# Manning - liveProject - Create a Web Application Scaffold Generator in Go


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


# Project implementation


## Setup Env


## Create github repository


## Create project in GoLand

### Link the project to the github repository


### Initiate a module main in the project
From terminal, in the project directory, execute the command:

    go mod init module scaffold_gen

### Compile the app and get the executable

From terminal, in the project directory, execute the command:

    go build -o scaffold_gen

If you have not previously created the module with the 
"go mod init scaffold_gen" command you will get an error message like:

    go: cannot find main module, but found .git/config in /Users/admin/work/go/Manning/liveProject/liveproject-create-a-web-application-scaffold-generator-in-go
        to create a module there, run:
        go mod init

### Test

From terminal, in the project directory, execute the command:

    go test

The output should report at the end something like:

    PASS
    ok      scaffold_gen    0.214s

To have more info on the tests execution, add -v as parameter:

    go test -v


## Milestone 1 - Use the Standard Library Packages to Create a Command-Line Application



## Milestone 2 - Create the Web Applications to Be Used as the Scaffold

In the main project directory there are two subdirectory containing 
the applications for both the static and the handler web systems:

- handler
- static

In both subdirectories you can build the relative application with the "go build" 
command and test it with the "go test" one.



## Milestone 3 - Update the Command-Line Application to Generate the Web Application Scaffold

### Objective

In this milestone, you will update the command-line application you wrote in Milestone 1 to generate the scaffold web application. The key work you will do in this milestone is to insert the templated versions of the web applications you constructed in Milestone 2 as strings in your code, and then render the customized templates to files (See Notes for an example).

### Workflow

Create the following functions in the main package:
- Write a new function to create the directory structure for the web applications. 
  For both categories of web applications, you will create the handlers subdirectory 
  inside the user-specified location. For the web application for the website, 
  create an additional subdirectory static with two subdirectories inside it, 
  js and css. Let’s call this function createScaffoldDirs.
- Write a new function to create and generate the following files: go.mod, server.go, handlers/setup.go, handlers/index.go, and handlers/healthcheck.go with the rendered contents. This step will be common for both the web application categories. Let’s call this function generateScaffoldCommon().
- Write a new function to generate the following files: static/js/index.js and static/css/styles.css. Let’s call this function generateScaffoldStatic().

Update the generateScaffold() function you wrote in Milestone 1 to call the functions createScaffoldDirs() and generateScaffoldCommon(). If the scaffold is being generated for a website, call the generateScaffoldStatic() function

### Deliverable

At the end of this milestone, you will have a command-line application which will generate a customized web application scaffold that you created in Milestone 2.

Assume that scaffold-gen is the name of the built application. First, we will not specify the -s flag:

    $ ./scaffold-gen -d ./project1 -n Project1 -r github.com/username/project1
    Generating scaffold for project Project1 in ./project1

Then, by navigating to the ./project1 directory, you should be able to run the web application using go run server.go. In a separate terminal/command-line window, use cURL to make requests to the web application:

    $ curl localhost:8080
    Hello, world!

HealthCheck call:

    $ curl localhost:8080/healthcheck
    ok

When a -s is specified, it will generate the scaffold for a website:

    $ ./scaffold-gen -d ./project2 -n Project2 -r github.com/username/project2  -s

    Generating scaffold for project Project2 in ./project2

Then, by navigating to the ./project2 directory, you should be able to run the web application using go run server.go. In a separate terminal/command-line window, use cURL to make requests to the web application:

    $ curl localhost:8080
    <html> <head> <title>This is the homepage for project MyProject2</title><link rel="stylesheet" href="static/css/styles.css" />&lt;script>&lt;/script></head><body> <h1>Hello, World!</h1> </body></html> 

HealthCheck call:

    $ curl localhost:8080/healthcheck
    ok



## Milestone 4 - Write End-to-End Tests

### Objective

In this milestone, you update the tests from Milestone 1 and add new ones to test the scaffold generated by the web application. Verify that the web application scaffold matches the expected directory structure on the filesystem, and that the different files have the correct contents.

### Workflow

Update the test for scaffold generation to check whether the directory structure as well as the files are created.

Add a new test for checking if the scaffold generated contains the expected contents. Instead of writing two tests for the scaffold with static assets and without static tests, I recommend only writing one test which verifies whether the directory structure and contents are as expected.

### Deliverable

Copy and paste the link in the field below under the heading "submit your work”. Once you save it, the Author’s solution and peer solutions will appear on the page for you to examine.

The next Milestone link in the ToC will be activated and you will be able to move on.

Note: if you need to, you can change your link. Just click on the white pencil icon in the blue circle and put in a new one.

Below you will find a link to the sample solution as a zip file. Your solution can look different. 

Your tests must check the following:

- Scaffold directory structure is created correctly for both website and API.
- The contents of the generated files match the expected data.

In addition, you must ensure that any temporary directories you create are automatically cleaned up after the test execution finishes.

