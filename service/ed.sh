if [ ! -d bkb ]
then
mkdir bkb
fi

path=$(basename $PWD)
Path=$(echo $path|awk -F"_" '{for(i=1;i<=NF;i++){$i=toupper(substr($i,1,1)) substr($i,2)}} 1' OFS="")
oldPath=$(grep  -m1 package *.go|awk '{print $2;exit}')

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
cp -vi ../../model/$oldPath.go ../../model/$path.go
echo "1
/^[ 	]*type
s/type.*struct/type $Path struct/
w
"|ed ../../model/$path.go

cp -v ../../../main.go ../../../main.go.bkb
echo "
1
/component\/$oldPath
t-
s/$oldPath/$path/
w
"|ed ../../../main.go

#Path=$(basename $PWD)
#path=$(echo $Path|tr '[:upper:]' '[:lower:]')
