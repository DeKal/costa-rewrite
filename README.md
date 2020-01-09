### Description
Analyzing data return from public API of Costa Rewrite

### Prerequisite
- Input CSV file
- Default CSV file name `example_input.csv`
- CSV file must have Header view sample file `example_input.csv`

#### Run with Default param
```
./run.sh 
```

#### Help
```
./run.sh --help

Usage of ./costa-rewrite:
  -country string
    	country filter for some specific word [ HK, ID, MY, PH, SG, TW] (default "SG")
  -inputName string
    	an Input name for reading data (default "example_input.csv")
  -outputName string
    	an Output name for writing data (default "output.csv")
  -rewriteHost string
    	rewrite host (default "http://localhost:9999")
```

#### Run with custom param 
```
./run.sh -country SG -inputName example_input.csv -outputName output.csv -rewriteHost rewrite-url-example.com
```