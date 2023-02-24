import http from 'k6/http';
import { check } from 'k6'

export const options = {
  duration: '1m',
  vus: 30,

  thresholds: {
    http_req_duration: ['p(95)<500'],
  },
};

export default function () {
  const res = http.get('http://server:8080');

  check(res, {
    'is status 200': (r) => r.status === 200
  })
}
