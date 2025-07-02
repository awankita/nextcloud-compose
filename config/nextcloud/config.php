<?php
// config/nextcloud/config.php
// Nextcloud Production Configuration

$CONFIG = array (
  'instanceid' => 'nextcloud_production_instance',
  'passwordsalt' => 'CHANGE_THIS_TO_RANDOM_STRING',
  'secret' => 'CHANGE_THIS_TO_RANDOM_STRING',
  'trusted_domains' => explode(',', getenv('TRUSTED_DOMAINS')),
  'datadirectory' => '/var/www/html/data',
  'dbtype' => 'mysql',
  'version' => '29.0.0.0',
  'overwrite.cli.url' => 'https://' . getenv('DOMAIN_NAME'),
  'overwriteprotocol' => 'https',
  'overwritehost' => getenv('DOMAIN_NAME'),
  'overwritewebroot' => '',
  'htaccess.RewriteBase' => '/',
  
  // Database configuration
  'dbname' => getenv('MYSQL_DATABASE'),
  'dbhost' => 'mariadb',
  'dbport' => '',
  'dbtableprefix' => 'oc_',
  'mysql.utf8mb4' => true,
  'dbuser' => getenv('MYSQL_USER'),
  'dbpassword' => getenv('MYSQL_PASSWORD'),
  
  // Redis configuration for caching and file locking
  'memcache.local' => '\OC\Memcache\APCu',
  'memcache.distributed' => '\OC\Memcache\Redis',
  'memcache.locking' => '\OC\Memcache\Redis',
  'redis' => 
  array (
    'host' => 'redis',
    'port' => 6379,
    'password' => getenv('REDIS_PASSWORD'),
    'timeout' => 0.0,
    'dbindex' => 0,
  ),
  
  // File locking
  'filelocking.enabled' => true,
  'filelocking.ttl' => 3600,
  
  // S3 Object Storage
  'objectstore' => array(
    'class' => '\\OC\\Files\\ObjectStore\\S3',
    'arguments' => array(
      'bucket' => getenv('S3_BUCKET'),
      'key' => getenv('S3_ACCESS_KEY'),
      'secret' => getenv('S3_SECRET_KEY'),
      'hostname' => getenv('S3_ENDPOINT'),
      'port' => (int)getenv('S3_PORT'),
      'use_ssl' => getenv('S3_SSL') === 'true',
      'region' => getenv('S3_REGION'),
      'use_path_style' => true,
      'autocreate' => true,
      'sse_c' => false,
    ),
  ),
  
  // Trusted proxies (Nginx)
  'trusted_proxies' => explode(',', getenv('TRUSTED_PROXIES')),
  'forwarded_for_headers' => 
  array (
    0 => 'HTTP_X_FORWARDED_FOR',
    1 => 'HTTP_FORWARDED',
  ),
  
  // Security settings
  'auth.bruteforce.protection.enabled' => true,
  'ratelimit.protection.enabled' => true,
  'check_data_directory_permissions' => true,
  'csrf.disabled' => false,
  'enable_previews' => true,
  'preview_max_x' => 2048,
  'preview_max_y' => 2048,
  'preview_max_scale_factor' => 1,
  'preview_max_filesize_image' => 50,
  
  // Email configuration
  'mail_smtpmode' => 'smtp',
  'mail_smtpsecure' => 'tls',
  'mail_sendmailmode' => 'smtp',
  'mail_from_address' => getenv('MAIL_FROM_ADDRESS'),
  'mail_domain' => getenv('MAIL_DOMAIN'),
  'mail_smtphost' => getenv('SMTP_HOST'),
  'mail_smtpport' => (int)getenv('SMTP_PORT'),
  'mail_smtpauth' => 1,
  'mail_smtpname' => getenv('SMTP_NAME'),
  'mail_smtppassword' => getenv('SMTP_PASSWORD'),
  
  // Performance and caching
  'default_phone_region' => 'US',
  'maintenance_window_start' => 2,
  'upgrade.disable-web' => false,
  'updater.release.channel' => 'stable',
  'has_rebuilt_cache' => true,
  'filesystem_check_changes' => 1,
  'part_file_in_storage' => false,
  'cytpe_detector' => '\OC\Files\Type\Detection',
  
  // App settings
  'appstoreenabled' => true,
  'appstore.experimental.enabled' => false,
  'apps_paths' => 
  array (
    0 => 
    array (
      'path' => '/var/www/html/apps',
      'url' => '/apps',
      'writable' => false,
    ),
    1 => 
    array (
      'path' => '/var/www/html/custom_apps',
      'url' => '/custom_apps',
      'writable' => true,
    ),
  ),
  
  // Logging
  'log_type' => 'file',
  'logfile' => '/var/log/nextcloud/nextcloud.log',
  'loglevel' => 1,
  'log_rotate_size' => 104857600,
  'logdateformat' => 'F d, Y H:i:s',
  'logtimezone' => 'UTC',
  'log_query' => false,
  
  // Session settings
  'session_lifetime' => 86400,
  'session_keepalive' => true,
  'auto_logout' => false,
  'remember_login_cookie_lifetime' => 60*60*24*15,
  'token_auth_enforced' => false,
  
  // Full-text search (Elasticsearch)
  'fulltextsearch.enabled' => true,
  'fulltextsearch_elasticsearch.host' => 'elasticsearch:9200',
  'fulltextsearch_elasticsearch.index' => 'nextcloud',
  
  // Additional security headers
  'htaccess.IgnoreFrontController' => false,
  'allow_local_remote_servers' => false,
  'check_for_working_wellknown_setup' => true,
  'check_for_working_htaccess' => true,
  
  // Performance tuning
  'simpleSignUpLink.shown' => false,
  'sharing.force_share_accept' => false,
  'sharing.enable_share_accept' => false,
  'sharing.managerFactory' => '\OC\Share20\ProviderFactory',
  'sharing.maxAutocompleteResults' => 25,
  'sharing.minSearchStringLength' => 0,
  'sharing.enable_share_mail' => false,
  
  // File handling
  'knowledgebaseenabled' => false,
  'enable_certificate_management' => false,
  'files_external_allow_create_new_local' => false,
  'blacklisted_files' => 
  array (
    0 => '.htaccess',
    1 => 'Thumbs.db',
    2 => 'thumbs.db',
  ),
  
  // Theme and customization
  'theme' => '',
  'default_locale' => 'en_US',
  'default_language' => 'en',
  'force_locale' => 'en_US',
  'force_language' => 'en',
  
  // Two-factor authentication
  'twofactor_enforced' => 'false',
  'twofactor_enforced_groups' => array(),
  'twofactor_enforced_excluded_groups' => array(),
  
  // Activity app settings
  'activity_expire_days' => 365,
  
  // Preview providers
  'enabledPreviewProviders' => 
  array (
    0 => 'OC\Preview\PNG',
    1 => 'OC\Preview\JPEG',
    2 => 'OC\Preview\GIF',
    3 => 'OC\Preview\BMP',
    4 => 'OC\Preview\XBitmap',
    5 => 'OC\Preview\MP3',
    6 => 'OC\Preview\PDF',
    7 => 'OC\Preview\TXT',
    8 => 'OC\Preview\MarkDown',
    9 => 'OC\Preview\OpenDocument',
    10 => 'OC\Preview\Movie',
  ),
  
  // Maintenance mode
  'maintenance' => false,
  'singleuser' => false,

  // Make config read-only for Nextcloud
  'config_is_read_only' => true,
);
?>