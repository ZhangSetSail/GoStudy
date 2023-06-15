# GoStudy
本仓库用于存放Go语言学习的各种Demo
```
CREATE TABLE console.service_security_context (
 service_id varchar(32) NULL,
 seccomp_profile varchar(1024) NULL,
 run_as_non_root BOOL NULL,
 allow_privilege_escalation BOOL NULL,
 run_as_user INTEGER NULL,
 run_as_group INTEGER NULL,
 capabilities LONGTEXT NULL,
 read_only_root_filesystem BOOL NULL
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_unicode_ci;
```
