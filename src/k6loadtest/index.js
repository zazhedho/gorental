const http = require('k6/http')
const metrics = require('k6/metrics')
const k6 = require('k6')

const myFailRate = new metrics.Rate('failed request')

module.exports.options = {
  stages: [
    { duration: '10s', target: 20 },
    { duration: '20s', target: 100 },
    { duration: '20s', target: 0 }
  ]
}

module.exports.default = function () {
  const res = http.get('https://go-vehiclerent.herokuapp.com/api/v1/vehicles')
  const checkRes = k6.check(res, {
    'status was 200': function (r) {
      return r.status == 200
    }
  })
  if (!checkRes) {
    myFailRate.add(1)
  }
  k6.sleep(1)
}
