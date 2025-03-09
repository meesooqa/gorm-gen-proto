# WIP: gorm-gen-proto

Copy `var/config.yml.example` to `var/config.yml`. Set
```yaml
path_maps: "var/data/maps"
path_tmpl: "var/data/templates"
```

These files should be created:
- `%path_maps%/import.json`,
- `%path_maps%/type.json`,
- `%path_tmpl%/proto3.tmpl`.

Add dependencies, e.g. `pb/proto/google/api`.