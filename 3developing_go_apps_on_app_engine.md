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


For more control over the underlying instance, you can use aetest.NewInstance instead.
```
- Writing Datastore and memcache tests
```

Testing code which uses the datastore or memcache is simple once you create a context.Context with the aetest package:
in your test call aetest.NewContext to create a context to pass to the function under test.
```




## Monitoring Your App
### Monitoring Latency with Cloud Trace
```
NOTE:cloud trace is beta
```

#### quickstarts(https://cloud.google.com/trace/)
```
Cloud Trace is a distributed tracing system for Google Cloud Platform that collects latency data from App Engine applications and displays it in near real time in the Google Cloud Platform Console.

It helps you understand
how long it takes your application to handle incoming requests from users or other applications,
and how long it takes to complete operations like RPC calls performed when handling the requests.

Currently, Cloud Trace collects end-to-end latency data for requests to App Engine URIs and additional data for round-trip RPC calls to App Engine services like Datastore, URL Fetch, and Memcache.


You can use Cloud Trace to:
* Quickly view a snapshot of last-day latency data for your application in the trace overview.
* Drill down to detailed latency data, including performance insights and detailed data for the most frequent application requests and RPC calls.
* Find latency data for individual requests and view latency details.
* Generate custom analysis reports that show an overview of latency data for all or a subset or requests, and allow you to compare two different sets of latency data.
** NOTE:IMPORTANT!
```

- why cloud trace?
```
Cloud Trace can help you answer the following questions:
* How long does it take my application to handle a given request?
* Why is it taking my application so long to handle a request?
* Why do some of my requests take longer than others?
* What is the overall latency of requests to my application?
* Has latency for my application increased or decreased over time?
* What can I do to reduce application latency?
```
- how does cloud trace work?
```
Cloud Trace works with all App Engine APIs, with the exception of Cloud SQL.
```


##### getting started
- enable cloud trace
skip.
- view the trace overview
skip.
- find a trace
!https://cloud.google.com/trace/images/cloud-trace-trace-list.png!
- view trace details
```
In the trace details view, you can:
* View the total latency for the request, aggregate latency for RPC calls, and other request details at the top of the page.
* Click the Log View link to view the log entry for the request in the console Logs Viewer.
* Click the root span in the Timeline tab to view overall latency details for the request.
* Click any subspan in the timeline to view latency details for a particular RPC call within the request.
* Click Summary to view a tabular summary of latency data for the request and RPC calls.
```
- create an analysis report
skip.
- view an analysis report
```
points
* Latency Density Distribution and Cumulative Distribution
* Latency panel below the report graphs
** View the distribution of average latencies for requests by percentage 
* Sample Traces column
** trace details for representative samples of requests at different percentiles.
* Bottlenecks panel
** any detected performance bottlenecks
* Sample Traces column
** trace details for representative samples of requests where reported bottlenecks were detected.
```



####guides
NOTE:skip all
##### trace overview
##### finding traces
##### viewing trace details
##### analysis reports
####apis & reference
##### cloud trace api
##### rest api
##### rpc api


### Detecting Outages and Downtime
####  Capabilities API Overview
```
With the Capabilities API, your application can detect outages and scheduled downtime for specific API capabilities.

You can use this API to reduce downtime in your application by detecting when a capability is unavailable and then bypassing it.
```
- using the capabilities api in go
```
The capability.Enabled function returns true if the provided API and capability are available.
```
- supported capabilities
```
The API currently supports the following capabilities:
* Availability of the blobstore
* Datastore reads
* Datastore writes
* Availability of the Mail service
* Availability of the Memcache service
* Availability of the Task Queue service
* Availability of the URL Fetch service
* Availability of the XMPP service
```

####  Capabilities API Reference
skip.


### Configuring Dashboards and Alerts with Cloud Monitoring
#### quickstarts
- what is google cloud monitoring?
```
** Receive alerts when issues occur. **
Receive alerts via email, SMS, PagerDuty, HipChat, and more.
Alert on individual metrics and thresholds or on aggregate group performance.


Integrate with common open source software.
Cloud Monitoring collects metrics from many common open source servers with minimal configuration.


```


- cloud monitoring concepts
-- resources and groups
```
A ** resource ** is an abstract object provided by certain products, platforms, or services within Google Cloud Platform.


A group is a collection of resources, such as "all VM instances in my project that are running Cassandra."
Cloud Monitoring can automatically detect resources that are related and group them together.
*** Groups are not presently available in App Engine-only projects. ***
```
-- metrics
```
** metric ** is a measured value that can be used to assess a system. 


** service metric ** or ** platform metric **
is a metric that is provided for you when you use a particular service or platform. 


You can define ** custom metrics ** to measure any aspect of your system.
Custom metrics might include cart checkouts, user logins, or business KPIs.
You are responsible for sending your custom metrics data to Cloud Monitoring using the
*** custom-metrics API. ***
```
-- charts and dashboards
```
A ** chart ** is a named, visual representation of one or more metrics

A ** dashboard ** is a collection of charts and other information such as event logs and incident lists.
```
-- uptime checks and events
```
** Uptime checks ** is a service of Cloud Monitoring.
You configure the service to check your system's health by sending requests to your applications, services, or URLs from various locations around the world.
*** You can use the results of the checks as conditions in your alert policies, so you will be notified if system health is degraded. ***



An event log is a time-ordered list of system events related to your application and to the platforms and services it uses.
Examples of events include downtime of the cloud infrastructure, notifications related to your alert policy, and code deployments.
You can also add your own events to the log.
Event logs are supported by Cloud Monitoring
    and are different from the logs supported by Google Cloud Logging.

```
-- alerts, notifications, and incidents
```
An ** alert policy ** is a set of rules that determine whether your resources or groups are operating normally.
The rules are logical conditions involving metric thresholds and uptime checks.
For example, you can create a rule that your web site's average response latency must not exceed five seconds over a period of two minutes.


An ** alert ** occurs when an alert policy's conditions are met, causing an ** incident ** to appear in the Incidents section of the Cloud Monitoring Console.
Incidents remain open until the alert policy rules are no longer in violation or until the incident is manually closed.


You can associate ** notifications ** with alert policies.
For example, alerts can send email or SMS notifications to people or services.
```
- google cloud monitoring and stackdriver
none.


#### getting started
- Start using Cloud Monitoring
-- Monitoring App Engine projects
-- Monitoring Google Compute Engine projects
- The Event Log
- Uptime checks
-- Create an uptime check
- Dashboards and charts
-- Create a dashboard and chart
-- View logs associated with charts
- Groups
-- Create a group
- Alerting policies
-- Create a policy
