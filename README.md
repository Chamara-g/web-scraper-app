# Web Scraper APP

Go REST API and React Frontend

## Demo

https://github.com/user-attachments/assets/3c183492-ca78-454a-b7d9-f23f1c1279fd

## Clone Application

1.  Clone the repository (main branch)

```
git clone https://github.com/Chamara-g/web-scraper-app
```

## Run REST API

1.  Go to backend directory

```
 cd backend
```

2.  Clean up and manage dependencies

```
 go mod tidy
```

3. Run go Backend

```
 go run cmd/api-server/main.go
```

## Run React Frontend

1.  Go to frontend directory

```
 cd frontend
```

2.  Copy env-sample file and create .env file to add environment data

3.  Install packages

```
 npm install
```

3. Run frontend

```
 npm start
```

## Dockerize APP

### run both frontend and backend using compose

```
docker-compose up --build

// access urls
React App - http://localhost:3000
Go REST API - http://localhost:8080
```

### Dockerize Go REST API

```
// build
docker build -t go-rest-api .

// run
docker run -p 8080:8080 go-rest-api
```

### Dockerize React APP

```
// build
docker build -t react-app .

// run
docker run -p 3000:3000 react-app
```

## For developers

### Branching strategy

- **main** : production
- **staging** : production ready
  - release/{version} : release
  - hotfix/{id}-{name} : hotfix
- **develop** : stable development
  - develop : production ready
  - feature/{id}-{name} : feature
  - bugfix/{id}-{name} : fix
  - refactor/{id}-{name} : refactor
  - docs/{id}-{name} : documentation
