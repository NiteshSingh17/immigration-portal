# Immigration Portal

Simple immigration portal application.

## How to run

### Prerequisites
- Go (1.21 or later)

### Run the app

```bash
# Clone the repo
git clone git@github.com:NiteshSingh17/immigration-portal.git
cd immigration-portal

# Checkout the web-2 branch
git checkout web-2

# Copy the env file
cp .env.example .env

# Run the app
go run main.go
```

The app will start locally. Visit `http://localhost:8000` in your browser (or whichever port your config uses).

### Notes
- `.env` is not tracked — copy from `.env.example` and fill in your values.
- Binaries (`goravel`, `goravel-app`, `tmp/`) and logs are ignored.
