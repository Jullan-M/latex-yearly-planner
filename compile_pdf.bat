set PLANNER_YEAR=2026
go run cmd/plannergen/plannergen.go --config "cfg/base.yaml,cfg/template_months_on_side.yaml,cfg/kscribe.mos.default.dailycal.yaml,cfg/kscribe.mos.default.dailycal.yaml"
cd out
xelatex "kscribe.mos.default.dailycal.tex"
xelatex "kscribe.mos.default.dailycal.tex"
cd ..