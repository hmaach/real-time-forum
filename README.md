# Real-Time Forum

A real-time web forum application built with Go and WebSockets, enabling live user communication through posts, comments, and private messaging.

## Authors

- Hamza Maach
- Yassine Elmach

## Features

### User Authentication
- Registration with comprehensive user details:
  - Nickname
  - Age
  - Gender
  - First Name
  - Last Name
  - Email
  - Password
- Login using nickname/email and password
- Secure session management
- Real-time online/offline status tracking

### Content Management
- Single Page Application (SPA) architecture
- Create and read posts
- Comment on posts
- Multiple category associations for posts

### Private Messaging System
- Real-time private messaging between users
- Online/offline user status display
- Message history
- Message sorting by recent activity
- Message format includes:
  - Timestamp
  - Sender information
  - Message content

## Project Structure

```
real-time-forum/
├── cmd/
│   └── main.go           # Application entry point and server initialization
├── server/
│   ├── api/              # RESTful API routes and WebSocket endpoint definitions
│   ├── config/           # Environment configurations and application settings
│   ├── controllers/      # HTTP handlers and WebSocket message handlers
│   ├── database/         # Database files (database.db, schema.sql, etc)
│   ├── middlewares/      # Request middleware (auth, rate limiting, etc)
│   ├── models/           # Data models and database query methods
│   ├── utils/            # Helper functions
│   └── validators/       # Request payload validation and sanitization
├── web/
│   ├── assets/
│   │   ├── css/           # Application styles and layout definitions
│   │   ├── js/            # Frontend logic and WebSocket client implementation
│   │   └── images/        # Static image resources
│   └── index.html      # Single page application entry point
├── dockerfile       # Container configuration for deployment
├── go.mod           # Go dependencies and module configuration
└── README.md        # Project documentation
```

## Technologies

### Backend
- Go 1.22+
- SQLite3 database
- Gorilla WebSocket for real-time communication
- bcrypt for password hashing
- UUID for session management

### Frontend
- HTML5 & CSS3
- Font Awesome icons
- Vanilla JavaScript (No frameworks)
- WebSocket API
- Single Page Application architecture

### Development & Deployment
- Docker containerization

## Technical Requirements

### WebSocket Implementation
- Real-time message delivery
- Online status updates
- Connection state management
- Error handling and reconnection logic

### Frontend Features
- Throttled/debounced scroll events for message loading
- Dynamic content rendering without page reloads
- Real-time UI updates
- Responsive design

### Database Schema
- Users table with extended profile information
- Messages table for private communications
- Online status tracking
- Session management
- Posts and comments with real-time capabilities

View the detailed database schema [here](https://drawsql.app/teams/zone-01/diagrams/real-time-forum).

## Getting Started

### Prerequisites
- Go 1.22 or higher
- SQLite3
- Docker (optional)

### Local Development

1. **Clone the Repository**
   ```bash
   git clone https://github.com/hmaach/real-time-forum
   cd real-time-forum
   ```

2. **Install Dependencies**
   ```bash
   go mod download
   ```

3. **Run the Application**
   ```bash
   go run cmd/main.go
   ```
   
   Access the forum at `http://localhost:8080`

### Docker Deployment

1. Build and run using Docker:
   ```bash
   docker build -t real-time-forum .
   docker run -p 8080:8080 real-time-forum
   ```

## API Documentation

### WebSocket Endpoints
- `/ws/chat` - Private messaging connection

### HTTP Endpoints
- POST `/api/register` - User registration
- POST `/api/login` - User authentication
- GET `/api/messages/:userId` - Fetch message history
- POST `/api/posts` - Create new post
- POST `/api/comments` - Create new comment
