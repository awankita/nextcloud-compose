package generator

const defaultPHPConfig = `memory_limit = 1024M
max_execution_time = 3600
max_input_time = 3600
max_input_vars = 10000
upload_max_filesize = 10G
post_max_size = 10G
file_uploads = On
session.lifetime = 86400
session.cookie_httponly = 1
session.cookie_secure = 1
session.use_strict_mode = 1
display_errors = Off
display_startup_errors = Off
log_errors = On
error_log = /var/log/nextcloud/php_errors.log
realpath_cache_size = 10M
realpath_cache_ttl = 120
expose_php = Off
allow_url_fopen = Off
allow_url_include = Off
date.timezone = UTC

opcache.enable = 1
opcache.enable_cli = 1
opcache.interned_strings_buffer = 16
opcache.max_accelerated_files = 10000
opcache.memory_consumption = 256
opcache.save_comments = 1
opcache.revalidate_freq = 60
opcache.fast_shutdown = 1
opcache.validate_timestamps = 0
opcache.jit_buffer_size = 100M
opcache.jit = 1235
opcache.max_wasted_percentage = 10
opcache.use_cwd = 1
opcache.validate_root = 1
opcache.file_update_protection = 2
opcache.revalidate_path = 0
opcache.enable_file_override = 0
opcache.optimization_level = 0x7FFFBFFF
opcache.inherited_hack = 1
opcache.dups_fix = 0
opcache.blacklist_filename=""
`

const defaultMariaDBConfig = `[mysqld]
user = mysql
pid-file = /var/run/mysqld/mysqld.pid
socket = /var/run/mysqld/mysqld.sock
port = 3306
basedir = /usr
datadir = /var/lib/mysql
tmpdir = /tmp
lc-messages-dir = /usr/share/mysql
skip-external-locking
character-set-server = utf8mb4
collation-server = utf8mb4_unicode_ci
init-connect = 'SET NAMES utf8mb4'
default-storage-engine = InnoDB
innodb_file_per_table = 1
innodb_buffer_pool_size = 1G
innodb_log_file_size = 256M
innodb_log_buffer_size = 64M
innodb_flush_log_at_trx_commit = 2
innodb_lock_wait_timeout = 120
innodb_flush_method = O_DIRECT
innodb_io_capacity = 200
innodb_io_capacity_max = 400
max_connections = 512
max_user_connections = 450
thread_cache_size = 50
open_files_limit = 65535
table_definition_cache = 4096
table_open_cache = 4096
key_buffer_size = 128M
max_allowed_packet = 1G
thread_stack = 256K
thread_cache_size = 8
myisam_recover_options = BACKUP
max_length_for_sort_data = 4096
sort_buffer_size = 2M
read_buffer_size = 2M
read_rnd_buffer_size = 8M
bulk_insert_buffer_size = 64M
query_cache_type = 0
query_cache_size = 0
general_log = 0
log_error = /var/log/mysql/error.log
slow_query_log = 1
slow_query_log_file = /var/log/mysql/slow.log
long_query_time = 2
log_queries_not_using_indexes = 0
log-bin = mysql-bin
binlog_format = ROW
binlog_expire_logs_seconds = 604800
max_binlog_size = 100M
sync_binlog = 1
skip-name-resolve
local-infile = 0
performance_schema = ON
performance_schema_max_table_instances = 12500
performance_schema_max_table_handles = 4000
interactive_timeout = 28800
wait_timeout = 28800
lock_wait_timeout = 60

[mysql]
default-character-set = utf8mb4

[client]
default-character-set = utf8mb4
port = 3306
socket = /var/run/mysqld/mysqld.sock

[mysqldump]
quick
quote-names
max_allowed_packet = 1G
default-character-set = utf8mb4

[isamchk]
key_buffer_size = 256M
`
