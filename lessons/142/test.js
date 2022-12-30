import http from 'k6/http';
import { sleep, group, check } from 'k6';

export const options = {
    stages: [
        { target: 20, duration: '20m' },
    ]
};

const url = 'http://localhost:8080/api/devices'

export default () => {

    group('Send GET requests', () => {
        const res = http.get(url);
        const checkRes = check(res, {
            'status is 200': (r) => r.status === 200,
        });
    });

    sleep(1);
};