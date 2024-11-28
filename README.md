### Brief application description
This application was made as a test task by Vitali Saroka in 2024 and is an example of a get query that returns the index of an element or an error if the index is not found. The original task description and requirements are at the end of the document.

### Installing and using the application


In order to install and run the application locally, you will need:
1) Download the source code from this repository
2) Install Go according to the official instructions from https://go.dev/doc/install (version 1.23.3 was used during development)
3) Optionally, create an env file and add the port and log level (Info, Error or Debug) in it:
`PORT=-3000`
`LOG_LEVEL=Debug`
4) Run the `make all` command
5) Check everything using a web client by making a request (for example http://localhost:9999/endpoint/90000)

In addition, you can run the following commands using *make*:
- *make about* - displays brief information about the application
- *make install* - installs the necessary libraries
- *make build* - builds and compiles the project
- *make run* - runs the project locally
- *make test* - runs tests
- *make all* - installs modules and runs the project locally after the tests
- *make help* -shows instructions for installation from scratch

### Original task description
We have created large file containing sorted numbers from 0 to 1000000, for example:

```
Value: 0 10 20 100 … 1000000
Index: 0 1  2  3   … 50000
```

We would like to be able to call designed endpoint with http `GET` method and send `value` that should be found in input file.
As a response we should get `index` number for given value and the corresponding value and optional message.

For example, we are sending GET for /endpoint/100 and as result we should receive 3. 

Remark: `As a requirement`, we want to load that file into `slice` once service starts.
So all search operations should be optimized for that particular slice.

- In case you’re not able to find index for given value, you can return `index` for any other existing value, assuming that conformation is at `10% level`, (for example, you were looking for `index` for value = `1150`, but in input file you have `1100` and `1200`, so in that case you can return index for `1100` or `1200`).
 
- In case you were not able to find valid `index` number, `error message` should be added into response.

`To summarize`:
- Design API for http `GET` method
- Implement functionality for searching `index` for `given` value (it should be the most efficient algorithm) 
- Add logging
- Add possibility to use configuration file where you can specify service port and log level (you should be able to choose between Info, Debug, Error)
- Add `unit tests` for created components
- Add `README.md` to describe your service
- Automate running tests with `make` file
- Remember that code structure matters
- Upload solution into `GitHub` account and share the link

Sample input file is added as `input.txt` file.
