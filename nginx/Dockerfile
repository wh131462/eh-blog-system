FROM nginx:1.25-alpine

# 删除默认配置
RUN rm /etc/nginx/conf.d/default.conf

# 复制自定义配置
COPY nginx.conf /etc/nginx/nginx.conf

# 创建缓存目录
RUN mkdir -p /var/cache/nginx

# 设置静态文件目录
RUN mkdir -p /usr/share/nginx/html
RUN mkdir -p /usr/share/nginx/admin

# 暴露端口
EXPOSE 80 443

# 健康检查
HEALTHCHECK --interval=30s --timeout=10s \
  CMD curl -f http://localhost/health || exit 1