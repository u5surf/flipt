# Configuration

There are two ways to configure Flipt: using a configuration file or through environment variables.

## Configuration File

The default way that Flipt is configured is with the use of a configuration file `default.yml`.

This file is read when Flipt starts up and configures several important properties for the backend and API.

You can edit any of these properties to your liking, and on restart Flipt will pick up the new changes.

!!! note
    These defaults are commented out in `default.yml` to give you an idea of what they are. To change them you'll first need to uncomment the property.

These properties are as follows:

| Property | Description | Default |
|---|---|---|
| host | The IP on which to serve the Flipt application | 0.0.0.0 |
| log.level | Level at which messages are logged (trace, debug, info, warn, error, fatal, panic) | info |
| api.port | The port on which to serve the Flipt REST API | 8080 |
| backend.port | The port on which to serve the Flipt Backend server | 9000 |
| db.name | The name given to the Flipt database (suffixed with .db) | flipt |
| db.path | Where to store the Flipt database | /var/opt/flipt |
| db.migrations.auto | If database migrations are run on Flipt startup | true |
| db.migrations.path | Where the Flipt database migration files are kept | /etc/flipt/config/migrations |

## Using Environment Variables

All options in the configuration file can be overridden using environment variables using the syntax:

```shell
FLIPT_<SectionName>_<KeyName>
```

!!! tip
    Using environment variables to override defaults is especially helpful when running with Docker as described in the [Installation](installation.md) documentation.

Everything should be upper case, `.` should be replaced by `_`. For example, given these configuration settings:

```yaml
backend:
  port: 9000

db:
  name: flipt
  path: /var/opt/flipt
```

You can override them using:

```shell
export FLIPT_BACKEND_PORT=9001
export FLIPT_DB_NAME=my-db
export FLIPT_DB_PATH=/tmp/db
```
