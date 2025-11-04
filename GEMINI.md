## Development Guidelines

### Project Layout

This project follows the standard Go project layout as described
in [https://github.com/golang-standards/project-layout](https://github.com/golang-standards/project-layout).

- `/cmd`: Main applications for this project.
- `/internal`: Private application and library code. This is the code you don't want others importing in their
  applications or libraries.
- `/pkg`: Library code that's ok to use by external applications.
- `/api`: OpenAPI specs, JSON schema files, protocol definition files.
- `/scripts`: Scripts to perform various build, install, analysis, etc operations.

### Design Patterns

Development should adhere to the principles of the Gang of Four (GoF) design patterns where applicable.

### API Documentation

The API documentation will be generated automatically using `swaggo`. All handlers should be annotated to provide the
necessary information for the documentation generation.

### Authentication

Authentication is handled using JWT. A valid token is required to access protected endpoints.

### Code Style

Follow the standard Go formatting and style guidelines. Use `gofmt` and `golint` to ensure consistency.
