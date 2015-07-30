#!/bin/bash
url="http://127.0.0.1:15353/metrics"
filter="skydns_dns_request_count,process_cpu_seconds_total,skydns_cache_total_size,skydns_dns_cache_miss_count,skydns_dns_dnssec_ok_count,skydns_dns_request_external_count"

/usr/bin/genmetrics -d ${url} -f ${filter}