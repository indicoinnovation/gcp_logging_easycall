# gcp_logging_easycall

[![pkg-img]][pkg-url]

It helps you log your telemetry to **google cloud logging** without much effort, just import and use it.

## Rationale
Deal with some google cloud libraries is sometimes painful and unclear code.
We were looking for some telemetry snippet that could allow us just import it and call function to log our actions.
The goal was to log our actions without code rewrites everytime we want.
This Library does this with google cloud logging, allowing users just call a function and pass arguments

## Dependencies
- Google Cloud SDK (google-cloud-logging)

## Features
- Simple API.
- Clean and tested code.
- Supports different payloads

## Install
Go version 1.14+
```
go get github.com/indicoinnovation/gcp_logging_easycall
```

##Example

```
import (
    gcpLogging "github.com/indicoinnovation/gcp_logging_easycall"
)

gcpProjectName := "your google cloud project"
loggerName := "the name to group your logs into"
message := &gcpLogging.Logger{
    User: "your user",
	Message: "your message",
	Reason: "A reason like err.Error()",
	RemoteIp: "the remote IP if applied",
	Method: "If applied the method used in this call",
	Urlpath: "",
	StatusCode: 200,
	RequestData: "json_data",
	ResponseData: "json_data",
	SessionId: "session_id"
}
severity := "alert" //If none of options applied then default severity will be logged

gcpLogging.Log(gcpProjectName, loggerName, message, severity)
```

## Documentation
See [these docs][pkg-url].

## License
[GNU General Public License v3.0](LICENSE)

[build-url]: https://github.com/indicoinnovation/gcp_logging_easycall/actions
[pkg-img]: https://pkg.go.dev/badge/indicoinnovation/gcp_logging_easycall
[pkg-url]: https://pkg.go.dev/github.com/indicoinnovation/gcp_logging_easycall
[version-url]: https://github.com/indicoinnovation/gcp_logging_easycall/releases
