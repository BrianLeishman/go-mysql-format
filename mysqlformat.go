package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func isAlpha(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

func isXdigit(b byte) bool {
	return (b >= '0' && b <= '9') || (b >= 'a' && b <= 'f') || (b >= 'A' && b <= 'F')
}

func isAlnum(b byte) bool {
	return (b >= '0' && b <= '9') || (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func isNumeric(b byte) bool {
	return (b >= '0' && b <= '9') || b == '.'
}

func isQuote(b byte) bool {
	return b == '\'' || b == '"' || b == '`'
}

type token struct {
	value     string
	tokenType int
}

const tokenString = 0
const tokenName = 1
const tokenNumeric = 2
const tokenOperator = 3
const tokenWord = 4
const tokenFunction = 5
const tokenBinary = 6
const tokenNewline = 7

var operators = map[byte]int{'!': 1, '%': 1, '&': 1, '*': 1, '+': 1, ',': 1, '-': 1, '.': 1,
	'/': 1, ':': 1, '<': 1, '=': 1, '>': 1, '^': 1, '|': 1, '~': 1, ';': 1, '(': 1, ')': 1}
var keywords = map[string]int{"accessible": 1, "account": 1, "action": 1, "add": 1, "after": 1, "against": 1, "aggregate": 1, "algorithm": 1,
	"all": 1, "alter": 1, "always": 1, "analyse": 1, "analyze": 1, "and": 1, "any": 1, "as": 1, "asc": 1, "ascii": 1, "asensitive": 1, "at": 1,
	"autoextend_size": 1, "auto_increment": 1, "avg": 1, "avg_row_length": 1, "backup": 1, "before": 1, "begin": 1, "between": 1, "bigint": 1,
	"binary": 1, "binlog": 1, "bit": 1, "blob": 1, "block": 1, "bool": 1, "boolean": 1, "both": 1, "btree": 1, "by": 1, "byte": 1, "cache": 1,
	"call": 1, "cascade": 1, "cascaded": 1, "case": 1, "catalog_name": 1, "chain": 1, "change": 1, "changed": 1, "channel": 1, "char": 1,
	"character": 1, "charset": 1, "check": 1, "checksum": 1, "cipher": 1, "class_origin": 1, "client": 1, "close": 1, "coalesce": 1, "code": 1,
	"collate": 1, "collation": 1, "column": 1, "columns": 1, "column_format": 1, "column_name": 1, "comment": 1, "commit": 1, "committed": 1,
	"compact": 1, "completion": 1, "compressed": 1, "compression": 1, "concurrent": 1, "condition": 1, "connection": 1, "consistent": 1,
	"constraint": 1, "constraint_catalog": 1, "constraint_name": 1, "constraint_schema": 1, "contains": 1, "context": 1, "continue": 1,
	"convert": 1, "cpu": 1, "create": 1, "cross": 1, "cube": 1, "current": 1, "current_date": 1, "current_time": 1, "current_timestamp": 1,
	"current_user": 1, "cursor": 1, "cursor_name": 1, "data": 1, "database": 1, "databases": 1, "datafile": 1, "date": 1, "datetime": 1,
	"day": 1, "day_hour": 1, "day_microsecond": 1, "day_minute": 1, "day_second": 1, "deallocate": 1, "dec": 1, "decimal": 1, "declare": 1,
	"default": 1, "default_auth": 1, "definer": 1, "delayed": 1, "delay_key_write": 1, "delete": 1, "desc": 1, "describe": 1, "des_key_file": 1,
	"deterministic": 1, "diagnostics": 1, "directory": 1, "disable": 1, "discard": 1, "disk": 1, "distinct": 1, "distinctrow": 1, "div": 1,
	"do": 1, "double": 1, "drop": 1, "dual": 1, "dumpfile": 1, "duplicate": 1, "dynamic": 1, "each": 1, "else": 1, "elseif": 1, "enable": 1,
	"enclosed": 1, "encryption": 1, "end": 1, "ends": 1, "engine": 1, "engines": 1, "enum": 1, "error": 1, "errors": 1, "escape": 1, "escaped": 1,
	"event": 1, "events": 1, "every": 1, "exchange": 1, "execute": 1, "exists": 1, "exit": 1, "expansion": 1, "expire": 1, "explain": 1,
	"export": 1, "extended": 1, "extent_size": 1, "false": 1, "fast": 1, "faults": 1, "fetch": 1, "fields": 1, "file": 1, "file_block_size": 1,
	"filter": 1, "first": 1, "fixed": 1, "float": 1, "float4": 1, "float8": 1, "flush": 1, "follows": 1, "for": 1, "force": 1, "foreign": 1,
	"format": 1, "found": 1, "from": 1, "full": 1, "fulltext": 1, "function": 1, "general": 1, "generated": 1, "geometry": 1,
	"geometrycollection": 1, "get": 1, "get_format": 1, "global": 1, "grant": 1, "grants": 1, "group": 1, "group_replication": 1,
	"handler": 1, "hash": 1, "having": 1, "help": 1, "high_priority": 1, "host": 1, "hosts": 1, "hour": 1, "hour_microsecond": 1,
	"hour_minute": 1, "hour_second": 1, "identified": 1, "if": 1, "ignore": 1, "ignore_server_ids": 1, "import": 1, "in": 1, "index": 1,
	"indexes": 1, "infile": 1, "initial_size": 1, "inner": 1, "inout": 1, "insensitive": 1, "insert": 1, "insert_method": 1, "install": 1,
	"instance": 1, "int": 1, "int1": 1, "int2": 1, "int3": 1, "int4": 1, "int8": 1, "integer": 1, "interval": 1, "into": 1, "invoker": 1,
	"io": 1, "io_after_gtids": 1, "io_before_gtids": 1, "io_thread": 1, "ipc": 1, "is": 1, "isolation": 1, "issuer": 1, "iterate": 1,
	"join": 1, "json": 1, "key": 1, "keys": 1, "key_block_size": 1, "kill": 1, "language": 1, "last": 1, "leading": 1, "leave": 1,
	"leaves": 1, "left": 1, "less": 1, "level": 1, "like": 1, "limit": 1, "linear": 1, "lines": 1, "linestring": 1, "list": 1, "load": 1,
	"local": 1, "localtime": 1, "localtimestamp": 1, "lock": 1, "locks": 1, "logfile": 1, "logs": 1, "long": 1, "longblob": 1, "longtext": 1,
	"loop": 1, "low_priority": 1, "master": 1, "master_auto_position": 1, "master_bind": 1, "master_connect_retry": 1, "master_delay": 1,
	"master_heartbeat_period": 1, "master_host": 1, "master_log_file": 1, "master_log_pos": 1, "master_password": 1, "master_port": 1,
	"master_retry_count": 1, "master_server_id": 1, "master_ssl": 1, "master_ssl_ca": 1, "master_ssl_capath": 1, "master_ssl_cert": 1,
	"master_ssl_cipher": 1, "master_ssl_crl": 1, "master_ssl_crlpath": 1, "master_ssl_key": 1, "master_ssl_verify_server_cert": 1,
	"master_tls_version": 1, "master_user": 1, "match": 1, "maxvalue": 1, "max_connections_per_hour": 1, "max_queries_per_hour": 1,
	"max_rows": 1, "max_size": 1, "max_statement_time": 1, "max_updates_per_hour": 1, "max_user_connections": 1, "medium": 1, "mediumblob": 1,
	"mediumint": 1, "mediumtext": 1, "memory": 1, "merge": 1, "message_text": 1, "microsecond": 1, "middleint": 1, "migrate": 1, "minute": 1,
	"minute_microsecond": 1, "minute_second": 1, "min_rows": 1, "mod": 1, "mode": 1, "modifies": 1, "modify": 1, "month": 1, "multilinestring": 1,
	"multipoint": 1, "multipolygon": 1, "mutex": 1, "mysql_errno": 1, "name": 1, "names": 1, "national": 1, "natural": 1, "nchar": 1, "ndb": 1,
	"ndbcluster": 1, "never": 1, "new": 1, "next": 1, "no": 1, "nodegroup": 1, "nonblocking": 1, "none": 1, "not": 1, "no_wait": 1,
	"no_write_to_binlog": 1, "null": 1, "number": 1, "numeric": 1, "nvarchar": 1, "offset": 1, "old_password": 1, "on": 1, "one": 1, "only": 1,
	"open": 1, "optimize": 1, "optimizer_costs": 1, "option": 1, "optionally": 1, "options": 1, "or": 1, "order": 1, "out": 1, "outer": 1,
	"outfile": 1, "owner": 1, "pack_keys": 1, "page": 1, "parser": 1, "parse_gcol_expr": 1, "partial": 1, "partition": 1, "partitioning": 1,
	"partitions": 1, "password": 1, "phase": 1, "plugin": 1, "plugins": 1, "plugin_dir": 1, "point": 1, "polygon": 1, "port": 1,
	"precedes": 1, "precision": 1, "prepare": 1, "preserve": 1, "prev": 1, "primary": 1, "privileges": 1, "procedure": 1, "processlist": 1,
	"profile": 1, "profiles": 1, "proxy": 1, "purge": 1, "quarter": 1, "query": 1, "quick": 1, "range": 1, "read": 1, "reads": 1,
	"read_only": 1, "read_write": 1, "real": 1, "rebuild": 1, "recover": 1, "redofile": 1, "redo_buffer_size": 1, "redundant": 1,
	"references": 1, "regexp": 1, "relay": 1, "relaylog": 1, "relay_log_file": 1, "relay_log_pos": 1, "relay_thread": 1, "release": 1,
	"reload": 1, "remove": 1, "rename": 1, "reorganize": 1, "repair": 1, "repeat": 1, "repeatable": 1, "replace": 1, "replicate_do_db": 1,
	"replicate_do_table": 1, "replicate_ignore_db": 1, "replicate_ignore_table": 1, "replicate_rewrite_db": 1, "replicate_wild_do_table": 1,
	"replicate_wild_ignore_table": 1, "replication": 1, "require": 1, "reset": 1, "resignal": 1, "restore": 1, "restrict": 1, "resume": 1,
	"return": 1, "returned_sqlstate": 1, "returns": 1, "reverse": 1, "revoke": 1, "right": 1, "rlike": 1, "rollback": 1, "rollup": 1,
	"rotate": 1, "routine": 1, "row": 1, "rows": 1, "row_count": 1, "row_format": 1, "rtree": 1, "savepoint": 1, "schedule": 1, "schema": 1,
	"schemas": 1, "schema_name": 1, "second": 1, "second_microsecond": 1, "security": 1, "select": 1, "sensitive": 1, "separator": 1,
	"serial": 1, "serializable": 1, "server": 1, "session": 1, "set": 1, "share": 1, "show": 1, "shutdown": 1, "signal": 1, "signed": 1,
	"simple": 1, "slave": 1, "slow": 1, "smallint": 1, "snapshot": 1, "socket": 1, "some": 1, "soname": 1, "sounds": 1, "source": 1,
	"spatial": 1, "specific": 1, "sql": 1, "sqlexception": 1, "sqlstate": 1, "sqlwarning": 1, "sql_after_gtids": 1, "sql_after_mts_gaps": 1,
	"sql_before_gtids": 1, "sql_big_result": 1, "sql_buffer_result": 1, "sql_cache": 1, "sql_calc_found_rows": 1, "sql_no_cache": 1,
	"sql_small_result": 1, "sql_thread": 1, "sql_tsi_day": 1, "sql_tsi_hour": 1, "sql_tsi_minute": 1, "sql_tsi_month": 1, "sql_tsi_quarter": 1,
	"sql_tsi_second": 1, "sql_tsi_week": 1, "sql_tsi_year": 1, "ssl": 1, "stacked": 1, "start": 1, "starting": 1, "starts": 1, "stats_auto_recalc": 1,
	"stats_persistent": 1, "stats_sample_pages": 1, "status": 1, "stop": 1, "storage": 1, "stored": 1, "straight_join": 1, "string": 1,
	"subclass_origin": 1, "subject": 1, "subpartition": 1, "subpartitions": 1, "super": 1, "suspend": 1, "swaps": 1, "switches": 1,
	"table": 1, "tables": 1, "tablespace": 1, "table_checksum": 1, "table_name": 1, "temporary": 1, "temptable": 1, "terminated": 1,
	"text": 1, "than": 1, "then": 1, "time": 1, "timestamp": 1, "timestampadd": 1, "timestampdiff": 1, "tinyblob": 1, "tinyint": 1,
	"tinytext": 1, "to": 1, "trailing": 1, "transaction": 1, "trigger": 1, "triggers": 1, "true": 1, "truncate": 1, "type": 1, "types": 1,
	"uncommitted": 1, "undefined": 1, "undo": 1, "undofile": 1, "undo_buffer_size": 1, "unicode": 1, "uninstall": 1, "union": 1, "unique": 1,
	"unknown": 1, "unlock": 1, "unsigned": 1, "until": 1, "update": 1, "upgrade": 1, "usage": 1, "use": 1, "user": 1, "user_resources": 1,
	"use_frm": 1, "using": 1, "utc_date": 1, "utc_time": 1, "utc_timestamp": 1, "validation": 1, "value": 1, "values": 1, "varbinary": 1,
	"varchar": 1, "varcharacter": 1, "variables": 1, "varying": 1, "view": 1, "virtual": 1, "wait": 1, "warnings": 1, "week": 1, "weight_string": 1,
	"when": 1, "where": 1, "while": 1, "with": 1, "without": 1, "work": 1, "wrapper": 1, "write": 1, "x509": 1, "xa": 1, "xid": 1, "xml": 1,
	"xor": 1, "year": 1, "year_month": 1, "zerofill": 1}
var functions = map[string]int{"abs": 1, "acos": 1, "adddate": 1, "addtime": 1, "aes_decrypt": 1, "aes_encrypt": 1, "and": 1, "any_value": 1,
	"area": 1, "asbinary": 1, "aswkb": 1, "ascii": 1, "asin": 1, "astext": 1, "aswkt": 1, "asymmetric_decrypt": 1, "asymmetric_derive": 1,
	"asymmetric_encrypt": 1, "asymmetric_sign": 1, "asymmetric_verify": 1, "atan": 1, "atan2": 1, "avg": 1, "benchmark": 1, "between": 1,
	"bin": 1, "binary": 1, "bit_and": 1, "bit_count": 1, "bit_length": 1, "bit_or": 1, "bit_xor": 1, "buffer": 1, "case": 1, "cast": 1,
	"ceil": 1, "ceiling": 1, "centroid": 1, "char": 1, "char_length": 1, "character_length": 1, "charset": 1, "coalesce": 1, "coercibility": 1,
	"collation": 1, "compress": 1, "concat": 1, "concat_ws": 1, "connection_id": 1, "contains": 1, "conv": 1, "convert": 1, "convert_tz": 1,
	"convexhull": 1, "cos": 1, "cot": 1, "count": 1, "crc32": 1, "create_asymmetric_priv_key": 1, "create_asymmetric_pub_key": 1,
	"create_dh_parameters": 1, "create_digest": 1, "crosses": 1, "curdate": 1, "current_date": 1, "current_time": 1, "current_timestamp": 1,
	"current_user": 1, "curtime": 1, "database": 1, "date": 1, "date_add": 1, "date_format": 1, "date_sub": 1, "datediff": 1, "day": 1,
	"dayname": 1, "dayofmonth": 1, "dayofweek": 1, "dayofyear": 1, "decode": 1, "default": 1, "degrees": 1, "des_decrypt": 1, "des_encrypt": 1,
	"dimension": 1, "disjoint": 1, "distance": 1, "div": 1, "elt": 1, "encode": 1, "encrypt": 1, "endpoint": 1, "envelope": 1, "equals": 1,
	"exp": 1, "export_set": 1, "exteriorring": 1, "extract": 1, "extractvalue": 1, "field": 1, "find_in_set": 1, "floor": 1, "format": 1,
	"found_rows": 1, "from_base64": 1, "from_days": 1, "from_unixtime": 1, "geomcollfromtext": 1, "geometrycollectionfromtext": 1,
	"geomcollfromwkb": 1, "geometrycollectionfromwkb": 1, "geometrycollection": 1, "geometryn": 1, "geometrytype": 1, "geomfromtext": 1,
	"geometryfromtext": 1, "geomfromwkb": 1, "geometryfromwkb": 1, "get_format": 1, "get_lock": 1, "glength": 1, "greatest": 1,
	"group_concat": 1, "gtid_subset": 1, "gtid_subtract": 1, "hex": 1, "hour": 1, "if": 1, "ifnull": 1, "in": 1, "inet_aton": 1,
	"inet_ntoa": 1, "inet6_aton": 1, "inet6_ntoa": 1, "insert": 1, "instr": 1, "interiorringn": 1, "intersects": 1, "interval": 1, "is": 1,
	"is_free_lock": 1, "is_ipv4": 1, "is_ipv4_compat": 1, "is_ipv4_mapped": 1, "is_ipv6": 1, "not": 1, "null": 1, "is_used_lock": 1,
	"isclosed": 1, "isempty": 1, "isnull": 1, "issimple": 1, "json_append": 1, "json_array": 1, "json_array_append": 1, "json_array_insert": 1,
	"json_arrayagg": 1, "json_contains": 1, "json_contains_path": 1, "json_depth": 1, "json_extract": 1, "json_insert": 1, "json_keys": 1,
	"json_length": 1, "json_merge": 1, "json_merge_patch": 1, "json_merge_preserve": 1, "json_object": 1, "json_objectagg": 1, "json_pretty": 1,
	"json_quote": 1, "json_remove": 1, "json_replace": 1, "json_search": 1, "json_set": 1, "json_storage_size": 1, "json_type": 1,
	"json_unquote": 1, "json_valid": 1, "last_day": 1, "last_insert_id": 1, "lcase": 1, "least": 1, "left": 1, "length": 1, "like": 1,
	"linefromtext": 1, "linestringfromtext": 1, "linefromwkb": 1, "linestringfromwkb": 1, "linestring": 1, "ln": 1, "load_file": 1,
	"localtime": 1, "localtimestamp": 1, "locate": 1, "log": 1, "log10": 1, "log2": 1, "lower": 1, "lpad": 1, "ltrim": 1, "make_set": 1,
	"makedate": 1, "maketime": 1, "master_pos_wait": 1, "match": 1, "max": 1, "mbrcontains": 1, "mbrcoveredby": 1, "mbrcovers": 1,
	"mbrdisjoint": 1, "mbrequal": 1, "mbrequals": 1, "mbrintersects": 1, "mbroverlaps": 1, "mbrtouches": 1, "mbrwithin": 1, "md5": 1,
	"microsecond": 1, "mid": 1, "min": 1, "minute": 1, "mlinefromtext": 1, "multilinestringfromtext": 1, "mlinefromwkb": 1,
	"multilinestringfromwkb": 1, "mod": 1, "month": 1, "monthname": 1, "mpointfromtext": 1, "multipointfromtext": 1, "mpointfromwkb": 1,
	"multipointfromwkb": 1, "mpolyfromtext": 1, "multipolygonfromtext": 1, "mpolyfromwkb": 1, "multipolygonfromwkb": 1, "multilinestring": 1,
	"multipoint": 1, "multipolygon": 1, "name_const": 1, "regexp": 1, "now": 1, "nullif": 1, "numgeometries": 1, "numinteriorrings": 1,
	"numpoints": 1, "oct": 1, "octet_length": 1, "old_password": 1, "or": 1, "ord": 1, "overlaps": 1, "password": 1, "period_add": 1,
	"period_diff": 1, "pi": 1, "point": 1, "pointfromtext": 1, "pointfromwkb": 1, "pointn": 1, "polyfromtext": 1, "polygonfromtext": 1,
	"polyfromwkb": 1, "polygonfromwkb": 1, "polygon": 1, "position": 1, "pow": 1, "power": 1, "procedure": 1, "analyse": 1, "quarter": 1,
	"quote": 1, "radians": 1, "rand": 1, "random_bytes": 1, "release_all_locks": 1, "release_lock": 1, "repeat": 1, "replace": 1,
	"reverse": 1, "right": 1, "rlike": 1, "round": 1, "row_count": 1, "rpad": 1, "rtrim": 1, "schema": 1, "sec_to_time": 1, "second": 1,
	"session_user": 1, "sha1": 1, "sha": 1, "sha2": 1, "sign": 1, "sin": 1, "sleep": 1, "soundex": 1, "sounds": 1, "space": 1, "sqrt": 1,
	"srid": 1, "st_area": 1, "st_asbinary": 1, "st_aswkb": 1, "st_asgeojson": 1, "st_astext": 1, "st_aswkt": 1, "st_buffer": 1,
	"st_buffer_strategy": 1, "st_centroid": 1, "st_contains": 1, "st_convexhull": 1, "st_crosses": 1, "st_difference": 1, "st_dimension": 1,
	"st_disjoint": 1, "st_distance": 1, "st_distance_sphere": 1, "st_endpoint": 1, "st_envelope": 1, "st_equals": 1, "st_exteriorring": 1,
	"st_geohash": 1, "st_geomcollfromtext": 1, "st_geometrycollectionfromtext": 1, "st_geomcollfromtxt": 1, "st_geomcollfromwkb": 1,
	"st_geometrycollectionfromwkb": 1, "st_geometryn": 1, "st_geometrytype": 1, "st_geomfromgeojson": 1, "st_geomfromtext": 1, "st_geometryfromtext": 1,
	"st_geomfromwkb": 1, "st_geometryfromwkb": 1, "st_interiorringn": 1, "st_intersection": 1, "st_intersects": 1, "st_isclosed": 1,
	"st_isempty": 1, "st_issimple": 1, "st_isvalid": 1, "st_latfromgeohash": 1, "st_length": 1, "st_linefromtext": 1, "st_linestringfromtext": 1,
	"st_linefromwkb": 1, "st_linestringfromwkb": 1, "st_longfromgeohash": 1, "st_makeenvelope": 1, "st_mlinefromtext": 1,
	"st_multilinestringfromtext": 1, "st_mlinefromwkb": 1, "st_multilinestringfromwkb": 1, "st_mpointfromtext": 1, "st_multipointfromtext": 1,
	"st_mpointfromwkb": 1, "st_multipointfromwkb": 1, "st_mpolyfromtext": 1, "st_multipolygonfromtext": 1, "st_mpolyfromwkb": 1,
	"st_multipolygonfromwkb": 1, "st_numgeometries": 1, "st_numinteriorring": 1, "st_numinteriorrings": 1, "st_numpoints": 1, "st_overlaps": 1,
	"st_pointfromgeohash": 1, "st_pointfromtext": 1, "st_pointfromwkb": 1, "st_pointn": 1, "st_polyfromtext": 1, "st_polygonfromtext": 1,
	"st_polyfromwkb": 1, "st_polygonfromwkb": 1, "st_simplify": 1, "st_srid": 1, "st_startpoint": 1, "st_symdifference": 1, "st_touches": 1,
	"st_union": 1, "st_validate": 1, "st_within": 1, "st_x": 1, "st_y": 1, "startpoint": 1, "std": 1, "stddev": 1, "stddev_pop": 1,
	"stddev_samp": 1, "str_to_date": 1, "strcmp": 1, "subdate": 1, "substr": 1, "substring": 1, "substring_index": 1, "subtime": 1,
	"sum": 1, "sysdate": 1, "system_user": 1, "tan": 1, "time": 1, "time_format": 1, "time_to_sec": 1, "timediff": 1, "timestamp": 1,
	"timestampadd": 1, "timestampdiff": 1, "to_base64": 1, "to_days": 1, "to_seconds": 1, "touches": 1, "trim": 1, "truncate": 1,
	"ucase": 1, "uncompress": 1, "uncompressed_length": 1, "unhex": 1, "unix_timestamp": 1, "updatexml": 1, "upper": 1, "user": 1,
	"utc_date": 1, "utc_time": 1, "utc_timestamp": 1, "uuid": 1, "uuid_short": 1, "validate_password_strength": 1, "values": 1,
	"var_pop": 1, "var_samp": 1, "variance": 1, "version": 1, "wait_for_executed_gtid_set": 1, "wait_until_sql_thread_after_gtids": 1,
	"week": 1, "weekday": 1, "weekofyear": 1, "weight_string": 1, "within": 1, "x": 1, "xor": 1, "y": 1, "year": 1, "yearweek": 1}

func main() {

	timestart := time.Now()

	mysql := ""
	chunkSize := 4 * 1024

	//nBytes, nChunks := int64(0), int64(0)
	r := bufio.NewReader(os.Stdin)
	buf := make([]byte, 0, chunkSize)
	for {
		n, err := r.Read(buf[:cap(buf)])

		if err != nil && err != io.EOF {
			panic(err.Error())
		}

		if n == 0 {
			if err == nil {
				continue
			}
			if err == io.EOF {
				break
			}
		}

		mysql += string(buf[:n])
	}

	tokens := make([]token, 0)

	var i int
	l := len(mysql)

	for i < l {

		switch b := mysql[i]; {
		case isAlpha(b):
			s := i
			for i+1 < l && (isAlpha(mysql[i+1]) || mysql[i+1] == '_') {
				i++
			}

			word := mysql[s : i+1]
			lword := strings.ToLower(word)

			isFunction := functions[lword] == 1

			tokenType := tokenName
			if isFunction && i+1 < l && mysql[i] == '(' {
				tokenType = tokenFunction
				word = lword
			} else if keywords[lword] == 1 {
				tokenType = tokenWord
				word = lword
			} else if isFunction {
				tokenType = tokenFunction
				word = lword
			}

			if tokenType != tokenWord || word != "as" {
				tokens = append(tokens, token{value: word, tokenType: tokenType})
			}
		case isQuote(b):
			q := mysql[i]
			i++
			s := i
			for i+1 < l && mysql[i] != q {
				if i+2 < l && mysql[i] == '\\' {
					i++
				}
				i++
			}

			tokenType := tokenString
			if mysql[i] == '`' {
				tokenType = tokenName
			}

			value := mysql[s:i]
			value = strings.Replace(value, "\\"+string(q), string(q), -1)

			tokens = append(tokens, token{value: value, tokenType: tokenType})
		case isNumeric(b):
			s := i

			tokenType := tokenNumeric

			if i+1 < l {
				if mysql[i+1] == 'x' {
					i++
					tokenType = tokenBinary
					for i+1 < l && isXdigit(mysql[i+1]) {
						i++
					}
				} else {
					for i+1 < l && isNumeric(mysql[i+1]) {
						i++
					}
				}
			}
			tokens = append(tokens, token{value: string(mysql[s : i+1]), tokenType: tokenType})
		case operators[b] == 1:
			tokens = append(tokens, token{value: string(mysql[i]), tokenType: tokenOperator})
		}

		i++
	}

	mysql = ""
	//p := 0

	oldTokens := tokens
	tokens = make([]token, 0)

	maxLineLength := 80

	lineLength := 0
	lastNewlineOptionIndex := 0
	lineLengthSinceLastNewlineOption := 0

	for i, t := range oldTokens {

		switch t.tokenType {
		case tokenString, tokenName:
			lineLength += len(t.value) + 2
			lineLengthSinceLastNewlineOption += len(t.value) + 2
		default:
			lineLength += len(t.value)
			lineLengthSinceLastNewlineOption += len(t.value)
		}

		if i > 0 {
			switch t.tokenType {
			case tokenWord:
				if t.value == "from" || t.value == "and" || t.value == "where" || t.value == "straight_join" ||
					t.value == "order" || t.value == "group" || t.value == "limit" || t.value == "select" ||
					t.value == "or" || t.value == "set" || t.value == "values" || t.value == "case" ||
					t.value == "when" || t.value == "end" ||
					t.value == "inner" || t.value == "cross" || t.value == "natural" ||
					(t.value == "join" && (oldTokens[i-1].tokenType != tokenWord ||
						(oldTokens[i-1].value != "inner" && oldTokens[i-1].value != "cross" &&
							oldTokens[i-1].value != "left" && oldTokens[i-1].value != "right" &&
							oldTokens[i-1].value != "outer" && oldTokens[i-1].value != "natural"))) ||
					((t.value == "left" || t.value == "right") && (oldTokens[i-1].tokenType != tokenWord || oldTokens[i-1].value != "natural")) ||
					(t.value == "outer" && (oldTokens[i-1].tokenType != tokenWord ||
						(oldTokens[i-1].value != "left" && oldTokens[i-1].value != "right"))) {
					lineLength = 0
					lastNewlineOptionIndex = 0
					lineLengthSinceLastNewlineOption = 0

					tokens = append(tokens, token{tokenType: tokenNewline})
				}
			case tokenOperator:
				if t.value == "," || t.value == "(" {
					fmt.Println("Found comma", t, "; i:", i,
						"; lastNewlineOptionIndex:", lastNewlineOptionIndex,
						"; lineLength:", lineLength,
						"; lineLengthSinceLastNewlineOption:", lineLengthSinceLastNewlineOption)
					if lastNewlineOptionIndex > 0 && lineLength > maxLineLength {
						tokens = append(tokens, token{})
						copy(tokens[lastNewlineOptionIndex:], tokens[lastNewlineOptionIndex-1:])
						tokens[lastNewlineOptionIndex] = token{tokenType: tokenNewline}

						lineLength = lineLengthSinceLastNewlineOption
						lineLengthSinceLastNewlineOption = 0
						lastNewlineOptionIndex = 0
					} else {
						lastNewlineOptionIndex = i
					}

					fmt.Println("change comma", t, "; i:", i,
						"; lastNewlineOptionIndex:", lastNewlineOptionIndex,
						"; lineLength:", lineLength,
						"; lineLengthSinceLastNewlineOption:", lineLengthSinceLastNewlineOption, "\n")
				}
			}
		}

		tokens = append(tokens, t)
		i++
	}

	for i, t := range tokens {
		switch t.tokenType {
		case tokenNewline:
			mysql += "\n"
		case tokenOperator:
			mysql += t.value
		case tokenWord, tokenFunction, tokenNumeric, tokenBinary:
			if i > 0 && (tokens[i-1].tokenType == tokenWord || tokens[i-1].tokenType == tokenFunction ||
				tokens[i-1].tokenType == tokenNumeric || tokens[i-1].tokenType == tokenBinary) {
				mysql += " "

				if tokens[i-1].tokenType == tokenBinary {
					tokens[i-1].value = strings.ToLower(tokens[i-1].value)
				}
			}

			mysql += t.value
		case tokenName:
			if i > 0 && tokens[i-1].tokenType == tokenName {
				mysql += " "
			}
			mysql += "`" + t.value + "`"
		case tokenString:
			if i > 0 && tokens[i-1].tokenType == tokenString {
				mysql += " "
			}
			mysql += "'" + strings.Replace(t.value, `'`, `\'`, -1) + "'"
		}
	}

	fmt.Println(mysql)

	fmt.Println(time.Now().Sub(timestart))
}
