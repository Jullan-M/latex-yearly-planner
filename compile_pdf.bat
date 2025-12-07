set PLANNER_YEAR=2026
go run cmd/plannergen/plannergen.go --config "cfg/base.yaml,cfg/template_breadcrumb.yaml,cfg/kscribe.breadcrumb.default.yaml,cfg/kscribe.breadcrumb.default.dailycal.yaml"
cd out
xelatex "kscribe.breadcrumb.default.dailycal.tex"
cd ..