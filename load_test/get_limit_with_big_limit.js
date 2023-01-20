import http from 'k6/http';
import { sleep } from 'k6';

const VUS = 1000;

export const options = {
    stages: [
        { duration: '30s', target: VUS },
        { duration: '30s', target: 0 },
    ]
}

export default function scenarioFunc() {
    http.get(`http://localhost/article?limit=10000`);
    sleep(1);
}
