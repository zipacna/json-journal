# JSON Journal
This is a minimal API taking HTTP Requests for 
Creating, Reading, Updating and Deleting JSON Data (by timestamp).

* Author: Jean Mattes
* Author-URI: https://risingcode.net/
* License: 2021 MIT (Code) and CC-BY-4.0 (Docs)

**Features:**
* Minimalistic API with JSON File IO
* Create => http://localhost:9876/create
* Read => http://localhost:9876/read
* Update => http://localhost:9876/update
* Delete => http://localhost:9876/delete

<!--TODO Snippet => <span style="color:red">TODO: </span>-->

## Build
For a Build targeting Windows append ``.exe`` to `json-journal`; macOS untested.
````shell
go build -o json-journal main.go jsonFile.go httpHandlers.go httpHelpers.go
````
For a Build targeting an operating system (OS) different from the building OS, 
temporarily set the Environment Variable GOOS to the target OS, 
like targeting Windows from Linux:
````shell
go env -w GOOS=windows; go build -o json-journal.exe main.go jsonFile.go httpHandlers.go httpHelpers.go; go env -w GOOS=linux; 
````

## Usage
Substitute your URI if running remotely. 
To check if it is working run the following lines (or something equivalent):

### Starting the API
Run ``json-journal`` or ``json-journal`` on Windows

Run ``./json-journal`` on Linux/macOS/PowerShell

### Create
````shell
curl -w "Status %{http_code}\t" -H "Content-Type: application/json" -d "{\"Data\": \"that\tEscape 4: 2 json 2 go\"}" http://localhost:9876/create
# Example Output: Status 201
````

### Read
````shell
curl -sb -H "Accept: application/json" http://localhost:9876/read
# Example Output: [{"Data":"that\tEscape","Timestamp":"2021-10-10T13:06:12.0823785+02:00"},{"Data":"that\tEscape","Timestamp":"2021-10-10T13:25:43.2787097+02:00"}]
````

### Update
````shell
curl -w "Status %{http_code}\t" -H "Content-Type: application/json" -d "{\"Data\": \"that\tEscape3\", \"Timestamp\": \"2021-10-10T13:25:43.2787097+02:00\"}" http://localhost:9876/update
# Example Output: Status 204
````

### Delete
````shell
curl -w "Status %{http_code}\t" -H "Content-Type: application/json" -d "{\"Timestamp\": \"2021-10-10T15:05:59.4745304+02:00\"}" http://localhost:9876/delete
# Example Output: Status 204
````

### Other Examples
Windows (for PowerShell replace curl with curl.exe)
````shell
curl -w "Status %{http_code}\t" -H "Content-Type: application/json" -d "{\"Data\": \"that\tEscape\"}" http://localhost:9876/create
:: Example Output: Status 200
type journal.json
:: Example Output (local): {"Data":"that\tEscape","Timestamp":"2021-10-10T12:01:33.7875075+02:00"}
curl -sb -H "Accept: application/json" http://localhost:9876/read
:: Example Output (remote): {"Data":"that\tEscape","Timestamp":"2021-10-10T12:01:33.7875075+02:00"}
````

Linux
````shell
curl -w "Status %{http_code}\t" -H "Content-Type: application/json" -d "{\"Data\": \"that\tEscape 2: electric curl\"}" http://localhost:9876
# Example Output: Status 200
cat journal.json
# Example Output: [{"Data":"that\tEscape","Timestamp":"2021-10-10T12:01:33.7875075+02:00"},{"Data":"that\tEscape 2: electric curl","Timestamp":"2021-10-10T12:11:33.7875075+02:00"}]
curl -sb -H "Accept: application/json" http://localhost:9876/read
# Example Output: [{"Data":"that\tEscape","Timestamp":"2021-10-10T12:01:33.7875075+02:00"},{"Data":"that\tEscape 2: electric curl","Timestamp":"2021-10-10T12:11:33.7875075+02:00"}]
````

*Disclaimer:* Either I committed some war crimes in the code,
Go (Golang/Gopher) with JetBrains GoLand is keeping errors at bay
or writing the update and delete features worked on first compile as designed.

## Changelog
### Version 1.0
* <span style="color:Gray">Planned</span>: Have Frontends in place
    * [ ] Python Flask/Gunicorn Web Server
    * [ ] Android Java App
    * [ ] C# .Net Framework 4.8 WPF App
### Version 0.8
* [ ] <span style="color:Gray">Planned</span>: Decide on Logging-Strategy
  => Flask (log every request) vs. Go-Auth-S (log zero)
### Version 0.7
* [ ] <span style="color:red">TODO</span>: fix Bash Test Script vs. unit/integration tests
### Version 0.6
* [x] Reduced Debug Messages and published on GitHub.
### Version 0.5
* [x] Major Refactor into multiple files and less Duplication.
### Version 0.4
* [x] Added Delete Operation
### Version 0.3
* [x] Added Update Operation
### Version 0.2
* [x] Added Read Operation
### Version 0.1
* [x] Added Create Operation
