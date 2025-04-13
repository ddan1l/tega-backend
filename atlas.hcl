data "external_schema" "gorm" {
  program = [
  "go",
  "run",
  "-mod=mod",
  "ariga.io/atlas-provider-gorm",
  "load",
   "--path", "./models",
   "--dialect", "postgres",
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url

  //dev = "postgres://tega:123@db:5432/tega?sslmode=disable&search_path=public"
  url = "postgres://tega:123@localhost:5432/tega?sslmode=disable&search_path=public"

  migration {
    dir = "file://migrations"
  }

  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
  exclude = ["atlas_schema_revisions"]

  diff {
    skip {
        // By default, none of the changes are skipped.
        drop_schema = true
        drop_table  = true
    }
  }
}