# shop

#### 介绍
网上商城

#### 软件架构
软件架构说明


#### 功能

首页
专题列表、专题详情
分类列表、分类详情
品牌列表、品牌详情
新品首发、人气推荐
优惠券列表、优惠券选择
搜索
商品详情、商品评价、商品分享
购物车
下单
订单列表、订单详情、订单售后
地址、收藏、足迹、意见反馈
客服

#### 使用说明

先导入数据库，sql文件下的sql语句

编译运行：
    go run main.go

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request


#### swagger接口文档

打开router中注释代码

bee run -gendoc=true -downdoc=true

http://localhost:8082/swagger/#!