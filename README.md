# go-design-performance-monitoring

### Overview

In this repo, we are trying to introduce one new desing pattern for microservices in golang. In this pattern, we are trying to avoid using locks (mutexes) anywhere in our codebase. We are trying to achieve this using serializing capability of the go channels. We can call this new pattern as `inputQ` pattern. So, overall we are trying to test the performance of `inputQ` pattern vs using `mutex`. 

### How to test?

The repo will host a grpcServer. Which can be called by multiple clients any number of times. Internally, depending on the `mode` in which this server runs, the request will be fulfilled by using one of the two patterns. To simulate the multiple user and concurrent requests, we are using [grafana/k6](https://k6.io/) and it's `VU` features.

### Using k6
* To use k6, I have gone by the docker approach. You can use any approach [listed here](https://k6.io/docs/get-started/installation/)
* We have a `Dockerfile` in the repo. The base image that it is using is that of `Debian`. You can change that accordingly and create the required image by issuing `docker build . -t go_perf_testing`. This will create one image with name `go_perf_testing`.
* Now, to run the container using this image, issue command :- `docker run -dit --rm --name go_perf_testing_container go_perf_testing /bin/bash -c "bash"`. You can also mount the local directory using something like this `docker run -dit --rm -v /home/${USER}:/home/${USER} -w="/home/${USER}" --name go_perf_testing_container go_perf_testing /bin/bash -c "bash"`. Mounting the local directory helps because you can do incremental changes and can test it. 
* Now if you attach to the container named `go_perf_testing_container`, you can confirm that `k6` command is available. You can do `which k6` to confirm that.

### Try it yourself!
* Run the docker container. Attach to the the docker container and follow the below steps.
* We have a `main.go` file which can be run as `go run main.go -mode <mutex/inputQ>`. The main file can be run in one of the two modes. This will kickstart the grpcServer.
* For, client simulation, we have to use k6. 
* The `testingscript.js` can be run using command `k6 run testingscript.js`. This will simulate the multiple client and give you the test results at the end.