# ontic

## What is Ontic?

Ontic is a dot file management tool. Ontic was and is inspired by the
[...](https://github.com/ingydotnet/...)
project.

## Getting started

First ontic needs to be installed. This can be done with go install:

```
    go install github.com/logie17/ontic
```

Next create a ontic configuration file:

```
    mkdir ~/.ontic; touch ~/.ontic.json;
```

A configuration file needs to be created with a location of your dot files,
an example of a conf file:

```
    {
        "dots": [
    	    {
    	        "path": "loop-dots",
    	        "repo": "git@github.com:logie17/loop-dots.git"
    	    },
    	    {
    	        "path": "logie-stuff",
    	        "repo": "git@github.com:logie17/logie-dots.git"
    	    }
        ]
    }

```

Run the following command to initialize ontic for the first time:

```
    ontic init
```

Next backup your current dot files:

```
    ontic backup
```

Finally install your new dot files specified in the configuration file:

```
    ontic install
```

