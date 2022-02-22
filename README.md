# BodyParser

It divides the string that comes with the request body and transfers it into a map, so that we can access the data in the body regularly.

## Getting Started


### Installing
import to your project
```go
import "github.com/Periyot/BodyParser"
```

```go
go get github.com/Periyot/BodyParser
```

### Quick start
We convert the data we read from `r.Body' to String and send it to body Parser. BodyParser returns us a Map
```go
func indexHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.NotFound(w, r)
		return
	}
	
	bodyByte, _ := ioutil.ReadAll(r.Body)
	bodyString := string(bodyByte)
	
	maps := BodyParser.Parser(bodyString)
	mapsJson, _ := json.Marshal(maps)
	fmt.Fprint(w, string(mapListJSON))
}
```
## License

This project is licensed under the [MIT] License - see the LICENSE.md file for details

