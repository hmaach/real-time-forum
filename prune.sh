# Stop and remove existing container if it exists
docker stop forum-con 
docker rm forum-con

# Remove the old image
docker rmi forum-img 

# Remove unused data
docker system prune -f