# Stage 1: Build
FROM node:20-alpine AS builder

WORKDIR /app

# Kopier package-filene og installer
COPY package*.json ./
RUN npm install

# Kopier resten og bygg
COPY . .
RUN npm run build

# Stage 2: Runtime (Nginx)
FROM nginx:stable-alpine

# Kopier de bygde filene fra dist-mappa til Nginx sin public mappe
COPY --from=builder /app/dist /usr/share/nginx/html

# Kopier nginx-config
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]