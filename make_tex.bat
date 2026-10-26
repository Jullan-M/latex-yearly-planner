set PLANNER_YEAR=2027
go run cmd/plannergen/plannergen.go --config "cfg/base.yaml,cfg/template_months_on_side.yaml,cfg/kscribe.mos.default.yaml,cfg/kscribe.mos.default.dailycal.yaml"
python add_daily_passages.py
cd out
xelatex "kscribe.mos.default.dailycal.tex"
xelatex "kscribe.mos.default.dailycal.tex"
cd ..