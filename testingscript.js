import http from 'k6/http';
import grpc from 'k6/net/grpc';
import { sleep } from 'k6';

const client = new grpc.Client();
client.load(['./proto'], "counter.proto")


export default () => {
    client.connect('localhost:50055', {
        plaintext: true
    });

    const data = { value: 1 };
    const response = client.invoke('counter.incrementCounter/increment', data);

    // check(response, {
    //   'status is OK': (r) => r && r.status === grpc.StatusOK,
    // });

    console.log(JSON.stringify(response.message));

    client.close();
    sleep(1);
};