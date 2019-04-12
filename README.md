# go-bigquery-acl

go-bigquery-acl lets you use a [YAML](http://yaml.org/) configuration file to
automatically apply Access Control List (ACL) on [BigQuery](https://cloud.google.com/bigquery/) datasets.

## Run via docker

```
docker run \
       --rm \
       -v "/path/to/your/application_default_credentials.json:/tmp/application_default_credentials.json:ro" \
       -v "/path/to/your/config.yaml:/app/configs/config.yaml:ro" \
       pcorbel/go-bigquery-acl:latest
```

## Install

```
go get -u github.com/pcorbel/go-bigquery-acl
```

## Usage

```
go-bigquery-acl -conf configs/config.yaml 
```

## Examples

### YAML configuration file

```
project: your-project-id

datasets:
  - name: your-dataset-id
    owner:
      group_by_email:
        - group.owner@company.com
      user_by_email:
        - user.owner@company.com
      special_group:
        - projectOwners
    writer:
      group_by_email:
        - group.writer@company.com
      user_by_email:
        - user.writer@company.com
      special_group:
        - projectWriters
    reader:
      group_by_email:
        - group.reader@company.com
      user_by_email:
        - user.reader@company.com
      special_group:
        - projectReaders
    view:
      - dataset_id: a-dataset-id
        view_id: authorized-view-id
```

### Execution

```
$ go-bigquery-acl -conf configs/config.yaml 

BigQuery update information
  Created at:           2019-01-01 12:00:00.000000 +0000 UTC
  Author:               pierre.corbel
  Credentials:          /Users/pierre.corbel/.config/gcloud/application_default_credentials.json
  Configuration file:   ./configs/config.yaml

Updating accesses for your-project-id:your-dataset-id
  Already up-to-date

Updating accesses for your-project-id:your-dataset-id-2
  +group.owner@company.com:OWNER
  -projectReaders:READER
  -projectWriters:WRITER
  -projectOwners:OWNER

BigQuery update result
  Status:               success
```