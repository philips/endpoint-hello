## Endpoint Hello

Example of using [discovery APIs][discovery] outside of Google App Engine.

[discovery]: https://developers.google.com/discovery/v1/using

### Running

```
export PATH=$PATH:$PWD
go install code.google.com/p/rsc/devweb
./bin/devweb -addr ":8080" github.com/philips/endpoint-hello/dev
```

- API: `http://localhost:8080/_ah/api/discovery/v1/apis/greeting/v1/rest`
- List: `http://localhost:8080/_ah/api/greeting/v1/greetings`
