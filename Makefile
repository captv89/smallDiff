run-comp:
	go run . --action=delta --old=./temp/old.txt --new=./temp/new.txt --patch=./temp/patch.txt

run-decomp:
	go run . --action=apply --old=./temp/old.txt --new=./temp/new-one.txt --patch=./temp/patch.txt.gz