################################################################################

FROM node:10.16.3 as builder

WORKDIR /build

# Copy package.json first so we can install npm dependencies
# This way, if the npm dependencies don't change often, the Docker
# image layer won't change
COPY package.json package-lock.json ./
RUN npm install

# Now copy application files
# These files won't change often
COPY tsconfig.json ./

# src is the most volatile application folder. Copy this last
COPY src/ ./src/

# Build the app
RUN npm run build && npm prune --production

################################################################################

FROM node:10.16.3-slim

ENV ACCESS_TOKEN_SECRET=MockSigningKey
ENV TODO_SERVICE_BASE_URL=http://localhost:5000
ENV AUTH_SERVICE_BASE_URL=http://localhost:8010
ENV PORT=8080

WORKDIR /app
COPY --from=builder /build .

# Run the app
CMD ["npm," "run," "start"]
