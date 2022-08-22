module arrowhead

go 1.18

replace example.com/password => ./password

require example.com/password v0.0.0-00010101000000-000000000000

require (
	go.etcd.io/bbolt v1.3.6 // indirect
	golang.org/x/sys v0.0.0-20220818161305-2296e01440c6 // indirect
)
