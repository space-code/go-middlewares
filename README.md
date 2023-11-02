<h1 align="center" style="margin-top: 0px;">go-middlewares</h1>

<p align="center">
<a href="https://github.com/space-code/go-middlewares/blob/main/LICENSE"><img alt="License" src="https://img.shields.io/badge/License-MIT-yellow.svg"></a>
<a href="https://github.com/space-code/go-middlewares"><img alt="CI" src="https://github.com/space-code/go-middlewares/actions/workflows/ci.yml/badge.svg?branch=main"></a>
</p>

## Description
`go-middlewares` is a collection of commonly used middlewares in Go.

- [Usage](#usage)
- [Installation](#installation)
- [Communication](#communication)
- [Contributing](#contributing)
- [Author](#author)
- [License](#license)

## Usage

1. The `ErrorHandlerMiddleware` is a middleware function designed to handle panics and errors that may occur during the execution of an HTTP handler.

Here's an example of how to use it in your code:

```go
// Create a new HTTP handler and wrap it with the ErrorHandlerMiddleware.
handler := ErrorHandlerMiddleware(yourHandler)

// Use the wrapped handler for your server.
http.Handle("/your-route", handler)

// Start your HTTP server.
http.ListenAndServe(":8080", nil)
```

2. The `LogHandlerMiddleware` is designed to log HTTP requests made to your web application. It captures various details about the request and logs them to your preferred output, which is especially useful for debugging and monitoring.

```go
// Create a new HTTP handler and wrap it with the LogHandlerMiddleware.
handler := LogHandlerMiddleware(yourHandler)

// Use the wrapped handler for your HTTP server.
http.Handle("/your-route", handler)

// Start your HTTP server.
http.ListenAndServe(":8080", nil)
```

## Installation

```
go get github.com/space-code/go-middleware
```

## Communication
- If you **found a bug**, open an issue.
- If you **have a feature request**, open an issue.
- If you **want to contribute**, submit a pull request.

## Contributing

Please feel free to help out with this project! If you see something that could be made better or want a new feature, open up an issue or send a Pull Request!

## Author
Nikita Vasilev, nv3212@gmail.com

## License
go-middlewares is available under the MIT license. See the LICENSE file for more info.