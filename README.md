# Web Observer

The application is made to observe selected websites using cURL and calculate the changes using ssdeep fuzzy hash. Should the changes exceed selected treshold, the application would record the diff in a file

## Getting Started

After installation, you may insert the websites you want to monitor in 'curlfile.txt'. The format goes like this:
```
Example Site 1
curl https://www.example.com
Example Site 2
curl https://www.example2.com
Example Site 3
curl https://www.example3.com
```

### Prerequisites

The application uses gohtml to beautify html and ssdeep to conduct fuzzy hashing.

```
go get github.com/yosssi/gohtml
go get github.com/glaslos/ssdeep
```

### Installing

A step by step series of examples that tell you how to get a development env running

Say what the step will be

```
Give the example
```

### TODOs
- Improve the current shitty JS beautify
- Include functionality to email a user when changes are observed

## Authors

* **Samuel Pua** - *Initial work* - [Github](https://github.com/telboon)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

