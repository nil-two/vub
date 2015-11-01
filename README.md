vub
===

Install Vim plugin to under the management of
[vim-unbundle](https://github.com/sunaku/vim-unbundle).

	#Install previm to under ~/.vim/bundle
	$ vub https://github.com/kannokanno/previm

	#Install vim-quickrun to under ~/.vim/bundle
	$ vub thinca/vim-quickrun

	#Install vim-go to under ~/.vim/ftbundle/go
	$ vub -f go fatih/vim-go

Usage
-----
	$ vub [OPTION]... URI

	URI:
	  sunaku/vim-unbundle                    # short URI
	  https://github.com/sunaku/vim-unbundle # full URI

	Options:
	  -f, --filetype=TYPE       installing under the ftbundle/TYPE
	  -l, --list                change the behavior to list packages
	  -r, --remove              change the behavior to remove
	  -u, --update              change the behavior to clean update
	      --help                show this help message
	      --version             output version information

Installation
------------

###compiled binary(not yet)

See [releases](https://github.com/kusabashira/vub/releases)

###go get

	go get github.com/kusabashira/vub

License
-------

MIT License

Author
------

kusabashira <kusabashira227@gmail.com>
