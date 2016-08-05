## watcher ##
watcher of multi machines  by swoole

## Requirements ##

* PHP 5.3.10 or later
* swoole 1.7.9 or later
* libssh2 1.4.2
* ssh2 pecl 0.12

## Installation ##

1. Install swoole via pecl
	
```
pecl install swoole
```

2. Install libssh2

```
wget https://github.com/libssh2/libssh2/archive/libssh2-1.4.2.tar.gz
tar -zvxf libssh2-1.4.2.tar.gz
cd libssh2-libssh2-1.4.2/
./buildconf
./configure
make && sudo make install
```

3. Install ssh2 via pecl

```
wget http://pecl.php.net/get/ssh2-0.12.tgz
tar -zvxf ssh2-0.12.tgz
cd ssh2-0.12
phpize
./configure --with-ssh2
make && make install
```
	
## Protocol ##

```
command get info from swoole websocket server
{
	command: 'get',
	data:	{
		// name of machine1
		name1: {
			// info of machine , such as ['cpu','memory']
			types: [{
					// name of type, like 'cpu'
					'type1' : 'string',
					// interval of crawling
					'interval': 'int',
					start: 'int',     // timestamp
					end: 'int',       // timestamp
				},
				{
					// name of type, like 'memory'
					'type2' : 'string',
					// interval of crawling
					'interval': 'int',
					start: 'int',     // timestamp
					end: 'int',       // timestamp
				}
			]
		},
		// name of machine2
		name2: {
			// info of machine , such as ['cpu','memory']
			types: [{
					// name of type, like 'cpu'
					'type1' : 'string',
					// interval of crawling
					'interval': 'int',
					start: 'int',     // timestamp
					end: 'int',       // timestamp
				},
				{
					// name of type, like 'memory'
					'type2' : 'string',
					// interval of crawling
					'interval': 'int',
					start: 'int',     // timestamp
					end: 'int',       // timestamp
				}
			]
		}
	}
}
data of swoole websocket server returned
{
	command:   'stat',
	code:      1,
	message:   '',
	data:	{
		// name of machine1
		name1: {
			types: {
				// like 'cpu'
				type1: [{
						'int', // timestamp
						'int', // cpu of this moment
					},
					{
						'int', // timestamp
						'int', // cpu of this moment
					}
				],
				// like 'memory'
				type2: [{
						'int', // timestamp
						'int', // memory of this moment
					},
					{
						'int', // timestamp
						'int', // memory of this moment
					},
					{
						'int', // timestamp
						'int', // memory of this moment
					}
				]	
			}
		},
		// name of machine2
		name2: {
			types: {
				// like 'cpu'
				type1: [{
						'int', // timestamp
						'int', // cpu of this moment
					},
					{
						'int', // timestamp
						'int', // cpu of this moment
					}
				],
			}
		}
	}
}

set interval of server crawling
{
	command: 'set',
	// second of server crawling
	data: {
		name1:{
			type1: {
				// interval of
				interval: 'int',
			}
		}
	}
}
common response of server
{
	command: 'return',
	code: 1,
	message: 'string'
}
```
