#!/bin/bash
url="127.0.0.1:2379/metrics"
filter="etcd_server_file_descriptors_used_total,etcd_rafthttp_message_sent_failed_total,etcd_store_reads_total,etcd_store_writes_total,etcd_server_proposal_failed_total,etcd_store_watchers,etcd_store_expires_total,go_gc_duration_seconds_count,go_gc_duration_seconds_count,process_cpu_seconds_total,process_open_fds"

/usr/bin/genmetrics -d ${url} -f ${filter}