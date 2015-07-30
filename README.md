# open-falcon-plugins
A place for open-falcon plugins 

Metrics format see open-falcon project

# genmetrics使用说明

pluginlibs 

	- MetricSingle Falcon接收格式，单条Metric记录
	
	- MetricsEcho  包含所有一次性输出到标准输出的metrics记录，其实就是一个MetricSingle的数组加上一个过滤器
	
	
metricEcho可以设置添加过滤器的列表，以metric的string为过滤条件，可以设置过滤器的模式，filterMode为false为白名单（只在列表内的才放行），true为黑名单（列表内的不放行）



genmetrics

    -获取HTTP API中的监控指标文本，经过promethues API解析出监控数据
    
	-对于除Counter、Gauge外的监控数据直接过滤，不处理
	
	-Counter的数据会把所有的标签下的数值相加，得到和
	
	-Gauge的数据直接将值上报
	
	-MetricEcho根据过滤器，将需要上报的添加到数组中，然后输出到标准输出
	
	
	
参数 -d 获取监控数据的HTTP API，输出的必须是promethues支持的V004文本格式

	-f 需要过滤掉的指标名，多个通过 ‘,’ 分割
	
	-m 过滤器模式，false为白名单，true为黑名单，默认是黑名单
