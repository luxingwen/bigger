rsync -vltzh -e ssh -r --include 'bigger' --exclude '*' ./ ss:/data/lxw/bigger/
rsync -vltzh -e ssh --delete -r $GOPATH/src/bigger/views/* ss:/data/lxw/bigger/views/
rsync -vltzh -e ssh --delete -r $GOPATH/src/bigger/static/* ss:/data/lxw/bigger/static/