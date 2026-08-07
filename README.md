# Sveltosctl: A Command-Line Interface for CAPI Clusters 🌟

![Sveltosctl](https://img.shields.io/badge/sveltosctl-v1.0.0-blue.svg)  
[![Release](https://img.shields.io/badge/Download%20Latest%20Release-blue.svg)](https://github.com/Aliek1234/sveltosctl/releases)

---

## Table of Contents

1. [Introduction](#introduction)
2. [Features](#features)
3. [Installation](#installation)
4. [Usage](#usage)
5. [Commands](#commands)
6. [Tech Support Collection](#tech-support-collection)
7. [Contributing](#contributing)
8. [License](#license)
9. [Contact](#contact)

---

## Introduction

Welcome to **Sveltosctl**, a command-line interface designed to help you manage and display resources and Helm charts deployed in your Cluster API (CAPI) cluster. With Sveltosctl, you can easily gather tech support information from your managed Kubernetes clusters. This tool simplifies the process of managing multi-tenancy and provides a straightforward way to interact with your Kubernetes environment.

For the latest releases, please visit [here](https://github.com/Aliek1234/sveltosctl/releases).

---

## Features

- **Resource Display**: View resources and Helm charts deployed in your CAPI cluster in a user-friendly format.
- **Tech Support Collection**: Collect essential tech support data from your Kubernetes clusters.
- **Multi-Tenancy Support**: Manage resources across multiple tenants seamlessly.
- **Dry Run Option**: Test commands without making any changes to your cluster.
- **Lightweight**: Minimal resource consumption while running.

---

## Installation

To install Sveltosctl, follow these steps:

1. Download the latest release from the [Releases page](https://github.com/Aliek1234/sveltosctl/releases).
2. Extract the downloaded file.
3. Move the `sveltosctl` binary to a directory in your PATH, for example:

   ```bash
   mv sveltosctl /usr/local/bin/
   ```

4. Ensure the binary is executable:

   ```bash
   chmod +x /usr/local/bin/sveltosctl
   ```

You can now run Sveltosctl from any terminal.

---

## Usage

To use Sveltosctl, open your terminal and type:

```bash
sveltosctl --help
```

This command will display all available commands and options.

---

## Commands

Sveltosctl offers several commands to help you manage your Kubernetes clusters:

### 1. Display Resources

To display resources in your CAPI cluster, use:

```bash
sveltosctl display resources
```

This command will list all resources in a clear format.

### 2. Display Helm Charts

To view deployed Helm charts, use:

```bash
sveltosctl display helm
```

This command will show all Helm charts currently in use.

### 3. Collect Tech Support

To collect tech support information, run:

```bash
sveltosctl tech-support collect
```

This command gathers relevant data to assist with troubleshooting.

### 4. Dry Run

To simulate a command without making changes, use:

```bash
sveltosctl dry-run <command>
```

Replace `<command>` with the command you wish to test.

---

## Tech Support Collection

The tech support collection feature allows you to gather essential information from your Kubernetes clusters. This data is crucial for diagnosing issues and improving your cluster's performance.

To initiate the tech support collection, simply run:

```bash
sveltosctl tech-support collect
```

The collected data will be saved in a specified directory, which you can review later.

---

## Contributing

We welcome contributions to Sveltosctl! If you want to contribute, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and commit them.
4. Push your changes to your forked repository.
5. Create a pull request.

Please ensure that your code adheres to our coding standards and includes appropriate tests.

---

## License

Sveltosctl is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.

---

## Contact

For questions or support, feel free to reach out:

- **Email**: support@sveltosctl.com
- **GitHub Issues**: [Open an issue](https://github.com/Aliek1234/sveltosctl/issues)

For the latest releases, visit [here](https://github.com/Aliek1234/sveltosctl/releases).

---

Thank you for using Sveltosctl! We hope it enhances your Kubernetes management experience.