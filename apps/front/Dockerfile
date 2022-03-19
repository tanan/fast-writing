FROM node:14-alpine as build-stage

ADD . /usr/src/front
WORKDIR /usr/src/front
RUN yarn
RUN yarn build

FROM nginx:latest as production-stage
COPY --from=build-stage /usr/src/front/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/nginx.conf
COPY default.conf /etc/nginx/conf.d/default.conf