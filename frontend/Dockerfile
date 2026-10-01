# syntax=docker/dockerfile:1

# ---------- build ----------
FROM node:22-alpine AS build

WORKDIR /src

# VITE_API_URL is baked into the bundle at build time
ARG VITE_API_URL=http://localhost:8080
ENV VITE_API_URL=$VITE_API_URL

COPY package.json package-lock.json ./
RUN npm ci

COPY . .
RUN npm run build

# ---------- runtime ----------
FROM nginx:1.27-alpine

COPY --from=build /src/dist /usr/share/nginx/html

# SPA fallback: serve index.html for client-side routes
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80
