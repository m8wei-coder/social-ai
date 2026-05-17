# Social AI

Social AI is a full-stack web application with a Go backend, React frontend, and Elasticsearch database.

## Tech Stack

### Backend

- Go
- Gorilla Mux
- Elasticsearch Go client

### Frontend

- React

### Database

- Elasticsearch 7.x

### Development Tools

- SSH
- VS Code Remote SSH
- Vim
- Java Runtime Environment

---

## Project Structure

```bash
social-ai/
├── backend/
│   ├── go.mod
│   ├── go.sum
│   └── ...
├── frontend/
│   └── ...
├── README.md
└── .gitignore
````

---

## Backend Setup

### Go Version

```bash
go version
```

Expected version:

```bash
go version go1.26.3 linux/amd64
```

> Note: Please make sure the Go version matches your local or remote environment.

---

### Install Go on Ubuntu

```bash
sudo apt-get update
sudo apt install -y software-properties-common
sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt-get update
sudo apt-get install -y golang-go
```

Verify installation:

```bash
go version
```

---

### Initialize Go Module

Go to the backend directory:

```bash
cd projects/socialai/backend
```

Initialize the Go module:

```bash
go mod init socialai
```

---

### Install Backend Dependencies

Install Gorilla Mux:

```bash
go get github.com/gorilla/mux
```

Install Elasticsearch Go client:

```bash
go get github.com/olivere/elastic/v7
```

```bash
go get github.com/pborman/uuid
go get cloud.google.com/go/storage
```

```bash
go get github.com/form3tech-oss/jwt-go
go get github.com/auth0/go-jwt-middleware
```

```bash
go get github.com/gorilla/handlers
```

After installing dependencies, run:

```bash
go mod tidy
```

---

## Frontend Setup

The frontend is built with React.

Go to the frontend directory:

```bash
cd projects/socialai/frontend
```

Install dependencies:

```bash
npm install
```

Start the development server:

```bash
npm start
```

Or, if using Vite:

```bash
npm run dev
```

---

## Elasticsearch Setup

### Install Elasticsearch 7.x on Ubuntu

Install required packages:

```bash
sudo apt install -y apt-transport-https
```

Add the Elasticsearch GPG key:

```bash
wget -qO - https://artifacts.elastic.co/GPG-KEY-elasticsearch | sudo apt-key add -
```

Add the Elasticsearch 7.x repository:

```bash
sudo sh -c 'echo "deb https://artifacts.elastic.co/packages/7.x/apt stable main" > /etc/apt/sources.list.d/elastic-7.x.list'
```

Update package list:

```bash
sudo apt update
```

Install Elasticsearch:

```bash
sudo apt install -y elasticsearch
```

---

### Configure Elasticsearch

Open the Elasticsearch config file:

```bash
sudo vim /etc/elasticsearch/elasticsearch.yml
```

Update or add the following configuration:

```yml
network.host: 0.0.0.0
http.port: 9200

discovery.type: single-node

xpack.security.enabled: true
```

Verify the effective configuration:

```bash
sudo cat /etc/elasticsearch/elasticsearch.yml | grep "^[^#;]"
```

---

### Start Elasticsearch

Enable Elasticsearch to start automatically when the GCE instance starts:

```bash
sudo systemctl enable elasticsearch
```

Start Elasticsearch:

```bash
sudo systemctl start elasticsearch
```

Check Elasticsearch status:

```bash
sudo systemctl status elasticsearch
```

Press `q` to exit the status page.

---

### Create Elasticsearch User

Create a new Elasticsearch user:

```bash
sudo /usr/share/elasticsearch/bin/elasticsearch-users useradd YOUR_NEW_USER_NAME -p YOUR_NEW_USER_PASSWORD -r superuser
```

Replace:

```bash
YOUR_NEW_USER_NAME
YOUR_NEW_USER_PASSWORD
```

with your actual username and password.

---

## SSH Setup

### Generate SSH Key

Run the following command on your local machine:

```bash
ssh-keygen -t rsa -f ~/.ssh/gcekey -C YOUR_USERNAME
```

Verify that both public and private keys are created:

```bash
ls ~/.ssh/
```

Print the public key:

```bash
cat ~/.ssh/gcekey.pub
```

The public key should be in this format:

```bash
ssh-rsa KEY_VALUE YOUR_USERNAME
```

---

## VS Code Remote SSH Setup

### Install Remote SSH Extension

1. Open VS Code.
2. Click the Extensions button.
3. Search for `Remote SSH`.
4. Install the first result.

---

### Add SSH Target

Open VS Code Remote Explorer.

Under the SSH section, click the plus button to create a new SSH target.

Use the following command:

```bash
ssh -i ~/.ssh/gcekey YOUR_USERNAME@YOUR_GCE_EXTERNAL_IP_ADDRESS
```

Replace:

```bash
YOUR_USERNAME
YOUR_GCE_EXTERNAL_IP_ADDRESS
```

with your actual GCE username and external IP address.

---

### Update SSH File Permissions

Go to the SSH directory:

```bash
cd ~/.ssh
```

Update permissions:

```bash
chmod 600 config
chmod 600 gcekey
```

---

## Common Tools

### Install Vim

```bash
sudo apt install -y vim
```

---

### Install Java Runtime

Elasticsearch requires Java.

Install Java Runtime Environment:

```bash
sudo apt install -y default-jre
```

Check Java version:

```bash
java -version
```

Expected major version:

```bash
11
```

---

## Useful Commands

### Backend

Go to backend directory:

```bash
cd projects/socialai/backend
```

Run Go app:

```bash
go run main.go
```

Format Go code:

```bash
go fmt ./...
```

Clean and verify dependencies:

```bash
go mod tidy
```

---

### Elasticsearch

Start Elasticsearch:

```bash
sudo systemctl start elasticsearch
```

Stop Elasticsearch:

```bash
sudo systemctl stop elasticsearch
```

Restart Elasticsearch:

```bash
sudo systemctl restart elasticsearch
```

Check status:

```bash
sudo systemctl status elasticsearch
```

Test Elasticsearch connection:

```bash
curl http://localhost:9200
```

If security is enabled:

```bash
curl -u YOUR_NEW_USER_NAME:YOUR_NEW_USER_PASSWORD http://localhost:9200
```

---

## Notes

* Keep backend code inside the `backend/` directory.
* Keep frontend code inside the `frontend/` directory.
* Do not commit sensitive files such as `.env`, private keys, or passwords.
* Make sure Elasticsearch is running before starting the backend service.
* If using GCE, make sure firewall rules allow access to the required ports.










