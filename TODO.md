# TODO List for GOTTH Stack Template

## Setup

- [ ] **Clone the Repository**
  - Clone the repository to your local machine.
  - ```bash
    git clone https://github.com/yourusername/go-fullstack.git
    ```

- [ ] **Install Dependencies**
  - Navigate to the project directory.
  - Initialize npm and install TailwindCSS:
    ```bash
    cd go-fullstack
    npm init -y
    npm install -D tailwindcss postcss autoprefixer
    npx tailwindcss init
    ```
  - [ ] Configure tailwind.config.js by setting your paths:

   ```javascript
   module.exports = {
     content: ["./internal/templates/**/*.templ"],
     theme: {
       extend: {},
     },
     plugins: [],
   }
   ```
   - [ ] Create an input CSS file for Tailwind in `/static/styles.css`:

     ```css
     @tailwind base;
     @tailwind components;
     @tailwind utilities;
     ```

   - The output of the compiled CSS will go into `/static/dist/output.css`.

   - [ ] Add the following script to your `package.json`:
   ```json
   "scripts": {
     "build:css": "tailwindcss -i ./static/styles.css -o ./static/dist/output.css --watch"
   }
   ```
   Then run:

   ```bash
   npm run build:css
   ```

- [ ] **Configure Templ**
  - Ensure Templ is installed:
    ```bash
    templ version
    ```
    - If not installed, follow the instructions at [Templ Installation](https://templ.guide/quick-start/installation).

- [ ] **Generate Templ Files**
  - Run the Templ command to generate `*_templ.go` files:
    ```bash
    templ generate
    ```

## Development

- [ ] **Implement Features**
  - Use the Echo framework to implement application routes and handlers.
  - Use the `slog` package for logging throughout the application.
  - Implement validation using the `validator` package.

- [ ] **Testing**
  - Write unit tests for handlers and services.
  - Test the application locally to ensure functionality.

## Configuration

- [ ] **Database Configuration**
  - Choose the database you want to use (SQLite, MySQL, PostgreSQL, etc.).
  - Update the configuration files in `/config` to match your database settings.

- [ ] **Cloud Storage Configuration**
  - Create appropriate subdirectories for cloud storage configurations (Google Cloud Bucket, AWS S3, etc.).
  - Modify configuration files to include cloud storage credentials and settings.

## Docker Setup

- [ ] **Create Dockerfile**
  - Ensure the Dockerfile is set up correctly (as provided).

- [ ] **.dockerignore Configuration**
  - Create a `.dockerignore` file to exclude unnecessary files from the Docker image.

## Deployment

- [ ] **Build Docker Image**
  - Build the Docker image:
    ```bash
    docker build -t yourapp:latest .
    ```

- [ ] **Deploy on Cloud Provider**
  - Choose a cloud provider (AWS, Google Cloud, etc.) and follow the appropriate steps:
    - **For AWS**
      - Set up an ECS or EKS cluster.
      - Push the Docker image to Amazon ECR.
      - Deploy the service using ECS.
    - **For Google Cloud**
      - Push the Docker image to Google Container Registry (GCR).
      - Deploy the service using Google Kubernetes Engine (GKE) or Cloud Run.

## Maintenance

- [ ] **Monitoring**
  - Set up monitoring for your application (e.g., using Prometheus, Grafana).
  
- [ ] **Logging**
  - Ensure that logs are being captured and stored appropriately.

- [ ] **Updates**
  - Regularly update dependencies and libraries.
  - Maintain documentation for any new features or configurations.

## Documentation

- [ ] **Update README.md**
  - Ensure the README.md is updated with the latest information on setup, usage, and features.
  
- [ ] **Create Deployment Documentation**
  - Export deployment information to `DEPLOYMENT.md`.

## Future Enhancements

- [ ] **User Authentication**
  - Implement user authentication and authorization (e.g., JWT, OAuth).

- [ ] **Testing Frameworks**
  - Consider using testing frameworks like Go testing, Postman for API tests, etc.

- [ ] **CI/CD Pipeline**
  - Set up a CI/CD pipeline for automated testing and deployment.

- [ ] **Feature Enhancements**
  - List any additional features or improvements you want to add in the future.


