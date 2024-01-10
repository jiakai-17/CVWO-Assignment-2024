# Start from a node base image
FROM node:20-alpine AS builder

WORKDIR /frontend

# Copy files over
COPY . .

# Install node packages, install serve, build the app, and remove dependencies at the end
RUN npm ci \
    && npm run build \
    && rm -rf node_modules

EXPOSE 3000

FROM nginx:mainline-alpine3.18-slim

COPY ./nginx/default.conf /etc/nginx/conf.d/default.conf

COPY --from=builder /frontend/dist /usr/share/nginx/html
