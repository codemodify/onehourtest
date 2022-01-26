# onehourtest

- `go build`
- `./onehourtest`

# remarks
- the `http://petstore-demo-endpoint.execute-api.com/petstore/pets` seems taken from amazon samples
	- for that reason the POST seems is not creating the objects no matter the combinations
- for the Docker, this seems to be a client app, no running services

# DELETE
- make a POST with the the DELETE HTTP verb on `http://petstore-demo-endpoint.execute-api.com/petstore/pets/{id}`
- headers can contain HTTP signature of the request to be validated on the server
	- for that we need some API/SECRET combination