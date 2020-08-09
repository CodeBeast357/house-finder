#syntax=docker/dockerfile:experimental
FROM node:14.10.1-alpine3.11 as dev

COPY package.json /app/package.json
WORKDIR /app

VOLUME /app/node_modules

RUN --mount=type=cache,target=/root/.npm,id=npm-cache npm install
COPY . /app/.

ENTRYPOINT [ "npm", "run", "serve" ]

FROM node:14.10.1-alpine3.11 as build

COPY package.json /app/package.json
WORKDIR /app
RUN --mount=type=cache,target=/root/.npm,id=npm-cache npm install

COPY . /app/.

RUN npm run build

FROM node:14.10.1-alpine3.11

RUN --mount=type=cache,target=/root/.npm,id=npm-cache npm install -g serve
COPY --from=build /app/dist /app

ENTRYPOINT [ "serve", "-s", "/app", "-l", "8080" ]

