package envoy

var collectList = map[string]interface{}{
	"envoy_cluster_manager_cluster_added":                       nil,
	"envoy_cluster_manager_cluster_modified":                    nil,
	"envoy_cluster_manager_cluster_removed":                     nil,
	"envoy_cluster_manager_cluster_updated":                     nil,
	"envoy_cluster_manager_cluster_updated_via_merge":           nil,
	"envoy_cluster_manager_update_merge_cancelled":              nil,
	"envoy_cluster_manager_update_out_of_merge_window":          nil,
	"envoy_filesystem_flushed_by_timer":                         nil,
	"envoy_filesystem_reopen_failed":                            nil,
	"envoy_filesystem_write_buffered":                           nil,
	"envoy_filesystem_write_completed":                          nil,
	"envoy_filesystem_write_failed":                             nil,
	"envoy_http1_dropped_headers_with_underscores":              nil,
	"envoy_http1_metadata_not_supported_error":                  nil,
	"envoy_http1_requests_rejected_with_underscores_in_headers": nil,
	"envoy_http1_response_flood":                                nil,
	"envoy_listener_admin_downstream_cx_destroy":                nil,
	"envoy_listener_admin_downstream_cx_overflow":               nil,
	"envoy_listener_admin_downstream_cx_total":                  nil,
	"envoy_listener_admin_downstream_global_cx_overflow":        nil,
	"envoy_listener_admin_downstream_pre_cx_timeout":            nil,
	"envoy_listener_admin_main_thread_downstream_cx_total":      nil,
	"envoy_listener_admin_no_filter_chain_match":                nil,
	"envoy_listener_manager_listener_added":                     nil,
	"envoy_listener_manager_listener_create_failure":            nil,
	"envoy_listener_manager_listener_create_success":            nil,
	"envoy_listener_manager_listener_in_place_updated":          nil,
	"envoy_listener_manager_listener_modified":                  nil,
	"envoy_listener_manager_listener_removed":                   nil,
	"envoy_listener_manager_listener_stopped":                   nil,
	"envoy_runtime_deprecated_feature_use":                      nil,
	"envoy_runtime_load_error":                                  nil,
	"envoy_runtime_load_success":                                nil,
	"envoy_runtime_override_dir_exists":                         nil,
	"envoy_runtime_override_dir_not_exists":                     nil,
	"envoy_server_debug_assertion_failures":                     nil,
	"envoy_server_dynamic_unknown_fields":                       nil,
	"envoy_server_envoy_bug_failures":                           nil,
	"envoy_server_main_thread_watchdog_mega_miss":               nil,
	"envoy_server_main_thread_watchdog_miss":                    nil,
	"envoy_server_static_unknown_fields":                        nil,
	"envoy_server_watchdog_mega_miss":                           nil,
	"envoy_server_watchdog_miss":                                nil,
	"envoy_cluster_manager_active_clusters":                     nil,
	"envoy_cluster_manager_warming_clusters":                    nil,
	"envoy_filesystem_write_total_buffered":                     nil,
	"envoy_listener_admin_downstream_cx_active":                 nil,
	"envoy_listener_admin_downstream_pre_cx_active":             nil,
	"envoy_listener_admin_main_thread_downstream_cx_active":     nil,
	"envoy_listener_manager_total_filter_chains_draining":       nil,
	"envoy_listener_manager_total_listeners_active":             nil,
	"envoy_listener_manager_total_listeners_draining":           nil,
	"envoy_listener_manager_total_listeners_warming":            nil,
	"envoy_listener_manager_workers_started":                    nil,
	"envoy_runtime_admin_overrides_active":                      nil,
	"envoy_runtime_deprecated_feature_seen_since_process_start": nil,
	"envoy_runtime_num_keys":                                    nil,
	"envoy_runtime_num_layers":                                  nil,
	"envoy_server_concurrency":                                  nil,
	"envoy_server_days_until_first_cert_expiring":               nil,
	"envoy_server_hot_restart_epoch":                            nil,
	"envoy_server_hot_restart_generation":                       nil,
	"envoy_server_live":                                         nil,
	"envoy_server_memory_allocated":                             nil,
	"envoy_server_memory_heap_size":                             nil,
	"envoy_server_memory_physical_size":                         nil,
	"envoy_server_parent_connections":                           nil,
	"envoy_server_state":                                        nil,
	"envoy_server_stats_recent_lookups":                         nil,
	"envoy_server_total_connections":                            nil,
	"envoy_server_uptime":                                       nil,
	"envoy_server_version":                                      nil,
	"envoy_listener_admin_downstream_cx_length_ms_sum":          nil,
	"envoy_listener_admin_downstream_cx_length_ms_count":        nil,
	"envoy_server_initialization_time_ms_sum":                   nil,
	"envoy_server_initialization_time_ms_count":                 nil,
}
