# Forum Application

A comprehensive web forum application built using Go that enables user communication through posts, comments, and reactions.

## Authors

- Abdelhamid Bouziani
- Hamza Maach
- Omar Ait Benhammou
- Mehdi Moulabbi
- Youssef Basta

## Features

### User Authentication
- Username-based registration and login
- Secure session management using cookies

### Content Management
- Create and read posts
- Comment on posts
- Multiple category associations for posts

### Interaction System
- Like/dislike functionality for posts and comments
- Comprehensive user engagement tools

### Content Discovery
- Filter posts by categories
- Filter posts by creation date

## Project Structure

```
forum/
├── cmd/
│   └── main.go           # Application entry point
├── server/
│   ├── config/           # Configuration management
│   ├── database/         # Database interaction logic
│   ├── controllers/      # Request handling and business logic
│   ├── models/           # Data structures and models
│   ├── routes/           # Application routing
│   └── utils/            # Shared utility functions
├── web/ 
│   ├── assets/           # Static resources (CSS, JS, images)
│   └── templates/        # HTML templates
├── dockerfile            # Docker containerization
├── commands.sh           # Docker build and deployment script
├── prune.sh              # remove unused objects
├── go.mod                # Go module dependencies
├── go.sum                # Dependency checksum
└── README.md             # Project documentation
```

## Database Schema

View the detailed database schema [here](https://drawsql.app/teams/zone-01/diagrams/forum-db).

### Key Tables
- **Users**: User authentication and profile information
- **Posts**: Forum post content
- **Comments**: Post responses and discussions
- **Categories**: Post classification
- **Categories_Posts**: Post-category relationships
- **Posts_Reactions**: Post interaction tracking
- **Comments_Reactions**: Comment interaction tracking
- **Sessions**: User authentication state management

## Technologies

### Backend
- Go 1.22+
- SQLite3 database
- bcrypt for password hashing

### Frontend
- HTML5 & CSS3
- JavaScript
- Font Awesome icons

### Development & Deployment
- Docker containerization

## Getting Started

### Prerequisites
- Go 1.22 or higher
- SQLite3
- Docker (optional)

### Local Development

1. **Clone the Repository**
   ```bash
   git clone https://github.com/hmaach/forum.git
   cd forum
   ```

2. **Install Dependencies**
   ```bash
   go mod download
   ```



3. **Run the Application**
   ```bash
   cd cmd
   go run .
   ```
   
   Access the forum at `http://localhost:8080`

### Docker Deployment

1. Make script executable:
   ```bash
   chmod +x commands.sh
   ```

2. Run deployment script:
   ```bash
   ./commands.sh
   ```

3. Access the forum at `http://localhost:8080`


4. To delete created images and containers, run the script:
   ```bash
   ./prune.sh
   ```

## Contributing

Please read our contributing guidelines before submitting pull requests or issues.
