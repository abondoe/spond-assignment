# Spond Assignment – Member Registration Wizard
 
A fullstack application for registering member through a multi-step wizard. Built with a **Go backend**, a **React frontend**, and **PostgreSQL** for persistence.
 
---
 
## Architecture
 
```
spond-assignment/
├── backend/     # Go REST API with hot-reloading (air) and TypeScript type generation (tygo)
└── frontend/    # React app built with Vite, Tailwind CSS, and Shadcn UI
```
 
The backend exposes a REST API consumed by the frontend. TypeScript types are generated directly from Go structs using `tygo`, keeping the two layers in sync without manual duplication.
 
---
 
## Getting started
 
### Prerequisites
 
| Tool | Version | Required |
|------|---------|----------|
| Docker & Docker Compose | Latest | ✅ |
| Mise-en-place | Latest | ✅ |
| Tygo | Latest | ✅ |
| Playwright | Latest | ✅ |
 
Find the installation instructions for mise (Mise-en-place) here: https://mise.jdx.dev/getting-started.html

After mise installation, run in project root:
```bash
mise trust
mise install
```

To install tygo globally, run:
```bash
go install github.com/gzuidhof/tygo@latest
```

To install playwright globally:
```bash
npx playwright install
```

### Start everything at once

Mprocs is one of the tools that are installed by mise. Run mprocs to spin up the backend, frontend, and db for development:

If it is the first time you run `mprocs` you must install all code dependencies first:
```bash
npm run install
 ```

Generate the shared types:
```bash
npm run generate:types
```

Then spin up everything:
```bash
mprocs
 ```
 Remember to migrate your DB before trying to register a new member:
 ```bash
npm run db:migrate
 ```

 > Frontend found on: http://localhost:5180/
 
 > Use http://localhost:5180/B171388180BC457D9887AD92B6CCFC86 to access the preconfigured form


 
### Start services individually
 
If you prefer to run services separately (in separate terminals):
 
```bash
# Start the database
npm run db:up
```
```bash
# Start the backend (with hot-reload)
npm run backend:run
```
```bash
# Start the frontend (with hot-reload)
npm run frontend:run
```
 
---
 
## Development
 
### Backend
 
The backend uses [`air`](https://github.com/air-verse/air) for hot-reloading — any changes to `.go` files will automatically restart the server.

```bash
npm run backend:run
```
 
To regenerate TypeScript types after modifying Go structs: 
```bash
npm run generate:types
```

To check for lints:
```bash
npm run backend:lint
# or
npm run backend:lint-fix
```

To format the code:
```bash
npm run backend:format
# or
npm run backend:format-check
```


 
Generated types are written to the frontend so both sides stay in sync.
 
### Frontend
 
The frontend is a React + Vite app styled with [Tailwind CSS](https://tailwindcss.com/) and [Shadcn UI](https://ui.shadcn.com/) components.
 
```bash
npm run frontend:install
npm run frontend:run
```

To check for lints:
```bash
npm run frontend:lint
# or
npm run frontend:lint-fix
```

To format the code:
```bash
npm run frontend:format
# or
npm run frontend:format-check
```
 
---
 
## Database
 
PostgreSQL runs in Docker. The connection is configured via environment variables, or fallbacks to values for local development.
 
```bash
# Start only the database
npm run db:up

# Start only the database as a daemon
npm run db:upd
 
# Stop and remove containers
npm run db:downd

# Migrate
npm run db:migrate
 
# Wipe data
npm run db:clear
```
 
---
 
## Testing
 
```bash
# Backend tests
npm run backend:test
 
# Frontend tests
npm run frontend:test

# End to end test
npm run e2e:upd
# Wait until the system is up
npm run e2e:test
```
---
## Assigment quirks
The ids used in appendix 1 was not standard formatted UUIDs. However, I took the assignment text literal and provided the UUIDs from the backend in this non standard way, but internal to the backend I've used standard UUID types so as to have a strong typed UUID both in the backend and the database. I`ve called the non standard UUIDs compact UUIds in the code, and I've implemented marshalling/unmarshalling to handle the non standard UUID as and ouptut and as an input to the endpoints. Notice, that I did this just to show that this was possible, but I would usually not do it like this in a real world project.

---
## AI usage
In this project I've only used AI sparsely as a search engine when there are things that I need to look up. At no time has an AI read the assignment text. I've also not used any in editor agentic AI, or alike.

---
## Future improvements
- CI with Github actions to prevent merging code until test complete successfully
- CD with Github actions for deployment as needed
- Backend integration tests
- More code documentation
- More unit tests all over