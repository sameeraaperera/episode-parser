# Episode Parser

Accepts a list of episodes in JSON format and returns a list of shows that is DRM enabled

&nbsp;
## Running the server

Set a valid port for the HTTP server to listen at as an environment variable
```
export PORT=8080
```
From the root directory run
```
make run
```
&nbsp;
## Sample Usage

```
POST http://domain.com/parse
```

Sample Request:
https://challengeaccepted.streamco.com.au/samples/request.json

Sample Response:
https://challengeaccepted.streamco.com.au/samples/response.json
