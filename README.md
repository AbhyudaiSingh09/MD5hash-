
# Go File Upload and MD5 Check Service

This Go application provides a simple web server that allows users to upload files and check the MD5 hash of files stored on the server. It features two main functionalities: uploading files through a POST request and retrieving the MD5 hash of a stored file via a GET request.

## Prerequisites

- Go (1.15 or later recommended)
- Basic knowledge of HTTP requests
- A modern web browser or a tool like `curl` for testing the API endpoints

## Installation

1. Ensure Go is installed on your system. You can download it from [https://golang.org/dl/](https://golang.org/dl/).

2. Clone the repository to your local machine (or download the source code):

   ```bash
   git clone https://github.com/Axs7941/MD5api.git
   ```

3. Navigate to the directory containing the source code.

4. Run the server:

   ```bash
   go run main.go
   ```

   This command compiles and starts the web server on port 8080.

## Usage

### Uploading a File

To upload a file, send a POST request to `/upload` with the file included in the request body. You can use a tool like `curl`:

```bash
curl -F "myfile=@path/to/your/file" http://localhost:8080/upload
```

Replace `path/to/your/file` with the actual file path you wish to upload.

### Checking File MD5 Hash

To get the MD5 hash of an uploaded file, send a GET request to `/checkmd5` with the `filename` query parameter:

```bash
curl "http://localhost:8080/checkmd5?filename=yourfilename"
```

Replace `yourfilename` with the name of the file you uploaded previously.

## Contributing

Contributions to this project are welcome! Here are a few ways you can help:

- Report bugs and issues
- Suggest new features or improvements
- Submit pull requests with bug fixes or new functionalities
