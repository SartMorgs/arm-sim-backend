# Armsim Backend

This repository was prepared to build backend block for simulator of arm cortex M0, which intend to provider a JSON out file with all step's executions of a program provided into some request. This repository are a block of another bigger project called [arm-sim](http://github.com/arm-sim).

The content of this abstract are:

- How to run it by yourself
- How it works
- How to report bugs and errors
- How to contribute

## How to run it by yourself

You will need docker to run it by yourself to guarantee that it will run like a base environemtn where it was builded. You chan check [this link](https://docs.docker.com/engine/install/) to see how to install and configure docker for your machine.

To run the docker environment you'll need to build a image configured inside of a dockerfile, and after this run this image generating a docker container. For build the image you need to run the bellow command:

```bash
docker build -f arm_backend_img -t docker/Dockerfile .
```

To run the image that you builded at last command you will use:

```bash
docker run --name arm_backend -p 9000:9000 -d arm_backend_img
```

Now you can make requests for http://localhost:9000 with assembly code into body to receive execution's steps in json format. The format of request is like this:

~IMG~

The response format is like this:

```json
teste
```

You can test requests and responses using postman or another way to make api requests.

## How it works

There is a little bit complex backend system, which was developed using goland and builded like most common backend systems: using POO. TDD was aplied in development process. So, bellow you can see class diagram for this system.

~IMG~

Basically, all intern components was implemented as a goland package in backend development and can be usable and testable separatebly.

## How to report bugs and errors

In case of bugs and problems with arm-sim execution, there can be opened a github's issue in this repository describing clarery the situation. You can find some advices to open a good issue.
