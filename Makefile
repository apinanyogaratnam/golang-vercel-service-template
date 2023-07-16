REPO := golang-vercel-service-template

edit:
	go mod edit -module github.com/apinanyogaratnam/${REPO}

vercel:
	vercel .
	vercel git connect --yes

setup:
	make edit
	make vercel
