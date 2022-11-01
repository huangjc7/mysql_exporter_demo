* 监控:
  * 可用性
  * 延迟
    * 请求消耗时间
    * 操作使用时间
* 错误次数
* 容量
  * 当前请求多少/总请求多少
  * 当前连接数量/总的连接数量
  
* mysql => exporter =>
    * 监控对象api => 获取指标信息(计算)
    * show global status
* mysql 可用性
  * 操作失败
    * select 1;
    * ping
* 容量
  * qps:
    * show global status where Variable_name like '%Queries%';
    * tps:
      * insert, update, delete *
      * com_insert
      * com_update
      * com_delete
      * com_select
      * com_replaca
    连接
      * Threads_running
      * show global variables where Variable_name like 'max_connections';