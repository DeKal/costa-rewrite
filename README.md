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
    	country filter for some specific word [HK, ID, MY, PH, SG, TW] (default "SG")
  -file1 string
    	file name 1 (default "output")
  -file2 string
    	file name 2 (default "output-2")
  -inputName string
    	an Input name for reading data (default "example_input.csv")
  -mode string
    	mode: [normal, compare] (default "normal")
  -outputName string
    	an Output name for writing data (default "output.csv")
  -rewriteHost string
    	rewrite host (default "http://localhost:9999")
```

#### Run with custom param 
```
./run.sh -country SG -inputName example_input.csv -outputName output.csv -rewriteHost rewrite-url-example.com
```

#### Run with compare mode 
```
./run.sh -mode compare -file1 output-local-normal.csv -file2 output-local-remove.csv -outputName output.csv
```