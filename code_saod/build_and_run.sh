set -x

go build -o ./bin/saod main.go

if [[ $? != 0 ]]; then
    echo "ERROR: build is failed"
    exit 1
fi

if [[ "$1" == "-b" ]]; then
  echo "only build, exit"
  exit 0
fi

if [[ "$1" == "-t" ]]; then
  echo "only test and exit"
  cd test && go test -v 
  exit 0
fi

./bin/saod $@

#dot -Tpng tree.gv &> tree.png
#dot -Tpng heap.gv &> heap.png
#dot -Tpng btree.gv &> btree.png
