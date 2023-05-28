# Gigachat

Welcome to Gigachat! This is a simple chat application built with React, PostgreSQL, Go.

## Coursework

This work was part of Nikita Sheverdov's coursework, so I advise you to study the [course report](https://www.overleaf.com/read/pncrvxygzjfp)

## Project Structure

The project is organized into the following directories:

- **DevOps**: Contains deployment and infrastructure-related files.
- **Core**: Contains the core code and modules for the Gigachat application.

## DevOps

The `DevOps` directory includes files and scripts related to deployment and infrastructure. It may contain files such as Kubernetes manifests, Dockerfiles, CI/CD 
configuration, and scripts for automating deployment.

## Core

The `Core` directory contains the core code and modules for the Gigachat application. It includes the frontend, backend, and other supporting modules.

- **Auth**: Microservice written in Go, responsible for user authentication.
- **Chat**:  Microservice written in Go, responsible for chat logic.
- **Frontend**:  Microservice written in React, responsible for frontend.
- **PostgreSQL**: Contains code and scripts for database operations.

## Getting Started

To get started with the Gigachat project, follow the instructions below:

1. Clone the repository: `git clone https://github.com/your-username/gigachat.git`
2. Expand each of the directories in Core in any convenient container registry
3. Modify the kubernetes charts and manifests in the DevOps helm directory (Helm and K8s-manifests directories) so that you can deploy your project in kubernetes

## Contributing

We welcome contributions to the Gigachat project! If you find any issues or have suggestions for improvements, please submit them through GitHub issues. If you'd like to 
contribute code, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature/bug fix: `git checkout -b feature-name`.
3. Make your changes and commit them: `git commit -am 'Add feature'`.
4. Push the changes to your fork: `git push origin feature-name`.
5. Submit a pull request.

Please make sure to follow our [Code of Conduct](./CODE_OF_CONDUCT.md) when contributing.

## License

The Gigachat project is licensed under the [MIT License](./LICENSE).

