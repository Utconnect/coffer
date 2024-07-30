# Coffer
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=Utconnect_coffer&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=Utconnect_coffer)

Coffer is an API service that simplifies secret management in a microservice architecture. Instead of making direct HTTP requests to HashiCorp Vault, services can request secrets from Coffer, streamlining and securing the process.

## Features

- **Centralized Secret Access**: Unified API for retrieving secrets across various services.
- **HashiCorp Vault Integration**: Securely manage and access secrets using Vault.
- **Simplified Microservice Communication**: Reduce direct dependencies on Vault within individual services.

## Getting Started

### Prerequisites

- Docker
- Docker Compose
- Go (for development)

### Installation

1. Clone the repository:
  ```sh
  git clone https://github.com/Utconnect/coffer.git
  cd coffer
  ```

2. Build and run the Docker container:
  ```sh
  docker-compose up --build
  ```

3. Enter the Docker container and run the unseal script:
  ```sh
  docker exec -it <container_id> /bin/sh
  source /vault/config.d/init.sh
  ```

## Usage

Once set up, services can call the Coffer API to fetch necessary secrets. Refer to the documentation for detailed API endpoints and usage instructions.

## Contributing

We welcome contributions! Please fork the repository and submit a pull request with your improvements or fixes.

## License

This project is licensed under the GNU General Public License v3.0. See the [LICENSE](LICENSE) file for more details.

## Contact

For questions or support, please open an issue in the repository.
