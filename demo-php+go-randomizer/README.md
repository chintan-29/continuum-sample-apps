### demo-php+go-randomizer

This app is used as an example in the job linking tutorial, and includes two components:

1. A simple REST API server written in Go that providers a random number
2. A PHP front-end that consumes the API and displays the random number along with information regarding the instance of both the Go and PHP jobs

To construct the app, perform the following:

1. Navigate to the demo-php+go-randomizer/go/ directory and run the following:
```
apc app create
```
2. Next, navigate to the demo-php+go-randomizer/php/ directory and run the following:
```
apc app create
```
3. Lastly:
```
apc job link randomizer-front-end -to randomizer-rest --name rand --port 0
```
4. Navigate to the URL provided from the PHP app staging process to view the final result.

Both the Go and PHP jobs will be built with 5 instances each as defined in their individual manifests. After the jobs are linked, the PHP front-end will use the 'RAND_URI' environment variable to communicate with the REST API. Each time the page is refreshed, it is possible to reach a different PHP+Go job instance as well as receive a different number.