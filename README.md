<h1 align="center">
  <a href="#">
    <picture>
      <source height="125" media="(prefers-color-scheme: dark)" srcset="https://github.com/hive-go/hive/raw/main/assets/logo.jpg" style="border-radius:15px">
      <img height="125" alt="Hive" src="https://github.com/hive-go/hive/raw/main/assets/logo.jpg" style="border-radius:15px">
    </picture>
  </a>
  <br>
  
  
</h1>
<p align="center">
  <em><b>Hive CLI</b> is an CLI that helps build faster in Hive framework
</p>

---

## ⚙️ Installation


```bash
go install github.com/hive-go/hive-cli/
```

## ⚡️ Quickstart

Getting started with Hive Cli is easy. Here's a basic example to create User Module

```bash
  hive-cli generate_resource user
```

Result in folder structure
```bash

 ├── src
 │   └── modules
 │       └── user
 │           ├── user.module.go
 │           ├── user.controller.go
 │           └── user.service.go
 ├── main.go
 ├── go.mod
 └── go.sum
```

 <a href="https://github.com/hive-go/example-project">
📚 Show more code examples
 </a>




