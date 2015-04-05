# Warren

[![Build
Status](https://travis-ci.org/warren-community/warren.svg)](https://travis-ci.org/warren-community/warren)
[![GoDoc](https://godoc.org/github.com/warren-community/warren?status.svg)](https://godoc.org/github.com/warren-community/warren)

Warren is a networked content-sharing site, allowing you to link both yours and
others' creations into a web of works.  It allows posting under multiple
different content types, with new ones being added all the time.

## Developing

Warren requires a working [go](http://golang.org/) environment of at least
version 1.2 with your GOPATH properly set.

You'll need to fetch dependencies; the `sysdeps` target will install Mongo,
ElasticSearch, and a few go dependencies not used in production:

    make sysdeps

The `deps` target will install all of the dependencies required for running
Warren and set them to the versions specified in the `dependencies.tsv` file:

    make deps

The `devel` target will run both gin (a go autobuilder that will build the go
application whenever a file is saved) and the coffeescript autobuilder.  This
way you can develop without having to restart any servers when files change.
Note that this is not an ideal way to run the production server, though.:

    make devel

## Deploying

Warren has been charmed for deploying with Juju, and may be deployed using the
[Warren charm](https://github.com/warren-community/warren-charm).  Warren also
requires Mongo and ElasticSearch, which may be deployed as well and related with
the Warren service:

    juju deploy local:trusty/warren-charm warren
    juju deploy mongodb
    juju deploy elasticsearch
    juju add-relation warren mongodb
    juju add-relation warren elasticsearch
    juju expose warren


Additionally you may use the [Warren
bundle](https://github.com/warren-community/warren-bundle) to deploy all the
services needed, plus haproxy for exposing Warren to the world.  You can deploy
this bundle using the GUI or Juju Quickstart\*:

    juju quickstart warren-bundle/bundle.yaml

\* NB: this currently will not work as local charms in bundles are not yet
supported; once Warren and its bundle wind up in the
[charmstore](http://jujucharms.com), this method of deploying will work.
