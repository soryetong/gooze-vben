<h1 align="center">Gooze-Vben-Admin</h1>

<p align="center"> 基于 vue-vben-admin (Element-plus) +后端 Go(Gin、gorm) 实现的管理后台</p>

<p align="center">实现了用户、菜单、角色、API 权限管理；后端接口可快速移植至老项目中</p>

<br>

## 快速开始

1. 下载

```bash
git clone https://github.com/soryetong/gooze-vben.git
```

2. 启动 admin

> 要求：`node` 版本 20 以上，`pnpm` 9.15 以上

> 我的 `nvm use 20.11`，`pnpm` 9.15.5

```bash
cd gooze-vben/gooze-vben-admin

# 使用项目指定的pnpm版本进行依赖安装
npm i -g corepack

# 安装依赖
pnpm install

# 启动
pnpm dev
```

3. 启动服务器

> 要求：go 1.24

```bash
cd gooze-vben/gooze-vben-api

# 导入 docs/sql/default.sql

# 修改 `configs/config.yaml` 文件，将数据库链接地址修改为你的数据库链接地址。

# ⚠️⚠️⚠️：一定要修改数据库链接地址

sh ./build/scripts/start_server.sh
```

4. 访问第 2 步输出的地址即可

<br>

## 介绍

> 后端介绍 [看这里](./gooze-vben-api/README.md)
>
> 前端直接使用的 `vue-vben-admin` [看这里](https://doc.vben.pro/)

<br>

## 演示地址

> 复制地址访问：http://8.137.16.100:5003/
>
> 默认账号密码: `admin` / `admin`
>
> 注意 📢：该账号下的数据都不可操作，你可以新建账号操作

<br>

## 技术栈

-   后端基于 Golang + Gin + Gorm

> 权限基于 `Casbin` 实现

-   前端基于 Vue3 + TypeScript + Element-plus + Vben Admin

<br>

## 功能列表

> 详细功能可访问演示站进行使用

-   用户登录

![用户登录](./images/login.png)

-   菜单管理

![菜单管理](./images/menu.png)

-   角色管理

![角色管理](./images/role.png)

-   接口管理

![接口管理](./images/api.png)

-   字典管理

![字典管理](./images/dict.png)

-   操作日志

![操作日志](./images/record.png)
