# Set the base image
FROM node:18-alpine

# Set the working directory
WORKDIR /app

# Copy the package.json and package-lock.json files to the container
COPY package*.json ./

# Install the dependencies
RUN npm install

# Copy the rest of the application code to the container
COPY . .


# Build the application
RUN npm run build

# Expose the port that the application runs on
EXPOSE 3000

# Start the application with astro
CMD ["npm", "run", "start"]

