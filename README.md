# Envok

Loads environment variables from yaml files.


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

