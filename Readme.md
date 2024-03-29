# URL Shortener Service

![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)

<div align="center">
  <img src="https://res.cloudinary.com/swaraj-cloud/image/upload/v1695746882/wahavuo3msy7yvwgteaa.png" alt="Banner" />
</div>


A Scalable URL shortener service implemented in Golang with Redis and MongoDB for storage, and Docker for containerization.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
  - [Installation](#installation)
  - [Configuration](#configuration)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Features

- Shorten long URLs into easy-to-share links
- Store URL mappings in both Redis and MongoDB for performance and durability
- Retrieve original URLs by visiting the shortened link
- Dockerized for easy deployment

## Prerequisites

Before you begin, ensure you have met the following requirements:

- [Golang](https://golang.org/) installed
- [Redis](https://redis.io/) server installed and running
- [MongoDB](https://www.mongodb.com/) server installed and running
- [Docker](https://www.docker.com/) installed (optional, for containerization)

## Getting Started

### Installation

1. Clone this repository:

   ```bash
   git clone https://github.com/swarajkumarsingh/ziplink.git

   ```

2. Navigate to the project directory:
   ```bash
   cd ziplink
   ```

3. Install the required Go packages:
   ```bash
   go mod tidy
   ```

### Benchmarks
* **POST requests (Create Short-URL):** Average 856.728106ms (shared cloud cluster)
* **GET requests (REDIRECT URL):** Average 5.33ms (Redis, MongoDB)

### Configuration
1. Mongodb client(local, atlas)
2. Redis client(local, cloud)

### Usage
- After installing docker, just run dev.bat(windows), linux and mac run make run(Note: Mongodb, and redis client should be running,)
**NOTE:** Add configs as shown in run.example.sh 
```bash
./run.sh
```


### Contributing
Contributions are welcome! If you'd like to contribute to this project, please follow these guidelines:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and test thoroughly.
4. Commit your changes with clear commit messages.
5. Create a pull request against the main branch.

### License
This project is licensed under the MIT License. See the LICENSE file for details.

