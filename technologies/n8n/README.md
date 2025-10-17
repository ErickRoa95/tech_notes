## What's n8n ? 

n8n is an open-source, workflow automation platform that allows users to connect different apps and automate tasks through a visual, node-based editor. It is a low-code tool, meaning you can build complex automations without extensive coding, though it also supports JavaScript for more advanced customization.

## What's a n8n workflow? 

An n8n workflow is a sequence of automated tasks created using n8n, a visual, node-based tool for connecting different applications and services. You build an n8n workflow by visually linking "nodes," which represent specific actions like sending an email or updating a spreadsheet, to create a chain of automation that is triggered by a specific event.

## What's a node ?
You can think of node as abuiliding blocks that serve different functions that, when put together, make up a functioning machin: **an automated workflow**. 

Types of nodes:
- **Action  Nodes**: Add, remove and edit data; request and send external data; and trigger events in other systems.
- **Trigger Nodes**: Start a workflow and supply the initial data.
- **Core Nodes**: Provide functionality such as logic scheduling or generic API calls.
- **Cluster Nodes**: Node groups that work together to provide functionality in a workflow.

## Data structure of n8n
n8n nodes function as an **Extract, Transform, Load(ETL)** tool. The nodes allow you to **access(extract)** data from multiples disparate sources, **modify(transform)** the data in particular way, and **pass(load)** it along to where it needs to be.  

The data that moves along from note to node in your wofklow must be in a format<structure> that can be recognized and interpreted by each node. In n8n, this required structure is an array of objects. 

Data sent from one node to another is send as an array of JSON objects. The elements in the collaction are called items. 

Example of returning object:
```
return [
  {
    json: {
      data: "some data"
    }
  }
]
```

You want to use the following expression to get access to incoming JSON formatted from another node:
```
# $json is a custom variable provided by n8n, to handle incoming json data. 
{{ $json.<parameter> }}
```

## Self-host n8n plataform with Docker.
On the official [n8n documentation](https://docs.n8n.io/hosting/installation/docker/#starting-n8n), provided docker-compose file to locally build and run n8n plataform using postgresql to save n8n credentials. 
- [Docker file](https://gist.github.com/ErickRoa95/e0bb4151140ff6d85ac5bfe8613ebf75) to start n8n platform.
- [.env template](https://gist.github.com/ErickRoa95/e0bb4151140ff6d85ac5bfe8613ebf75) file for database credentials.

## Example workflows

> **NOTE**: All the files saved on workflows folders are JSON configuration files. This file can be imported to n8n plataform.

- [Merge Workflow](./workflows/Merge_Workflow.json) : Workflow that showcase the merge action of 2 node's outputs into 1 Node. 
- [LoopBatch workflow](./workflows/LoopBatch_Workflow.json): Workflow that showcase the Loop/Batch action, to handle multiple batches of information. 
- [DataSet Workflow](./workflows/DataSet_workflow.json): Workflow that showcase the process of adding a new data/value into the format of a JSON. 
- [BinaryFile Workflow](./workflows/BinaryFile_Workflow.json): Workflow that showcase the behavior of Files in n8n. 
- [ETL Workflow](./workflows/Pokemon_ETL_workflow.json): Workflow that showcase the Steps of ETL: Extract information from Get Endpoint, Transform information to required format and Load/showcase the end result of the information. 
- [Customer's report Workflow](./workflows/Nathan_workflow.json): Real life scenarios, where you consumed a real database with credentials and use IF nodes to handle scenarios to format and split information.

## Notes
This information and examples where obtained directly from the official documentation from n8n platform. 
