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


## Performance comparison results:

**Test case:-** 10 Virtual users (VUs) for 30 seconds

Parameters          |     InputQ pattern                                                           |        Mutex pattern
-------------       |---------------------------                                                   |---------------------
data_received       |23 MB  754 kB/s                                                               |24 MB  798 kB/s
data_sent           |29 MB  956 kB/s                                                               |30 MB  1.0 MB/s
grpc_req_duration   |avg=719.94µs min=184.45µs med=687.39µs max=10.71ms  p(90)=1.16ms p(95)=1.37ms |avg=712.01µs min=179.46µs med=672.36µs max=9.41ms p(90)=1.15ms p(95)=1.36ms
iteration_duration  |avg=2.19ms   min=768.88µs med=1.97ms   max=781.69ms p(90)=2.89ms p(95)=3.34ms |avg=2.07ms   min=766.71µs med=1.97ms   max=13.3ms p(90)=2.85ms p(95)=3.29ms
iterations          |135980 4532.286646/s                                                          |144028 4800.542003/s


**Test case:-** 100 Virtual users (VUs) for 20 seconds

Parameters          |     InputQ pattern                                                           |        Mutex pattern
-------------       |---------------------------                                                   |---------------------
data_received       |36 MB  1.8 MB/s                                                               |32 MB  1.6 MB/s
data_sent           |45 MB  2.3 MB/s                                                               |41 MB  2.0 MB/s
grpc_req_duration   |avg=1.63ms min=188.12µs med=791.59µs max=27.96ms  p(90)=4.19ms  p(95)=5.81ms  |avg=1.63ms  min=179.07µs med=778.16µs max=25.06ms  p(90)=4.23ms  p(95)=5.89ms
iteration_duration  |avg=9.35ms min=863.84µs med=8.25ms   max=164.59ms p(90)=15.38ms p(95)=17.92ms |avg=10.33ms min=883.24µs med=8.17ms   max=300.16ms p(90)=16.03ms p(95)=19.01ms
iterations          |213401 10664.646039/s                                                         |193205 9656.439518/s

**Test case:-** 200 Virtual users (VUs) for 20 seconds

Parameters          |     InputQ pattern                                                           |        Mutex pattern
-------------       |---------------------------                                                   |---------------------
data_received       |37 MB  1.7 MB/s                                                               |38 MB  1.8 MB/s
data_sent           |46 MB  2.2 MB/s                                                               |48 MB  2.3 MB/s
grpc_req_duration   |avg=1.88ms  min=186.67µs med=820.17µs max=31.36ms p(90)=5ms     p(95)=7.32ms  |avg=1.82ms min=195.28µs med=812.78µs max=46.31ms p(90)=4.76ms p(95)=6.94ms
iteration_duration  |avg=18.42ms min=940.36µs med=11.26ms  max=3.06s   p(90)=21.28ms p(95)=24.98ms |avg=17.8ms min=864.02µs med=11.16ms  max=3.05s   p(90)=20.6ms p(95)=24.09ms
iterations          |218389 10398.382496/s                                                         |226902 10809.455794/s

**Test case:-** 300 Virtual users (VUs) for 20 seconds

Parameters          |     InputQ pattern                                                           |        Mutex pattern
-------------       |---------------------------                                                   |---------------------
data_received       |37 MB  1.7 MB/s                                                               |37 MB  1.8 MB/s
data_sent           |47 MB  2.1 MB/s                                                               |46 MB  2.2 MB/s
grpc_req_duration   |avg=1.92ms  min=193.8µs  med=824.72µs max=40.62ms p(90)=5.03ms  p(95)=7.58ms  |avg=1.87ms  min=190.44µs med=817.61µs max=70.83ms p(90)=4.77ms  p(95)=7.32ms
iteration_duration  |avg=27.03ms min=946.79µs med=11.55ms  max=7.27s   p(90)=21.99ms p(95)=26.5ms  |avg=27.77ms min=945.71µs med=11.74ms  max=7.19s   p(90)=21.96ms p(95)=26.78ms
iterations          |223305 10109.286962/s                                                         |218096 10453.813865/s

**Test case:-** 500 Virtual users (VUs) for 20 seconds

Parameters          |     InputQ pattern                                                           |        Mutex pattern
-------------       |---------------------------                                                   |---------------------
data_received       |37 MB  1.7 MB/s                                                               |37 MB  1.7 MB/s
data_sent           |47 MB  2.1 MB/s                                                               |47 MB  2.2 MB/s
grpc_req_duration   |avg=1.88ms  min=192.69µs med=799.54µs max=57.02ms p(90)=4.75ms  p(95)=7.33ms  |avg=1.83ms  min=185.09µs med=756.21µs max=60.66ms p(90)=4.57ms p(95)=7.22ms
iteration_duration  |avg=46.17ms min=954.63µs med=12.08ms  max=7.19s   p(90)=23.86ms p(95)=31.52ms |avg=45.75ms min=998.44µs med=11.82ms  max=7.21s   p(90)=24ms   p(95)=32.36ms
iterations          |221074 10160.735101/s                                                         |220961 10266.287413/s

**Test case:-** 1000 Virtual users (VUs) for 20 seconds

Parameters          |     InputQ pattern                                                           |        Mutex pattern
-------------       |---------------------------                                                   |---------------------
data_received       |37 MB  1.5 MB/s                                                               |36 MB  1.6 MB/s
data_sent           |47 MB  1.9 MB/s                                                               |47 MB  2.0 MB/s
grpc_req_duration   |avg=2.31ms  min=178.27µs med=1.06ms  max=485.17ms p(90)=5.77ms p(95)=8.81ms   |avg=2.27ms  min=185.06µs med=1.05ms  max=454.45ms p(90)=5.5ms   p(95)=8.39ms
iteration_duration  |avg=91.98ms min=950.8µs  med=12.93ms max=15.44s   p(90)=26.4ms p(95)=1.02s    |avg=93.12ms min=952.92µs med=12.76ms max=21.1s    p(90)=26.23ms p(95)=1.01s
iterations          |223862 8924.532635/s                                                          |220310 9430.541952/s
