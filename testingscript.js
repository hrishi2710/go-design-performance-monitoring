import http from 'k6/http';
import grpc from 'k6/net/grpc';
import { sleep } from 'k6';

export const options = {
    vus: 10,
    duration: '30s',
};

const client = new grpc.Client();
client.load(['./proto'], "counter.proto")


export default () => {
    client.connect('localhost:50055', {
        plaintext: true
    });

    const data = { value: 1 };
    const response = client.invoke('counter.incrementCounter/increment', data);

    console.log(JSON.stringify(response.message));

    client.close();
    sleep(1);
};