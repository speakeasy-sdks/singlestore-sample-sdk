
<div align="center">
    <a href="https://codespaces.new/speakeasy-sdks/singlestore-sample-sdk.git/tree/main"><img src="https://github.com/codespaces/badge.svg" /></a>
</div>
<br>

> **Remember to shutdown a GitHub Codespace when it is not in use!**

# Dev Containers Quick Start

The default location for usage snippets is the `samples` directory.

## Running a Usage Sample

A sample usage example has been provided in a `root.go` file. As you work with the SDK, it's expected that you will modify these samples to fit your needs. To execute this particular snippet, use the command below.

```
go run root.go
```

## Generating Additional Usage Samples

The speakeasy CLI allows you to generate more usage snippets. Here's how:

- For a specific OperationID, use:

```
speakeasy generate usage -s ./openapi.yaml -l go -i operation_id -o ./samples
```

- To generate samples for an entire namespace (like a tag or group name), use:

```
speakeasy generate usage -s ./openapi.yaml -l go -n namespace -o ./samples
```