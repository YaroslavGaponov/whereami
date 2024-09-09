whereami
=====
Service for determining the nearest city by latitude and longitude


# Demo

## Build docker image

```sh
docker build -t YaroslavGaponov/whereamid:latest .
```

## Run docker image

```sh
docker run -p 8080:8080 YaroslavGaponov/whereamid:latest
```

## Test

### Request

```sh
curl 'http://localhost:8080/whereami?lat=44.060522&lng=15.345933'
```

### Response

```sh
{"id":"1191160875","lat":44.1194,"lng":15.2319,"distance":11.21630046370274,"took":1058629,"city":"Zadar","country":"Croatia"}
```

