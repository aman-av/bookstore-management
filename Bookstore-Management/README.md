# Go Bookstore Management Project

This is a simple Go project for managing books. It provides RESTful APIs for creating, retrieving, updating, and deleting books.

## Getting Started

Follow these instructions to get a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- [Go](https://golang.org/dl/) installed on your machine
- [Git](https://git-scm.com/) for version control

### Installing

1. Clone the repository to your local machine
2. This project uses environment variables for configuration

## Configuration

This project uses environment variables for configuration. Before running the application, create a `.env` file in the `config` folder with the following mysql server details:

```env
# config/.env

# Database Configuration
DB_USERNAME=your_database_username
DB_PASSWORD=your_database_password
DB_HOST=your_database_host
DB_PORT=your_database_port
DB_NAME=your_database_name