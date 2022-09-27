# Dapr + Wasm Experiments
This repository is an experiment in getting Dapr and Wasm to play together better with the goal of eliminating the need for Dapr as an application sidecar.

Dapr has proven to be wildly popular by providing a common, cloud-native development API eliminating the need for developers to build there applications with dependencies on vendor specific APIs and SDKs. Dapr enables developers to build applications against vendor-agnostic APIs, which enables these applications to be portable across a wide range of environments.

Dapr is deployed as a process that runs next to an application process (sidecar process). This means for each instance of an application process, there is a Dapr process running too. This is not so bad on a small scale, but as the scale of the application increases, the consumption of resources grows quickly. This problem is further exacerbated by the fact that Dapr includes all vendor-specific implementations as part of the compiled binary (>150MB!).

Since Dapr runs as a sidecar proces, there is interprocess communication that needs to take place between the application process and the Dapr process. This communication is done either over HTTP or via gRPC. This introduces network communication overhead.

## Enter Wasm
WebAssembly (Wasm) offers a secure VM for running portable binaries within the process boundary of a hosting application. Wasm can also take advantage of system resources as allowed by the host of the Wasm module through WASI, and can be extended with additional APIs using the WebAssembly Component Model described using WebAssembly Interface Types (WIT).

What if we could use Wasm to allow Dapr to be lighter weight? What if we could compile Dapr capabilities to Wasm and use them directly within an application? Could we eliminate the network boundary between the application process and the sidecar? Could we reduce the resource consumption to provide users of Dapr better return on their cost of goods (CoGs) investment?

## Areas of research

### Dapr components as Wasm modules
Imagine if Dapr components were able to compile to Wasm. A user application would be able to includes these capabilities within its own process space and eliminate the need to run the Dapr sidecar. This would reduce CoGs and increase performance.

#### Wasm State Go Plugin
As an experiment, we implemented a simple state store Wasm plugin and interface using the [Go Plugin System over WebAssembly](https://github.com/knqyf263/go-plugin) to illustrate how Dapr or another host could embed a Go Dapr capability as a Wasm module.

In the [wasm-state-go-plugin](./wasm-state-go-plugin) you will find the result of the experiment. The readme in that directory explains the details of the experiment.

### Dapr as a Wasm host (Future)
Imagine if Dapr was a host for a Wasm application and the Wasm application was able to use the Dapr capabilities via the component model interface. This would enable a user application to run within the Dapr process space, eliminating the network boundary and the need of having a sidecar. This would reduce CoGs and increase performance.

## Disclaimer
This is an experiment and is not intended for production usage.


