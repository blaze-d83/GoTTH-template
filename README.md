# GOTTH Stack Template

This repository serves as a fullstack application template built with Go, Templ, TailwindCSS, HTMX, and other modern technologies. It provides a solid foundation for building scalable web applications with server-side rendering, frontend styling, and easy interactivity.

## Project Directory

```plaintext
GOTTH-stack
├── bin/                          # Binary outputs after build
├── cmd/                          # Main application entry points
│   └── main.go                   # Main Go application file
├── config/                       # Configuration files for different environments
├── Dockerfile                    # Docker configuration file for building and running the app
├── internal/                     # Core internal modules (handlers, routes, templates)
│   ├── handlers/                 # API and web route handlers
│   ├── routes/                   # Route definitions
│   └── templates/                # .templ files for server-side rendering (using Templ)
├── node_modules/                 # Frontend dependencies (installed via npm)
├── pkg/                          # Shared packages
│   ├── htmx/                     # HTMX integration and helper functions
│   ├── logger/                   # Logging utility using slog
│   ├── middleware/               # Custom middlewares for Echo framework
│   ├── services/                 # Business logic and services layer
│   ├── types/                    # Custom types and models
│   ├── utils/                    # Utility functions and helpers
│   └── validator/                # Input validation utilities
├── static/                       # Static assets (e.g., JS, CSS, images)
│   ├── styles.css                # Main input CSS file for Tailwind
│   └── dist/
│       └── output.css            # Compiled CSS output (after Tailwind build)
├── store/                        # Database connection and setup (SQLite, MySQL, PostgreSQL)
├── utils/                        # Additional utility functions
├── .gitignore                    # Files and directories to ignore in git
├── Makefile                      # Makefile for building, running, and testing the app
├── go.mod                        # Go module definition
├── go.sum                        # Go module dependencies lock file
└── README.md                     # Project documentation
```

## Features

- **Go backend**: Fast, scalable, and lightweight using the Echo framework.
- **Templ**: Templating language for server-side rendering with `.templ` files.
- **TailwindCSS**: Utility-first CSS framework for fast and modern styling.
- **HTMX**: Enables partial updates and enhances interactivity without JavaScript frameworks.
- **SQLite/MySQL/PostgreSQL support**: Default database setup with SQLite, easily configurable to use MySQL, PostgreSQL, or NoSQL.
- **Viper for CLI configuration**: Environment-based configuration management.
- **Logging with slog**: Structured logging for better observability.
- **Validation**: Input validation using `validator` for form and API request safety.
- **Cloud Storage Scalability**: Extendable to include Google Cloud, AWS S3, or any other cloud provider for static asset and data storage.
- **Docker Support**: Containerized application for easy deployment on cloud platforms like AWS and Google Cloud.

## How to Use

### 1. Clone the repository

```bash
git clone https://github.com/yourusername/GOTTH-stack.git
cd GOTTH-stack
```

### 2. Initialize and Configure TailwindCSS

Since the `package.json` and `tailwind.config.js` files are not pushed, you will need to initialize the project and set up TailwindCSS:

1. **Initialize npm**:

   ```bash
   npm init -y
   ```

2. **Install TailwindCSS**:

   ```bash
   npm install tailwindcss postcss autoprefixer
   ```

3. **Create a TailwindCSS configuration file**:

   ```bash
   npx tailwindcss init
   ```

4. **Configure `tailwind.config.js`** by setting your paths:

   ```javascript
   module.exports = {
     content: ["./internal/templates/**/*.templ"],
     theme: {
       extend: {},
     },
     plugins: [],
   }
   ```

5. **Set up your CSS files**:

   - Create an input CSS file for Tailwind in `/static/styles.css`:

     ```css
     @tailwind base;
     @tailwind components;
     @tailwind utilities;
     ```

   - The output of the compiled CSS will go into `/static/dist/output.css`.

6. **Run the Tailwind build process**:

   Add the following script to your `package.json`:

   ```json
   "scripts": {
     "build:css": "tailwindcss -i ./static/styles.css -o ./static/dist/output.css --watch"
   }
   ```

   Then run:

   ```bash
   npm run build:css
   ```

### 3. Install Go dependencies

```bash
go mod tidy
```

### 4. Run the application

- **Development**:

  ```bash
  make run
  ```

- **Build for production**:

  ```bash
  make build
  ./bin/yourapp
  ```

### 5. Docker Setup

The project includes a `Dockerfile` to easily containerize the application.

1. **Build the Docker image**:

   ```bash
   docker build -t yourapp:latest .
   ```

2. **Run the Docker container locally**:

   ```bash
   docker run -p 8080:8080 yourapp:latest
   ```

### 6. Deploying Docker Image to AWS

To deploy the Docker image on **AWS** using **Amazon Elastic Container Service (ECS)** or **Amazon Elastic Container Registry (ECR)**:

1. **Login to AWS**:

   ```bash
   aws configure
   ```

2. **Push Docker image to ECR**:

   - Create a repository in ECR:
   
     ```bash
     aws ecr create-repository --repository-name yourapp
     ```

   - Authenticate Docker to ECR:

     ```bash
     aws ecr get-login-password --region your-region | docker login --username AWS --password-stdin <your-account-id>.dkr.ecr.<your-region>.amazonaws.com
     ```

   - Tag your image:

     ```bash
     docker tag yourapp:latest <your-account-id>.dkr.ecr.<your-region>.amazonaws.com/yourapp:latest
     ```

   - Push the image to ECR:

     ```bash
     docker push <your-account-id>.dkr.ecr.<your-region>.amazonaws.com/yourapp:latest
     ```

3. **Deploy to ECS**:
   - Use **Amazon ECS** (Elastic Container Service) with **Fargate** for running your containerized application.
   - You can use the ECS console to create a task definition, configure services, and deploy the container image.

### 7. Deploying Docker Image to Google Cloud

To deploy the Docker image on **Google Cloud** using **Google Cloud Run** or **Google Kubernetes Engine (GKE)**:

1. **Login to Google Cloud**:

   ```bash
   gcloud auth login
   ```

2. **Push Docker image to Google Container Registry**:

   - Tag your image:

     ```bash
     docker tag yourapp:latest gcr.io/your-project-id/yourapp:latest
     ```

   - Push the image to Google Container Registry:

     ```bash
     docker push gcr.io/your-project-id/yourapp:latest
     ```

3. **Deploy to Google Cloud Run**:

   - Deploy the image to **Google Cloud Run** (a fully managed platform for containerized applications):

     ```bash
     gcloud run deploy yourapp --image gcr.io/your-project-id/yourapp:latest --platform managed
     ```

   - Follow the prompts to configure the deployment region and service options.

4. **Deploy to Google Kubernetes Engine (GKE)**:

   - Create a GKE cluster:

     ```bash
     gcloud container clusters create your-cluster-name --zone your-zone
     ```

   - Deploy your container to GKE:

     ```bash
     kubectl create deployment yourapp --image=gcr.io/your-project-id/yourapp:latest
     ```

   - Expose the deployment:

     ```bash
     kubectl expose deployment yourapp --type=LoadBalancer --port 8080
     ```

---

