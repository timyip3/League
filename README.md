# Interviews from League
It consists of 3 rounds: 1. take-home assessment 2. technical interview 3. architectural interview

## Take-home assessment from League:

To run the code, please execute it with the command "go run ." to run the main method on main.go under the directory League.

To run the functions, please send the request(s) with:
/echo:
        curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"

/invert:
        curl -F 'file=@/path/matrix.csv' "localhost:8080/invert"

/flatten:
        curl -F 'file=@/path/matrix.csv' "localhost:8080/flatten"

/sum:
        curl -F 'file=@/path/matrix.csv' "localhost:8080/sum"
        
/multiply:
        curl -F 'file=@/path/matrix.csv' "localhost:8080/multiply"

## 1st Round Challenge
This session will meet 2 engineers who would ask you questions related to microservices
e.g.: fault-handling on service communication, idempotencies on HTTP methods

## 2nd Round Challenge
This session involves system design based on the question provided before meeting:

```
For the backend architecture challenge, we are asking you to provide a system diagram for a claims submission system. We don’t
expect you to spend more than an hour on this challenge.
We'll start the interview by asking you to show us on your diagram (by sharing your screen) how systems interact to support the claims
submission use cases.
In the bulk of the remaining time we’ll introduce some new requirements or modify the system’s operating environment. We’ll discuss if or
how your initial design would have to change to support that.
Please also write 4-5 sentences explaining how you arrived at your solution. Your description should be concise and speak to the work
performed. This write up will be assessed as part of the overall architecture challenge
The Architecture Challenge: Claims Submission
You are building a system that processes health insurance claims, where people insured under a benefit plan (members) are submitting
a receipt (claim) for a medical expense.
The steps in the system are:
A member uploads an image of the receipt into the system.
OCR (Optical Character Recognition) is run on the image to extract information like the procedure and the cost.
A League Claims Agent reviews the claim.
If approved, money is sent to the person who submitted the claim.
Deliverables
A system diagram that shows system components and how they interact with each other. System components include things
such as databases and services.
Example of a system diagram: https://commons.wikimedia.org/wiki/File:Content_persistence.system_architecture.
diagram.svg
Diagramming tools: diagrams.net, http://excalidraw.com , or a tool of your choice.
A high level data model on what information is stored.
A list of assumptions made.
A list of future considerations, if any.
The above can be stored in a single Google doc and shared with us at least 24 hours before your interview.
System objectives
The primary goal of the system is to review claims and make payouts to the user as fast as possible, regardless of OCR success
or failure. OCR could take several minutes.
You can assume hundreds of claim submissions per hour and dozens of agents available to keep reviewing the claims.
Additional things to consider
How does the design handle scaling?
What are the security risks?
How are errors handled?
How can the system be changed to handle varying requirements?
```