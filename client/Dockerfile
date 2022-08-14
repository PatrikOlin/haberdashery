FROM node:18 AS build-env

WORKDIR /app

COPY package.json package-lock.json ./
RUN npm install --no-optional
COPY . ./
RUN npm run build

FROM nginx:1.19-alpine
COPY --from=build-env /app/dist /usr/share/nginx/html
