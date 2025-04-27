# Envok

Loads environment variables from yaml file.


## Usage

Here is an example yaml file with environment variables:

```yaml
# my.yaml
test_sqlite:
  PMG__DATABASE: sqlite://sqlite.db
  PMG__MEDIA_ROOT: /tmp/test/root/

test_pg:
  PMG__DATABASE: postgres://user:pass@host
  PMG__MEDIA_ROOT: /tmp/test/root

dev_pg:
  PMG__DATABASE: postgres://user2:pass2@host
  PMG__MEDIA_ROOT: /tmp/dev/root
```

To list all profiles available use command:

```bash
$ envok -f my.yaml list
```

Each of the above `test_sqlite`, `test_pg`, `dev_pg` top level keys is
called "profile". You can list content of current profile with:

```bash
$ envok -f my.yaml -p dev_pg show
```

Profile's environment variables will rendered in red if they differ
from what is set in current environment.

To export current profile:

```bash
$ envok -f my.yaml -p dev_pg export
```

Exporting means that `envok` will print environment variables from
respective profile in format `export VAR=value`

## Repeating Blocks

You can reuse repeating block by using yaml's anchors.
Here is an example:

```yaml
# my.yaml
x-common: &common
  PMG__REDIS__URL: redis://127.0.0.1:6734/0
  PMG__MAIN__SECRET_KEY: 1234

profile_1:
  <<: *common
  PMG__DATABASE: sqlite://sqlite.db
  PMG__MEDIA_ROOT: /tmp/test/root/

profile_2:
  <<: *common
  PMG__DATABASE: postgres://user:pass@host
  PMG__MEDIA_ROOT: /tmp/test/root

profile_3:
  PMG__DATABASE: postgres://user2:pass2@host
  PMG__MEDIA_ROOT: /tmp/dev/root
```

In above example, `profile_1` and `profile_2` feature four environment
variables:

```bash
$ envok -f my.yaml -p profile_1 export

export PMG__MEDIA_ROOT=/tmp/test/root/
export PMG__REDIS__URL=redis://127.0.0.1:6734/0
export PMG__MAIN__SECRET_KEY=1234
export PMG__DATABASE=sqlite://sqlite.db
```

```bash
$ envok -f my.yaml -p profile_2 export

export PMG__DATABASE=postgres://user:pass@host
export PMG__MEDIA_ROOT=/tmp/test/root
export PMG__REDIS__URL=redis://127.0.0.1:6734/0
export PMG__MAIN__SECRET_KEY=1234
```

Profile 3 has only two environment variables:
```bash
$ envok -f my.yaml -p profile_3 export

export PMG__DATABASE=postgres://user2:pass2@host
export PMG__MEDIA_ROOT=/tmp/dev/root
```

## Multiple Repeating Blocks

You can have multiple repeating blocks. Here is a realistic example:

```yaml
x-common-be: &common-be
  PAPERMERGE__MAIN__API_PREFIX: /api
  PAPERMERGE__DATABASE__URL: postgresql://coco:jumbo@127.0.0.1:5432/db
  PAPERMERGE__MAIN__MEDIA_ROOT: /home/eugen/tmp/data4roles

x-common-fe: &common-fe
  VITE_BASE_URL: http://localhost:8000
  VITE_KEEP_UNUSED_DATA_FOR: 1

dev_pg_admin:
  <<: [*common-be, *common-fe]
  VITE_REMOTE_USER: admin
  VITE_REMOTE_USER_ID: 656b72c5-45ed-4206-9d36-26d683d0bf68

lila:
  <<: [*common-be, *common-fe]
  VITE_REMOTE_USER: lila
  VITE_REMOTE_USER_ID: 47f1d7c4-d646-4459-b22c-6462af44a227
  VITE_REMOTE_ROLES: employee

hana:
  <<: [*common-be, *common-fe]
  VITE_REMOTE_USER: hana
  VITE_REMOTE_USER_ID: 6b8a0668-3476-48f2-befd-96aa22174bfd
  VITE_REMOTE_ROLES: employee

mark:
  <<: [*common-be, *common-fe]
  VITE_REMOTE_USER: mark
  VITE_REMOTE_USER_ID: 34b4895c-071c-4ce7-a827-e578d2ed400c
  VITE_REMOTE_ROLES: employee

david:
  <<: [*common-be, *common-fe]
  VITE_REMOTE_USER: david
  VITE_REMOTE_USER_ID: f42881bc-7d71-471e-b4e9-13095f6ada99
  VITE_REMOTE_ROLES: employee

dev_pg_eugen:
  <<: [*common-be, *common-fe]
  VITE_REMOTE_USER: eugen
  VITE_REMOTE_USER_ID: 1bcb3491-b161-4225-bddc-2f4bc5f9b30e

dev_pg_wife:
  <<: [*common-be, *common-fe]
  VITE_REMOTE_USER: wife
  VITE_REMOTE_USER_ID: 104b3bca-6f26-4dd2-918a-be72c36b9af5

test_pg:
  PAPERMERGE__DATABASE__URL: postgresql://coco:kesha@127.0.0.1:5432/db4test
  PAPERMERGE__MAIN__MEDIA_ROOT: /home/eugen/tmp/data4test
  PAPERMERGE__MAIN__API_PREFIX: /api

test_sqlite:
  PAPERMERGE__DATABASE__URL: sqlite:///test.sqlite3
  PAPERMERGE__MAIN__MEDIA_ROOT: /home/eugen/tmp/media_root_1
  PAPERMERGE__MAIN__API_PREFIX: /api

dev_sqlite:
  PAPERMERGE__DATABASE__URL: sqlite:///dev.sqlite3
  PAPERMERGE__MAIN__MEDIA_ROOT: /home/eugen/tmp/media_root_1
  PAPERMERGE__MAIN__API_PREFIX: /api
```