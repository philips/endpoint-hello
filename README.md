## Endpoint Hello

Example of using discovery APIs outside of Google App Engine.

### Running

```
export PATH=$PATH:$PWD
go install code.google.com/p/rsc/devweb
./bin/devweb -addr ":8080" github.com/philips/endpoint-hello/dev
```

Now visit `http://localhost:8080/\_ah/api/discovery/v1/apis/greeting/v1/rest`