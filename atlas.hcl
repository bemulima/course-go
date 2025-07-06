
data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "ariga.io/atlas-provider-gorm",
    "load",
    "--path", "./modules/user/models",
    "--dialect", "postgres"
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url

  dev = "postgres://postgres:postgres@localhost:5432/dev_course?sslmode=disable"
  url = "postgres://postgres:postgres@localhost:5432/course?sslmode=disable"

  migration {
    dir = "file://modules/user/migrations"
  }

  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}
