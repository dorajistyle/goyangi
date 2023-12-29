# Stage 0, "build-stage", based on Node.js, to build and compile the frontend
# inspired by followings
# * https://snyk.io/blog/10-best-practices-to-containerize-nodejs-web-applications-with-docker/
# * https://dev.to/otomato_io/how-to-optimize-production-docker-images-running-nodejs-with-yarn-504b
# * https://github.com/pyb4430/full-stack-fastapi-postgresql/blob/master/%7B%7Bcookiecutter.project_slug%7D%7D/frontend/Dockerfile
FROM node:16-alpine as build-stage

WORKDIR /app

COPY package*.json /app/

COPY yarn.lock /app/

RUN yarn install --immutable --immutable-cache --check-cache .

COPY ./ /app/

ARG FRONTEND_ENV=production

ENV VUE_APP_ENV=${FRONTEND_ENV}

RUN yarn build


FROM nginx:1.24-alpine

COPY --from=build-stage /app/dist/ /usr/share/nginx/html

COPY ./nginx.conf /etc/nginx/conf.d/default.conf
COPY ./nginx-backend-not-found.conf /etc/nginx/extra-conf.d/backend-not-found.conf