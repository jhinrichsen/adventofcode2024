GO ?= CGO_ENABLED=0 go

REPORTS := coverage.xml gl-code-quality-report.json govet.json govulncheck.sarif junit.xml staticcheck.json test.log

.PHONY: all
ifdef CI
all: dependencies reports
else
all: tidy test
endif

.PHONY: clean
clean:
	$(GO) clean # remove test results from previous runs so that tests are executed
	-rm $(REPORTS)

.PHONY: bench
bench:
	$(GO) test -bench=. -run="" -benchmem

.PHONY: tidy
tidy: | dep-staticcheck
	test -z $(gofmt -l .)
	$(GO) vet
	staticcheck

.PHONY: test
test:
	$(GO) test -run=Day -short -vet=all

.PHONY: dependencies
dependencies: \
	dep-gitlab-code-quality-report \
	dep-go-junit-report \
	dep-gocover-cobertura \
	dep-govulncheck \
	dep-staticcheck

.PHONY: dep-gitlab-code-quality-report
dep-gitlab-code-quality-report:
	which gitlab-code-quality-report || $(GO) install gitlab.com/jhinrichsen/gitlab-code-quality-report@latest
	gitlab-code-quality-report -version

.PHONY: dep-go-junit-report
dep-go-junit-report:
	which go-junit-report || $(GO) install github.com/jstemmer/go-junit-report/v2@latest
	go-junit-report -version

.PHONY: dep-gocover-cobertura
dep-gocover-cobertura:
	which gocover-cobertura || $(GO) install github.com/boumenot/gocover-cobertura@latest
	# not supported gocover-cobertura -version

.PHONY: dep-govulncheck
dep-govulncheck:
	which govulncheck || $(GO) install golang.org/x/vuln/cmd/govulncheck@latest
	govulncheck -version

.PHONY: dep-staticcheck
dep-staticcheck:
	which staticcheck || $(GO) install honnef.co/go/tools/cmd/staticcheck@latest
	staticcheck -version

.PHONY: reports
reports: $(REPORTS)
	@echo "created Gitlab reports $(REPORTS)"

coverage.txt test.log &:
	-$(GO) test -coverprofile=coverage.txt -covermode count -short -v | tee test.log
#
# Gitlab coverage report
coverage.xml: coverage.txt
	gocover-cobertura < $< > $@

# go vet prints complaints to stderr
govet.json:
	$(GO) vet -json 2> $@

# Gitlab dependency report
govulncheck.sarif: | dep-govulncheck
	govulncheck -format=sarif ./... > $@

# Gitlab test report
junit.xml: test.log
	go-junit-report < $< > $@

# Gitlab code quality report
gl-code-quality-report.json: govet.json staticcheck.json
	gitlab-code-quality-report -govet $< -staticcheck staticcheck.json > $@

staticcheck.json:
	-staticcheck -f json > $@

.SUFFIXES: .peg .go

.peg.go:
	peg -noast -switch -inline -strict -output $@ $<

peg: grammar.go

cpu.profile:
	$(GO) test -run=^$ -bench=Day10Part1$ -benchmem -memprofile mem.profile -cpuprofile $@

.PHONY: prof
prof: cpu.profile
	$(GO) tool pprof $^

.SUFFIXES: .peg .go

.peg.go:
	peg -noast -switch -inline -strict -output $@ $<

peg: grammar.go

.PHONY: totalruntime
totalruntime:
	go test -run=^$$ -bench=Day -benchmem | tee benches/all.txt
	awk -f total.awk < benches/all.txt
