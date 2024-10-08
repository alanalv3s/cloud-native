# Use Node.js LTS version for building the app
FROM node:20.18.0-alpine AS build-ts

WORKDIR /app

# Copy package.json and package-lock.json for installing dependencies
COPY package*.json ./
# Install dependencies
RUN npm ci
# Copy the source code
COPY . .
# Build the TypeScript code
RUN npm run build

# Use a lightweight Node.js image for running the app
FROM node:20.18.0-alpine AS build-deps

WORKDIR /app

COPY package*.json ./
RUN npm ci --only=production

FROM node:20.18.0-alpine AS final

USER node
WORKDIR /home/node
ENV NODE_ENV=production \
    PORT=3000

COPY --chown=node:node --from=build-ts  /app/dist ./
COPY --chown=node:node --from=build-deps  /app/node_modules/ ./node_modules

EXPOSE $PORT

CMD ["node", "index.js"]
