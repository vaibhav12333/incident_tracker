# Incident Dashboard (React + GO + Postgres)

A simple React dashboard to submit and view incidents, with integration to a Go backend and Cypress E2E tests.

## Features
- Submit new incidents via a form
- View all incidents in a dashboard
- Modal form for adding incidents
- E2E tests with Cypress

## Getting Started

### Prerequisites
- Node.js (v16+ recommended)
- npm
- Go backend running at `http://localhost:8080`

### Setup
```sh
cd tracker
npm install
```

### Running the App
```sh
npm start
```
The app will run at [http://localhost:3000](http://localhost:3000)

### Running Cypress E2E Tests
1. Start the React app (`npm start`) and ensure the backend is running.
2. In another terminal:
   ```sh
   npx cypress open
   ```
3. In the Cypress UI, select a test from `cypress/e2e/` to run.

### Project Structure
- `src/Components/` - React components (Dashboard, IncidentForm)
- `cypress/e2e/` - Cypress E2E tests
- `main.css` - App and modal styling

## Customization
- Update API endpoints in the code if your backend runs on a different host/port.
- Add more fields or validation as needed.
- Update Connection string in the code acc to the path of your database

