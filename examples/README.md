# examples

## example.go

will read test.md and render it as xhtml 

execute it by simply

```
go run example.go
```

## confluence.go

reads test.md and uses [confluence-go-api](https://github.com/cseeger-epages/confluence-go-api) to create a new confluence page using test.md markdown

you have to change the api initialisation to use the correct credentials
