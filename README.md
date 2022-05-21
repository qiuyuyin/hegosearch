# HEGO-SEARCH

这里是Hego-search的go版本实现

1. 首先需要将数据集放到dataset目录下,通过执行程序将数据导入
2. 再执行搜索层中的计划

已完成：
 - 分词
 - 索引存储
 - 文档存储
 - 文本搜索

代办：
 - 接入rpc
 - 关联词语
 - 计时
 - 。。。



```text
├─cmd                       
│  └─import 导入数据                
├─data                      
│  ├─dataset 数据集位置               
│  ├─db KV数据库位置
│  │  ├─doc
│  │  └─index
│  ├─doc doc实现
│  ├─index index实现
│  ├─model model层
│  ├─stopwords 停用词
│  ├─storage 存储层
│  └─tokenize 分词层
├─service
│  └─search 搜索层
├─test 测试
└─util 工具
```