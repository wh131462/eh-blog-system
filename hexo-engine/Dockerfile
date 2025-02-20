FROM node:18-alpine

# 禁用Hexo自动更新检查
ENV HEXO_NO_UPDATE_CHECK=1

WORKDIR /app

# 先复制依赖文件
COPY package.json .
COPY _config.yml .
COPY themes/butterfly/package.json themes/butterfly/

# 分阶段安装依赖
RUN npm install -g hexo-cli && \
    npm install && \
    cd themes/butterfly && \
    npm install --production

# 复制全部文件
COPY . .

# 构建优化
RUN hexo clean && \
    hexo generate --bail

EXPOSE 4000
CMD ["hexo", "server", "-p", "4000"]