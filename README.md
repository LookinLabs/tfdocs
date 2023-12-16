# Code Docs Generator

This project is a command-line tool named `docs_generator` that generates documentation from tf resources

! This code is not production ready, it's in a development stage

# Example (POC)

Run command `./tfdocs --type tf --path ./terraform_resources --output ./docs` will generate 

```
↪ tree docs/
docs/
├── acm
│   └── main.md
├── route53
│   └── main.md
└── vpc
    └── main.md
```

As an example, vpc main.md will look like:

#### VPC

| VPC Name | VPC Subnet |
| -------- | ---------- |
| insurance_vpc | 10.1.1.0/28 |

In code block:

```

#### VPC

| VPC Name | VPC Subnet |
| -------- | ---------- |
| insurance_vpc | 10.1.1.0/28 |
```


## Getting started

1. Clone repository

```
git clone git@github.com:LHV-Kindlustus-AS/tfdocs.git
```

2. Build application

```
make build
```

3. Test application

```
./tfdocs --type tf --path ./terraform_resources --output ./docs
```

# License

You can read about License [here](LICENSE)