# go-cleanup

run cleanup functions when server shutdown.

# install
```
go get -u github.com/liguangsheng/go-cleanup
```

# example

```go
func main() {
	defer cleanup.Run()
	
	resource := NewResource()
	cleanup.Add(func() {
		resource.close()
	})
	
    http.ListenAndServe(":8000")
}
```