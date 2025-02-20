# 定制化Dockerfile
FROM node:18-alpine
RUN npm install -g hexo-cli
COPY ./source /app/source
COPY ./themes /app/themes
WORKDIR /app
RUN npm install && hexo generate
CMD ["hexo", "server", "-p", "4000", "--draft"]