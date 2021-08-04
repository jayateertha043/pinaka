<h1 align="center">PINAKA</h1>

>Simple DOS tool (under development)


## REQUIREMENTS AND INSTALLATION

Build Pinaka:
```
git clone https://github.com/jayateertha043/pinaka.git
cd pinaka
go build pinaka.go
```

or

Install using Go get:

```
go get github.com/jayateertha043/pinaka
```

Run pinaka:

```
.\pinaka -h
```


>Note:Ensure you have git version>1.8

## USAGE:
```
Usage of pinaka:
  -data string
        To send custom post data)
  -headers string
        To use Custom Headers headers.json file
  -maxrequest uint
        Enter max requests to make (default 100000)
  -post
        To send post request
  -proxy string
        Use custom proxy [http://ip:port or https://ip:port]
  -t int
        Enter amount of threads (default 100)
  -timeout int
        Enter request timeout in seconds (default 3)
  -url string
        Input Url(https://www.example.com)
```

## Author

👤 **Jayateertha G**

* Twitter: [@jayateerthaG](https://twitter.com/jayateerthaG)
* Github: [@jayateertha043](https://github.com/jayateertha043)

