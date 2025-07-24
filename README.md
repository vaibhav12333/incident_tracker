# Incident Dashboard (React + Go + Postgres)

A simple dashboard to submit and view incidents, with a React frontend, Go backend, and PostgreSQL database. Includes E2E tests with Cypress.

## Features
- Submit new incidents via a form
- View all incidents in a dashboard
- Modal form for adding incidents
- E2E tests with Cypress

## Project Structure
- `server/` — Go backend (API server)
- `client/tracker/` — React frontend
- `schema.sql` — Database schema (PostgreSQL)

## Prerequisites
- Node.js (v16+ recommended)
- npm
- Go (v1.18+ recommended)
- PostgreSQL (v15+ recommended)

## Setup

### 1. Install dependencies
```sh
cd client/tracker
npm install
```

### 2. Configure the Database
- Create a PostgreSQL database (e.g., `postgres`).
- Update the connection string in `server/main.go`:
  ```go
  db, err := sql.Open("postgres", "postgres://<user>:<password>@localhost:<port>/<dbname>?sslmode=disable")
  ```

### 3. Running Both Client and Server
From the project root (`kafkaP`):
```sh
make start
```
- This will start the Go server and React client together.
- The React app runs at [http://localhost:3000](http://localhost:3000)
- The Go API server runs at [http://localhost:8080](http://localhost:8080)

## Database Schema Management

### Export the current schema
To export your PostgreSQL schema to `schema.sql` in the project root:
```sh
PGPASSWORD='your_password' pg_dump -U your_user -h localhost -p your_port -d your_db -s > schema.sql
```
Example:
```sh
PGPASSWORD='Vgarg@123' pg_dump -U uday -h localhost -p 7455 -d postgres -s > schema.sql
```

### Apply the schema to a new database
```sh
psql -U your_user -h localhost -p your_port -d your_db -f schema.sql
```

## Running Cypress E2E Tests
1. Start both the client and server (`make start`).
2. In another terminal:
   ```sh
   npx cypress open
   ```
3. In the Cypress UI, select a test from `client/tracker/cypress/e2e/` to run.

## More React Scripts
See [`client/tracker/README.md`](client/tracker/README.md) for available React scripts (build, test, etc).

## Customization
- Update API endpoints in the code if your backend runs on a different host/port.
- Add more fields or validation as needed.
- Update the connection string in the Go server as needed for your database.

