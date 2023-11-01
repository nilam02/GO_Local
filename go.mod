module example.com/sankey-golang-common-lib

go 1.21.3

replace example.com/sankey-golang-common-lib/components => ./components

replace example.com/sankey-golang-common-lib/config => ./config

require github.com/go-sql-driver/mysql v1.7.1 // indirect

replace example.com/sankey-golang-common-lib/validations => ./validations

replace example.com/sankey-golang-common-lib/http => ./http
