if [ ! -d bkb ]
then
mkdir bkb
fi
#Path=$(basename $PWD)
path=$(basename $PWD)
#path=$(echo $Path|tr '[:upper:]' '[:lower:]')
Path=$(echo $path|awk -F"_" '{for(i=1;i<=NF;i++){$i=toupper(substr($i,1,1)) substr($i,2)}} 1' OFS="")
echo $Path
cat <<EOF
fsfsfs
EOF

for i in *.go
do
	cp -b $i bkb/$i.bkb
	{
		cat <<-EOF
			1
			/^[ 	]*package
			s/^.*$/package $path
			w
		EOF
	} |ed $i
done

{ 
	cat <<-EOF
		1
		/Path
		s/.*/Path = "\/$path"/
		/model\.
		s/.*/type  $Path model.$Path
		/Proto
		s/p *[^)]*/p $Path/
		/Protos
		s/p *[^)]*/p \[\]$Path/
		w

	EOF
} |ed model.go
