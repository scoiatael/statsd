# statsd

## Get from source (or download from releases)
```
$ go get github.com/scoiatael/statsd
```

## Run (and send some metrics to `localhost:8125`)
```
$ statsd                                                                                                                       
[2017-07-28 19:14:50.137949881 +0200 CEST] StatsD Metric: app1.foo 1|c
[2017-07-28 19:15:13.427263263 +0200 CEST] StatsD Metric: app2.bar 205|ms
```
