FROM nginx:1.20.1-alpine
WORKDIR /myapp
COPY . .
ADD . .
RUN true
COPY ./nginx.conf /etc/nginx/nginx.conf
