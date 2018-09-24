## Description

More at https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html

### How To Run This Project

> Make Sure you have run the db.sql in your mysql

```bash
#move to directory
cd $GOPATH/src/github.com/adriacidre

# Clone into YOUR $GOPATH/src
git clone https://github.com/adriacidre/go-clean-arch.git

#move to project
cd go-clean-arch

# Install Dependencies
make run
```

# Running tests

Simply run `make test`

## REST actions

You have a helper to this actions on the make file, just run `make` to see the help

**Fetch a resource by id**
`curl http://localhost:9090/payment/1`

**Create a resource**
`curl -d '{"payment_id":"supu","organisation_id":"tupu"}' -H "Content-Type: application/json" -X POST http://localhost:9090/payment`

**Update a resource**
`curl -d '{"payment_id":"supu","organisation_id":"modified"}' -H "Content-Type: application/json" -X PATCH http://localhost:9090/payment`

**List a collection of payment resources**
`curl http://localhost:9090/payment`

**Delete a resource**
`curl -X "DELETE" http://localhost:9090/payment/8`
