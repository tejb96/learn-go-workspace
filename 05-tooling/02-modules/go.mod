module github.com/yourname/go-course/tooling/moduleslesson

go 1.22

require github.com/yourname/go-course/tooling/moduleschild v0.0.0

replace github.com/yourname/go-course/tooling/moduleschild => ./examples/child
