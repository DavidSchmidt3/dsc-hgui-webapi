param (
    $command
)

if (-not $command)  {
    $command = "start"
}

$ProjectRoot = "${PSScriptRoot}/.."

$env:HGUI_API_ENVIRONMENT="Development"
$env:HGUI_API_PORT="8080"
$env:HGUI_API_MONGODB_USERNAME="root"
$env:HGUI_API_MONGODB_PASSWORD="neUhaDnes"

switch ($command) {
    "start" {
        go run ${ProjectRoot}/cmd/dsc-hgui-api-service
    }
    "openapi" {
        docker run --rm -ti -v ${ProjectRoot}:/local openapitools/openapi-generator-cli generate -c /local/scripts/generator-cfg.yaml
    }
    "docker" {
        docker build -t davidschmidt38/hgui-webapi-webapi:local-build -f ${ProjectRoot}/build/docker/Dockerfile .
    }
    "test" {
        go test -v ./...
    }
    default {
        throw "Unknown command: $command"
    }
}