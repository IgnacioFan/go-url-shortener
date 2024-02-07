FROM node:16.3.0-alpine AS build

WORKDIR /front

COPY package*.json /front

RUN npm install

COPY . /front

RUN chmod +x /front/nginx/nginx.conf

RUN npm run build

FROM nginx:alpine

WORKDIR /usr/local/bin

COPY --from=build /front/dist /usr/share/nginx/html

COPY --from=build /front/nginx/nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
