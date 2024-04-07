# AlexaRecetardasPiolas

This project contains the source code and dataset for the Recetardas Piolas skill for Amazon Alexa

## Infrastructure

The project has been developed with the following architecture:

### Amazon Lambda Function

This Amazon service provides a cloud execution of the Go code

### Amazon Alexa Skill

The recipeSkill developed with Alexa Skill Kit

### MongoDB Atlas Cluster

The database is hosted in this service, allowing everyone to connect and use the dataset

## Dataset

First of all, in the dataset folder you will find a simple Python script to upload the data into the database.
It will read files in JSON format and upload to the Atlas cluster.
The dataset used was extracted from https://huggingface.co/datasets/Frorozcol/recetas-cocina/viewer/default/train

## Source Code

The source code was developed in Golang, using the official MongoDB Driver and ASK driver from https://pkg.go.dev/github.com/arienmalec/alexa-go@v0.0.0-20181025212142-975687393e90
Don't blame my code, it was my first experience <3
