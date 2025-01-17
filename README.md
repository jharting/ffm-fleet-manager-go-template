# Fleet Manager Golang Template

This project is an example fleet management service. Fleet managers govern service 
instances across a range of cloud provider infrastructure and regions. They are 
responsible for service placement, service lifecycle including blast radius aware 
upgrades,control of the operators handling each service instance, DNS management, 
infrastructure scaling and pre-flight checks such as quota entitlement, export control, 
terms acceptance and authorization. They also provide the public APIs of our platform 
for provisioning and managing service instances.


To help you while reading the code the example service implements a simple collection
of _dinosaurs_ and their provisioning, so you can immediately know when something is 
infrastructure or business logic. Anything that talks about dinosaurs is business logic, 
which you will want to replace with your our concepts. The rest is infrastructure, and you
will probably want to preserve without change.

For a real service written using the same fleet management pattern see the
[kas-fleet-manager](https://github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager).

To contact the people that created this template go to [zulip](https://bf2.zulipchat.com/).

## Prerequisites
* [Golang 1.16+](https://golang.org/dl/)
* [Docker](https://docs.docker.com/get-docker/) - to create database
* [ocm cli](https://github.com/openshift-online/ocm-cli/releases) - ocm command line tool
* [Node.js v12.20+](https://nodejs.org/en/download/) and [npm](https://docs.npmjs.com/downloading-and-installing-node-js-and-npm)

## Using the template for the first time
The [implementation](./docs/implementation.md) talks about the main components of this template. 
To bootstrap your application, after cloning the repository. 

1. Replace _dinosaurs_ placeholder with your own business entity / objects
1. Implement code that have TODO comments
   ```go
   // TODO
   ```

## Running Fleet Manager for the first time in your local environment
Please make sure you have followed all of the prerequisites above first.  

1. Follow the [populating configuration guide](docs/populating-configuration.md)
   to prepare Fleet Manager with its needed configurations

1. Compile the Fleet Manager binary
```
make binary
```
1. Create and setup the Fleet Manager database
    - Create and setup the database container and the initial database schemas
    ```
    make db/setup && make db/migrate
    ```
    - Optional - Verify tables and records are created
    ```
    # Login to the database to get a SQL prompt
    make db/login
    ```
    ```
    # List all the tables
    serviceapitests# \dt
    ```
    ```
    # Verify that the `migrations` table contains multiple records
    serviceapitests# select * from migrations;
    ```

1. Start the Fleet Manager service in your local environment
    ```
    ./fleet-manager serve
    ```

    This will start the Fleet Manager server and it will expose its API on
    port 8000 by default

    >**NOTE**: The service has numerous feature flags which can be used to enable/disable certain features 
    of the service. Please see the [feature flag](./docs/feature-flags.md) documentation for more information.
1. Verify the local service is working
    ```
    curl -H "Authorization: Bearer $(ocm token)" http://localhost:8000/api/dinosaurs_mgmt/v1/dinosaurs
   {"kind":"DinosaurRequestList","page":1,"size":0,"total":0,"items":[]}
    ```
   >NOTE: Change _dinosaur_ to your own rest resource

   >NOTE: Make sure you are logged in to OCM through the CLI before running
          this command. Details on that can be found [here](./docs/populating-configuration.md#interacting-with-the-fleet-manager-api)

## Using the Fleet Manager service

### Interacting with Fleet Manager's API

See the [Interacting with the Fleet Manager API](docs/populating-configuration.md#interacting-with-the-fleet-manager-api)
subsection in the [Populating Configuration](docs/populating-configuration.md)
documentation

### Viewing the API docs

```
# Start Swagger UI container
make run/docs

# Launch Swagger UI and Verify from a browser: http://localhost:8082

# Remove Swagger UI conainer
make run/docs/teardown
```

### Running additional CLI commands

In addition to starting and running a Fleet Manager server, the Fleet Manager
binary provides additional commands to interact with the service (i.e. cluster
creation/scaling, Dinosaur creation, Errors list, etc.) without having to use a
REST API client.

To use these commands, run `make binary` to create the `./fleet-manager` binary.

Then run `./fleet-manager -h` for information on the additional available
commands.

### Fleet Manager Environments

The service can be run in a number of different environments. Environments are
essentially bespoke sets of configuration that the service uses to make it
function differently. Environments can be set using the `OCM_ENV` environment
variable. Below are the list of known environments and their
details.

- `development` - The `staging` OCM environment is used. Sentry is disabled.
   Debugging utilities are enabled. This should be used in local development.
   This is the default environment used when directly running the Fleet
   Manager binary and the `OCM_ENV` variable has not been set.
- `testing` - The OCM API is mocked/stubbed out, meaning network calls to OCM
   will fail. The auth service is mocked. This should be used for unit testing.
- `integration` - Identical to `testing` but using an emulated OCM API server
   to respond to OCM API calls, instead of a basic mock. This can be used for
   integration testing to mock OCM behaviour.
- `production` - Debugging utilities are disabled, Sentry is enabled.
   environment can be ignored in most development and is only used when
   the service is deployed.

The `OCM_ENV` environment variable should be set before running any Fleet
Manager binary command or Makefile target

## Additional documentation
- [Adding new endpoint](docs/adding-a-new-endpoint.md)
- [Adding new CLI flag](docs/adding-new-flags.md)
- [Automated testing](docs/automated-testing.md)
- [Deploying fleet manager via Service Delivery](docs/onboarding-with-service-delivery.md)
- [Requesting credentials and accounts](docs/getting-credentials-and-accounts.md)
- [Data Plane Setup](docs/data-plane-osd-cluster-options.md)
- [Access Control](docs/access-control.md)
- [Quota Management](docs/quota-management-list-configuration.md)
- [Running the Service on an OpenShift cluster](./docs/deploying-fleet-manager-to-openshift.md)
- [Explanation of JWT token claims used across the fleet-manager](docs/jwt-claims.md)

## Contributing
See the [contributing guide](CONTRIBUTING.md) for general guidelines on how to
contribute back to the template.