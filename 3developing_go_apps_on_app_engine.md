# Developing Go Apps on App Engine
## Install the App Engine SDK for Go
```
none.
```

## Testing and Debugging
### Using the Local Development Server
```
Note that when your application is running in the development server,
you can still make remote API calls to the production infrastructure using Google APIs HTTP endpoints.
```

- Using the goapp tool
```
For help using the tool, run goapp help serve for a full description of its options.

The goapp serve command wraps the underlying dev_appserver.py python tool provided with the SDK. Invoking the python tool directly gives greater control over the more esoteric features of the development server
```
- Running the development web server
```
To change which port the web server uses, use the --port option:

dev_appserver.py --port=9999 myapp
```
- Application IDs in the development web server
```
If you need to access your App ID, for example to spoof an email address, use the appengine.AppID function. To get the hostname of the running app, use the appengine.DefaultVersionHostname function.
```
- Using the Datastore
```
This datastore persists between invocations of the web server, so data you store will still be available the next time you run the web server.

To clear the local datastore for an application, use the --clear_datastore=yes option when you start the web server:
dev_appserver.py --clear_datastore=yes myapp


To change the location used for the datastore file, use the --datastore_path option:

dev_appserver.py --datastore_path=/tmp/myapp_datastore myapp


```
-- Browsing the local datastore
```
```
https://cloud.google.com/appengine/docs/go/tools/devserver#Go_The_Development_Console
click datastore viewer in the left navigation pane to view your local datastore contents.
```
-- specifying the automatic id allocation policy
```
IMPORTANT KEYWORD:automatic entiry ID

* development server's automatic entity id policy
** sequential: IDs are assigned from the sequence of consecutive integers
** scattered:  IDs are assigned from a non-repeating sequence of approximately uniformly distributed integers


Note: The auto ID assignment policies for the production server are completely different than those used by the development server.
The default production server policy is similar to the scattered policy but not the same.
There is no policy that corresponds to sequential.
Your app should make no assumptions about the sequence of automatic IDs assigned in production.



Note: The default local test configuration uses the sequential policy, but will change to scattered in a future release.
You may wish to run your tests with the scattered policy to prepare for this
```
- The Users service in the development web server
```
While running under the development web server,
the LoginURL and LogoutURL functions return URLs for /_ah/login and /_ah/logout on the local server.
```
- Using Mail
```
To enable email support, the web server must be given options that specify a mail server to use. The web server can use an SMTP server, or it can use a local installation of Sendmail.


dev_appserver.py --smtp_host=smtp.example.com --smtp_port=25 \
    --smtp_user=ajohnson --smtp_password=k1tt3ns myapp



```
- Using URL Fetch
```
NONE
```
- The Development Console
```
The development web server includes a console web application.
With the console, you can browse the local datastore,
and interact with the application by submitting Python code to a web form.

To access the console, visit the URL http://localhost:8000 on your server.
```
- Command-line arguments
```
--help


--log_level=...
    The lowest logging level at which logging messages will be written to the console;
    messages of the specified logging level or higher will be output.
    Possible values are debug, info, warning, error, and critical.


--logs_path=LOGS_FILE
    By default, development server logs are stored in memory only.

    This option turns on disk storage of logs at the location specified by LOGS_FILE,
    making the logs available across server restarts.
    
    for example
        dev_appserver.py --logs_path=/home/logs/boglogs bog

--storage_path=...
    Path at which all local files (such as the Datastore, Blobstore files, Google Cloud Storage Files, logs, etc) will be stored, unless overridden by --datastore_path, --blobstore_path, --logs_path, etc.
```



### Local Unit Testing
```
App Engine provides testing utilities that use local implementations of datastore and other App Engine services.
This means you can exercise your code's use of these services locally, without deploying your code to App Engine.

Any entity stored during a datastore unit test is stored locally and is deleted after the test run.
```

- Introducing the Go testing package
```
    goapp test
```
- Introducing the aetest package
```
Many function calls to App Engine services require a context.Context as an argument.
The appengine/aetest package provided with the SDK allows you to create a fake context.Context to run your tests using the services provided in the development environment.



```
- Writing Datastore and memcache tests
```
```




## Monitoring Your App
### Monitoring Latency with Cloud Trace
### Detecting Outages and Downtime
####  Capabilities API Overview
####  Capabilities API Reference
### Configuring Dashboards and Alerts with Cloud Monitoring
