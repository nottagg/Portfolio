# Stage 1: Build the React app
FROM node:20 AS build

WORKDIR /app

COPY package*.json ./
RUN yarn install

COPY . .
RUN yarn build

# Stage 2: Serve the build with a lightweight web server
FROM node:20

WORKDIR /app

# Install serve to serve the production build
RUN yarn global add serve

# If using Vite (output is dist), update as follows:
COPY --from=build /app/dist ./dist

EXPOSE 3000

CMD ["serve", "-s", "dist", "-l", "3000"]