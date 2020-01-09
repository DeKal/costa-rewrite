## Just a Project to analyze data return from public API of Costa Rewrite

### Run

#### Run with Default param
```
./run.sh 
```

#### Help
```
./run.sh --help
_______________________________________________________________________
Usage of ./costa-rewrite:
  -inputName string
    	an Input name for reading data (default "example_input.csv")
  -outputName string
    	an Output name for writing data (default "output.csv")
  -rewriteUrl string
    	rewrite url (default "http://localhost:9999/_c/v1/search/rewrite/?q=%s&lang=en&segment=women")
```

#### Run with custom param 
```
./run.sh -inputName example_input.csv -outputName output.csv -rewriteUrl rewrite-url-example.com
```