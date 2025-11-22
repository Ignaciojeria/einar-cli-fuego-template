# Einar Framework Coding Rules

## Overview
This documentation outlines the coding standards and architectural patterns for the **Einar** framework used in this project. Strict adherence to these rules is required to ensure compatibility with the Einar CLI and the project's dependency injection system.

## Documentation Index

*   **[Controller Creation](controllers_creation.md)**: Detailed rules for creating HTTP controllers, including file placement, naming conventions, code structure, and configuration updates.
*   **[HTTP Server Infrastructure](http_server.md)**: Documentation for the core `httpserver.Server` dependency and its automatic generation rule.

## General Principles

1.  **Virtual CLI Behavior:** When creating components manually or via AI, you must follow the exact same conventions as the Einar CLI.
2.  **IoC Container:** All components must be registered with the `einar-ioc` container in their `init()` functions.
3.  **Configuration:** The `.einar.cli.json` file serves as the registry for project components and must be kept in sync with the codebase.

## 4. Mandatory Main Imports

> [!IMPORTANT]
> **When generating controllers or the HTTP server infrastructure, the LLM MUST inspect (or assume the existence of) `main.go`.**

If the required blank imports do not exist, the LLM **MUST** instruct the user to add them, or must generate the updated `main.go` including:

```go
_ "<module-name>/app/shared/infrastructure/httpserver"
_ "<module-name>/app/adapter/in/fuegoapi"
_ "<module-name>/app/shared/configuration"
```

**Note:** Replace `<module-name>` with the actual module name defined in `go.mod`.

These imports are mandatory for IoC registration and must always be present.

These imports are mandatory for IoC registration and must always be present.
