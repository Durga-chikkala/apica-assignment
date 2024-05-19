# In-Memory Cache with Real-Time Updates

This repository contains code for implementing an in-memory cache with real-time updates. The application is divided into two main components: the frontend, developed using React, and the backend, written in Go.

## Features

- **LRU Cache**: Implements an in-memory cache with Least Recently Used (LRU) eviction policy.
- **Real-Time Updates**: Utilizes WebSocket to provide real-time updates to the frontend whenever the cache is modified.
- **RESTful API**: Exposes RESTful endpoints to interact with the cache, including setting, getting, and deleting cache entries.
- **WebSocket Integration**: Enables bidirectional communication between the backend and frontend for instantaneous updates.

## Frontend

The frontend component is responsible for the user interface and client-side logic of the application.

### Technologies Used:

- React: For building interactive user interfaces.
- Axios: Handles HTTP requests to the backend API.
- WebSocket: Facilitates real-time communication with the backend server.

### Getting Started:

1. **Navigate to the `frontend` directory:**
   ```bash
   cd frontend
   ```

2. Install dependencies:
    ```bash
    npm install
    ```
3. Start the development server:
    ```bash
    npm start
    ```
### Folder Structure:
- `src/`: Contains the source code for the React application.
- `public/`: Holds static assets and the HTML file where the React app is rendered.
  
## Backend
The backend component handles the core logic of the application, including cache management and communication with the front end.

### Technologies Used:
- Go: Programming language for backend development.
- Gorilla WebSocket: Library for WebSocket implementation in Go.
- Gorilla Mux: Router for HTTP request handling.

### Getting Started:
1. **Navigate to the `backend` directory:**
    ```bash
    cd backend
     ```
2. **Install dependencies using Go modules:**
    ```bash
    go mod tidy
    ```
3. **Run the backend server:**
    ```bash
    go run main.go
    ```
### Folder Structure:
- `main.go`: Entry point for the backend application.
- Other Go files: Contains the backend logic




