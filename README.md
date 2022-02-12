Nomad Await Dependency Plugin
==========

Build the plugin.

```sh
$ make build
```

## Deploying Driver Plugins in Nomad

The initial version of the skeleton is a simple task that outputs a greeting.
You can try it out by starting a Nomad agent and running the job provided in
the `example` folder:

```sh
$ make build
$ nomad agent -dev -config=./example/agent.hcl -plugin-dir=$(pwd)

# in another shell
$ nomad run ./example/example.nomad
$ nomad logs <ALLOCATION ID>
```

Code Organization
-------------------
Follow the comments marked with a `TODO` tag to implement your driver's logic.
For more information check the
[Nomad documentation on plugins](https://www.nomadproject.io/docs/internals/plugins/index.html).
