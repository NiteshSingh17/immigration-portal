# Northway Immigration Website

A responsive immigration website built with Goravel, Tailwind CSS, and MySQL.

## Features

- Immigration service website
- Contact/enquiry form
- Enquiries stored in MySQL
- Protected admin panel
- Responsive layout

## Requirements

- Go 1.25 or newer
- Node.js and npm
- MySQL database

## Setup

1. Clone the repository:

```bash
git clone <repository-url>
cd imegration-website
```

2. Copy the environment file:

```bash
cp .env.example .env
```

3. Add your configuration to `.env`:

```env
APP_HOST=127.0.0.1
APP_PORT=4001

DB_CONNECTION=mysql
DB_HOST=your-mysql-host
DB_PORT=3306
DB_DATABASE=your_database
DB_USERNAME=your_username
DB_PASSWORD=your_password

ADMIN_USERNAME=admin
ADMIN_PASSWORD=your_password
```

4. Generate the application key:

```bash
./artisan key:generate
```

5. Install dependencies:

```bash
go mod tidy
npm install
```

6. Build the Tailwind stylesheet:

```bash
npm run build:css
```

7. Create the database tables:

```bash
./artisan migrate
```

8. Start the website:

```bash
go build -o main .
./main
```

Open: http://127.0.0.1:4001

## Admin Panel

Open: http://127.0.0.1:4001/admin

Sign in using the `ADMIN_USERNAME` and `ADMIN_PASSWORD` values from `.env`.

## Development

Watch Tailwind changes:

```bash
npm run dev:css
```

Run the Goravel server:

```bash
go run .
```

## Database

Enquiries are stored in the `enquiries` MySQL table.

The table contains:

- Name
- Email
- Phone
- Service
- Message
- Created date

## Important

Do not commit `.env`. It contains database credentials and admin login details.
