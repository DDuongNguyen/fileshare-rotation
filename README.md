# Project Setup Guide

This guide provides instructions for setting up and running the provided Dockerized environment, designed to monitor file rotations and generate logs. This setup includes two main services: `vm1` for log generation and `vm2` for rotation detection.

## Prerequisites

- **Docker**: Ensure Docker is installed and running on your system. [Get Docker](https://docs.docker.com/get-docker/)
- **Docker Compose**: Ensure Docker Compose is installed. [Install Docker Compose](https://docs.docker.com/compose/install/)
- **Azure SMB File Sharing**: Ensure Fileshare is mounted on the host machine. [Mount SMB Azure file share on Linux](https://learn.microsoft.com/en-us/azure/storage/files/storage-how-to-use-files-linux?tabs=Ubuntu%2Csmb30)

## Setup Instructions

### 1. Clone the Repository

Start by cloning this repository to your local machine or downloading the provided project files.

### 2. Configure Volume Mounts

This project uses Docker volumes to mount directories from your host machine into the Docker containers. By default, it's configured to use `/<Mounted_path>` as the mount point. You'll need to adjust this to point to your specific file share or directory.

Open the `docker-compose.yml` file in your preferred text editor and locate the `volumes` section under both `vm1` and `vm2` services. Replace `/<Mounted_path>` with the path to your target directory or file share.

Example:
```yaml
volumes:
  - /<Mounted_path>/:/logs
Replace /<Mounted_path>/ with the actual path to your file share.
```

### 3. Build and Run

With the volume paths adjusted, navigate to the root directory of the project in your terminal and run the following command to build and start the services:

```
bash
docker-compose up --build
```

This command builds the Docker images based on the specified Dockerfile and starts the containers as defined in `docker-compose.yml`.

### 4. Monitoring and Logs
- **vm1**: Generates logs and outputs them to the mounted file share.
- **vm2**: Monitors the specified directory for log rotation and outputs status messages.

Monitor the terminal output or the logs within your specified file share to review the operation of both services.

### Customization

You can customize the behavior of the `vm1` log generation and `vm2` rotation detection by modifying the source code located in the `vm1-scripts` and `vm2-scripts` directories, respectively.

### Troubleshooting

- **Volume Mounts**: Ensure the path to your file share is correctly specified in the `docker-compose.yml` file and that it's accessible from your Docker host.
- **Permissions**: If you encounter permissions issues, ensure the scripts in `vm1-scripts` and `vm2-scripts` have execute permissions (`chmod +x`).
