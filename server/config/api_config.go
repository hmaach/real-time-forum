package config

var APIs = map[string]string{
	"GET  /api":               "List all available API routes and their descriptions.",
	"GET  /api/posts":         "Retrieve a list of all posts.",
	"GET  /api/posts/{id}":    "Retrieve details of a specific post. Parameters: {id} (integer) - the ID of the post.",
	"POST /api/posts/create": "Create a new post. Parameters: title (string), content (string), authorID (integer).",
	"POST /api/posts/react":  "React to a post. Parameters: postID (integer), reaction (string).",
	"POST /api/comment/create": "Create a new comment on a post. Parameters: postID (integer), content (string), authorID (integer).",
	"POST /api/comment/react":  "React to a comment. Parameters: commentID (integer), reaction (string).",
	"POST /api/register":     "Register a new user. Parameters: username (string), email (string), password (string).",
	"POST /api/login":        "Log in an existing user. Parameters: email (string), password (string).",
	"POST /api/logout":       "Log out the current user. No parameters required.",
}
