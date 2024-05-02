# Work task Backend Engineering [edit]
Hello Martin, and welcome to our Yaak Backend Engineering challenge!

## Background

The purpose of this challenge is to evaluate how you approach and solve a problem. While we want to see how you work hands-on within your field, interested in all artifacts you produce, it is just as important for us to understand your way of thinking and the steps you take to complete the task. While code speaks words, please document it where it's worthwile.

Please don’t spend more than 2 hours on the challenge. There is no hard time limit, but don’t spend too long on this task; definitely no more than a handful of hours in total. If you hit the time limit just commit what you have until that point and follow step 5, we’ll get back to you in every case!

## Your Challenge

We are providing you with some moving parts and you are required to fill in the blanks in between:

You should find a baseline setup of a golang service running
a mini API on port `9595` check it out locally and compile it e.g. via `go build -o server ./cmd/server && ./server`, one test is also added.

### Task 1
* Add an endpoint `POST /task/add` to the service which adds up two integers provided in a JSON
  in the form `{"a": 1, "b": 5}` and returns `{"result": n}` and a `HTTP 200 OK`.
* Add some basic error handling for missing arguments or malformed json with appropriate HTTP response codes.
* Add one test to verify the baseline functionality of your endpoint

### Task 2
* Add an endpoint `/task/drive/:id` which takes in a `GET` request where `:id` is any integer.
* Respond to all requests with a JSON response in this format
```json
{"id": "<id-from-the-request>", "route": "<see-below>" }
``` 
* When receiving a request, synchronously ask a live service on
```
https://q9bd0ylk42.execute-api.eu-central-1.amazonaws.com/sandbox/worktask/mobile/route
```
for a JSON response and add it to the initial response as well, grouped under a field called `route`.
(To reduce the scope, do not fully parse the response into a struct or similar with individual fields).
Note that the live service is flaky and might return `429`s, in this case please retry a sensible number of times until you get a `200 OK`.

Authentication needed to access the work task endpoint mentioned above (provide the credentials in the header):
```
Key: x-api-key
Value: KzCwHmiE5S7JfmPfOHgd5aRNqAMIuvCj7xHY2Ugx

Key: x-worktask-auth
Value: 5isLxi6bw5nJXt5bVCvtWvCkLeb8ZPs3
```

* Add a very simple test, mocking the downstream services behavior.
* Dockerize the setup: Create a (minimal) Dockerfile which runs the service we provided and you extended, and expose port `9696` locally for taking in requests.

You can send us your challenge back or create a repository with a PR and let us know when you are done.
Lean back, have a drink of your choice and wait for us to get back to you :)

### Things to keep in mind
Code quality exceeds speed, write your code in a way that it is easily understandable by the next coder reading it.

### Questions we might try to answer together during the interview
* When taking in responses from the service and the response fails with a 429 specifically - what would be an important thing to consider?
* When we are dealing with 10 instead of 2 downstream services, what can we do architecturally to ease handling all the services?
* What would be your initial approach to properly monitor the service and to catch elevated error rates from the downstream services?


