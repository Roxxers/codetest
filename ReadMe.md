# **REDACTED** Backend Coding Challenge

## Purpose

The purpose of this exercise is to produce a body of work that can be evaluated, reviewed and discussed as a practical part of our job interview process.

The code will be compiled, evaluated and code reviewed by members of the **REDACTED** development team before the interview and a semi-formal code review will make-up at least part of the first interview.

While the goal of the exercise is to have a fairly substantial piece of code to evaluate and discuss in an interview, the task itself is intended to show your ability to build a microservice, so we will be looking at how you structure your code, incorporating things like logging, configuration and error handling.


## Overview

Our repo contains folders for two Go microservices which communicate with each other.

1. One service is a Watcher Node: each instance of this service watches a folder and keeps track of the list of files within it. There will be multiple instances of these servers launched, each watching a different folder. We are providing an implementation of this – you are free to modify it if you wish, such as to fit with your Aggregation Server implementation, but we don’t expect you to do so, and we won’t automatically rate a submission more or less highly based on whether you choose to.

1. The other service is an Aggregation Server that communicates with all of the Watcher Nodes and, in turn, produces an aggregated, sorted list of files in all the folders, on request, in JSON over HTTP. We want you to write this service.

We will query the Aggregation Server directly over HTTP from a browser to test how it responds as files are added, removed or otherwise changed.

 
### Watcher Nodes

* Each Watcher Node will watch a single folder and keep an up-to-date list of the files within it
* Our provided Watcher Node spots changes within the folder quickly, using filesystem APIs.
* Watcher Nodes promptly report any changes in the list to the Aggregation Server, by sending an HTTP PATCH request
* Watcher Nodes respond to an HTTP GET request to return their current list
* Watcher Nodes are configured via command line flags. These are documented in the Watcher Node's readme file.
* Each Watcher Node will notify the file aggregator that it exists by sending a `hello` message to the Aggregation Server address after it starts up, and periodically thereafter. It will also send a `bye` message when it shuts down.
* To reiterate – you are free to modify the code as you wish, but are not expected to do so.

### Aggregation Server
* The Aggregation Server should communicate with the Watcher Nodes to receive updated lists of the files in each folder
* The Aggregation Server should keep a single, sorted list of files
* The Aggregation Server will accept an HTTP request and return a simple JSON file representing the full list of files 
  * It should be consistent with the following:

`GET http://localhost:12345/files/`

```
{"files":[{"filename":"cat.jpg"},{"filename":"dog.jpg"}]}
```

## Client

No client is required – the JSON data will be queried and displayed directly in a browser

## Code

* The server should be written in Go
* You can use any libraries you wish but we advise you to stick to ones you know well and have experience of using before
* You can make any coding decisions you wish, however you should be ready to discuss (and possibly justify) those decisions in a code review
* The Watcher implementation we provide implies a certain architecture – pulling initial data and then processing pushed changes over HTTP – you are free to change this if you wish, but you are not expected to do so
* You will be judged on the quality of your code, including its correctness, handling of race conditions, error checking, robustness and maintainability. You will not be judged on the prettiness of output nor on extraneous extra code.
* If you find a bug in the Watcher, please let us know, and feel free to fix it if it gets in your way – but this is not a hidden challenge, and we will not be judging you on anything that you do or do not find.

## Timekeeping

You will be given an agreed length of time to complete the task after a quick discussion with **REDACTED**. We are hoping that this is a task that can be completed in a few evenings in a week. We will not be judging you on your ability, or keenness, to work in your spare time so we will look to find a suitable timeline for you to find those hours.

## Keeping in Touch

This is not an exam so you do not have to complete the task in isolation or silence, and we actively encourage you to get in touch to discuss ideas or technical designs before spending a lot of time implementing them. Software engineering is a team effort, not done in a vacuum, and part of that is making use of the expertise of other team members. You will have access to a Slack channel with a developer from **REDACTED** to discuss with.

## Notes

* This task is designed to give you an opportunity to demonstrate code structure, concurrency and clarity, and has been left quite open-ended so that if there is an ability that you would like to show us, you have the flexibility to do so.
* By all means, use any framework or tool you like – do beware of large, complex frameworks if you don’t already know them well, however, as they can take more time and effort to master than they give back for a relatively small project.
* As mentioned earlier, there is an implied architecture from the Watcher implementation that we provide. You’re free to change that as you wish, but we would suggest that you keep it to minor amendments unless it is an important part of what you want to show us.
* You can add extra fields to the JSON if it helps you debug your own code (e.g. timestamps). 
* We don’t expect perfection in the time you have allocated; rather we will be looking for sensible, pragmatic decisions.
* To reiterate: please talk to us about your thoughts and ideas, and ask for help if you get stuck. We want to see how you could contribute in real life, not to waste a lot of time banging your head against a metaphorical wall!


## Makefile

For ease of testing, the included makefile runs two watcher-node instances, each instance watches one of the folders in the directory `sample-folders`.

### To run
`make run`

This will launch watchers on ports `4000` and `4001`, which attempt to connect to an aggregator at `http://127.0.0.1:8000`

To pass a different aggregator address, you can call:
`make AGGREGATOR_ADDRESS=http://host.name:port run`

### To stop
`make stop`

### Adding your aggregation server
You do not have to support this via the makefile at all; if you find it convenient you may either add it as a new make target or extend the existing run/stop definitions. This is completely up to you, and is not being assessed.


