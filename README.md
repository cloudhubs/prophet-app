
# Installation

```bash
brew install go
```

# Build

## Run the main app
```go
go build
```

# Run

```go
./main
```

# Files

* main.go: main server
* model.go: data structures
* requests.go: validates number of requests per day
* decode.go: decodes response
* utilsHttp.go: makes http call to code analyzer
* 

## Main Endpoint

Analyzing repository `https://github.com/cloudhubs/tms` by running:

```bash
curl --request POST \
  --url http://localhost:4000/ \
  --header 'content-type: application/json' \
  --data '{
    "repositories": [
        {
            "organization":"cloudhubs",
            "repository":"tms",
            "isMonolith":"true"
        },
    ]
}'
```

* validate number of requests per day
* validate size of the repository < 100 MB
* make call to:

```bash
curl --request POST \
  --url http://localhost:5000/ \
  --header 'content-type: application/json' \
  --data '{
    "repositories": [
        {
            "organization":"cloudhubs",
            "repository":"tms",
            "isMonolith":"true"
        },
    ]
}'
```

Error calls:

* 403 if requests per day exhausted
* 401 if repository size exceeds 100 MB