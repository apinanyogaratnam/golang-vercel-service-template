REPO := golang-vercel-service-template

edit:
	go mod edit -module github.com/apinanyogaratnam/${REPO}
	find . -name "*.go" -type f -exec sed -i 's/github.com\/apinanyogaratnam\/golang-vercel-service-template/github.com\/apinanyogaratnam\/${REPO}/g' {} \;

vercel:
	vercel .
	vercel git connect --yes

setup:
	make edit
	make edit-imports
	make vercel

start:
	vercel dev --listen 8000
