FROM node:21-alpine AS builder

WORKDIR /app

# Copy package.json and package-lock.json (or yarn.lock)
COPY src/package.json  ./

# Install dependencies
RUN npm install

# Copy the rest of the application code

COPY service.codefly.yaml .

COPY src .


# Build the application
RUN npm run build

# Production environment
FROM node:21-alpine AS production

# Set working directory
WORKDIR /app


# Copy package.json and package-lock.json (or yarn.lock) for production dependencies
COPY --from=builder /app/package*.json ./

# Install only production dependencies
RUN npm install --omit=dev

# Set up a non-root user for better security
RUN adduser -D appuser
USER appuser

# Copy the built Next.js application and other necessary files from the builder stage
COPY --chown=appuser --from=builder /app/.next ./.next
COPY --chown=appuser --from=builder /app/service.codefly.yaml ./service.codefly.yaml

# Environment variable to specify production mode
ENV NODE_ENV production

EXPOSE 3000


COPY service.codefly.yaml .

COPY src .


# Command to start the Next.js application in production mode
CMD ["npm", "run", "start"]
