# 文件夹对应功能

本项目在 go.mod 中定义为 gomall，导入库时请加： ```import gomall/***```

## dal

数据库操作

## service

各项服务，将数据库操作封装成一个个服务提供给 handler 模块

## handler

处理请求，调用 service 模块，返回响应

## router

路由，将请求分发到对应的 handler

## template

前端模板，不要动

## static

前端所需文件，不要动
