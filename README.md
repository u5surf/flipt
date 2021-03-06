# Flipt

[![Build Status](https://travis-ci.com/markphelps/flipt.svg?token=TBiDDmnBkCmRa867CqCG&branch=master)](https://travis-ci.com/markphelps/flipt)
[![CodeClimate Coverage](https://img.shields.io/codeclimate/coverage/markphelps/flipt.svg?style=flat)](https://codeclimate.com/github/markphelps/flipt)
[![Go Report Card](https://goreportcard.com/badge/github.com/markphelps/flipt)](https://goreportcard.com/report/github.com/markphelps/flipt)
[![GitHub Release](https://img.shields.io/github/release/markphelps/flipt.svg?style=flat)](https://github.com/markphelps/flipt/releases)

A self contained feature flag solution.

![Flipt](docs/assets/images/flipt.png)

## Documentation

Documentation is hosted at: [https://markphelps.github.io/flipt/](https://markphelps.github.io/flipt/)

## What is Flipt

Flipt is an open source, self contained application that enables you to use feature flags and experiment (A/B test) across services, running in **your** environment.

This means that you can deploy Flipt within your existing infrastructure and not have to worry about your information being sent to a third party, or the latency required to communicate across the internet.

Flipt includes native client SDKs as well as a REST API so you can choose how to best integrate Flipt with your applications.

For more on Flipt and it's concepts, take a look at the [Concepts](https://markphelps.github.io/flipt/concepts/) documentation.

## Flipt Features

Flipt enables you to add [feature flag](https://martinfowler.com/bliki/FeatureToggle.html) support to your existing applications, with a simple, single UI and API.

This can range from simple on/off feature flags to more advanced use cases where you want to be able to rollout different versions of a feature to percentages of your users.

Flipt features include:

* Fast. Written in Go. Optimized for performance
* Stand alone, easy to run server with no external dependencies
* Ability to create advanced distribution rules to target segments of users
* Native GRPC client SDKs to integrate with your applications
* Simple REST API
* Modern UI and debug console

## Why Flipt

Many organizations understand the benefit of using feature flags in production, so they choose to implement them themselves in their main application or monolith.

As their organization grows, so does their application/infrastructure and functionality makes it's way into a multitude of other services. Many times those services aren't even implemented in the same language.

This is where their original feature flag solution tends to break down as it cannot be easily adapted to those services or languages. This results in:

1. Not being able to use feature flags in a subset of services.
1. Having multiple sources of truth for feature flags depending on the service/implementation which leads to unpredictability.

Flipt solves all of these issues and more, enabling you to focus on your applications without having to worry about implementing your own feature flag solution that works across your infrastructure.

On top of this, Flipt provides a nice, modern UI so that you can always monitor the state of your feature flags and experiments in a single place.

## Running Flipt

Flipt is a single, self contained binary that you run on your own servers or cloud infrastructure. There are a multitude of benefits to running Flipt yourself, including:

* **Security**. No data leaves your servers and you don't have to open your systems to the outside world to communicate with Flipt. It all runs within your existing infrastructure.
* **Speed**. Since Flipt is co-located with your existing services, you do not have to communicate across the internet to another application running on the other side of the world which can add excessive latency and slow down your applications.
* **Simplicity**. Flipt is a single binary with no external dependencies. This means there is no database to manage or connect to, no clusters to configure, and data backup is as simple as copying a single file.

### Try It

```bash
❯ docker run -p 8080:8080 -p 9000:9000 markphelps/flipt:latest
```

## What's Next

To see Flipt in action, checkout an [example](example/).

Want to get up and running with Flipt? See [Getting Started](https://markphelps.github.io/flipt/getting_started/).

For a more detailed guide on how to setup and run Flipt, checkout the [Installation](https://markphelps.github.io/flipt/installation/) documentation.

To learn how Flipt works behind the curtain, read up on it's [Architecture](https://markphelps.github.io/flipt/architecture/).

For information on how to integrate Flipt with your existing applications, see the [Integration](https://markphelps.github.io/flipt/integration/) guide.

## Author

* Website: [https://markphelps.me](https://markphelps.me)
* Twitter: [@mark_a_phelps](https://twitter.com/mark_a_phelps)
* Email: _mark.aaron.phelps at gmail.com_

## TODO/Contributing

I would love your help! Before submitting a PR, please read over the [Contributing](.github/contributing) guide.

Here's a couple of areas that could use some love:

* Caching - Evaluation speed could greatly be improved with the help of caching flags/segments/rules/etc in memory.
* Documentation - Typo? Does something not make sense? Could it be worded better? Please help!
* Examples - More examples on how to use Flipt.
* Test Coverage - Would love to get all coverage over 80%.
* Javascripts - I'm no JS wizz, I'm sure the Javascript code in [ui/src](ui/src) could be improved/simplified/tested.

## Licensing

There are currently two types of licenses in place for Flipt:

1. Client License(s)
2. Server License

### Client License(s)

All of the code required to generate GRPC clients in other languages as well as the existing GRPC Go client are licensed under the [MIT License](https://spdx.org/licenses/MIT.html).

This code exists in the [proto/](proto/) directory.

The client code is the code that you would integrate into your applications, which is why a more permissive license is used.

### Server License

The server code is licensed under the [GPL 3.0 License](https://spdx.org/licenses/GPL-3.0.html).

If there are any concerns about the use of this license for the server, please [open an issue](https://github.com/markphelps/flipt/issues/new) on GitHub so that we can discuss publicly.

## Pro Version

My plan is to soon start working on a Pro Version of Flipt for enterprise. Along with support, some of the planned features include:

* User management/login
* Permissions
* HTTPS
* Auditing
* Streaming updates
* Metrics
* HA support

If you or your organization would like to help beta test a Pro version of Flipt, please get in touch with me:

* Twitter: [@mark_a_phelps](https://twitter.com/mark_a_phelps)
* Email: _mark.aaron.phelps at gmail.com_
