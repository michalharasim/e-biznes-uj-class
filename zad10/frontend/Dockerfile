FROM node:20 AS build
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build
FROM node:20-alpine
RUN npm install -g serve
COPY --from=build /app/build /build
EXPOSE 80
CMD ["serve", "-s", "/build", "-l", "80"]