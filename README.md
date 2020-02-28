
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

### Endpoint

Make a post request with data containing project name after the slash in `github.com`, e.g. `github.com/cloudhubs/tms`,
so `cloudhubs/tms` is what we are interested in.

```bash
curl --request POST \
  --url http://localhost:8080/ \
  --header 'content-type: application/json' \
  --data '{
    "url":"cloudhubs/tms"
}'
```